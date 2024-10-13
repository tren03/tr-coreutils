package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func byteCount(data []byte) int {
	return len(data)
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

		fmt.Printf("bytes=%d %s \n", info.Size(), item)
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

func execCommand() {
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

func main() {

	// CHECKING STDIN
	content, err := io.ReadAll(os.Stdin)
	if err != nil {
		if err == io.EOF {
			log.Println("reached end")
		} else {
			log.Println("err reading from std", err)
		}
	}
	if len(content) != 0 {
		if len(os.Args) == 1 {
			fmt.Println("bytes", byteCount(content))
			fmt.Println("word", wordCount(content))
			fmt.Println("line", lineCount(content))
			return
		} else if len(os.Args) == 2 {
			command := os.Args[1]

			switch command {
			case "-c":
				fmt.Printf("bytes=%d\n",byteCount(content))
			case "-l":
				fmt.Printf("lines=%d\n",lineCount(content))
            case "-w":
				fmt.Printf("words=%d\n",wordCount(content))
			default:
				fmt.Println("Invalid command")
				fmt.Println("-c for bytes")
				fmt.Println("-l for lines")
				fmt.Println("-w for words")

			}
		}else{
            fmt.Println("too many commands")
            return
        }


		return
	}

	// FILE OPERATIONS
	if len(os.Args) == 1 {
		fmt.Println("wecome to trwc :)")
		fmt.Println("-c for bytes")
		fmt.Println("-l for lines")
		fmt.Println("-w for words")
		return
	}
	if len(os.Args) == 2 {
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
		fmt.Println("bytes", byteCount(data))
		fmt.Println("word", wordCount(data))
		fmt.Println("line", lineCount(data))
		return
	}

    if len(os.Args) > 2{
        fmt.Println("too many commands")
        return
    }
	execCommand()

}
