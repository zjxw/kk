package main

import (
	"yizuo/manager"
	_ "yizuo/routers"
)

/*
   主执行函数
*/
func main() {
	manager.Run()
}
