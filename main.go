package main

import (
	"crawler/engine"
	"crawler/parse/douban"
	"crawler/scheduler"
)

func main() {
	//seeds := engine.Request{"https://www.zhenai.com/zhenghun", truelove.ParseCityList}
	seeds := engine.Request{"https://book.douban.com/tag/?view=cloud", douban.ParseHotTag}
	concurrentEngine := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkCount: 2,
	}
	concurrentEngine.Run(seeds)

}
