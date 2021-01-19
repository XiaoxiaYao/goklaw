package main

import (
	"goklaw/engine"
	"goklaw/scheduler"
	"goklaw/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 2,
	}

	e.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun/", ParseFunc: parser.ParseCityList})
}
