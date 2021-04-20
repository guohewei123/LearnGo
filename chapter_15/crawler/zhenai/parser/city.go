package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
	"strings"
)

const userProfile = `<table><tbody><tr><th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]+>([^<]+)</a></th></tr> <tr><td[^>]+><span[^>]+>性别：</span>([^<]+)</td> <td><span[^>]+>居住地：</span>([^<]+)</td></tr> <tr><td[^>]+><span[^>]+>年龄：</span>([^<]+)</td>  <td><span class="grayL">[^<]+</span>([^<]+)</td></tr> <tr><td[^>]+><span[^>]+>婚况：</span>([^<]+)</td> <td[^>]+><span[^>]+>身   高：</span>([^<]+)</td></tr></tbody></table>`

func ParseCity(contents []byte) engine.ParserResult {

	re := regexp.MustCompile(userProfile)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, match := range matches {
		profile := model.Profile{}
		profile.Name = string(match[2])
		profile.Gender = string(match[3])
		profile.Residence = string(match[4])
		retInt, err := strconv.Atoi(string(match[5]))
		if err == nil {
			profile.Age = retInt
		}
		profile.IncomeOrEducation = string(match[6])
		profile.Marriage = string(match[7])
		height, err := strconv.Atoi(string(match[8]))
		if err == nil {
			profile.Height = height
		}
		result.Items = append(result.Items, profile)
		result.Requests = append(result.Requests, engine.Request{
			Url:        strings.Replace(string(match[1]), `\u002F`, "/", -1),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
