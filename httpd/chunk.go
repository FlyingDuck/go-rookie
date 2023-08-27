package httpd

import (
	"bufio"
	"errors"
	"io"
)

type ChunkReader struct {
	bufr *bufio.Reader

	// 当前正在处理的块中还剩多少字节未读
	n int
	// 记录报文主体是否读取完毕
	completed bool
	// 用来读取\r\n
	crlf [2]byte
}

func newChunkReader(bufr *bufio.Reader) *ChunkReader {

	return &ChunkReader{
		bufr: bufr,
	}
}

// 每一分块中包含两部分：
//   - 第一部分为chunk size，代表该块chunk data的长度，利用16进制表示。
//   - 第二部分为chunk data，该区域存储有效载荷，实际欲传输的数据存储在这部分。
//
// chunk size与chunk data之间都利用\r\n作为分割，通过0\r\n\r\n来标记报文主体的结束。
//
// # 以下为body
// 17\r\n							#chunk size
// hello, this is chunked \r\n		#chunk data
// D\r\n							#chunk size
// data sent by \r\n				#chunk data
// 7\r\n							#chunk size
// client!\r\n						#chunk data
// 0\r\n\r\n						#end
func (r *ChunkReader) Read(p []byte) (int, error) {
	if r.completed {
		return 0, io.EOF
	}
	if r.n == 0 {
		n, err := r.getNextChunkSize()
		if err != nil {
			return 0, err
		}
		r.n = n
	}

	if r.n == 0 {
		r.completed = true
		//将最后的CRLF消费掉，防止影响下一个http报文的解析
		err := r.discardCRLF()
		return 0, err
	}

	if len(p) <= r.n { // 当前块剩余的数据大于欲读取的长度
		n, err := r.bufr.Read(p)
		r.n -= n
		return n, err
	} else { // 当前块剩余的数据不够欲读取的长度，将剩余的数据全部取出返回
		n, err := io.ReadFull(r.bufr, p)
		r.n = 0

		err = r.discardCRLF()
		if err != nil {
			return 0, err
		}
		return n, nil
	}
}

func (r *ChunkReader) getNextChunkSize() (chunkSize int, err error) {
	line, err := readLine(r.bufr)
	if err != nil {
		return 0, nil
	}
	//将16进制换算成10进制
	for i := 0; i < len(line); i++ {
		switch {
		case 'a' <= line[i] && line[i] <= 'f':
			chunkSize = chunkSize*16 + int(line[i]-'a') + 10
		case 'A' <= line[i] && line[i] <= 'F':
			chunkSize = chunkSize*16 + int(line[i]-'A') + 10
		case '0' <= line[i] && line[i] <= '9':
			chunkSize = chunkSize*16 + int(line[i]-'0')
		default:
			return 0, errors.New("illegal hex number")
		}
	}
	return chunkSize, nil
}

func (r ChunkReader) discardCRLF() error {
	if _, err := io.ReadFull(r.bufr, r.crlf[:]); err == nil {
		if r.crlf[0] != '\r' || r.crlf[1] != '\n' {
			return errors.New("unsupported encoding format of chunk")
		}
	}
	return nil
}
