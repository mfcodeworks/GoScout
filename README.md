# GoScout-Server-Monitor

[![GitHub license](https://img.shields.io/github/license/mfsoftworks/GoScout-Server-Monitor.svg)](https://github.com/mfsoftworks/GoScout-Server-Monitor/blob/master/LICENSE.md)

## Purpose

The executable, a port of the Python [Server Monitoring Script](https://github.com/mfsoftworks/server-monitoring-script), is designed to be run as a cronjob on every boot to run in the background.
The script will gather information:

- UUID (Unique for each system to avoid overlapping hostname for multi-network monitoring)
- Hostname
- CPU
- Memory
- Network Usage
- Network Cards
- Hard Drives
- System OS
- System Uptime
- UTC Timestamp

The executable will produce a JSON output at set intervals for use with any software or server accepting a JSON input.
Example:

```json
{
    "hostname": "HOME-LAPTOP1",
    "system": {
        "name": "Windows",
        "version": "10"
    },
    "uptime" : 231199,
    "cpu_count": 4,
    "cpu_usage": 17.9,
    "memory_total": 8440942592,
    "memory_used": 6244225024,
    "memory_used_percent": 74.0,
    "drives": [
        {
            "name": "C:",
            "mount_point": "C:",
            "type": "NTFS",
            "total_size": 536224985088,
            "used_size": 167306108928,
            "percent_used": 31.2
        },
        {
            "name": "D:",
            "mount_point": "D:",
            "type": "NTFS",
            "total_size": 463332921344,
            "used_size": 49498419200,
            "percent_used": 10.7
        }
    ],
    "network_up": 54,
    "network_down": 4150,
    "network_cards": [
        {
            "address": "127.0.0.1",
            "address6": "::1",
            "mac": "",
            "name": "Loopback Pseudo-Interface 1",
            "netmask": "255.0.0.0"
        },
        {
            "address": "10.15.62.112",
            "address6": "fe80::844d:a87:54ea:2100",
            "mac": "1C-39-47-A6-4C-5E",
            "name": "Ethernet",
            "netmask": "255.255.0.0"
        }
    ],
    "timestamp" : "2018-10-10T01:41:21+00:00",
    "uuid" : 180331603484325
}
```

The script includes a function to POST JSON to a remote server.

This script can be installed on several machines that report to a central monitoring server.

The destination, checking interval, sending attempts after failure and timeout between resending attempts can be set through arguments, use `goscout -h` for more.

## Usage

Download the latest [release](https://github.com/mfsoftworks/GoScout-Server-Monitor/releases).

To test the script output run with `goscout` or `goscout.exe` in the terminal.

### **Linux Autostart**

Create a cron job to run the script on every boot.

Edit cron with `crontab -e`.

Add the script at the bottom of the cron list as `@reboot /path/to/script/goscout`.

### **Windows Autostart**

To create an autostart for Windows open start menu and search for `Task Scheduler`.

Select `Create Task`.

Enter a name e.g. monitor.

Add new trigger at login.

Add a new action and select the `goscout.exe` executable.

**Unselect** Stop task if it runs longer than 3 days.

Select Okay.

A new task will be created that runs the python monitoring script in the background every time a user starts the system and logs in.

## Author

MF Softworks <mf@nygmarosebeauty.com>

mf.nygmarosebeauty.com
