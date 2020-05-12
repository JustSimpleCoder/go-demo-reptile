package main

import (
	"go-demo-reptile/bussiness/engine"
	"go-demo-reptile/bussiness/parser"
	"go-demo-reptile/bussiness/persist"
	"go-demo-reptile/bussiness/scheduler"
)

func main() {

	e := &engine.ConcurrentEngine{
		WorkCount: 10,
		Scheduler: &scheduler.SimpleScheduler{},
		ItemChan:  persist.ItemSaver(),
	}

	e.Run(engine.Request{Url: "https://www.sina.com.cn/", ParserFunc: parser.ParseBaseUrl})

}
