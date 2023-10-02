package main

import (
	"fmt"
	"github.com/yanyiwu/gojieba"
	"strings"
)

func main() {
	var seg = gojieba.NewJieba()
	defer seg.Free()

	var useHmm = true
	var separator = "|"

	var resWords []string
	var sentence = "关于golang如何处理中文分词的具体实现"

	resWords = seg.CutAll(sentence)
	fmt.Printf("%s\t全模式：%s\n", sentence, strings.Join(resWords, separator))

	resWords = seg.Cut(sentence, useHmm)
	fmt.Printf("%s\t精确模式：%s\n", sentence, strings.Join(resWords, separator))

	var addWord = "具体实现"
	seg.AddWord(addWord)
	fmt.Printf("添加新词：%s\n", addWord)
	resWords = seg.Cut(sentence, useHmm)
	fmt.Printf("%s\t精确模式：%s\n", sentence, strings.Join(resWords, separator))

	sentence = "关于Redis中持久化机制，请阐述RDB（快照）的具体原理"
	resWords = seg.Cut(sentence, useHmm)
	fmt.Printf("%s\t新词识别:%s\n", sentence, strings.Join(resWords, separator))

	sentence = "关于Redis中持久化机制，请阐述RDB（快照）的具体原理"
	resWords = seg.CutForSearch(sentence, useHmm)
	fmt.Println(sentence, "\t搜索引擎模式：", strings.Join(resWords, separator))

	sentence = "东北师范大学"
	resWords = seg.Tag(sentence)
	fmt.Println(sentence, "\t词性标注：", strings.Join(resWords, separator))

	words := seg.Tokenize(sentence, gojieba.SearchMode, !useHmm)
	fmt.Println(sentence, "\tTokenize Default Mode 搜索引擎模式：", words)

	words = seg.Tokenize(sentence, gojieba.DefaultMode, !useHmm)
	fmt.Println(sentence, "\tTokenize Default Mode 搜索引擎模式：", words)
}
