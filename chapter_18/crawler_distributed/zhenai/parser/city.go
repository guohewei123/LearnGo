package parser

import (
	"crawler_distributied/engine"
	"crawler_distributied/model"
	"regexp"
	"strconv"
	"strings"
)

var ProfileRe = regexp.MustCompile(
	`<table><tbody><tr><th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]+>([^<]+)</a></th></tr> <tr><td[^>]+><span[^>]+>性别：</span>([^<]+)</td> <td><span[^>]+>居住地：</span>([^<]+)</td></tr> <tr><td[^>]+><span[^>]+>年龄：</span>([^<]+)</td>  <td><span class="grayL">[^<]+</span>([^<]+)</td></tr> <tr><td[^>]+><span[^>]+>婚况：</span>([^<]+)</td> <td[^>]+><span[^>]+>身   高：</span>([^<]+)</td></tr></tbody></table>`)
var cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([0-9]+)`)

func ParseCity(contents []byte) engine.ParserResult {

	matches := ProfileRe.FindAllSubmatch(contents, -1)
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
		itemProfile := engine.Item{
			Url: string(match[1]),
			//Id:      string(idUrlRe.FindSubmatch(match[1])[1]),
			Id:      extractString(match[1], idUrlRe),
			Type:    "zhenai",
			Payload: profile,
		}
		result.Items = append(result.Items, itemProfile)
		//result.Requests = append(result.Requests, engine.Request{
		//	Url:        strings.Replace(string(match[1]), `\u002F`, "/", -1),
		//	ParserFunc: engine.NilParser,
		//})
	}

	urlMatches := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, match := range urlMatches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    strings.Replace(string(match[1]), `\u002F`, "/", -1),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}
	return result
}
