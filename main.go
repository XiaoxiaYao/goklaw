package main

import (
	"goklaw/engine"
	"goklaw/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun/", ParseFunc: parser.ParseCityList})
}
