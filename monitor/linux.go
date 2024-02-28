package monitor

import (
	"fmt"
	"log"
	"time"

	"io/ioutil"
	"os"
	"strings"

	"github.com/LightningDev1/go-foreground"
)

// getAppName retrieves the application name from its PID on Linux.
func getAppName(pid uint32) (string, error) {
	// Check if the process exists
	_, err := os.Stat(fmt.Sprintf("/proc/%d", pid))
	if os.IsNotExist(err) {
		return "", fmt.Errorf("process with PID %d does not exist", pid)
	}

	// Read the process's cmdline file to get the command-line arguments
	cmdlineBytes, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid))
	if err != nil {
		return "", fmt.Errorf("failed to read cmdline for PID %d: %v", pid, err)
	}

	// Extract the application name from the cmdline
	cmdline := string(cmdlineBytes)
	parts := strings.Fields(cmdline)
	if len(parts) > 0 {
		appName := parts[0]
		// Extract the executable name from the path
		if strings.Contains(appName, "/") {
			appName = appName[strings.LastIndex(appName, "/")+1:]
		}
		return appName, nil
	}

	return "", fmt.Errorf("unable to determine application name for PID %d", pid)
}

func RunLinux() {
	for {

		pid, _ := foreground.GetForegroundPID()
		title, _ := foreground.GetForegroundTitle()

		name, _ := getAppName(pid)

		fmt.Printf("PID: %d\ntitle: %s\n", pid, title)
		log.Println(name)

		time.Sleep(3 * time.Second)
	}
}
