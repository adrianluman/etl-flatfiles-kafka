package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "io"
	// "log"
)

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <filepath>\n",
			os.Args[0])
		os.Exit(1)
	}

	filename := os.Args[1]
	// filename = "./testfile"
	input, _ := os.Open(filename)
	defer input.Close()
	// output_non, _ := os.OpenFile("./star.txt", os.O_WRONLY, 0666)
	// output_hashtag, _ := os.OpenFile("./hashtag.txt", os.O_WRONLY, 0666)
	if !Exists("./star.txt") {
		os.Create("./star.txt")
	}
	output_star, _ := os.OpenFile("./star.txt", os.O_WRONLY, 0666)
	defer output_star.Close()
	// defer output_non.Close()
	// defer output_hashtag.Close()

	i := 0
	line := ""
	j := 0
	var err error

	scanner := bufio.NewScanner(input)
	buffWriter_star := bufio.NewWriter(output_star)
	// buffWriter_non := bufio.NewWriter(output_non)
	// buffWriter_hashtag := bufio.NewWriter(output_hashtag)

	for scanner.Scan() {
		line = scanner.Text()
		s := strings.Split(line, "|")
		if s[0] == "IDX" {
			switch s[4] {
				case "2": {
					j, err = buffWriter_star.WriteString(line + "\n")
					// fmt.Println(s[5:len(s)-1])
				}
			}
			// fmt.Println(s[0])
			// fmt.Println(line)
		}
		// fmt.Println(j)
		buffWriter_star.Flush()
		// buffWriter.Reset(buffWriter)
		// buffWriter = bufio.NewWriter(output)

		// fmt.Println(line)
		i++
		// fmt.Println(i)
	}
	if scanner.Err() != nil {
		fmt.Println("error: %s", scanner.Err())
		fmt.Println(err)
	}

	fmt.Println(i)
	fmt.Println(j)
}
