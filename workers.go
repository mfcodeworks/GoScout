package main

// func runCheck() {
// 	// Hostname info
// 	hostname := Hostname()
// 	fmt.Printf("Hostname: %v\n", hostname)

// 	// CPU info
// 	cpuCount, cpuUsage := CPUInfo()
// 	fmt.Printf("CPU:\n\tCount: %v\n\tUsage: %v\n", cpuCount, cpuUsage)

// 	// Memory info
// 	memoryTotal, memoryUsed, memoryUsedPercent := MemoryInfo()
// 	fmt.Printf("Memory:\n\tPercent: %v\n\tTotal: %v MB\n\tUsed: %v MB\n", memoryUsedPercent, memoryTotal / 1e+6, memoryUsed / 1e+6)

// 	// Disk Info
// 	disks := DiskInfo()
// 	for _, disk := range disks {
// 		fmt.Printf("Disk:\n\tName: %v\n\tMount Point: %v\n\tType: %v\n\tSize: %v\n\tUsage: %v\n\tPercent Used: %v\n", disk.Name, disk.MountPoint, disk.Type, disk.TotalSize / 1e+9, disk.UsedSize / 1e+9, disk.PercentUsed)
// 	}

// 	// Bandwidth Info
// 	networkDown, networkUp := NetworkBandwidthInfo()
// 	fmt.Printf("Network:\n\tTraffic in: %v\n\tTraffic out: %v\n", networkDown / 1e+3, networkUp / 1e+3)

// 	// Network Info
// 	nics := NetworkInterfaceInfo()
// 	for _, nic := range nics {
// 		fmt.Printf("NIC:\n\t%v\n", nic)
// 	}

// 	// Platform Info
// 	system := OSInfo()
// 	fmt.Printf("OS:\n\t %v %v\n", system.Name, system.Version)

// 	// Time Info
// 	timestamp := time.Now().Format(time.RFC3339)
// 	uptime := UptimeInfo()
// 	fmt.Printf("System Uptime: %v\nTimestamp: %v\n", uptime, timestamp)

// 	// System UUID
// 	fmt.Printf("UUID: %v\n", systemUUID)

// 	device := Device{
// 		Hostname: hostname,
// 		System: system,
// 		Uptime: uptime,
// 		CPUCount: cpuCount,
// 		CPUUsage: cpuUsage,
// 		MemoryTotal: memoryTotal,
// 		MemoryUsed: memoryUsed,
// 		MemoryUsedPercent: memoryUsedPercent,
// 		Drives: disks,
// 		NetworkUp: networkUp,
// 		NetworkDown: networkDown,
// 		NetworkCards: nics,
// 		Timestamp: timestamp,
// 		UUID: systemUUID,
// 	}

// 	fmt.Printf("Data: %v\n", device)

// 	send(device)
// }

// func hostnameWorker(result &chan string) {
// 	// Hostname info
// 	hostname := Hostname()
// 	fmt.Printf("Hostname: %v\n", hostname)
// }