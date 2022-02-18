package truelove

import (
	"crawler/engine"
	"fmt"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+") target="_blank">([^<]+)</a>`

func ParseCity(content []byte) engine.ParseResult {
	compile := regexp.MustCompile(cityRe)
	matches := compile.FindAllSubmatch(content, -1)
	fmt.Println(string(content))
	for _, match := range matches {
		name := string(match[2])
		result.Items = append(result.Items, match[2])
		result.Request = append(result.Request, engine.Request{string(match[1]), func(bytes []byte) engine.ParseResult {
			return ParseUser(bytes, name)
		}})

	}
	return result
}
