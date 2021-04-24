/*-------------------------------------------------------------------
| Project   : goland
| Author    : 今夕何夕
| QQ/Email  : 184566608<qingyueheji@qq.com>
| Time      : 2020-10-07 18:47
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
QComboBox是允许用户从下拉列表中进行选择的控件。

示例中展示了一个QComboBox与一个QLabel，QComboBox控件中有5个选项(Linux系统的几个发行版名称)。
QLabel控件会显示QComboBox中选中的某个选项。
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

	// 标签
	lbl := widgets.NewQLabel2("Ubuntu", widget, 0)
	lbl.Move2(50, 150)

	// 下拉列表
	combo := widgets.NewQComboBox(widget)
	combo.AddItem("Ubuntu", core.NewQVariant15("Ubuntu"))
	combo.AddItem("Mandriva", core.NewQVariant15("Mandriva"))
	combo.AddItem("Fedora", core.NewQVariant15("Fedora"))
	combo.AddItem("Arch", core.NewQVariant15("Arch"))
	combo.AddItem("Gentoo", core.NewQVariant15("Gentoo"))

	combo.Move2(50, 50)
	// 当选中某个条目时会调用方法。
	combo.ConnectActivated2(func(text string) {
		// 在方法中我们将QLabel控件的内容设置为选中的条目，然后调整它的尺寸。
		lbl.SetText(text)
		// 滚动长度
		lbl.AdjustSize()
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


