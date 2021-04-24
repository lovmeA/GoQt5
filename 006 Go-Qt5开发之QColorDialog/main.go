/*-------------------------------------------------------------------
| Project   : goland
| Author    : 今夕何夕
| QQ/Email  : 184566608<qingyueheji@qq.com>
| Time      : 2020-10-07 18:12
| Describe  : main
+------------------------------------------------------------------*/

package main

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"os"
)

/*
QColorDialog 显示一个用于选择颜色值的对话框。

这个应用程序显示一个按钮和一个QFrame。
QFrame的背景为黑色。
通过QColorDialog,我们可以改变它的背景。

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

	// 初始化QFrame的颜色为黑色
	col := gui.NewQColor3(0, 0, 0, 0)

	frm := widgets.NewQFrame(widget, 0)
	frm.SetGeometry2(130, 22, 100, 100)
	frm.SetStyleSheet("QWidget { background-color: " + col.Name() + " }")

	btn := widgets.NewQPushButton2("Dialog", widget)
	btn.Move2(20, 20)
	btn.ConnectClicked(func(checked bool) {
		col2 := widgets.NewQColorDialog(widget).GetColor(
			col,
			widget,
			"调色板",
			widgets.QColorDialog__ShowAlphaChannel)

		// 我们要先检查col的值。如果点击的是Cancel按钮，返回的颜色值是无效的。
		//当颜色值有效时，我们通过样式表(style sheet)来改变背景颜色。
		if col2.IsValid() {
			fmt.Println("调色： ", col2.Name())
			frm.SetStyleSheet("QWidget { background-color: " + col2.Name() + " }")
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
