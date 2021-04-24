// WARNING! All changes made in this file will be lost!
package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/gui"
)

type UIFormWindow struct {
	HorizontalLayout4 *widgets.QHBoxLayout
	HorizontalLayout *widgets.QHBoxLayout
	HorizontalSpacer2 *widgets.QSpacerItem
	Widget2 *widgets.QWidget
	VerticalLayout *widgets.QVBoxLayout
	VerticalSpacer *widgets.QSpacerItem
	Widget3 *widgets.QWidget
	HorizontalLayout3 *widgets.QHBoxLayout
	HorizontalSpacer5 *widgets.QSpacerItem
	Label12 *widgets.QLabel
	HorizontalSpacer6 *widgets.QSpacerItem
	Widget5 *widgets.QWidget
	FormLayout *widgets.QFormLayout
	Label *widgets.QLabel
	LineEdit *widgets.QLineEdit
	Label2 *widgets.QLabel
	LineEdit2 *widgets.QLineEdit
	Label3 *widgets.QLabel
	LineEdit3 *widgets.QLineEdit
	Label4 *widgets.QLabel
	LineEdit4 *widgets.QLineEdit
	Label5 *widgets.QLabel
	LineEdit6 *widgets.QLineEdit
	Label6 *widgets.QLabel
	LineEdit5 *widgets.QLineEdit
	Label7 *widgets.QLabel
	LineEdit7 *widgets.QLineEdit
	Label8 *widgets.QLabel
	LineEdit8 *widgets.QLineEdit
	Label9 *widgets.QLabel
	LineEdit9 *widgets.QLineEdit
	Label10 *widgets.QLabel
	LineEdit10 *widgets.QLineEdit
	Label11 *widgets.QLabel
	LineEdit11 *widgets.QLineEdit
	Widget *widgets.QWidget
	HorizontalLayout2 *widgets.QHBoxLayout
	HorizontalSpacer3 *widgets.QSpacerItem
	PushButton2 *widgets.QPushButton
	PushButton *widgets.QPushButton
	HorizontalSpacer4 *widgets.QSpacerItem
	VerticalSpacer2 *widgets.QSpacerItem
	HorizontalSpacer *widgets.QSpacerItem
}

