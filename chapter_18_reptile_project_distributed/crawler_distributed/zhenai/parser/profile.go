package parser

import (
	"crawler_distributied/engine"
	"crawler_distributied/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([\d]+)岁</div>`)
var marriageRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([^<]+)</div>`)

func ParseProfile(contents []byte) engine.ParserResult {
	profile := model.Profile{}
	profile.Age = extractInt(contents, ageRe)

	profile.Marriage = extractString(contents, marriageRe)
	return engine.ParserResult{
		//Items: []interface{}{profile},
	}
}

func extractString(c []byte, re *regexp.Regexp) string {
	//return string(re.FindSubmatch(c)[1])
	match := re.FindSubmatch(c)
	if match != nil && len(match) > 1 {
		return string(match[1])
	}
	return ""
}

func extractInt(c []byte, re *regexp.Regexp) int {
	match := re.FindSubmatch(c)
	if match != nil && len(match) > 2 {
		retInt, err := strconv.Atoi(string(match[1]))
		if err != nil {
			return retInt
		}
	}
	return -1
}
