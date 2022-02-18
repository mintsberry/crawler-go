package main

import (
	"bufio"
	"crawler/engine"
	"crawler/parse/douban"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io"
)

func main() {
	//seeds := engine.Request{"https://www.zhenai.com/zhenghun", truelove.ParseCityList}
	seeds := engine.Request{"https://book.douban.com/tag/?view=cloud", douban.ParseHotTag}
	engine.Run(seeds)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	peek, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(peek, "")
	return e
}
