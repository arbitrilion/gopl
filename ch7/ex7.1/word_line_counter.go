// 练习7.1:
// 使用类似 ByteCounter 的想法，实现单词和行的计数器。
// 实现时考虑使用 bufio.ScanWords
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var text = "hello hello\nlovely word"

// ByteCounter ...
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

// WordCounter ...
type WordCounter int

func (c *WordCounter) Write(input []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Count the words.
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	// fmt.Printf("%d\n", count)
	*c += WordCounter(count) // convert int to WordCounter
	return count, nil
}

// LineCounter ...
type LineCounter int

func (c *LineCounter) Write(input []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanLines)
	// Count the words.
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	// fmt.Printf("%d\n", count)
	*c += LineCounter(count) // convert int to WordCounter
	return count, nil
}

func main() {
	fmt.Println("text:", text)
	var b ByteCounter
	b.Write([]byte(text))
	fmt.Println("byte count:", b)

	var w WordCounter
	w.Write([]byte(text))
	fmt.Println("word count:", w)

	var l LineCounter
	l.Write([]byte(text))
	fmt.Println("line count:", l)
}
