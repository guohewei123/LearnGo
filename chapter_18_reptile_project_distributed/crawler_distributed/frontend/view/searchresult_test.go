package view

import (
	"crawler/engine"
	"crawler/frontend/model"
	common "crawler/model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template.html")
	out, err := os.Create("template.test.html")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:     "https://album.zhenai.com/u/1876503328",
		Id:      "1876503328",
		Type:    "zhenai",
		Payload: common.Profile{
			Name:              "张三",
			Gender:            "男",
			Residence:         "北京",
			Age:               18,
			IncomeOrEducation: "3000-4000",
			Marriage:          "未婚",
			Height:            180,
		},
	}

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	if err != nil {
		t.Error(err)
	}

	// TODO: verify contents in template.test.html
}
