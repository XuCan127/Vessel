package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <listen-addr>\n", os.Args[0])
		os.Exit(1)
	}

	listenAddr := os.Args[1]

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %s", listenAddr, err)
	}
	defer ln.Close()

	fmt.Printf("Listening on %s\n", listenAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %s", err)
			continue
		}

		go func() {
			defer conn.Close()

			// 读取第一行输入作为可执行文件路径
			reader := bufio.NewReader(conn)
			cmdPath, err := reader.ReadString('\n')
			if err != nil {
				log.Printf("Failed to read command path: %s", err)
				return
			}
			cmdPath = strings.TrimSpace(cmdPath)

			// 执行可执行文件，并将输入输出重定向给用户
			cmd := exec.Command(cmdPath)
			cmd.Stdin = conn
			cmd.Stdout = conn
			cmd.Stderr = conn
			if err := cmd.Run(); err != nil {
				log.Printf("Failed to run command %s: %s", cmdPath, err)
				return
			}
		}()
	}
}
