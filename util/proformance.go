package util

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type SystemInfo struct {
	CPUUsage    float64
	MemoryUsage float64
	NetworkTX   uint64
	NetworkRX   uint64
	DiskRead    uint64
	DiskWrite   uint64
	Hostname    string
	OS          string
	Platform    string
	PlatformVer string
	KernelArch  string
	Uptime      uint64
}

func GetSystemInfo() SystemInfo {
	var info SystemInfo
	// CPU usage
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		fmt.Println("Error getting CPU usage:", err)
	} else {
		info.CPUUsage = cpuPercent[0]
	}

	// Memory usage
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error getting memory usage:", err)
	} else {
		info.MemoryUsage = memInfo.UsedPercent
	}

	// Network IO
	netIO, err := net.IOCounters(true)
	if err != nil {
		fmt.Println("Error getting network IO:", err)
	} else {
		info.NetworkRX = netIO[0].BytesRecv
		info.NetworkTX = netIO[0].BytesSent
	}

	// Disk IO
	diskIO, err := disk.IOCounters()
	if err != nil {
		fmt.Println("Error getting disk IO:", err)
	} else {
		info.DiskRead = diskIO["C:"].ReadBytes
		info.DiskWrite = diskIO["C:"].WriteBytes
	}

	// Host info
	hostInfo, err := host.Info()
	if err != nil {
		fmt.Println("Error getting host info:", err)
	} else {
		info.Hostname = hostInfo.Hostname
		info.OS = hostInfo.OS
		info.Platform = hostInfo.Platform
		info.PlatformVer = hostInfo.PlatformVersion
		info.KernelArch = hostInfo.KernelArch
		info.Uptime = hostInfo.Uptime
	}

	return info
}
