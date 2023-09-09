package main

import (
	"bufio"
	"fmt"
	"io"
	"sync"
	"time"
)

func main() {
	pipe()
}

func pipe() {
	reader, writer := io.Pipe()

	var wg sync.WaitGroup

	// 一边 写
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			data := fmt.Sprintf("%d: %v \n", i, time.Now())
			writer.Write([]byte(data))
			time.Sleep(1 * time.Second)
		}
		writer.Close()
	}()

	// 一边 读
	wg.Add(1)
	go func() {
		defer wg.Done()
		bufReader := bufio.NewReader(reader)

		for {
			line, isPrefix, err := bufReader.ReadLine()
			if err != nil {
				if err == io.EOF {

				} else {
					break
				}
			}
			fmt.Println("prefix: ", isPrefix, " line: ", string(line))
			if err == io.EOF {
				break
			}
		}

		reader.Close()
	}()

	wg.Wait()
}
