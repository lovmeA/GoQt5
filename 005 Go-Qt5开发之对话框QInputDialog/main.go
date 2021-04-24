/*-------------------------------------------------------------------
| Project   : goland
| Author    : 今夕何夕
| QQ/Email  : 184566608<qingyueheji@qq.com>
| Time      : 2020-10-07 18:07
| Describe  : main
+------------------------------------------------------------------*/

package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"os"
)


/*
QInputDialog：
提供了一种简单方便的对话框从用户得到一个值。
输入值可以是字符串,一个数字,或一个项目从一个列表。

这个例子显示一个按钮和一个文本框，用户点击按钮显示一个输入框，用户输入信息会显示在文本框中。
*/

func InitUi() *widgets.QMainWindow {
	// 创建窗口
	app := widgets.NewQMainWindow(nil, 0)

	// 设置窗口的标题
	app.SetWindowTitle("Qt 教程")

	// 设置窗口的位置和大小
	app.SetGeometry2(300, 300, 300, 220)

	// 设置窗口的图标，引用当前目录下的web.png图片
	app.SetWindowIcon(gui.NewQIcon5("images/app.ico"))



	// 布局窗口组件载体
	widget := widgets.NewQWidget(app, core.Qt__Widget)
	//widget.SetGeometry(core.NewQRect4(300, 300, 300, 220))
	widget.SetGeometry2(0, 0, 300, 220)
	app.SetCentralWidget(widget)

	btn := widgets.NewQPushButton2("Dialog", app)
	btn.Move2(20, 20)


	le := widgets.NewQLineEdit(app)
	le.Move2(130, 22)

	btn.ConnectClicked(func(checked bool) {
		var flag = false
		input := widgets.NewQInputDialog(widget, 0)

		text := input.GetText(
			widget,
			"Input Dialog",
			"Enter your name:",
			widgets.QLineEdit__Normal,
			"", &flag,0 , 0)
		// 对话框收到的文本消息会显示在文本框中
		if flag {
			le.SetText(text)
		}
	})


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


