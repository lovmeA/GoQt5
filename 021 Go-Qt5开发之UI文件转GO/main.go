/*-------------------------------------------------------------------
| Project   : goland
| Author    : 今夕何夕
| QQ/Email  : 184566608<qingyueheji@qq.com>
| Time      : 2020-10-07 21:15
| Describe  : main
+------------------------------------------------------------------*/

package main

import (
	"encoding/json"
	"fmt"
	"github.com/therecipe/qt/widgets"
	"os"
	"reflect"
	"strings"
	"utils"
)

type UsersInfo struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Age      string `json:"age"`
	Birthday string `json:"birthday"`
	Pp       string `json:"pp"`
}

func (this *UIMainwindowMainWindow) Msg(text string) {
	widgets.QMessageBox_Question(nil, "提示信息", text, widgets.QMessageBox__Yes, widgets.QMessageBox__Yes)
}

func (this *UIMainwindowMainWindow) data() []UsersInfo {
	infos := []UsersInfo{}
	s, m := utils.ReadFileString("data.db")
	if s {
		_, _ = utils.WriteFIleString("data.db", "[{\"id\":\"1\",\"name\":\"刘亦菲\",\"gender\":\"女\",\"age\":\"20\",\"birthday\":\"5月20\",\"pp\":\"520次\"},{\"id\":\"2\",\"name\":\"迪丽热巴\",\"gender\":\"女\",\"age\":\"18\",\"birthday\":\"5月18\",\"pp\":\"518次\"},{\"id\":\"3\",\"name\":\"杨幂\",\"gender\":\"女\",\"age\":\"23\",\"birthday\":\"8月20\",\"pp\":\"820次\"},{\"id\":\"4\",\"name\":\"赵丽颖\",\"gender\":\"女\",\"age\":\"22\",\"birthday\":\"7月22\",\"pp\":\"722次\"},{\"id\":\"5\",\"name\":\"袁冰妍\",\"gender\":\"女\",\"age\":\"23\",\"birthday\":\"9月28\",\"pp\":\"928次\"}]")
		return nil
	} else {
		err := json.Unmarshal([]byte(m), &infos)
		if err != nil {
			this.Msg(m)
			return nil
		}
		return infos
	}
}

func (this *UIMainwindowMainWindow) Header() []string {
	s, m := utils.ReadFileString("header.db")
	if s {
		return nil
	} else {
		headers := strings.Split(m, "|")
		headers = append(headers, "")
		return headers
	}
}

func (this *UIMainwindowMainWindow) initTableHeader() {
	headers := this.Header()
	if headers == nil {
		_, _ = utils.WriteFIleString("header.db", "编号|姓名|性别|年龄|生日|爱爱")
		this.Msg("表头加载失败")
	}
	this.Table.SetColumnCount(len(headers))
	for index, header := range headers {
		this.Table.SetHorizontalHeaderItem(index, widgets.NewQTableWidgetItem2(header, index))
	}

	infos := this.data()
	this.Table.ClearContents()
	this.Table.SetRowCount(len(infos))
	var sortingEnabled bool
	sortingEnabled = this.Table.IsSortingEnabled()

	for index, obj := range infos {
		field := reflect.TypeOf(obj)
		value := reflect.ValueOf(obj)
		for i := 0; i < field.NumField(); i++ {
			this.Table.SetItem(index, i, widgets.NewQTableWidgetItem2(value.Field(i).String(), i))
		}
	}

	this.Table.SetSortingEnabled(sortingEnabled)
}

func (this *UIMainwindowMainWindow) onClickAddItem(bool2 bool) {
	fmt.Println("点击")
}

func (this *UIMainwindowMainWindow) init() {
	this.PushButton.ConnectClicked(this.onClickAddItem)
	this.Table.ConnectItemSelectionChanged(func() {
		row := this.Table.SelectedItems() 				// 或许我用的不对，到这儿就直接报错，pyqt5可以这样子用
		fmt.Println(row)
	})
}

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	// 初始化窗口
	mainWindow := widgets.NewQMainWindow(nil, 0)
	uiMain := UIMainwindowMainWindow{}
	uiMain.SetupUI(mainWindow)
	uiMain.initTableHeader()
	uiMain.init()

	// mainWindow.SetStyleSheet(utils.Qss())

	qr := mainWindow.FrameGeometry()
	cp := widgets.NewQDesktopWidget().AvailableGeometry2(mainWindow).Center()
	qr.MoveCenter(cp)

	mainWindow.Move(qr.TopLeft())

	mainWindow.Show()

	widgets.QApplication_Exec()
}
