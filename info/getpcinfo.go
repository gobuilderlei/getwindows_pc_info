package info

import (
	"fmt"
	"github.com/StackExchange/wmi"
	"net"
	"strings"
	"syscall"
	"unsafe"
)

var (
	advapi = syscall.NewLazyDLL("Advapi32.dll")
	kernel = syscall.NewLazyDLL("Kernerl32.dll")
	PcInfo map[string]interface{}
)

func GetPcInfo() map[string]interface{} {
	PcInfo = make(map[string]interface{})
	//PcInfo["info"], _ = host.Info()
	//PcInfo["cpu"], _ = cpu.Info()
	//PcInfo["mem"], _ = mem.VirtualMemory()
	PcInfo["network"] = GetIntfs()
	//PcInfo["username"] = GetUserName()
	//PcInfo["versionnow"] = runtime.GOOS
	//PcInfo["pcosversion"] = GetSystemVersion()
	//PcInfo["bios"] = GetBiosInfo()
	//PcInfo["motherboard"] = GetMotherboardInfo()
	//PcInfo["Cpu"] = GetCpuInfo()
	//PcInfo["Gpu"] = GetGPUInfo()

	//fmt.Println(GetUserName())
	return PcInfo
}

// 获取用户名
func GetUserName() string {
	var size uint32 = 128
	var buffer = make([]uint16, size)
	user := syscall.StringToUTF16Ptr("USERNAME")
	domain := syscall.StringToUTF16Ptr("USERDOMAIN")
	r, err := syscall.GetEnvironmentVariable(user, &buffer[0], size)
	if err != nil {
		return ""
	}
	buffer[r] = '@'
	old := r + 1
	if old > size {
		return syscall.UTF16ToString(buffer[:r])
	}
	r, err1 := syscall.GetEnvironmentVariable(domain, &buffer[old], size-old)
	if err1 != nil {
		return ""
	}
	return syscall.UTF16ToString(buffer[:old+r])
}

// 系统版本
func GetSystemVersion() string {
	version, err := syscall.GetVersion()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%d.%d (%d)", byte(version), uint8(version>>8), version>>16)
}

type diskusage struct {
	Path  string `json:"path"`
	Total uint64 `json:"total"`
	Free  uint64 `json:"free"`
}

func usage(getDiskFreeSpaceExW *syscall.LazyProc, path string) (diskusage, error) {
	lpFreeBytesAvailable := int64(0)
	var info = diskusage{Path: path}
	diskret, _, err := getDiskFreeSpaceExW.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(info.Path))),
		uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
		uintptr(unsafe.Pointer(&(info.Total))),
		uintptr(unsafe.Pointer(&(info.Free))))
	if diskret != 0 {
		err = nil
	}
	return info, err
}

// 硬盘信息
func GetDiskInfo() (infos []diskusage) {
	GetLogicalDriveStringsW := kernel.NewProc("GetLogicalDriveStringsW")
	GetDiskFreeSpaceExW := kernel.NewProc("GetDiskFreeSpaceExW")
	lpBuffer := make([]byte, 254)
	diskret, _, _ := GetLogicalDriveStringsW.Call(
		uintptr(len(lpBuffer)),
		uintptr(unsafe.Pointer(&lpBuffer[0])))
	if diskret == 0 {
		return
	}
	for _, v := range lpBuffer {
		if v >= 65 && v <= 90 {
			path := string(v) + ":"
			if path == "A:" || path == "B:" {
				continue
			}
			info, err := usage(GetDiskFreeSpaceExW, string(v)+":")
			if err != nil {
				continue
			}
			infos = append(infos, info)
		}
	}
	return infos
}

// CPU信息
// 简单的获取方法fmt.Sprintf("Num:%d Arch:%s\n", runtime.NumCPU(), runtime.GOARCH)
type CpuInfo struct {
	Name       string
	Numbers    uint32
	TreadCount uint32
}

