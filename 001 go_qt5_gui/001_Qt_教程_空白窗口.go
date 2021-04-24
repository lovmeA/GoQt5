/*-------------------------------------------------------------------
| Project   : goland
| Author    : 今夕何夕
| QQ/Email  : 184566608<qingyueheji@qq.com>
| Time      : 2020-10-07 02:17
| Describe  : main
+------------------------------------------------------------------*/

package main

import (
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"os"
)


// 关闭窗口
func closeWindow(app *widgets.QMainWindow) {
	// 创建按钮
	btn := widgets.NewQPushButton2("关闭窗口", app)
	// 设置按钮大小
	btn.Resize(btn.SizeHint())
	// 移动按钮位置
	btn.Move2(100, 100)
	btn.SetToolTip("This is a <b>QPushButton</b> widget")
	// 按钮点击事件
	btn.ConnectClicked(func(bool) {
		app.Close()
	})
}

// 控制窗口显示在屏幕中心的方法
func center(app *widgets.QMainWindow) {
	// 获得窗口
	qr := app.FrameGeometry()

	//  获得屏幕中心点
	cp := widgets.NewQDesktopWidget().AvailableGeometry2(app).Center()
	// 显示到屏幕中心
	qr.MoveCenter(cp)
	app.Move(qr.TopLeft())
}

// 初始化窗口
func InitUi() *widgets.QMainWindow {
	// 创建窗口
	app := widgets.NewQMainWindow(nil, 0)

	// 设置窗口的标题
	app.SetWindowTitle("Qt 教程")

	// 设置窗口的位置和大小
	app.SetGeometry2(300, 300, 300, 220)

	// 设置窗口的图标，引用当前目录下的web.png图片
	app.SetWindowIcon(gui.NewQIcon5("images/app.ico"))

	// 关闭窗口
	closeWindow(app)

	// 更改窗口关闭默认事件
	app.ConnectCloseEvent(func(event *gui.QCloseEvent) {
		reply := widgets.QMessageBox_Question(nil, "Message", "Are you sure to quit?", widgets.QMessageBox__Yes|widgets.QMessageBox__No, widgets.QMessageBox__Yes)
		if reply == widgets.QMessageBox__Yes {
			event.Accept()
		} else {
			event.Ignore()
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

	// 窗口居中
	center(app)

	// 显示组件
	app.Show()

	// 确保应用程序干净的退出
	widgets.QApplication_Exec()
}