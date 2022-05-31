package persist

import (
	"context"
	"crawler_distributied/engine"
	"crawler_distributied/model"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"testing"
)

func TestSave(t *testing.T) {

	expected := engine.Item{
		Url:     "https://album.zhenai.com/u/1876503328",
		Id:      "1876503328",
		Type:    "zhenai",
		Payload: model.Profile{
			Name:              "张三",
			Gender:            "男",
			Residence:         "北京",
			Age:               18,
			IncomeOrEducation: "3000-4000",
			Marriage:          "未婚",
			Height:            180,
		},
	}

	// save expected item
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.10.53:9200"),
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	err = save(client, "profile_test", expected)
	if err != nil {
		panic(err)
	}
	// TODO: try to start up elastic serch
	// here using docker go client
	// fetch save item
	resp, err := client.Get().
		Index("profile_test").
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	var actual engine.Item
	err = json.Unmarshal(resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, err := model.FromJsonObj(actual.Payload)
	if err != nil {
		panic(err)
	}
	actual.Payload = actualProfile
	// verify result
	if actual != expected {
		t.Errorf("got %v;\n expected: %v", actual, expected)
	}

}
