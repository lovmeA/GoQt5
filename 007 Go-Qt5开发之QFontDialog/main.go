/*-------------------------------------------------------------------
| Project   : goland
| Author    : 今夕何夕
| QQ/Email  : 184566608<qingyueheji@qq.com>
| Time      : 2020-10-07 18:14
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
QFontDialog 对话框用以选择字体
在这个例子中，我们创建了一个按钮和一个标签，通过QFontDialog来改变标签的字体
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

	vbox := widgets.NewQVBoxLayout2(widget)

	lbl := widgets.NewQLabel2("Knowledge only matters", widget, 0)
	lbl.Move2(130, 20)
	vbox.AddWidget(lbl, 0, 0)

	btn := widgets.NewQPushButton2("Dialog", widget)
	//btn.SetSizePolicy(widgets.QSizePolicy__Fixed)
	btn.Move2(20, 20)
	vbox.AddWidget(btn, 0, 0)
	btn.ConnectClicked(func(checked bool) {
		// 这一行代码弹出字体选择对话框，GetFont2()方法返回字体名称和ok参数，
		//如果用户点击了ok他就是True,否则就是false
		flag := false
		font := widgets.NewQFontDialog(widget).GetFont2(&flag, widget)
		// 如果我们点击了ok，标签的字体就会被改变
		if flag {
			lbl.SetFont(font)
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
