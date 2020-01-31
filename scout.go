package main

import (
	"fmt"
	"flag"
	"time"
	"net/http"
	"goscout/monitor"
	"github.com/denisbrodbeck/machineid"
)

var (
	dest string
	interval, attempts, timeout int
	systemUUID, _ = machineid.ID() // System UUID
)

func init() {
	// Define arg constants
	const (
		destDefault = "http://localhost:8080/"
		destUsage = "API Endpoint for Monitoring Data"
		intervalDefault = 5
		intervalUsage = "Interval between checks in seconds"
		attemptsDefault = 30
		attemptsUsage = "Attempts to send data when sending failes"
		timeoutDefault = 60
		timeoutUsage = "Timeout between resend attempts in seconds (If attempts is reached script will die)"
	)
	flag.StringVar(&dest, "dest", destDefault, destUsage)
	flag.StringVar(&dest, "d", destDefault, "dest (shorthand)")
	flag.IntVar(&interval, "interval", intervalDefault, intervalUsage)
	flag.IntVar(&interval, "i", intervalDefault, "interval (shorthand)")
	flag.IntVar(&attempts, "attempts", attemptsDefault, attemptsUsage)
	flag.IntVar(&attempts, "a", attemptsDefault, "attempts (shorthand)")
	flag.IntVar(&timeout, "timeout", timeoutDefault, timeoutUsage)
	flag.IntVar(&timeout, "t", timeoutDefault, "timeout (shorthand)")
	flag.Parse()

	// Adjust interval for bandwidth & cpu checking
	interval -= 2
}

func main() {
	// Print begin string and parsed arguments
	fmt.Println("GoScout Begin")

	// Run monitor on loop
	for {
		runCheck()
		fmt.Println("-----------------------------------------------------------------")
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func runCheck() {
	// Hostname info
	hostname := monitor.Hostname()
	fmt.Printf("Hostname: %v\n", hostname)

	// CPU info
	cpuCount, cpuUsage := monitor.CPUInfo()
	fmt.Printf("CPU:\n\tCount: %v\n\tUsage: %v\n", cpuCount, cpuUsage)

	// Memory info
	memoryTotal, memoryUsed, memoryUsedPercent := monitor.MemoryInfo()
	fmt.Printf("Memory:\n\tPercent: %v\n\tTotal: %v MB\n\tUsed: %v MB\n", memoryUsedPercent, memoryTotal / 1e+6, memoryUsed / 1e+6)

	// Disk Info
	disks := monitor.DiskInfo()
	for _, disk := range disks {
		fmt.Printf("Disk:\n\tName: %v\n\tMount Point: %v\n\tType: %v\n\tSize: %v\n\tUsage: %v\n\tPercent Used: %v\n", disk.Name, disk.MountPoint, disk.Type, disk.TotalSize / 1e+9, disk.UsedSize / 1e+9, disk.PercentUsed)
	}

	// Bandwidth Info
	networkDown, networkUp := monitor.NetworkBandwidthInfo()
	fmt.Printf("Network:\n\tTraffic in: %v\n\tTraffic out: %v\n", networkDown / 1e+3, networkUp / 1e+3)

	// Network Info
	nics := monitor.NetworkInterfaceInfo()
	for _, nic := range nics {
		fmt.Printf("NIC:\n\t%v\n", nic)
	}

	// Platform Info
	system := monitor.OSInfo()
	fmt.Printf("OS:\n\t %v %v\n", system.Name, system.Version)

	// Time Info
	timestamp := time.Now().Format(time.RFC3339)
	uptime := monitor.UptimeInfo()
	fmt.Printf("System Uptime: %v\nTimestamp: %v\n", uptime, timestamp)

	// System UUID
	fmt.Printf("UUID: %v\n", systemUUID)

	device := monitor.Device{
		Hostname: hostname,
		System: system,
		Uptime: uptime,
		CPUCount: cpuCount,
		CPUUsage: cpuUsage,
		MemoryTotal: memoryTotal,
		MemoryUsed: memoryUsed,
		MemoryUsedPercent: memoryUsedPercent,
		Drives: disks,
		NetworkUp: networkUp,
		NetworkDown: networkDown,
		NetworkCards: nics,
		Timestamp: timestamp,
		UUID: systemUUID,
	}

	fmt.Printf("Data: %v\n", device)

	send(device)
}

func send(data monitor.Device) {
	for i := 0; i < attempts; i++ {
		fmt.Printf("Posting: %v\n", data.ToJSON())
		response, err := http.Post(dest, "application/json", data.ToJSON())
		if err == nil && response.StatusCode == 200 {
			fmt.Printf("\nPOST:\nResponse: %v\nHeaders: %v\nContent: %v\n", response.StatusCode, response.Header, response.Body)
			return
		}
		fmt.Printf("\nPOST ERROR: %v\n", err)
		time.Sleep(time.Duration(timeout) * time.Second)
	}
}
