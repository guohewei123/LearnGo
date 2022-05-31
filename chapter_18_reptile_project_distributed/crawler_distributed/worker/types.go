package worker

import (
	"crawler_distributied/config"
	"crawler_distributied/engine"
	"crawler_distributied/zhenai/parser"
	"errors"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Item    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParserResult) ParseResult {
	result := ParseResult{
		Item: r.Items,
	}

	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func DeserializeParseResult(r ParseResult) engine.ParserResult {
	result :=  engine.ParserResult{
		Items:    r.Item,
	}

	for _, req := range r.Requests{
		enginReq, err := DeserializeRequest(req)
		if err != nil {
			//return engine.ParserResult{}, err
			log.Printf("error deserializing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, enginReq)
	}
	return result
}


func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("unknown parser name")
	}
}
