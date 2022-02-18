package douban

import (
	"crawler/engine"
	"regexp"
)

var hotTagRex = regexp.MustCompile(`<a href="(/tag/[^\"]+)">([^<]+)</a>`)

const baseUrl = "https://book.douban.com"

func ParseHotTag(content []byte) engine.ParseResult {
	matches := hotTagRex.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}

	limit := 0
	for _, match := range matches {
		limit++
		if limit > 2 {
			break
		}
		result.Items = append(result.Items, match[2])
		result.Request = append(result.Request, engine.Request{Url: baseUrl + string(match[1]), ParserFunc: ParseBookList})
	}

	return result
}
