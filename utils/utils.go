package utils

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/vikash-parashar/asset-locator/logger"
	"golang.org/x/crypto/ssh"
)

func executeCommand(command string, shell string) ([]byte, error) {
	var cmd *exec.Cmd

	switch shell {
	case "bash":
		cmd = exec.Command("bash", "-c", command)
	case "sh":
		cmd = exec.Command("sh", "-c", command)
	default:
		return nil, fmt.Errorf("unsupported shell type: %s", shell)
	}

	return cmd.CombinedOutput()
}

func FetchAndSendDisksInfo() ([]byte, error) {
	hostServer := os.Getenv("S_SERVER")
	stringPort := os.Getenv("S_PORT")
	username := os.Getenv("S_USER")
	password := os.Getenv("S_PASS")

	port, _ := strconv.Atoi(stringPort)

	// Establish an SSH connection
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", hostServer, port), &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	})
	if err != nil {
		logger.ErrorLogger.Printf("Failed to dial: %s\n", err)
		return nil, err
	}
	defer client.Close()

	// Run a command on the remote server using the executeCommand function
	cmd := "echo | format - Disk"
	output, err := executeCommand(cmd, "sh")
	if err != nil {
		logger.ErrorLogger.Printf("Failed to run command: %s\n", err)
		return nil, err
	}

	logger.InfoLogger.Println("Command output:", string(output))
	return output, nil
}

func commandHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Fetch and send disks info using SSH command
		output, err := FetchAndSendDisksInfo()
		if err != nil {
			logger.ErrorLogger.Printf("Error fetching disks info: %s\n", err)
			http.Error(w, fmt.Sprintf("Error fetching disks info: %s", err), http.StatusInternalServerError)
			return
		}

		logger.InfoLogger.Println("Disks info fetched successfully.")
		fmt.Fprintf(w, "Disks Info:\n%s", output)
		return
	}

	logger.WarningLogger.Println("Invalid request method.")
	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}
