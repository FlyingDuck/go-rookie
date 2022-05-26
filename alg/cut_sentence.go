package alg

import (
	"context"
	"strings"
)

const SPACE = " " // 空格
const NUMBERSIGN = "#"
const PLACEHOLDER = ""

func ConvertSentence2ASCII(ctx context.Context, sentence string) (asciiSentence string, replaces map[int]rune) {
	runeStc := []rune(sentence)
	if len(runeStc) == len(sentence) {
		return sentence, nil
	}

	var stcBuilder strings.Builder
	replaces = make(map[int]rune)
	maxByte := 0XFF
	for i, r := range runeStc {
		if r > rune(maxByte) {
			stcBuilder.WriteString(NUMBERSIGN)
			replaces[i] = r
		} else {
			stcBuilder.WriteRune(r)
		}
	}

	return stcBuilder.String(), replaces
}

func doCutSentence(ctx context.Context, sentence string, hlWordIdx []int32, minLen int32) (cuttedSentence string, newHL []int32, keepHead, keepTail bool) {

	asciiSentence, replaces := ConvertSentence2ASCII(ctx, sentence)

	type Piece struct {
		text       string
		start, end int32 // [s, e)
		highLight  bool  // 是否是高亮或挖空的块
	}

	parts := strings.Split(asciiSentence, SPACE)
	pieces := make([]*Piece, 0, len(parts))
	head, tail := -1, -1
	for i, part := range parts {
		if i == 0 {
			piece := &Piece{
				text:      part,
				start:     0,
				end:       int32(len(part)),
			}
			pieces = append(pieces, piece)
		} else {
			piece := &Piece{
				text:      part,
				start:     pieces[i-1].end + int32(len(SPACE)),
				end:       pieces[i-1].end + int32(len(SPACE)) + int32(len(part)),
			}
			pieces = append(pieces, piece)
		}

		if (hlWordIdx[0] >= pieces[i].start && hlWordIdx[0] < pieces[i].end) ||
			(hlWordIdx[1] >= pieces[i].start && hlWordIdx[1] < pieces[i].end) {
			pieces[i].highLight = true
			if head == -1 {
				head, tail = i, i
			} else {
				tail = i
			}
		}
	}

	for head > 0 || tail < len(pieces)-1 {
		selectedLen := pieces[tail].end - pieces[head].start
		if selectedLen >= minLen {
			break
		}
		if (tail-head)&0x1 == 0 {
			if head > 0 {
				head--
			} else if tail < len(pieces)-1 {
				tail++
			}
		} else {
			if tail < len(pieces)-1 {
				tail++
			} else if head > 0 {
				head--
			}
		}
	}

	if head == 1 {
		head = 0
	} else if head > 0 && strings.TrimSpace(pieces[head].text) == "" { // todo 或判定 head 为标点，也可以做一下处理
		head--
	}
	if tail == len(pieces)-2 {
		tail = len(pieces) - 1
	}

	offset := pieces[head].start

	// 新的高亮位置信息
	newHL = []int32{hlWordIdx[0]-offset, hlWordIdx[1]-offset}

	// 新的截断字符串 ASCII
	// ASCII → Runes
	var stcBuilder strings.Builder
	for i := head; i <= tail; i++ {
		stcBuilder.WriteString(pieces[i].text)
		if i != tail {
			stcBuilder.WriteString(SPACE)
		}
	}
	runeSentence := []rune(stcBuilder.String())
	for idx, r := range replaces {
		newIdx := idx - int(offset)
		if newIdx >= len(runeSentence) || newIdx < 0 {
			continue
		}
		runeSentence[newIdx] = r
	}

	keepHead = head == 0
	keepTail = tail == len(pieces)-1
	return string(runeSentence), newHL, keepHead, keepTail
}
