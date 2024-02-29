package main

import (
	"gomonitor/monitor"
	req "gomonitor/request"
	"gomonitor/utils"
	"log"
)

func main() {
	utils.GetIpNPcName()
	req.Post()
	// this function
	if utils.IsLinux() {
		log.Println("Hello linux guys")
		go monitor.RunLinux()
		select {}
	}

	if utils.IsWindows() {
		log.Println("hello windows user")
		go monitor.RunWindow()
		select {}
	}

	if utils.IsMacOs() {
		log.Println("hello macos user")
		go monitor.RunMacos()
		select {}
	}

}
