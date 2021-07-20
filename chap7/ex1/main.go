package main

import (
	"bufio"
	"fmt"
)

type ByteCounter int
type TermCounter int
type LineCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func (c *TermCounter) Write(p []byte) (int, error) {
	n := len(p)

	for len(p) > 0 {
		advance, token, err := bufio.ScanWords(p, true)

		if err != nil {
			return 0, err
		}
		if token != nil {
			*c++
		}

		p = p[advance:]
	}

	return n, nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	n := len(p)

	for len(p) > 0 {
		advance, token, err := bufio.ScanLines(p, true)

		if err != nil {
			return 0, err
		}
		if token != nil {
			*c++
		}

		p = p[advance:]
	}

	return n, nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Keita"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
