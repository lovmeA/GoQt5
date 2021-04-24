/*-------------------------------------------------------------------
| Project   : goland
| Author    : 今夕何夕
| QQ/Email  : 184566608<qingyueheji@qq.com>
| Time      : 2020-10-07 18:01
| Describe  : main
+------------------------------------------------------------------*/

package main

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"os"
	"reflect"
)

/*
在示例中我们重新实现了 keyPressEvent()事件处理器。

我们按下Escape键打印输出.

在Qt5中常通过重新实现事件处理器来处理事件。

不支持标准事件！
*/

func keyPressEvent(event *gui.QKeyEvent) {
	fmt.Println(reflect.TypeOf(event.Key()), event.Key())
	fmt.Println(reflect.TypeOf(core.Qt__Key_Escape), core.Qt__Key_Escape, int(core.Qt__Key_Escape))
	if event.Key() == 32 {
		fmt.Println("空格")
	}
}

func InitUi() *widgets.QMainWindow {
	// 创建窗口
	app := widgets.NewQMainWindow(nil, 0)

	// 设置窗口的标题
	app.SetWindowTitle("Qt 教程")

	// 设置窗口的位置和大小
	app.SetGeometry2(300, 300, 300, 220)

	// 设置窗口的图标，引用当前目录下的web.png图片
	app.SetWindowIcon(gui.NewQIcon5("images/app.ico"))
	app.ConnectKeyPressEvent(keyPressEvent)

	return app
}

func main() {
	// 创建一个应用程序对象
	// sys.argv参数是一个列表，从命令行输入参数
	widgets.NewQApplication(len(os.Args), os.Args)

	// 初始化窗口
	app := InitUi()

	// 显示组件
	app.Show()

	// 确保应用程序干净的退出
	widgets.QApplication_Exec()
}
