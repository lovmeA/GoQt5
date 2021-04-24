package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"os"
)

/*
QBrush是一个基本的图形对象。它用于油漆的背景图形形状,如矩形、椭圆形或多边形。
三种不同类型的刷可以:一个预定义的刷,一个梯度,或纹理模式。

示例中绘制九个不同的矩形
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

		//我们定义了一个笔刷对象，然后将它设置给QPainter对象，
		//并调用painter的drawRect()方法绘制矩形。
		brush := gui.NewQBrush2(core.Qt__SolidPattern)
		qp.SetBrush(brush)
		qp.DrawRect2(10, 15, 90, 60)

		brush.SetStyle(core.Qt__Dense1Pattern)
		qp.SetBrush(brush)
		qp.DrawRect2(130, 15, 90, 60)

		brush.SetStyle(core.Qt__Dense2Pattern)
		qp.SetBrush(brush)
		qp.DrawRect2(250, 15, 90, 60)

		brush.SetStyle(core.Qt__DiagCrossPattern)
		qp.SetBrush(brush)
		qp.DrawRect2(10, 105, 90, 60)

		brush.SetStyle(core.Qt__Dense5Pattern)
		qp.SetBrush(brush)
		qp.DrawRect2(130, 105, 90, 60)

		brush.SetStyle(core.Qt__Dense6Pattern)
		qp.SetBrush(brush)
		qp.DrawRect2(250, 105, 90, 60)

		brush.SetStyle(core.Qt__HorPattern)
		qp.SetBrush(brush)
		qp.DrawRect2(10, 195, 90, 60)

		brush.SetStyle(core.Qt__VerPattern)
		qp.SetBrush(brush)
		qp.DrawRect2(130, 195, 90, 60)

		brush.SetStyle(core.Qt__BDiagPattern)
		qp.SetBrush(brush)
		qp.DrawRect2(250, 195, 90, 60)

		//我们创建一个QPen对象。颜色是黑色的。
		//宽度设置为2像素,这样我们可以看到笔风格之间的差异。Qt.SolidLine是预定义的钢笔样式。
		//pen := gui.NewQPen3(gui.NewQColor2(core.Qt__black))
		pen := gui.NewQPen4(gui.NewQBrush2(core.Qt__SolidPattern), 2, core.Qt__SolidLine, core.Qt__FlatCap, core.Qt__MiterJoin)
		qp.SetPen(pen)
		qp.DrawLine3(20, 40, 250, 40)

		pen.SetStyle(core.Qt__DashDotDotLine)
		qp.SetPen(pen)
		qp.DrawLine3(20, 80, 250, 80)

		pen.SetStyle(core.Qt__DashDotLine)
		qp.SetPen(pen)
		qp.DrawLine3(20, 120, 250, 120)

		pen.SetStyle(core.Qt__DotLine)
		qp.SetPen(pen)
		qp.DrawLine3(20, 160, 250, 160)

		pen.SetStyle(core.Qt__DashDotDotLine)
		qp.SetPen(pen)
		qp.DrawLine3(20, 200, 250, 200)

		//这里我们定义了一个画笔风格。
		//我们设置了Qt.CustomDashLine并调用了setDashPattern()方法，
		//它的参数(一个数字列表)定义了一种风格，必须有偶数个数字；
		//其中奇数表示绘制实线，偶数表示留空。数值越大，直线或空白就越大。
		//这里我们定义了1像素的实线，4像素的空白，5像素实线，4像素空白。。。
		pen.SetStyle(core.Qt__CustomDashLine)
		qp.SetPen(pen)
		pen.SetDashPattern([]float64{1, 4, 5, 4})
		qp.DrawLine3(20, 240, 250, 240)

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
