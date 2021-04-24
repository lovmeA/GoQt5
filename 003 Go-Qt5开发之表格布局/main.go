/*-------------------------------------------------------------------
| Project   : goland
| Author    : 今夕何夕
| QQ/Email  : 184566608<qingyueheji@qq.com>
| Time      : 2020-10-07 17:19
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
这是一个使用信号槽的Qt5例子。

这个例子中展示了一个QtGui.QLCDNumber和QtGui.QSlider。lcd的值会随着滑块的拖动而改变。
在这里我们将滚动条的valueChanged信号连接到lcd的display插槽。
sender是发出信号的对象。receiver是接收信号的对象。
slot(插槽)是对信号做出反应的方法。

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

	// 布局窗口组件
	layoutWidget := widgets.NewQWidget(app, core.Qt__Widget)
	//layoutWidget.SetGeometry(core.NewQRect4(300, 300, 300, 220))
	layoutWidget.SetGeometry2(0, 0, 300, 220)

	// 数字
	lcd := widgets.NewQLCDNumber(layoutWidget)
	// 滑块
	sld := widgets.NewQSlider2(core.Qt__Horizontal, layoutWidget)

	// 信号连接槽
	sld.ConnectValueChanged(func(value int) {
		lcd.Display2(value)
	})

	vbox := widgets.NewQVBoxLayout2(layoutWidget)
	vbox.AddWidget(lcd,0, 0)
	vbox.AddWidget(sld,0,0)

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

