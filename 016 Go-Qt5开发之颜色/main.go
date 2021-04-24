package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"os"
)

/*
颜色是一个对象代表红、绿、蓝(RGB)强度值。有效的RGB值的范围从0到255。
我们可以用不同的方法定义了一个颜色。最常见的是RGB十进制或十六进制值的值。
我们也可以使用一个RGBA值代表红色,绿色,蓝色,透明度。我们添加一些额外的信息透明度。
透明度值255定义了完全不透明,0是完全透明的,例如颜色是无形的。


实例中我们绘制了3个不同颜色的矩形
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

		col := gui.NewQColor3(0, 0, 0, 0)
		// 定义一个使用十六进制符号颜色。
		col.SetNamedColor("#d4d4d4")
		qp.SetPen2(col)

		// 设置画刷颜色
		// 我们为QPainter设置了一个笔刷(Bursh)对象并用它绘制了一个矩形。
		//笔刷是用于绘制形状背景的基本图形对象。
		qp.SetBrush2(core.Qt__Dense2Pattern)
		// DrawRect2()方法接受四个参数，前两个是起点的x,y坐标，后两个是矩形的宽和高。
		//这个方法使用当前的画笔与笔刷对象进行绘制。
		qp.DrawRect2(10, 15, 90, 60)

		qp.SetBrush2(core.Qt__Dense4Pattern)
		qp.DrawRect2(130, 15, 90, 60)

		qp.SetBrush2(core.Qt__BDiagPattern)
		qp.DrawRect2(250, 15, 90, 60)

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
