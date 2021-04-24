package main

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
	"os"
	"unsafe"
)

/*
通过结构体实现类一般的体验。
*/

type UIMainWindow struct {
	CentralWidget *widgets.QWidget
	Menubar       *widgets.QMenuBar
	Statusbar     *widgets.QStatusBar
	pushButton    *widgets.QPushButton
	PushButton2   *widgets.QPushButton
	QLabel        *widgets.QLabel
}

func (w *UIMainWindow) SetupUI(MainWindow *widgets.QMainWindow) {
	MainWindow.SetObjectName("Window")
	// 设置窗口的标题
	MainWindow.SetWindowTitle("Qt 教程")
	// 设置窗口的位置和大小
	MainWindow.SetGeometry2(300, 300, 400, 220)
	// 设置窗口的图标
	MainWindow.SetWindowIcon(gui.NewQIcon5("images/app.ico"))

	// 状态栏
	//w.Statusbar = widgets.NewQStatusBar(MainWindow)
	//w.Statusbar.SetObjectName("Statusbar")
	//MainWindow.SetStatusBar(w.Statusbar)

	// 	添加按钮
	w.PushButton2 = widgets.NewQPushButton2("添加的按钮", MainWindow)
	w.PushButton2.Move2(30, 50)
	w.PushButton2.ConnectClicked(w.ButtonClicked2)

	// 默认按钮
	pushButtonObj := MainWindow.FindChild("pushButton", core.Qt__FindChildrenRecursively)
	if pushButtonName := pushButtonObj.ObjectName(); pushButtonName != "" {
		w.pushButton = (*widgets.QPushButton)(unsafe.Pointer(pushButtonObj))
		w.pushButton.ConnectClicked(w.ButtonClicked)
	}

	// 标签
	labelObj := MainWindow.FindChild("label", core.Qt__FindChildrenRecursively)
	if labelName := labelObj.ObjectName(); labelName != "" {
		w.QLabel = (*widgets.QLabel)(unsafe.Pointer(labelObj))
	}

}

func (w *UIMainWindow) ButtonClicked(checked bool) {
	fmt.Println("ButtonClicked")
	fmt.Println(w.QLabel)

	if w.QLabel != nil {
		w.QLabel.SetText("默认按钮信息")
	}
}

func (w *UIMainWindow) ButtonClicked2(checked bool) {
	fmt.Println("ButtonClicked2")
	if w.QLabel != nil {
		w.QLabel.SetText("添加按钮信息")
	}
}

func (w *UIMainWindow) RetranslateUi(MainWindow *widgets.QMainWindow) {
	_translate := core.QCoreApplication_Translate
	MainWindow.SetWindowTitle(_translate("Window", "QT 教程", "", -1))
	w.PushButton2.SetText(_translate("Window", "添加按钮", "", -1))
}

func main() {
	// 创建一个应用程序对象
	// sys.argv参数是一个列表，从命令行输入参数
	widgets.NewQApplication(len(os.Args), os.Args)
	//core.NewQCoreApplication(len(os.Args), os.Args).SetAttribute(core.Qt__AA_ShareOpenGLContexts, false)

	uiFile := core.NewQFile2("11_加载UI文件/untitled.ui")
	uiFile.Open(core.QIODevice__ReadOnly)
	loader := uitools.NewQUiLoader(nil)
	MainWindowObj := loader.Load(uiFile, nil)
	uiFile.Close()

	MainWindow := (*widgets.QMainWindow)(unsafe.Pointer(MainWindowObj))
	//fmt.Println(reflect.TypeOf(MainWindow), MainWindow)
	//fmt.Println(reflect.ValueOf(MainWindow))

	uiMain := UIMainWindow{}
	uiMain.SetupUI(MainWindow)
	uiMain.RetranslateUi(MainWindow)

	MainWindowObj.Show()

	// 确保应用程序干净的退出
	widgets.QApplication_Exec()
}
