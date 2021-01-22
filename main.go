package main

import (
	"goklaw/engine"
	"goklaw/persisit"
	"goklaw/scheduler"
	"goklaw/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 2,
		ItemChan:    persisit.ItemSaver(),
	}

	e.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun/", ParseFunc: parser.ParseCityList})
}
