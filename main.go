package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"gofyne_test/gui"
	"os"
)

//	func init() {
//		//fontPaths, err := findfont.Find("")
//		//if err != nil {
//		//	fmt.Println("err::::======", err)
//		//}
//		//fontData, err1 := ioutil.ReadFile(fontPaths)
//		//if err1 != nil {
//		//	fmt.Println("err1::::======", err1)
//		//}
//		//font, err2 := truetype.Parse(fontData)
//		//if err2 != nil {
//		//	fmt.Println("err2::::======", err2)
//		//}
//		//fmt.Println(font)
//		fontpa := findfont.List()
//		for _, path := range fontpa {
//			fmt.Println(path)
//			if strings.Contains(path, "simkai.ttf") {
//				os.Setenv("GOFYNE_TEST", path)
//				defer os.Unsetenv("GOFYNE_TEST")
//				break
//			}
//		}
//	}
func main() {

	//for _, v := range info.GetPcInfo() {
	//	//fmt.Println(k)
	//	fmt.Println(v)
	//}
	//fmt.Println(info.GetUserName())
	//fonts := &myfont.MyTheme{}
	//fonts.SetFonts("./assets/msyh.ttc", "")

	os.Setenv("FYNE_FONT", "./assets/msyh.ttc")
	a := app.New()
	//app.Settings().SetTheme(fonts)
	w := a.NewWindow("系统信息获取记录器")
	w.Resize(fyne.NewSize(800, 600))
	//toolbar := widget.NewToolbar(widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
	//	//写入保存信息;
	//}),
	//	//widget.NewToolbarSeparator(),
	//	//widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
	//	//widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
	//	//widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
	//	//widget.NewToolbarSpacer(),
	//	widget.NewToolbarAction(theme.HelpIcon(), func() {
	//		//写入帮助
	//	}),
	//)
	//toolscontent := container.NewBorder(toolbar, nil, nil, nil, widget.NewLabel("123"))
	//wordlabel.MinSize()
	//主菜单注册
	w.SetMainMenu(fyne.NewMainMenu(gui.MenuFile(w), gui.MenuAbout(w)))

	//w.SetContent(container.NewHBox(gui.AppTalbeMenuFile(w)))
	w.SetContent(container.NewVBox(gui.GetSysteminfo(), gui.GetInterntinfo(), gui.SetOfficeAdr(&gui.OiffceAdandName{})))
	w.ShowAndRun()

}
