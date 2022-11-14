package myexcel

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"gofyne_test/model"
)

func SaveInfomation(exc *model.ExcelInfo) error {

	////mo := gui.NetInfoP
	//fmt.Println("执行了保存了.!")
	//fmt.Printf("最后保存时候看看内存地址%p", exc)
	//exc := &model.ExcelInfo{}
	//hah := ExcelInfo{}
	//fmt.Println(hah.ComputerName)
	//fmt.Println("电脑名称归谁的:", exc.UserName)

	f := excelize.NewFile()
	index := f.NewSheet("Sheet1")
	//f.SetCellValue()
	f.SetCellValue("Sheet1", "A1", "所属人")
	f.SetCellValue("Sheet1", "B1", "办公室")
	f.SetCellValue("Sheet1", "C1", "系统版本")
	f.SetCellValue("Sheet1", "D1", "版本号")
	f.SetCellValue("Sheet1", "E1", "Cpu型号")
	f.SetCellValue("Sheet1", "F1", "CPU核心数")
	f.SetCellValue("Sheet1", "G1", "CPU主频")
	f.SetCellValue("Sheet1", "H1", "内存大小(GB)")
	f.SetCellValue("Sheet1", "I1", "硬盘大小(GB)")
	f.SetCellValue("Sheet1", "J1", "网卡名称")
	f.SetCellValue("Sheet1", "K1", "MAC")
	f.SetCellValue("Sheet1", "L1", "IPV4")

	//写入数据
	f.SetCellValue("Sheet1", "A2", exc.UserName)
	f.SetCellValue("Sheet1", "B2", exc.OfficeName)
	f.SetCellValue("Sheet1", "C2", exc.ComputerName)
	f.SetCellValue("Sheet1", "D2", exc.SystemVersion)
	f.SetCellValue("Sheet1", "E2", exc.CpuName)
	f.SetCellValue("Sheet1", "F2", exc.CpuCount)
	f.SetCellValue("Sheet1", "G2", exc.CpuHz)
	f.SetCellValue("Sheet1", "H2", exc.MemeoryInfo)
	f.SetCellValue("Sheet1", "I2", exc.DiskInfo)
	var netname string
	var netmac string
	var netip4 string
	for _, v := range exc.Netinfo {
		netname += v.Netname + " "
		netmac += v.MacInfo + " "
		netip4 += v.IpV4info + ""
	}
	f.SetCellValue("Sheet1", "J2", netname)
	f.SetCellValue("Sheet1", "K2", netmac)
	f.SetCellValue("Sheet1", "L2", netip4)
	f.SetActiveSheet(index)
	err := f.SaveAs("./" + exc.UserName + "_" + exc.OfficeName + ".xlsx")
	if err != nil {
		return err
	}

	return nil
}
