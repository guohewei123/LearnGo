package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%s\n", contents)

	contents, err := ioutil.ReadFile("citylist_test_data.txt")
	if err != nil {
		panic(err)
	}
	const cityCont = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	result := ParseCityList(contents)

	if len(result.Requests) != cityCont {
		t.Errorf("resoult should have %d requests, but had %d", cityCont, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s, but was %s", i, url, result.Requests[i].Url)
		}
	}
}
