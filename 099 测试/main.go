/*-------------------------------------------------------------------
| Project   : goland
| Author    : 今夕何夕
| QQ/Email  : 184566608<qingyueheji@qq.com>
| Time      : 2020-10-07 20:08
| Describe  : main
+------------------------------------------------------------------*/

package main

import "fmt"

type UserInfo struct {
}

func (this UserInfo) Name() []*string {
	a := "a"
	b := "b"
	c := "c"
	return []*string{&a, &b, &c}
}

func main() {
	u := UserInfo{}
	names := u.Name()
	for _, name := range names {
		fmt.Println(*name)
	}
}
