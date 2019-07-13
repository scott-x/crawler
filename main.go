package  main

import (
	"github.com/scott-x/crawler/engine"
	"github.com/scott-x/crawler/zhenai/parser"
)

func main(){
	engine.Run(engine.Request{
		Url:"https://www.zhenai.com/zhenghun",
		ParseFunc:parser.CityListParser,
	})
}
