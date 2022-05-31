package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	//contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun/aba")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%s\n", contents)

	contents, err := ioutil.ReadFile("city_test_data.txt")
	if err != nil {
		panic(err)
	}
	const itemsCont = 12
	result := ParseCity(contents)

	if len(result.Items) != itemsCont {
		t.Errorf("resoult should have %d items, but had %d", itemsCont, len(result.Items))
	}
	for _, item := range result.Items {
		fmt.Println(item)
	}
}
