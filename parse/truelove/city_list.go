package truelove

import (
	"crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z]+)" data-v-[^>]*>([^<]+)</a>`

func ParseCityList(content []byte) engine.ParseResult {
	compile := regexp.MustCompile(cityListRe)
	matches := compile.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	limit := 0
	for _, match := range matches {
		if limit > 9 {
			break
		}
		result.Items = append(result.Items, match[2])
		result.Request = append(result.Request, engine.Request{string(match[1]), ParseCity})
		limit++

	}
	return result
}
