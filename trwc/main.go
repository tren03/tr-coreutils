package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func byteCount(f fs.FileInfo) int {
	return int(f.Size())
}

func wordCount(data []byte) int {
	str_data := string(data)
	words := strings.Fields(str_data)
	return len(words)
}

func lineCount(data []byte) int {
	lines := 0
	for _, char := range data {
		if char == '\n' {
			lines++
		}
	}
	return lines
}

func getBytes() {
	total_size := 0
	for i, item := range os.Args {
		if i == 0 || i == 1 {
			continue
		}
		info, err := os.Stat(item)
		if err != nil {
			fmt.Println("invalid file path")
			continue
		}
		if info.IsDir() {
			fmt.Printf("item : %s, Cannot process dir \n", item)
			continue
		}

		nbytes := byteCount(info)

		fmt.Printf("bytes=%d %s \n", nbytes, item)
		total_size += int(info.Size())
	}
	if len(os.Args) > 3 {
		fmt.Println("Total size : ", total_size)
	}

}

func getLines() {
	total_lines := 0
	for i, item := range os.Args {
		if i == 0 || i == 1 {
			continue
		}
		info, err := os.Stat(item)
		if err != nil {
			fmt.Println("invalid file path")
			continue
		}
		if info.IsDir() {
			fmt.Printf("item : %s, Cannot process dir \n", item)
			continue
		}
		// main fmtic
		data, err := os.ReadFile(item)
		if err != nil {
			fmt.Println("some err")
		}

		lines := lineCount(data)
		fmt.Printf("lines=%d %s \n", lines, item)
		total_lines += lines

	}
	if len(os.Args) > 3 {
		fmt.Println("Total lines : ", total_lines)
	}

}
func getWords() {
	total_words := 0
	for i, item := range os.Args {
		if i == 0 || i == 1 {
			continue
		}
		info, err := os.Stat(item)
		if err != nil {
			fmt.Println("invalid file path")
			continue
		}
		if info.IsDir() {
			fmt.Printf("item : %s, Cannot process dir \n", item)
			continue
		}
		// main fmtic
		data, err := os.ReadFile(item)
		if err != nil {
			fmt.Println("some err")
		}
		words := wordCount(data)
		fmt.Printf("words=%d %s \n", words, item)
		total_words += words

	}
	if len(os.Args) > 3 {
		fmt.Println("Total words : ", total_words)
	}

}

func main() {
	if len(os.Args) == 2 {
		// will do default operation later
		item := os.Args[1]
		info, err := os.Stat(item)
		if err != nil {
			fmt.Println("invalid file path")
			return
		}
		if info.IsDir() {
			fmt.Printf("item : %s, Cannot process dir \n", item)
			return
		}
		data, err := os.ReadFile(item)
		fmt.Println(info.Name())
		fmt.Println("bytes",byteCount(info))
		fmt.Println("word",wordCount(data))
		fmt.Println("line",lineCount(data))
		return
	}
	if len(os.Args) == 1 {
		fmt.Println("wecome to trwc :)")
		fmt.Println("-c for bytes")
		fmt.Println("-l for lines")
		fmt.Println("-w for words")
        return
	}

	command := os.Args[1]

	switch command {
	case "-c":
		getBytes()
	case "-l":
		getLines()
	case "-w":
		getWords()
	default:
		fmt.Println("Invalid command")
		fmt.Println("-c for bytes")
		fmt.Println("-l for lines")
		fmt.Println("-w for words")

	}

}
