/*-------------------------------------------------------------------
| Project   : goland
| Author    : 今夕何夕
| QQ/Email  : 184566608<qingyueheji@qq.com>
| Time      : 2020-10-07 18:27
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
ToggleButton是QPushButton的一种特殊模式。
它是一个有两种状态的按钮：按下与未按下。
通过点击在这两种状态间来回切换。这种功能在某些场景会很实用。


代码中我们创建了三个ToggleButton与一个QWidget。
我们将QWidget的背景色设为黑色。
ToggleButton会切换颜色值中的红色、绿色与蓝色部分。
QWidget的背景颜色依赖于按下的按钮。

这是一个信号使用及其不自然的一个案例。

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

	// 这是初始黑颜色的值。
	col := gui.NewQColor3(0, 0, 0, 0)

	frm := widgets.NewQFrame(widget, 0)
	frm.SetGeometry2(150, 20, 100, 100)
	frm.SetStyleSheet("QWidget { background-color: " + col.Name() + " }")

	// 我们创建一个QPushButton并通过其SetCheckable()方法来得到一个ToggleButton。
	redb := widgets.NewQPushButton2("Red", widget)
	redb.SetCheckable(true)
	redb.Move2(10, 10)
	//redb.Text()

	greenb := widgets.NewQPushButton2("Green", widget)
	greenb.SetCheckable(true)
	greenb.Move2(10, 60)

	blueb := widgets.NewQPushButton2("Blue", widget)
	blueb.SetCheckable(true)
	blueb.Move2(10, 110)

	setColor := func(checked bool) {
		val := 0
		if checked {
			val = 255

		}

		if redb.Pointer() == redb.Sender().Pointer() {
			col.SetRed(val)
		} else if greenb.Pointer() == greenb.Sender().Pointer() {
			col.SetGreen(val)
		} else {
			col.SetBlue(val)
		}

		// 信号使用失败
		//if source.text() == "Red" {
		//	col.SetRed(val)
		//} else if source.text() == "Green" {
		//	col.SetGreen(val)
		//} else {
		//	col.SetBlue(val)
		//}

		frm.SetStyleSheet("QWidget { background-color: " + col.Name() + " }")

	}

	// 将clicked信号连接到用户自定义的方法。我们通过clicked信号操作一个布尔值。
	redb.ConnectClicked(setColor)
	greenb.ConnectClicked(setColor)
	blueb.ConnectClicked(setColor)

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
