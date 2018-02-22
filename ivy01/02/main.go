package main

import (
	"log"
	"os"

	_ "github.com/silentbalanceyh/go-ivy/ivy01/02/matchers"
	"github.com/silentbalanceyh/go-ivy/ivy01/02/search"
)
// init在main之前被自动调用
func init(){
	// 日志标准输出
	log.SetOutput(os.Stdout)
}

// main主程序入口
func main(){
	// 用特定的项做搜索
	search.Run("president");
}