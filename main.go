package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	dFlag  = flag.String("d", "", "目标文件夹")
	suffix = flag.String("o", "", "查找的文件后缀")
	newsuf = flag.String("n", "", "重命名的文件后缀")
)

func init() {
	fmt.Println(`使用方式：rename.exe -d=C:\xx\ss\ty\ -o=do.下载 -n=gif`)
}

func main() {
	flag.Parse()

	if *dFlag == "" {
		fmt.Println("未获取到目标文件夹")
		return
	}
	// fmt.Println(*dFlag)

	if *suffix == "" {
		fmt.Println("未获取到要查找的文件后缀")
		return
	}
	// fmt.Println(*suffix)

	if *newsuf == "" {
		fmt.Println("未获取到要重命名的后缀")
		return
	}
	// fmt.Println(*newsuf)

	files, err := getTxtFile(*dFlag, *suffix)
	if err != nil {
		fmt.Println("请检查目录是否存在")
		return
	}
	if len(files) == 0 {
		fmt.Printf(`未获能从目录 "%s" 中找到后缀是 "%s" 的文件\n`, *dFlag, *suffix)
		return
	}

	for _, v := range files {
		err := os.Rename(v, strings.Replace(v, *suffix, *newsuf, 1))
		if err != nil {
			fmt.Println("faild", err)
		}
	}

	fmt.Println("finished")
}

// 读取文件夹下的特定后缀名文件
func getTxtFile(dir string, suffix string) ([]string, error) {
	var txtFiles []string

	e := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), suffix) {
			txtFiles = append(txtFiles, path)
		}
		return nil
	})

	return txtFiles, e
}
