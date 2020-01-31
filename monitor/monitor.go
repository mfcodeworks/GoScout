package monitor

import (
	"fmt"
	"net"
	"time"
	"bytes"
	"strings"
	"runtime"
	"encoding/json"
    "github.com/shirou/gopsutil/host"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/disk"
    psnet "github.com/shirou/gopsutil/net"
)

type System struct {
	Name string `json:"name"`
	Version string `json:"version"`
}
type Nic struct {
	Name string `json:"name"`
	Mac string `json:"mac"`
	Address string `json:"address"`
	Address6 string `json:"address6"`
	Netmask string `json:"netmask"`
}
type Drive struct {
	Name string `json:"name"`
	MountPoint string `json:"mount_point"`
	Type string `json:"type"`
	TotalSize uint64 `json:"total_size"`
	UsedSize uint64 `json:"used_size"`
	PercentUsed float64 `json:"percent_used"`
}
type Device struct {
	Hostname string `json:"hostname"`
	System System `json:"system"`
	Uptime uint64 `json:"uptime"`
	CPUCount int `json:"cpu_count"`
	CPUUsage float64 `json:"cpu_usage"`
	MemoryTotal uint64 `json:"memory_total"`
	MemoryUsed uint64 `json:"memory_used"`
	MemoryUsedPercent float64 `json:"memory_used_percent"`
	Drives []Drive `json:"drives"`
	NetworkUp uint64 `json:"network_up"`
	NetworkDown uint64 `json:"network_down"`
	NetworkCards []Nic `json:"network_cards"`
	Timestamp string `json:"timestamp"`
	UUID string `json:"uuid"`
}

func (d *Device) ToJSON() *bytes.Buffer {
	json, _ := json.Marshal(d)
	return bytes.NewBuffer(json)
}

var hostinfo, _ = host.Info()

func ipv4Mask(m []byte) (string) {
    if len(m) != 4 {
        panic("ipv4Mask: len must be 4 bytes")
    }

    return fmt.Sprintf("%d.%d.%d.%d", m[0], m[1], m[2], m[3])
}

func getBandwidth() (uint64, uint64) {
	// Get net I/O
	io, _ := psnet.IOCounters(false)
	net1Out := io[0].BytesSent
	net1In := io[0].BytesRecv

	time.Sleep(1 * time.Second)

	// Get net I/O
	io2, _ := psnet.IOCounters(false)
	net2Out := io2[0].BytesSent
	net2In := io2[0].BytesRecv

	return net2In - net1In, net2Out - net1Out
}

func Hostname() (string) {
	return hostinfo.Hostname
}

func CPUInfo() (int, float64) {
	count := runtime.NumCPU()
	usage, _ := cpu.Percent(time.Second, false)
	return count, usage[0]
}

func MemoryInfo() (uint64, uint64, float64) {
	// Memory info
	memoryStats, _ := mem.VirtualMemory()
	return memoryStats.Total, memoryStats.Used, memoryStats.UsedPercent
}

func DiskInfo() ([]Drive) {
	disks := []Drive{}
	diskInfo, _ := disk.Partitions(false)

	for _, v := range diskInfo {
		x, _ := disk.Usage(v.Mountpoint)
		disk := Drive{
			v.Device,
			v.Mountpoint,
			v.Fstype,
			x.Total,
			x.Used,
			x.UsedPercent,
		}
		disks = append(disks, disk)
	}

	return disks
}

func NetworkBandwidthInfo() (uint64, uint64) {
	down, up := getBandwidth()
	return down, up
}

func NetworkInterfaceInfo() ([]Nic) {
	nics := []Nic{}
	interfaces, _ := net.Interfaces()

	for _, v := range interfaces {
		nic := Nic{
			Name: v.Name,
			Mac: v.HardwareAddr.String(),
		}

		snicArray, _ := v.Addrs()
		for _, snic := range snicArray {
			if len(strings.Split(snic.String(), ".")) == 4 {
				_, ipv4, _ := net.ParseCIDR(snic.String())
				nic.Address = ipv4.IP.String()
				nic.Netmask = ipv4Mask(ipv4.Mask)
			} else {
				_, ipv6, _ := net.ParseCIDR(snic.String())
				nic.Address6 = ipv6.IP.String()
			}
		}

		nics = append(nics, nic)
	}

	return nics
}

func OSInfo() (System) {
	return System{
		hostinfo.Platform,
		hostinfo.PlatformVersion,
	}
}

func UptimeInfo() (uint64) {
	return hostinfo.Uptime
}