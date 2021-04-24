package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"os"
)

/*
我们先以窗体内Unicode文本的绘制为例。

在我们的示例中,我们绘制一些Cylliric文本。文本垂直和水平对齐。
相信大多 数人和我一样，刚开始的时候都认为 drawText() 的 x, y 是字符串左上角的坐标，其实不然，它是字符串的第一个字符的 origin 的坐标，y 是字体的 base line 的 y 坐标
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
	//widget.SetGeometry2(0, 0, 300, 220)
	app.SetCentralWidget(widget)

	// 状态栏
	app.StatusBar()

	widget.ConnectPaintEvent(func(event *gui.QPaintEvent) {
		qp := gui.NewQPainter2(widget)
		//qp.Begin(widget)

		qp.SetRenderHint(gui.QPainter__Antialiasing, true)
		// 定义一个画笔和一个字体用于绘制文本。
		qp.SetFont(gui.NewQFont2("Times", 100, 0, true))
		qp.SetPen2(gui.NewQColor3(0, 0, 0, 0))

		//qp.DrawText3(150, 150, "jEh")

		// 方法将文本绘制在窗体，显示在中心
		qp.DrawText4(event.Rect(), int(core.QTextStream__AlignCenter), "text", widget.Rect())

		qp.End()

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
