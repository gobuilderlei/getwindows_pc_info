package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gofyne_test/myexcel"
)

func AppTalbeMenuFile(w fyne.Window) *container.AppTabs {
	tables := container.NewAppTabs(
		container.NewTabItemWithIcon("file", theme.FileIcon(), widget.NewButton("", func() {

		})),
		container.NewTabItemWithIcon("save", theme.DocumentSaveIcon(), widget.NewButton("", func() {

		})),
	)
	tables.SetTabLocation(container.TabLocationLeading)
	return tables
}

func MenuFile(w fyne.Window) *fyne.Menu {
	lbmenu := fyne.NewMenu("文件[File]",
		fyne.NewMenuItem("打开[Open]", func() {
		}),
		fyne.NewMenuItem("保存[Save]", func() {
			//s := myexcel.ExcelInfo{UserName: "ssss", OfficeName: "基地啊", ComputerName: "zero-oo"}
			mo := NetInfoP
			err := myexcel.SaveInfomation(mo)
			if err != nil {
				fmt.Println("", err)
			}
		}),
	)

	return lbmenu
}

func MenuAbout(w fyne.Window) *fyne.Menu {
	lbmenu := fyne.NewMenu("关于[About]",
		fyne.NewMenuItem("关于[About]", func() {

		}),
		fyne.NewMenuItem("版本[Version]", func() {

		}),
	)
	return lbmenu
}

// 弹出窗口
func POPMenuList(w fyne.Window) *fyne.Container {
	popfile := widget.NewPopUpMenu(MenuFile(w), w.Canvas())
	btn := widget.NewButton("File", func() {
		popfile.Show()
	})
	popabout := widget.NewPopUpMenu(MenuFile(w), w.Canvas())
	btnabout := widget.NewButton("About", func() {
		popabout.ShowAtPosition(fyne.Position{60, 0})
	})
	return container.NewHBox(btn, btnabout)
}