func (this *UIFormWindow) SetupUI(FormWindow *widgets.QWidget) {
	FormWindow.SetObjectName("FormWindow")
	FormWindow.SetGeometry(core.NewQRect4(0, 0, 800, 584))
	this.HorizontalLayout4 = widgets.NewQHBoxLayout2(FormWindow)
	this.HorizontalLayout4.SetObjectName("horizontalLayout_4")
	this.HorizontalLayout4.SetContentsMargins(0, 0, 0, 0)
	this.HorizontalLayout4.SetSpacing(0)
	this.HorizontalLayout = widgets.NewQHBoxLayout2(FormWindow)
	this.HorizontalLayout.SetObjectName("horizontalLayout")
	this.HorizontalLayout.SetContentsMargins(0, 0, 0, 0)
	this.HorizontalLayout.SetSpacing(0)
	this.HorizontalSpacer2 = widgets.NewQSpacerItem(40, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	this.HorizontalLayout.AddItem(this.HorizontalSpacer2)
	this.Widget2 = widgets.NewQWidget(FormWindow, core.Qt__Widget)
	this.Widget2.SetObjectName("Widget2")
	var sizePolicy *widgets.QSizePolicy
	sizePolicy = widgets.NewQSizePolicy2(widgets.QSizePolicy__Preferred, widgets.QSizePolicy__Preferred, widgets.QSizePolicy__DefaultType)
	sizePolicy.SetHorizontalStretch(1)
	sizePolicy.SetVerticalStretch(0)
	sizePolicy.SetHeightForWidth(this.Widget2.SizePolicy().HasHeightForWidth())
	this.Widget2.SetSizePolicy(sizePolicy)
	this.VerticalLayout = widgets.NewQVBoxLayout2(this.Widget2)
	this.VerticalLayout.SetObjectName("verticalLayout")
	this.VerticalLayout.SetContentsMargins(0, 0, 0, 0)
	this.VerticalLayout.SetSpacing(0)
	this.VerticalSpacer = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	this.VerticalLayout.AddItem(this.VerticalSpacer)
	this.Widget3 = widgets.NewQWidget(this.Widget2, core.Qt__Widget)
	this.Widget3.SetObjectName("Widget3")
	this.HorizontalLayout3 = widgets.NewQHBoxLayout2(this.Widget3)
	this.HorizontalLayout3.SetObjectName("horizontalLayout_3")
	this.HorizontalLayout3.SetContentsMargins(0, 0, 0, 0)
	this.HorizontalLayout3.SetSpacing(0)
	this.HorizontalSpacer5 = widgets.NewQSpacerItem(40, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	this.HorizontalLayout3.AddItem(this.HorizontalSpacer5)
	this.Label12 = widgets.NewQLabel(this.Widget3, core.Qt__Widget)
	this.Label12.SetObjectName("Label12")
	var font *gui.QFont
	font = gui.NewQFont()
	font.SetPointSize(30)
	font.SetWeight(75)
	this.Label12.SetFont(font)
	this.HorizontalLayout3.AddWidget(this.Label12, 0, 0)
	this.HorizontalSpacer6 = widgets.NewQSpacerItem(40, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	this.HorizontalLayout3.AddItem(this.HorizontalSpacer6)
	this.VerticalLayout.AddWidget(this.Widget3, 0, 0)
	this.Widget5 = widgets.NewQWidget(this.Widget2, core.Qt__Widget)
	this.Widget5.SetObjectName("Widget5")
	this.FormLayout = widgets.NewQFormLayout(this.Widget5)
	this.FormLayout.SetObjectName("formLayout")
	this.FormLayout.SetContentsMargins(9, 0, 0, 11)
	this.FormLayout.SetSpacing(0)
	this.FormLayout.SetHorizontalSpacing(11)
	this.FormLayout.SetVerticalSpacing(6)
	this.Label = widgets.NewQLabel(this.Widget5, core.Qt__Widget)
	this.Label.SetObjectName("Label")
	this.FormLayout.SetWidget(0, widgets.QFormLayout__LabelRole, this.Label)
	this.LineEdit = widgets.NewQLineEdit(this.Widget5)
	this.LineEdit.SetObjectName("LineEdit")
	this.FormLayout.SetWidget(0, widgets.QFormLayout__FieldRole, this.LineEdit)
	this.Label2 = widgets.NewQLabel(this.Widget5, core.Qt__Widget)
	this.Label2.SetObjectName("Label2")
	this.FormLayout.SetWidget(1, widgets.QFormLayout__LabelRole, this.Label2)
	this.LineEdit2 = widgets.NewQLineEdit(this.Widget5)
	this.LineEdit2.SetObjectName("LineEdit2")
	this.FormLayout.SetWidget(1, widgets.QFormLayout__FieldRole, this.LineEdit2)
	this.Label3 = widgets.NewQLabel(this.Widget5, core.Qt__Widget)
	this.Label3.SetObjectName("Label3")
	this.FormLayout.SetWidget(2, widgets.QFormLayout__LabelRole, this.Label3)
	this.LineEdit3 = widgets.NewQLineEdit(this.Widget5)
	this.LineEdit3.SetObjectName("LineEdit3")
	this.FormLayout.SetWidget(2, widgets.QFormLayout__FieldRole, this.LineEdit3)
	this.Label4 = widgets.NewQLabel(this.Widget5, core.Qt__Widget)
	this.Label4.SetObjectName("Label4")
	this.FormLayout.SetWidget(3, widgets.QFormLayout__LabelRole, this.Label4)
	this.LineEdit4 = widgets.NewQLineEdit(this.Widget5)
	this.LineEdit4.SetObjectName("LineEdit4")
	this.FormLayout.SetWidget(3, widgets.QFormLayout__FieldRole, this.LineEdit4)
	this.Label5 = widgets.NewQLabel(this.Widget5, core.Qt__Widget)
	this.Label5.SetObjectName("Label5")
	this.FormLayout.SetWidget(4, widgets.QFormLayout__LabelRole, this.Label5)
	this.LineEdit6 = widgets.NewQLineEdit(this.Widget5)
	this.LineEdit6.SetObjectName("LineEdit6")
	this.FormLayout.SetWidget(4, widgets.QFormLayout__FieldRole, this.LineEdit6)
	this.Label6 = widgets.NewQLabel(this.Widget5, core.Qt__Widget)
	this.Label6.SetObjectName("Label6")
	this.FormLayout.SetWidget(5, widgets.QFormLayout__LabelRole, this.Label6)
	this.LineEdit5 = widgets.NewQLineEdit(this.Widget5)
	this.LineEdit5.SetObjectName("LineEdit5")
	this.FormLayout.SetWidget(5, widgets.QFormLayout__FieldRole, this.LineEdit5)
	this.Label7 = widgets.NewQLabel(this.Widget5, core.Qt__Widget)
	this.Label7.SetObjectName("Label7")
	this.FormLayout.SetWidget(6, widgets.QFormLayout__LabelRole, this.Label7)
	this.LineEdit7 = widgets.NewQLineEdit(this.Widget5)
	this.LineEdit7.SetObjectName("LineEdit7")
	this.FormLayout.SetWidget(6, widgets.QFormLayout__FieldRole, this.LineEdit7)
	this.Label8 = widgets.NewQLabel(this.Widget5, core.Qt__Widget)
	this.Label8.SetObjectName("Label8")
	this.FormLayout.SetWidget(7, widgets.QFormLayout__LabelRole, this.Label8)
	this.LineEdit8 = widgets.NewQLineEdit(this.Widget5)
	this.LineEdit8.SetObjectName("LineEdit8")
	this.FormLayout.SetWidget(7, widgets.QFormLayout__FieldRole, this.LineEdit8)
	this.Label9 = widgets.NewQLabel(this.Widget5, core.Qt__Widget)
	this.Label9.SetObjectName("Label9")
	this.FormLayout.SetWidget(8, widgets.QFormLayout__LabelRole, this.Label9)
	this.LineEdit9 = widgets.NewQLineEdit(this.Widget5)
	this.LineEdit9.SetObjectName("LineEdit9")
	this.FormLayout.SetWidget(8, widgets.QFormLayout__FieldRole, this.LineEdit9)
	this.Label10 = widgets.NewQLabel(this.Widget5, core.Qt__Widget)
	this.Label10.SetObjectName("Label10")
	this.FormLayout.SetWidget(9, widgets.QFormLayout__LabelRole, this.Label10)
	this.LineEdit10 = widgets.NewQLineEdit(this.Widget5)
	this.LineEdit10.SetObjectName("LineEdit10")
	this.FormLayout.SetWidget(9, widgets.QFormLayout__FieldRole, this.LineEdit10)
	this.Label11 = widgets.NewQLabel(this.Widget5, core.Qt__Widget)
	this.Label11.SetObjectName("Label11")
	this.FormLayout.SetWidget(10, widgets.QFormLayout__LabelRole, this.Label11)
	this.LineEdit11 = widgets.NewQLineEdit(this.Widget5)
	this.LineEdit11.SetObjectName("LineEdit11")
	this.FormLayout.SetWidget(10, widgets.QFormLayout__FieldRole, this.LineEdit11)
	this.VerticalLayout.AddWidget(this.Widget5, 0, 0)
	this.Widget = widgets.NewQWidget(this.Widget2, core.Qt__Widget)
	this.Widget.SetObjectName("Widget")
	this.HorizontalLayout2 = widgets.NewQHBoxLayout2(this.Widget)
	this.HorizontalLayout2.SetObjectName("horizontalLayout_2")
	this.HorizontalLayout2.SetContentsMargins(0, 0, 0, 0)
	this.HorizontalLayout2.SetSpacing(12)
	this.HorizontalSpacer3 = widgets.NewQSpacerItem(40, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	this.HorizontalLayout2.AddItem(this.HorizontalSpacer3)
	this.PushButton2 = widgets.NewQPushButton(this.Widget)
	this.PushButton2.SetObjectName("PushButton2")
	this.HorizontalLayout2.AddWidget(this.PushButton2, 0, 0)
	this.PushButton = widgets.NewQPushButton(this.Widget)
	this.PushButton.SetObjectName("PushButton")
	this.HorizontalLayout2.AddWidget(this.PushButton, 0, 0)
	this.HorizontalSpacer4 = widgets.NewQSpacerItem(40, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	this.HorizontalLayout2.AddItem(this.HorizontalSpacer4)
	this.VerticalLayout.AddWidget(this.Widget, 0, 0)
	this.VerticalSpacer2 = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	this.VerticalLayout.AddItem(this.VerticalSpacer2)
	this.HorizontalLayout.AddWidget(this.Widget2, 0, 0)
	this.HorizontalSpacer = widgets.NewQSpacerItem(40, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	this.HorizontalLayout.AddItem(this.HorizontalSpacer)
	this.HorizontalLayout4.AddLayout(this.HorizontalLayout, 0)


    this.RetranslateUi(FormWindow)

}

func (this *UIFormWindow) RetranslateUi(FormWindow *widgets.QWidget) {
    _translate := core.QCoreApplication_Translate
	FormWindow.SetWindowTitle(_translate("FormWindow", "Widget", "", -1))
	this.Label12.SetText(_translate("FormWindow", "XXX表单", "", -1))
	this.Label.SetText(_translate("FormWindow", "TextLabel", "", -1))
	this.Label2.SetText(_translate("FormWindow", "TextLabel", "", -1))
	this.Label3.SetText(_translate("FormWindow", "TextLabel", "", -1))
	this.Label4.SetText(_translate("FormWindow", "TextLabel", "", -1))
	this.Label5.SetText(_translate("FormWindow", "TextLabel", "", -1))
	this.Label6.SetText(_translate("FormWindow", "TextLabel", "", -1))
	this.Label7.SetText(_translate("FormWindow", "TextLabel", "", -1))
	this.Label8.SetText(_translate("FormWindow", "TextLabel", "", -1))
	this.Label9.SetText(_translate("FormWindow", "TextLabel", "", -1))
	this.Label10.SetText(_translate("FormWindow", "TextLabel", "", -1))
	this.Label11.SetText(_translate("FormWindow", "TextLabel", "", -1))
	this.PushButton2.SetText(_translate("FormWindow", "PushButton", "", -1))
	this.PushButton.SetText(_translate("FormWindow", "PushButton", "", -1))
}
