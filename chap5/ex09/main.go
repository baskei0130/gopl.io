package main

import (
	"bufio"
	"fmt"
	"strings"
)

const str = `My name is Akiko. I am a Japanese high school student. I came to live in America with my family two
weeks ago. One day Mother asked me about my school life. I answered that ①I enjoyed it very much. But
that was not true. I didn't want Mother to worry about me.
I always said to myself, I live in two worlds, one at home and the other at school. They are very ( ).
At home I speak Japanese and live a happy Japanese life with my family. ②I feel that school is far away
when I am at home. In class at school teachers sometimes speak English so fast that I don't understand
them. I want to talk with a friend about that, but I don't have any friends. Every day I wait for a student who
will talk to me.`

func main() {
	repS := expand(str, func(word string) string {
		newW := ""
		for _, r := range word {
			newW += string(r + 1)
		}
		return newW
	})
	fmt.Println(repS)
}

func expand(s string, f func(string) string) string {
	var repS string = ""

	input := bufio.NewScanner(strings.NewReader(s))
	input.Split(bufio.ScanWords)

	for input.Scan() {
		repS += f(input.Text())
		repS += " "
	}

	return repS
}
