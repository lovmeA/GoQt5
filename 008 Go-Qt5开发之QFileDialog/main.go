/*-------------------------------------------------------------------
| Project   : goland
| Author    : 今夕何夕
| QQ/Email  : 184566608<qingyueheji@qq.com>
| Time      : 2020-10-07 18:15
| Describe  : main
+------------------------------------------------------------------*/

package main

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

/*
QFileDialog是一个让用户选择文件和目录的对话框，可用以选择打开或保存文件

这个例子展示了一个菜单栏，中部TextEdit控件和一个状态栏。
菜单项Open会显示用于选择文件的QtGui.QFileDialog对话框。
选定文件的内容会加载到TextEdit控件中。
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

	// 文档框
	textEdit := widgets.NewQTextEdit(widget)
	app.SetCentralWidget(textEdit)

	// 子按钮
	openFile := widgets.NewQAction3(gui.NewQIcon5("images/app.ico"), "&Open", widget)
	// 快捷键，自定义
	openFile.SetShortcut(gui.NewQKeySequence2("Ctrl+O", gui.QKeySequence__NativeText))
	// 提示语
	openFile.SetStatusTip("Open new File")
	// 事件触发
	openFile.ConnectTriggered(func(checked bool) {
		filename := widgets.NewQFileDialog(widget, 0).GetOpenFileName(
			widget,
			"Open file",
			"/home",
			"",
			"",
			widgets.QFileDialog__ShowDirsOnly)

		if filename != "" {
			// 读取了选择的文件并将文件内容显示到了TextEdit控件。
			if contents, err := ioutil.ReadFile(filename); err == nil {
				//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
				result := strings.Replace(string(contents), "\n", "", 1)
				textEdit.SetText(result)
			}
		} else {
			fmt.Println(reflect.TypeOf(filename), filename)
		}
	})

	var actions []*widgets.QAction
	actions = append(actions, openFile)

	menubar := app.MenuBar()
	fileMenu := menubar.AddMenu2("&File")
	fileMenu.AddActions(actions)

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
