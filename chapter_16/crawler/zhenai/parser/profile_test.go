package parser

import (
	"crawler/fetcher"
	"fmt"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := fetcher.Fetch("https://album.zhenai.com/u/1876503328")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", contents)

	//contents, err := ioutil.ReadFile("citylist_test_data.txt")
	//if err != nil {
	//	panic(err)
	//}
	//const cityCont = 470
	//expectedUrls := []string{
	//	"http://www.zhenai.com/zhenghun/aba",
	//	"http://www.zhenai.com/zhenghun/akesu",
	//	"http://www.zhenai.com/zhenghun/alashanmeng",
	//}
	//expectedCities := []string{
	//	"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	//}
	//result := ParseCityList(contents)
	//
	//if len(result.Requests) != cityCont {
	//	t.Errorf("resoult should have %d requests, but had %d", cityCont, len(result.Requests))
	//}
	//
	//for i, url := range expectedUrls {
	//	if result.Requests[i].Url != url {
	//		t.Errorf("expected url #%d: %s, but was %s", i, url, result.Requests[i].Url)
	//	}
	//}
	//
	//if len(result.Items) != cityCont {
	//	t.Errorf("resoult should have %d items, but had %d", cityCont, len(result.Items))
	//}
	//
	//for i, city := range expectedCities {
	//	if result.Items[i].(string) != city {
	//		t.Errorf("expected city #%d: %s, but was %s", i, city, result.Items[i])
	//	}
	//}

}
