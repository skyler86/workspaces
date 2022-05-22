package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
)

func main() {
	//建立SSH客户端连接
	client, err := ssh.Dial("tcp", "192.168.31.150:22", &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password("rainbow")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		log.Fatalf("SSH dial error: %s", err.Error())
	}

	// 建立新会话
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("new session error: %s", err.Error())
	}
	defer session.Close()

	var buffer bytes.Buffer
	session.Stdout = &buffer
	if err := session.Run("ls /"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(buffer.String())
}
