/*-------------------------------------------------------------------
| Project   : goland
| Author    : 今夕何夕
| QQ/Email  : 184566608<qingyueheji@qq.com>
| Time      : 2020-10-07 03:59
| Describe  : 水平垂直布局
+------------------------------------------------------------------*/

package main

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"io"
	"os"
)

func InitUI() *widgets.QMainWindow {
	// NewQMainWindow 没有布局功能，这只是一个窗口
	app := widgets.NewQMainWindow(nil, 0)
	app.SetGeometry2(0, 0, 500, 400)

	// 注册一个控件挂 NewQMainWindow(app) 上 管理全局
	centralWidget := widgets.NewQWidget(app, core.Qt__Widget)

	// 注册一个布局控件并设置 centralWidget 的布局(水平布局)
	verticalLayout := widgets.NewQVBoxLayout2(centralWidget)
	// 控件的间距
	verticalLayout.SetSpacing(0)
	verticalLayout.SetContentsMargins(0, 0, 0, 0)

	// 注册一个 QWidget 控件 并挂载到 centralWidget
	topLayoutWidget := widgets.NewQWidget(centralWidget, core.Qt__Widget)
	//topLayoutWidget.SetGeometry2(0, 0, 300, 220)	// 设置控件位置及大小
	/////////////////////////////////////注释了的等于没注释的///////////////////////////////////////////////////////////
	topLayoutWidget.SetSizePolicy2(widgets.QSizePolicy__Preferred, widgets.QSizePolicy__Expanding)
	//// 设置控件填充
	//sizePolicy := widgets.NewQSizePolicy2(widgets.QSizePolicy__Preferred, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__DefaultType)
	//// 设置水平
	//sizePolicy.SetHorizontalPolicy(0)
	//// 设置垂直
	//sizePolicy.SetVerticalStretch(0)
	//// 设置高
	//sizePolicy.SetHeightForWidth(topLayoutWidget.SizePolicy().HasHeightForWidth())
	//// 给控件设置填充
	//topLayoutWidget.SetSizePolicy(sizePolicy)
	////////////////////////////////////////////////////////////////////////////////////////////////
	// 设置 verticalLayout 的子控件 并受父控件布局影响
	verticalLayout.AddWidget(topLayoutWidget, 0, 0)

	// 注册一个 QWidget 控件 并挂载到 centralWidget
	bottomLayoutwidget := widgets.NewQWidget(centralWidget, core.Qt__Widget)
	// 设置 verticalLayout 的子控件 并受父控件布局影响
	verticalLayout.AddWidget(bottomLayoutwidget, 0, 0)

	// 注册一个布局控件并设置 topLayoutWidget 的布局(垂直布局)
	hbox := widgets.NewQHBoxLayout2(bottomLayoutwidget)

	// 注册一个伸缩因子
	spacerItem := widgets.NewQSpacerItem(40, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)

	// 在 ok 之前添加一个伸缩因子
	hbox.AddItem(spacerItem)

	// 按钮 (ok)
	ok := widgets.NewQPushButton2("OK", bottomLayoutwidget)
	// 添加子控件
	hbox.AddWidget(ok, 0, 0)

	// 按钮 (cancel)
	cancel := widgets.NewQPushButton2("Cancel", bottomLayoutwidget)
	hbox.AddWidget(cancel, 0, 0)
	// 在 cancel 之后添加一个伸缩因子
	hbox.AddItem(spacerItem)

	// 设置 centralWidget 控件铺满窗口
	app.SetCentralWidget(centralWidget)
	return app
}

func qss() string {
	path, _ := os.Getwd()
	path = path + "/src/099 测试/theme/theme.css"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
	}
	defer file.Close()
	var content []byte
	var buf [1024]byte

	for {
		n, err := file.Read(buf[:])
		if err == io.EOF || err != nil {
			break
		}
		if err != nil {
			fmt.Println("文件读取失败", err.Error())
		}
		content = append(content, buf[:n]...)
	}
	return string(content)
}

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	app := InitUI()
	app.SetWindowIcon(gui.NewQIcon5("images/app.ico"))

	app.SetWindowTitle("水平垂直布局")

	app.SetStyleSheet(qss())

	// 获取窗口 取得屏幕中心坐标
	qr := app.FrameGeometry()
	cp := widgets.NewQDesktopWidget().AvailableGeometry2(app).Center()
	qr.MoveCenter(cp)

	// 设置窗口移动位置
	app.Move(qr.TopLeft())

	app.Show()

	widgets.QApplication_Exec()
}
