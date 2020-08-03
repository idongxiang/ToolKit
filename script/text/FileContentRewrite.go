package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const DataFile string = "script/text/files/t.txt"

func main() {
	file, err := os.Open(DataFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)

	chunks := make([]byte, 1024, 1024)

	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	content := string(chunks)

	elements := strings.Split(content, ",")

	count := 0
	filterContent := ""
	for _, e := range elements {
		if strings.HasPrefix(e, "G") || strings.HasPrefix(e, "D") || strings.HasPrefix(e, "C") {
			fmt.Println(e)
			count++
			filterContent = filterContent + "," + e
		}
	}
	fmt.Println(count)
	fmt.Println(filterContent)
}
