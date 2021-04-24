// WARNING! All changes made in this file will be lost!
package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type UIMainwindowMainWindow struct {
	Centralwidget *widgets.QWidget
	VerticalLayout *widgets.QVBoxLayout
	Table *widgets.QTableWidget
	Widget *widgets.QWidget
	GridLayout *widgets.QGridLayout
	Widget2 *widgets.QWidget
	HorizontalLayout *widgets.QHBoxLayout
	PushButton *widgets.QPushButton
	PushButton2 *widgets.QPushButton
	LineEdit *widgets.QLineEdit
	Label *widgets.QLabel
	Label2 *widgets.QLabel
	Label3 *widgets.QLabel
	DateEdit *widgets.QDateEdit
	Label4 *widgets.QLabel
	Label5 *widgets.QLabel
	ComboBox *widgets.QComboBox
	LineEdit5 *widgets.QLineEdit
	LineEdit3 *widgets.QLineEdit
}

func (this *UIMainwindowMainWindow) SetupUI(MainWindow *widgets.QMainWindow) {
	MainWindow.SetObjectName("MainWindow")
	MainWindow.SetGeometry(core.NewQRect4(0, 0, 729, 492))
	this.Centralwidget = widgets.NewQWidget(MainWindow, core.Qt__Widget)
	this.Centralwidget.SetObjectName("Centralwidget")
	this.VerticalLayout = widgets.NewQVBoxLayout2(this.Centralwidget)
	this.VerticalLayout.SetObjectName("verticalLayout")
	this.VerticalLayout.SetContentsMargins(0, 0, 0, 0)
	this.VerticalLayout.SetSpacing(0)
	this.Table = widgets.NewQTableWidget(this.Centralwidget)
	this.Table.SetObjectName("Table")
	this.Table.SetFrameShape(widgets.QFrame__NoFrame)
	this.Table.SetLineWidth(2)
	this.Table.SetEditTriggers(widgets.QAbstractItemView__NoEditTriggers)
	this.Table.SetAlternatingRowColors(true)
	this.Table.SetSelectionMode(widgets.QAbstractItemView__SingleSelection)
	this.Table.SetSelectionBehavior(widgets.QAbstractItemView__SelectRows)
	this.Table.SetShowGrid(false)
	this.Table.SetSortingEnabled(true)
	this.Table.HorizontalHeader().SetStretchLastSection(true)
	this.Table.VerticalHeader().SetVisible(false)
	this.Table.VerticalHeader().SetDefaultSectionSize(20)
	this.VerticalLayout.AddWidget(this.Table, 0, 0)
	this.Widget = widgets.NewQWidget(this.Centralwidget, core.Qt__Widget)
	this.Widget.SetObjectName("Widget")
	this.GridLayout = widgets.NewQGridLayout(this.Widget)
	this.GridLayout.SetObjectName("gridLayout")
	this.GridLayout.SetContentsMargins(12, 6, 12, 12)
	this.GridLayout.SetSpacing(6)
	this.Widget2 = widgets.NewQWidget(this.Widget, core.Qt__Widget)
	this.Widget2.SetObjectName("Widget2")
	this.HorizontalLayout = widgets.NewQHBoxLayout2(this.Widget2)
	this.HorizontalLayout.SetObjectName("horizontalLayout")
	this.HorizontalLayout.SetContentsMargins(0, 0, 0, 0)
	this.HorizontalLayout.SetSpacing(0)
	this.PushButton = widgets.NewQPushButton(this.Widget2)
	this.PushButton.SetObjectName("PushButton")
	this.HorizontalLayout.AddWidget(this.PushButton, 0, 0)
	this.PushButton2 = widgets.NewQPushButton(this.Widget2)
	this.PushButton2.SetObjectName("PushButton2")
	this.HorizontalLayout.AddWidget(this.PushButton2, 0, 0)
	this.GridLayout.AddWidget3(this.Widget2, 1, 5, 1, 1, 0)
	this.LineEdit = widgets.NewQLineEdit(this.Widget)
	this.LineEdit.SetObjectName("LineEdit")
	this.LineEdit.SetClearButtonEnabled(true)
	this.GridLayout.AddWidget3(this.LineEdit, 0, 1, 1, 1, 0)
	this.Label = widgets.NewQLabel(this.Widget, core.Qt__Widget)
	this.Label.SetObjectName("Label")
	this.GridLayout.AddWidget3(this.Label, 0, 0, 1, 1, 0)
	this.Label2 = widgets.NewQLabel(this.Widget, core.Qt__Widget)
	this.Label2.SetObjectName("Label2")
	this.GridLayout.AddWidget3(this.Label2, 0, 2, 1, 1, 0)
	this.Label3 = widgets.NewQLabel(this.Widget, core.Qt__Widget)
	this.Label3.SetObjectName("Label3")
	this.GridLayout.AddWidget3(this.Label3, 0, 4, 1, 1, 0)
	this.DateEdit = widgets.NewQDateEdit(this.Widget)
	this.DateEdit.SetObjectName("DateEdit")
	this.DateEdit.SetDate(core.NewQDate3(2000, 1, 1))
	this.GridLayout.AddWidget3(this.DateEdit, 1, 1, 1, 1, 0)
	this.Label4 = widgets.NewQLabel(this.Widget, core.Qt__Widget)
	this.Label4.SetObjectName("Label4")
	this.GridLayout.AddWidget3(this.Label4, 1, 0, 1, 1, 0)
	this.Label5 = widgets.NewQLabel(this.Widget, core.Qt__Widget)
	this.Label5.SetObjectName("Label5")
	this.GridLayout.AddWidget3(this.Label5, 1, 2, 1, 1, 0)
	this.ComboBox = widgets.NewQComboBox(this.Widget)
	this.ComboBox.SetObjectName("ComboBox")
	this.ComboBox.AddItem("", core.NewQVariant())
	this.ComboBox.AddItem("", core.NewQVariant())
	this.GridLayout.AddWidget3(this.ComboBox, 0, 3, 1, 1, 0)
	this.LineEdit5 = widgets.NewQLineEdit(this.Widget)
	this.LineEdit5.SetObjectName("LineEdit5")
	this.LineEdit5.SetClearButtonEnabled(true)
	this.GridLayout.AddWidget3(this.LineEdit5, 1, 3, 1, 1, 0)
	this.LineEdit3 = widgets.NewQLineEdit(this.Widget)
	this.LineEdit3.SetObjectName("LineEdit3")
	this.LineEdit3.SetMaxLength(2)
	this.LineEdit3.SetClearButtonEnabled(true)
	this.GridLayout.AddWidget3(this.LineEdit3, 0, 5, 1, 1, 0)
	this.VerticalLayout.AddWidget(this.Widget, 0, 0)
	MainWindow.SetCentralWidget(this.Centralwidget)
	this.Label.SetBuddy(this.LineEdit)
	this.Label2.SetBuddy(this.ComboBox)
	this.Label3.SetBuddy(this.LineEdit3)
	this.Label4.SetBuddy(this.DateEdit)
	this.Label5.SetBuddy(this.LineEdit5)

    this.RetranslateUi(MainWindow)
	this.ComboBox.SetCurrentIndex(0)
}

