package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// getServiceName returns the systemd service name with .service suffix
func getServiceName(filename string) string {
	parts := strings.Split(filename, "/")
	name := parts[len(parts)-1]
	return name + ".service"
}

// createServiceFile generates a .service file for systemd
func createServiceFile(filename string, extraArgs []string) {
	extraArgsJoined := strings.Join(extraArgs, " ")
	serviceContent := fmt.Sprintf(`[Unit]
Description=Service automatically created using createservice
After=network-online.target

[Service]
ExecStart=%s %s
Restart=always

[Install]
WantedBy=multi-user.target
`, filename, extraArgsJoined)

	file, err := os.Create(getServiceName(filename))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err = file.WriteString(serviceContent); err != nil {
		panic(err)
	}
}

// enableService enables the service using systemctl
func enableService(serviceFile string) {
	out, err := exec.Command("systemctl", "enable", serviceFile).CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		panic(err)
	}
}

// startService starts the service using systemctl
func startService(serviceName string) {
	out, err := exec.Command("systemctl", "start", serviceName).CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		panic(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./createservice <path-to-executable> [args...]")
		return
	}

	filename := os.Args[1]
	extraArgs := os.Args[2:]

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("Error: Executable not found: %s\n", filename)
		return
	}

	fmt.Println("Creating systemd service file...")
	createServiceFile(filename, extraArgs)

	fmt.Println("🔗 Enabling service to start on boot...")
	enableService(getServiceName(filename))

	fmt.Println("🚀 Starting service now...")
	startService(getServiceName(filename))

	fmt.Printf("Service '%s' created, enabled, and started successfully!\n", getServiceName(filename))
}
