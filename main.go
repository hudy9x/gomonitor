package main

import (
	req "gomonitor/request"
	"gomonitor/utils"
	"log"
)

func main() {
	utils.GetIpNPcName()
	req.Post()
	if utils.IsLinux() {
		log.Println("Hello linux guys")
		// go monitor.RunLinux()
		// select {}
	}

}
