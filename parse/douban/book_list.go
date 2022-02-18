package douban

import (
	"crawler/engine"
	"regexp"
)

var bookListParse = regexp.MustCompile(`<a href="(https://book.douban.com/subject/[\d]+/)" title="[^\"]+"[\s]+onclick="[^\"]+">[\s]*([^\s]+)[\s]*</a>`)

func ParseBookList(content []byte) engine.ParseResult {
	matches := bookListParse.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	for _, match := range matches {
		result.Items = append(result.Items, match[2])
		result.Request = append(result.Request, engine.Request{Url: string(match[1]),
			ParserFunc: ParseBook})
	}

	return result
}
