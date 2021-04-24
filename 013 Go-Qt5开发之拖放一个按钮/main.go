package main

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"os"
)

/*
在下面的示例中我们将演示如何对一个按钮控件进行拖放。

在窗口显示一个QPushButton 。
如果用鼠标左键点击这个按钮会在控制台中输出’press’消息。鼠标右击进行拖动。
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

	widget.SetAcceptDrops(true)

	button := widgets.NewQPushButton2("Button", widget)
	button.Move2(100, 65)
	button.ConnectMouseMoveEvent(func(event *gui.QMouseEvent) {
		if event.Buttons() != core.Qt__RightButton {
			return
		}

		mimeData := core.NewQMimeData()

		// QDrag提供了对基于MIME的拖放的数据传输的支持。
		drag := gui.NewQDrag(widget)
		drag.SetMimeData(mimeData)

		//drag.SetHotSpot()
		// 用于启动拖放操作。
		drag.Exec(core.Qt__MoveAction)
	})
	button.ConnectMousePressEvent(func(event *gui.QMouseEvent) {
		//鼠标左击按钮时我们会在控制台打印‘press’。
		//注意我们也调用了父按钮的mousePressEvent()方法。否则会看不到按钮的按下效果。
		//widgets.QPushButton.MousePressEvent(button, event)
		if event.Button() == core.Qt__LeftButton {
			fmt.Println("press")
		}
	})

	widget.ConnectDragEnterEvent(func(event *gui.QDragEnterEvent) {
		event.AcceptProposedAction()
	})
	widget.ConnectDropEvent(func(event *gui.QDropEvent) {
		// 释放右键后调用dropEvent()方法中，即找出鼠标指针的当前位置，并将按钮移动过去。
		position := event.Pos()
		button.Move(position)

		// 我们可以对指定的类型放弃行动。在我们的例子中它是一个移动动作。
		event.SetDropAction(core.Qt__MoveAction)
		event.Accept()
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
