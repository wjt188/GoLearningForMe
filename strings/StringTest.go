package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRandomStr(n string) string {
	var l = len(n)
	b := make([]byte, l)
	for i := range b {
		rand.Seed(time.Now().UnixNano())
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func PlusConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}
func sprintfConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", s, str)
	}
	return s
}
func builderConcat(n int, str string) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}
func bufferConcat(n int, s string) string {
	buf := new(bytes.Buffer)
	for i := 0; i < n; i++ {
		buf.WriteString(s)
	}
	return buf.String()
}
func byteConcat(n int, str string) string {
	buf := make([]byte, 0)
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}
func main() {
	var str = GenerateRandomStr("abcdefg")
	fmt.Println(str)
	fmt.Println("+拼接", PlusConcat(3, str))
	fmt.Println("fmt.Sprintf:", sprintfConcat(3, str))
	fmt.Println("strings.Builder:", builderConcat(3, str))
	fmt.Println("bytes.Buffer:", bufferConcat(3, str))
	fmt.Println("[]byte:", byteConcat(3, str))

}
