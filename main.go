package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 获取程序运行目录
	clientRunDir, getClientRunDirErr := filepath.Abs(filepath.Dir(os.Args[0]))
	if getClientRunDirErr != nil {
		fmt.Println(getClientRunDirErr)
		return
	}
	fmt.Println("clientRunDir: ", clientRunDir)
}
