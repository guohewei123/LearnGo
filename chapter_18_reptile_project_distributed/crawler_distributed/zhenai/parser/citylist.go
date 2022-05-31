package parser

import (
	"crawler_distributied/engine"
	"regexp"
	"strings"
)

const cityListRe = `{"linkContent":"([^"]+)","linkURL":"([^"]+)"}`
//const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z]+)"[^>]+>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParserResult {

	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, match := range matches {
		//result.Items = append(result.Items, "City " + string(match[1]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        strings.Replace(string(match[2]), `\u002F`, "/", -1),
			//Url:        string(match[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}
	return result
}
