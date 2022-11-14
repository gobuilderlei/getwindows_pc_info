package model

type ExcelInfo struct {
	UserName      string          //用户信息
	OfficeName    string          //办公室编号
	ComputerName  string          //用户名
	SystemVersion string          //系统版本
	CpuName       string          //CPU型号
	CpuCount      string          //CPU核心数
	CpuHz         string          //CPU主频
	MemeoryInfo   string          //内存大小
	DiskInfo      string          //硬盘大小
	Netinfo       []Netinfomation //网卡信息

}

type Netinfomation struct {
	Netname  string //网卡名称
	MacInfo  string //mac 地址
	IpV4info string //ipv4
}
