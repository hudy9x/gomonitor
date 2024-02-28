package main

import (
	"gomonitor/utils"
	"log"
)

func main() {
	utils.GetIpNPcName()
	if utils.IsLinux() {
		log.Println("Hello linux guys")
		// go monitor.RunLinux()
		// select {}
	}

}
