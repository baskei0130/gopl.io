package main

import (
	"bytes"
	"fmt"
	"strings"
)

// 負ではない10進表記整数文字列にカンマを挿入します
func comma1(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma1(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string {
	var buf bytes.Buffer
	n := len(s)
	for i, v := range s {
		if i > 0 && (n-i)%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteString(string(v))
	}
	return buf.String()
}

func commaFloat(s string) string {
	dot := strings.LastIndex(s, ".")
	return comma1(s[:dot]) + s[dot:]
}

func main() {
	fmt.Println(comma1("420482"))
	fmt.Println(comma2("420482"))
	fmt.Println(comma2("1391420482"))
	fmt.Println(commaFloat("1391420482.43298"))
}
