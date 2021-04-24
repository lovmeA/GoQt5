/*-------------------------------------------------------------------
| Project   : goland
| Author    : 今夕何夕
| QQ/Email  : 184566608<qingyueheji@qq.com>
| Time      : 2020-10-07 18:28
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
QSlider是一个带有简单滑块的控件。滑块可以前后拖动。我们可以通过拖动选择一个特定的值。
有时使用滑动条比直接输入数字或使用旋转框更加自然。

在下面的例子中，我们会显示一个滑动条与一个标签，标签用于显示图片，并通过滑动条控件图片的显示 。
注意这里图片只能用ico格式的，png格式图片会显示不了。

例子中我们模拟了一个音量控制。通过拖动滑块来改变标签上的图像。
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

	//创建了一个QLabel控件并为它设置了一个初始音量图像。
	label := widgets.NewQLabel(widget, 0)
	//label.SetPixmap(gui.QPixmap_FromImage2(gui.NewQImage9("images/audio.png", ""), 0))

	// 图片缩放无效
	//pix := gui.NewQPixmap3("images/audio.png", "",core.Qt__AutoColor)
	//pix.Scaled(label.Size(),core.Qt__IgnoreAspectRatio, core.Qt__FastTransformation)
	//pix.Scaled2(10,10,core.Qt__IgnoreAspectRatio, core.Qt__FastTransformation)
	//label.SetPixmap(pix)

	label.SetScaledContents(true)
	label.SetPixmap(gui.NewQPixmap3("images/audio.png", "", core.Qt__AutoColor))

	label.SetGeometry2(160, 40, 50, 50)

	//创建一个水平滑块
	sld := widgets.NewQSlider2(core.Qt__Horizontal, widget)
	sld.SetFocusPolicy(core.Qt__NoFocus)
	sld.SetGeometry2(30, 40, 100, 30)
	//我们根据滑动条的值来设置标签的图像。
	//当滑动条的值为0时我们为标签设置audio.ico图像。
	sld.ConnectValueChanged(func(value int) {
		if value == 0 {
			label.SetText("audio")
		} else if value > 0 && value <= 30 {
			label.SetText("audio_min")
		} else if value > 30 && value < 80 {
			label.SetText("audio_med")
		} else {
			label.SetText("audio_max")
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
