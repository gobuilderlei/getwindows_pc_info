package gui

import (
	"bytes"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"gofyne_test/info"
	"gofyne_test/model"
	"image/color"
)

type OiffceAdandName struct {
	Name     string
	OfficeAd string
}

func insertNth(s string, n int) string {
	var buffer bytes.Buffer
	var n_1 = n - 1
	var l_1 = len(s) - 1
	for i, rune := range s {
		buffer.WriteRune(rune)
		if i%n == n_1 && i != l_1 {
			buffer.WriteRune('\n')
		}
	}
	return buffer.String()
}

var (
	NetInfoP *model.ExcelInfo
	NetInfo  = model.ExcelInfo{}
)

//制作主窗口

// 系统信息
func GetSysteminfo() *fyne.Container {
	//CPU信息
	CPU, err := cpu.Info()
	if err != nil {
		fmt.Println("获取错误", err)
	}
	//for _, v := range CPU {
	//	//fmt.Println(v)
	//}
	//计算机信息
	info1, e1 := host.Info()
	if e1 != nil {
	}
	//var pnetinfo *model.ExcelInfo

	//物理内存
	infoMemeroy, _ := mem.VirtualMemory()
	infoMemeroyint := infoMemeroy.Total / 1024 / 1024 / 1000

	//netinfo := &model.ExcelInfo{}
	NetInfo.ComputerName = info1.Hostname
	NetInfo.SystemVersion = info1.PlatformVersion
	NetInfo.MemeoryInfo = fmt.Sprint(infoMemeroyint)
	fmt.Println("结构体ComputerName为", NetInfo.ComputerName)
	versionLable := canvas.NewText("系统版本:"+info1.Platform, color.Black)
	versionvalue := canvas.NewText("版本号:"+info1.PlatformVersion, color.Black)
	Cpuinfo := canvas.NewText("CPU信息:", color.Black)
	var Cpuname *canvas.Text
	if len(CPU) > 0 {
		Cpuname = canvas.NewText("CPU型号:"+CPU[0].ModelName+" 核心数:"+
			fmt.Sprint(CPU[0].Cores)+" 主频:"+fmt.Sprint(CPU[0].Mhz),
			color.Black)
		NetInfo.CpuName = CPU[0].ModelName
		NetInfo.CpuCount = fmt.Sprint(CPU[0].Cores)
		NetInfo.CpuHz = fmt.Sprint(CPU[0].Mhz)
	} else {
		Cpuname = canvas.NewText("CPU信息获取失败:", //+
			//fmt.Sprint(CPU[0].Numbers)+"线程数:"+fmt.Sprint(CPU[0].TreadCount),
			color.Black)
	}
	NetInfoP = &NetInfo
	fmt.Printf("首次打印NetInfoP的地址:%p \t", NetInfoP)
	systeminfo := container.NewVBox(
		container.New(layout.NewVBoxLayout(),
			canvas.NewText("用户名:"+info1.Hostname, color.Black),
			versionLable, versionvalue,
			canvas.NewText("内  存:"+fmt.Sprint(infoMemeroyint)+"GB", color.Black),
			Cpuinfo, Cpuname,
			GetDiskInfo(),
		))

	return systeminfo
}
func GetInterntinfo() *fyne.Container {
	//var p_netinfo *model.ExcelInfo
	//hha := &model.ExcelInfo{}
	content := container.NewVBox(widget.NewLabel("网卡信息"))
	infomat := info.GetIntfs()
	//var value string
	for _, v := range infomat {
		NameLable := canvas.NewText(v.Name, color.Black)
		MACad := canvas.NewText(" MAC地址:"+v.Macad, color.Black)
		IPv4Value := canvas.NewText(" IP地址:"+fmt.Sprint(v.Ipv4), color.Black)
		//将数据写入到 Netinfo       []Netinfomation 结构体数组
		NetInfo.Netinfo = append(NetInfo.Netinfo,
			model.Netinfomation{Netname: v.Name, MacInfo: v.Macad, IpV4info: fmt.Sprint(v.Ipv4)})
		content.Add(container.New(layout.NewHBoxLayout(), NameLable, MACad, IPv4Value))
	}
	NetInfoP = &NetInfo
	//fmt.Printf(p_netinfo.CpuName)
	fmt.Printf("再次打印NetInfoP的地址:%p \t", NetInfoP)
	return content
}

// 磁盘信息
func GetDiskInfo() *fyne.Container {
	content := container.NewVBox(widget.NewLabel("磁盘信息:"))
	//磁盘信息
	//infodisk, _ := disk.IOCounters()
	//fmt.Println(infodisk)
	infodisk := info.GetStorageInfo()
	//var big uint64
	//for i, v := range infodisk {
	//	big = v.ReadBytes + v.WriteBytes
	//	big = big / 1024 / 1024
	content.Add(container.New(layout.NewHBoxLayout(), canvas.NewText("磁盘大小: "+infodisk+"GB", color.Black)))

	NetInfo.DiskInfo = infodisk
	NetInfoP = &NetInfo
	return content
}

func SetOfficeAdr(na *OiffceAdandName) *fyne.Container {
	//hah := &model.ExcelInfo{}
	content := container.NewVBox(widget.NewLabel("输入您的信息:"))
	//officenamelabel := widget.NewLabel("办公室编号")
	//officenamelabel.Wrapping = fyne.TextWrapBreak
	inputofficename := widget.NewEntry()
	inputofficename.TextStyle = fyne.TextStyle{Bold: true}
	inputofficename.SetPlaceHolder("输入您所在办公室楼层所属部门及房号,例如基地设备科402室")
	inputofficename.Validator = func(s string) error {
		if s == "" || s == "输入您所在办公室楼层所属部门及房号,例如基地设备科402室" {
			return errors.New("错误")
		}
		fmt.Println(s)
		//else {
		//na.OfficeAd = s
		NetInfo.OfficeName = s
		NetInfoP = &NetInfo
		return nil
		//}

	}

	//content := container.NewVBox(widget.NewLabel("所属人:"))
	//namelabel := widget.NewLabel("电脑所属人姓名:")
	//var nameoo *model.ExcelInfo
	//fmt.Printf("在输入框内部的用户名的指针为%p", nameoo)
	nameentry := widget.NewEntry()
	nameentry.TextStyle = fyne.TextStyle{Bold: true}
	nameentry.SetPlaceHolder("输入电脑归属人的姓名,例如张三")
	nameentry.Validator = func(s string) error {
		if s == "" || s == "输入电脑归属人的姓名,例如张三" {
			return errors.New("错误")
		} else {
			//na.Name = s
			NetInfo.UserName = s
			//nameoo.UserName = s
			//fmt.Printf("在输入框内部的用户名的指针为 %p", nameoo)
			return nil
		}

	}
	content.Add(container.New(layout.NewVBoxLayout(), inputofficename, nameentry))
	return content
}
func SetUserNameInfo() *fyne.Container {
	content := container.NewVBox(widget.NewLabel("所属人:"))
	namelabel := widget.NewLabel("电脑所属人姓名:")
	nameentry := widget.NewEntry()
	nameentry.TextStyle = fyne.TextStyle{Bold: true}
	content.Add(container.New(layout.NewVBoxLayout(), namelabel, nameentry))
	return content
}
