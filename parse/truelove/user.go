package truelove

import (
	"crawler/engine"
	"regexp"
)

var MarriageRe = regexp.MustCompile(`<div data-v-[\w]+="" class="m-btn purple">(离异)</div>`)

func ParseUser(content []byte, name string) engine.ParseResult {
	marriage := MarriageRe.Find(content)

	result := engine.ParseResult{
		Items: []interface{}{marriage},
	}
	return result
}
