package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/crypto/ssh"
)

func FetchAndSendDisksInfo() ([]byte, error) {
	host_server := os.Getenv("S_SERVER")
	string_port := os.Getenv("S_PORT")
	username := os.Getenv("S_USER")
	password := os.Getenv("S_PASS")

	port, _ := strconv.Atoi(string_port)

	// Establish an SSH connection
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host_server, port), &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	})
	if err != nil {
		// log.Fatalf("Failed to dial: %s\n", err)
		log.Printf("Failed to dial: %s\n", err)
		return nil, err
	}
	defer client.Close()

	// Run a command on the remote server
	session, err := client.NewSession()
	if err != nil {
		// log.Fatalf("Failed to create session: %s", err)
		log.Printf("Failed to create session: %s\n", err)
		return nil, err
	}
	defer session.Close()

	// Example command: "echo | format - Disk"
	cmd := "echo | format - Disk"

	
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		// log.Fatalf("Failed to run command: %s", err)
		log.Printf("Failed to run command: %s\n", err)
		return nil, err
	}

	fmt.Println("Command output:", string(output))
	return output, nil
}
