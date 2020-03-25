package main

import (
	"flag"
	"fmt"
	"log"
	"qztc-config/param"
)

var (
	tomlFile = flag.String("config", "qa.toml", "config file")
)



func main() {
	flag.Parse()
	// 解析配置文件
	tomlConfig, err := config.UnmarshalConfig(*tomlFile)
	if err != nil {
		log.Println("解析文件失败", err)
		return
	}

	fmt.Print("tomlConfig.GetListenAddr() -> ")
	fmt.Print(tomlConfig.GetListenAddr())
}