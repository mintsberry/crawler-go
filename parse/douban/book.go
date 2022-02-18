package douban

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

var nameReg = regexp.MustCompile(`<span property="v:itemreviewed">([^<]+)</span>`)
var authorReg = regexp.MustCompile(`<a class="" href="/author/[\d]+">([^<]+)</a>`)
var publisherReg = regexp.MustCompile(`<span class="pl">出版社:</span>[\s]*([^<]+)`)
var dataReg = regexp.MustCompile(`<span class="pl">出版年:</span>[\s]*([^<]+)`)
var scoreReg = regexp.MustCompile(`<strong class="ll rating_num " property="v:average">[\s]*([\d.]*)[\s]*</strong>`)
var numReg = regexp.MustCompile(`<a href="comments" class="rating_people"><span property="v:votes">([^<]+)</span>人评价</a>[\s]+</span>`)

func ParseBook(content []byte) engine.ParseResult {
	result := engine.ParseResult{}

	name := extraString(nameReg, content)
	author := extraString(authorReg, content)
	publisher := extraString(publisherReg, content)
	data := extraString(dataReg, content)
	num := extraString(numReg, content)
	score, _ := strconv.ParseFloat(extraString(scoreReg, content), 0)
	gradeNum, _ := strconv.Atoi(string(num))
	book := model.BooK{
		Name:        string(name),
		Author:      string(author),
		Publisher:   string(publisher),
		PublishData: string(data),
		GradeNum:    gradeNum,
		Score:       score,
	}

	result.Items = []interface{}{book}

	return result
}

func extraString(regex *regexp.Regexp, content []byte) string {
	subMatch := regex.FindSubmatch(content)
	if len(subMatch) >= 2 {
		return string(subMatch[1])
	}
	return ""
}
