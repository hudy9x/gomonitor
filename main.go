package main

import (
	"gomonitor/monitor"
	"gomonitor/utils"
	"log"
	"os"
)

func main() {
	Home, _ := os.UserHomeDir()
	log.Println("home user", Home)

	utils.GetIpNPcName()
	// req.Post()

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