func (this *UIMainwindowMainWindow) RetranslateUi(MainWindow *widgets.QMainWindow) {
    _translate := core.QCoreApplication_Translate
	MainWindow.SetWindowTitle(_translate("MainWindow", "MainWindow", "", -1))
	this.Table.SetColumnCount(0)
	this.Table.SetRowCount(0)
	this.PushButton.SetText(_translate("MainWindow", "添加", "", -1))
	this.PushButton2.SetText(_translate("MainWindow", "删除", "", -1))
	this.LineEdit.SetPlaceholderText(_translate("MainWindow", "请输入姓名", "", -1))
	this.Label.SetText(_translate("MainWindow", "姓名", "", -1))
	this.Label2.SetText(_translate("MainWindow", "性别", "", -1))
	this.Label3.SetText(_translate("MainWindow", "年龄", "", -1))
	this.DateEdit.SetDisplayFormat(_translate("MainWindow", "MM月dd", "", -1))
	this.Label4.SetText(_translate("MainWindow", "生日", "", -1))
	this.Label5.SetText(_translate("MainWindow", "啪啪", "", -1))
	this.ComboBox.SetCurrentText(_translate("MainWindow", "女", "", -1))
	this.ComboBox.SetItemText(0, _translate("MainWindow", "女", "", -1))
	this.ComboBox.SetItemText(1, _translate("MainWindow", "男", "", -1))
	this.LineEdit5.SetPlaceholderText(_translate("MainWindow", "啪啪次数", "", -1))
	this.LineEdit3.SetText(_translate("MainWindow", "18", "", -1))
	this.LineEdit3.SetPlaceholderText(_translate("MainWindow", "请输入年龄", "", -1))
}