func GetCpuInfo() []CpuInfo {
	//var size uint32 = 128
	//var buffer = make([]uint16, size)
	//var index = uint32(copy(buffer, syscall.StringToUTF16("Num:")) - 1)
	//nums := syscall.StringToUTF16Ptr("NUMBER_OF_PROCESSORS")
	//arch := syscall.StringToUTF16Ptr("PROCESSOR_ARCHITECTURE")
	//r, err := syscall.GetEnvironmentVariable(nums, &buffer[index], size-index)
	//if err != nil {
	//	return ""
	//}
	//index += r
	//index += uint32(copy(buffer[index:], syscall.StringToUTF16(" Arch:")) - 1)
	//r, err = syscall.GetEnvironmentVariable(arch, &buffer[index], size-index)
	//if err != nil {
	//	return syscall.UTF16ToString(buffer[:index])
	//}
	//index += r
	//return syscall.UTF16ToString(buffer[:index+r])
	var cpuinfo []CpuInfo
	err := wmi.Query("select *form Win32_Processor", &cpuinfo)
	if err != nil {
		fmt.Println("获取CPU失败", err)
	}
	return cpuinfo
}

// 获取GPU信息
type gpuInfo struct {
	Name string
}

func GetGPUInfo() string {
	var gpuinfo []gpuInfo
	err := wmi.Query("select * from Win32_videoContorller", &gpuinfo)
	if err != nil {
		return ""
	}
	return gpuinfo[0].Name
}

type memoryStatusEx struct {
	cbSize                  uint32
	dwMemoryLoad            uint32
	ullTotalPhys            uint64 // in bytes
	ullAvailPhys            uint64
	ullTotalPageFile        uint64
	ullAvailPageFile        uint64
	ullTotalVirtual         uint64
	ullAvailVirtual         uint64
	ullAvailExtendedVirtual uint64
}

// 内存信息
func GetMemory() string {
	GlobalMemoryStatusEx := kernel.NewProc("GlobalMemoryStatusEx")
	var memInfo memoryStatusEx
	memInfo.cbSize = uint32(unsafe.Sizeof(memInfo))
	mem, _, _ := GlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memInfo)))
	if mem == 0 {
		return ""
	}
	return fmt.Sprint(memInfo.ullTotalPhys / (1024 * 1024))
}

type intfInfo struct {
	Name  string
	Macad string
	Ipv4  []string
	Ipv6  []string
}

// 网卡信息
func GetIntfs() []intfInfo {
	intf, err := net.Interfaces()
	if err != nil {
		return []intfInfo{}
	}
	//for index, value := range intf {
	//	//fmt.Println(index)
	//	//fmt.Println(value)
	//}
	var is = make([]intfInfo, len(intf))
	for i, v := range intf {
		ips, err := v.Addrs()
		if err != nil {
			continue
		}

		is[i].Name = v.Name
		is[i].Macad = fmt.Sprint(v.HardwareAddr)
		for _, ip := range ips {
			if strings.Contains(ip.String(), ":") {
				is[i].Ipv6 = append(is[i].Ipv6, ip.String())
			} else {
				is[i].Ipv4 = append(is[i].Ipv4, ip.String())
			}
		}
	}
	return is
}

// 主板信息
func GetMotherboardInfo() string {
	var s = []struct {
		Product string
	}{}
	err := wmi.Query("SELECT  Product  FROM Win32_BaseBoard WHERE (Product IS NOT NULL)", &s)
	if err != nil {
		return ""
	}
	return s[0].Product
}

// BIOS信息
func GetBiosInfo() string {
	var s = []struct {
		Name string
	}{}
	err := wmi.Query("SELECT Name FROM Win32_BIOS WHERE (Name IS NOT NULL)", &s) // WHERE (BIOSVersion IS NOT NULL)
	if err != nil {
		return ""
	}
	return s[0].Name
}

type Storage struct {
	Name       string
	FileSystem string
	Total      uint64
	Free       uint64
}

type storageInfo struct {
	Name       string
	Size       uint64
	FreeSpace  uint64
	FileSystem string
}

func GetStorageInfo() string {
	var storageinfo []storageInfo
	var loaclStorages []Storage
	err := wmi.Query("Select * from Win32_LogicalDisk", &storageinfo)
	if err != nil {
		return ""
	}
	var size uint64
	for _, storage := range storageinfo {
		info := Storage{
			Name:       storage.Name,
			FileSystem: storage.FileSystem,
			Total:      storage.Size,
			Free:       storage.FreeSpace,
		}
		size += storage.Size
		loaclStorages = append(loaclStorages, info)
	}
	fmt.Printf("localStorages:=", loaclStorages)

	return fmt.Sprint(size / 1000 / 1000 / 1000)
}
