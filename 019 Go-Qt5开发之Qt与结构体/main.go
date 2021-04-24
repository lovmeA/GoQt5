package main

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"os"
)

/*
通过结构体实现类一般的体验。
*/

type UIMainWindow struct {
	CentralWidget *widgets.QWidget
	Menubar       *widgets.QMenuBar
	Statusbar     *widgets.QStatusBar
	PushButton1   *widgets.QPushButton
	PushButton2   *widgets.QPushButton
}

func (w *UIMainWindow) SetupUI(MainWindow *widgets.QMainWindow) {
	MainWindow.SetObjectName("MainWindow")
	// 设置窗口的标题
	MainWindow.SetWindowTitle("Qt 教程")
	// 设置窗口的位置和大小
	MainWindow.SetGeometry2(300, 300, 300, 220)
	// 设置窗口的图标
	MainWindow.SetWindowIcon(gui.NewQIcon5("images/app.ico"))

	// 中心窗口组件载体
	w.CentralWidget = widgets.NewQWidget(MainWindow, core.Qt__Widget)
	// w.CentralWidget.SetGeometry(core.NewQRect4(300, 300, 300, 220))
	w.CentralWidget.SetGeometry2(0, 0, 300, 220)
	MainWindow.SetCentralWidget(w.CentralWidget)

	// 状态栏
	w.Statusbar = widgets.NewQStatusBar(MainWindow)
	w.Statusbar.SetObjectName("Statusbar")
	MainWindow.SetStatusBar(w.Statusbar)

	// 按钮
	w.PushButton1 = widgets.NewQPushButton2("PushButton1", w.CentralWidget)
	w.PushButton1.Move2(30, 50)
	w.PushButton1.ConnectClicked(w.ButtonClicked)

}

func (w *UIMainWindow) ButtonClicked(checked bool) {
	fmt.Println("ButtonClicked", checked)
	w.Statusbar.ShowMessage("sender"+" was pressed", 0)
}

func (w *UIMainWindow) RetranslateUi(MainWindow *widgets.QMainWindow) {
	_translate := core.QCoreApplication_Translate
	MainWindow.SetWindowTitle(_translate("MainWindow", "QT 教程", "", -1))
	w.PushButton1.SetText(_translate("MainWindow", "PushButton", "", -1))
}

func main() {
	// 创建一个应用程序对象
	// sys.argv参数是一个列表，从命令行输入参数
	widgets.NewQApplication(len(os.Args), os.Args)

	// 初始化窗口
	MainWindow := widgets.NewQMainWindow(nil, 0)

	uiMain := UIMainWindow{}
	uiMain.SetupUI(MainWindow)
	uiMain.RetranslateUi(MainWindow)

	// 显示组件
	MainWindow.Show()

	// 确保应用程序干净的退出
	widgets.QApplication_Exec()
}
