/*-------------------------------------------------------------------
| Project   : goland
| Author    : 今夕何夕
| QQ/Email  : 184566608<qingyueheji@qq.com>
| Time      : 2020-10-07 22:03
| Describe  : utils
+------------------------------------------------------------------*/

package utils

import (
	"fmt"
	"io"
	"os"
)

func WriteFIleString(path, data string) (bool, string) {
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return true, "文件创建失败"
	}
	_, err = f.WriteString(data)
	if err != nil {
		return true, "信息写入失败"
	}
	return false, ""
}

func ReadFileString(path string) (bool, string) {
	file, err := os.Open(path)
	if err != nil {
		return true, "文件打开失败"
	}
	defer file.Close()
	var content []byte
	var buf [1024]byte

	for {
		n, err := file.Read(buf[:])
		if err == io.EOF || err != nil {
			break
		}
		if err != nil {
			return true, "文件读取失败"
		}
		content = append(content, buf[:n]...)
	}
	return false, string(content)
}

func Qss() string {
	file, err := os.Open("theme/theme.css")
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
	}
	defer file.Close()
	var content []byte
	var buf [1024]byte

	for {
		n, err := file.Read(buf[:])
		if err == io.EOF || err != nil {
			break
		}
		if err != nil {
			fmt.Println("文件读取失败", err.Error())
		}
		content = append(content, buf[:n]...)
	}
	return string(content)
}
