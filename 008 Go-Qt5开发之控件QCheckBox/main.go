/*-------------------------------------------------------------------
| Project   : goland
| Author    : 今夕何夕
| QQ/Email  : 184566608<qingyueheji@qq.com>
| Time      : 2020-10-07 18:23
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
QCheckBox复选框控件，它有两个状态:打开和关闭，他是一个带有文本标签（Label）的控件。
复选框常用于表示程序中可以启用或禁用的功能。

在我们的示例中,我们将创建一个复选框,将切换窗口标题。
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

	// 状态栏
	app.StatusBar()

	// 复选框组件
	cb := widgets.NewQCheckBox2("Show title", widget)
	cb.Move2(20, 20)
	// 我们有设置窗口标题,所以我们也必须检查复选框。
	//默认情况下,没有设置窗口标题和也没有勾选复选框。
	cb.Toggle()
	// 这个方法会切换窗体的标题。
	//复选框的状态经由state参数传入方法。
	//在勾选复选框时设置窗体标题，取消勾选时就将标题设为空字符串。
	cb.ConnectStateChanged(func(state int) {
		//fmt.Println(reflect.TypeOf(state), state)
		//fmt.Println(reflect.TypeOf(core.Qt__Checked), core.Qt__Checked)
		if state == int(core.Qt__Checked) {
			app.SetWindowTitle("QCheckBox")
		} else {
			app.SetWindowTitle(" ")
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
