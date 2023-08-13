package containerCenter

import (
	"Vessel/src/common/jsonStruct"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func DeFile() {
	filePath := "vesselfile"
	lines, err := readLines(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	baseImage := ""
	mem := ""
	cpu := ""
	net := ""
	var portMappings []jsonStruct.PortMapping
	commands := []string{}
	execCommand := ""

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue // 忽略空行和注释行
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			log.Fatalf("Invalid line: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "Base":
			baseImage = value
		case "Mem":
			mem = value
		case "Cpu":
			cpu = value
		case "Net":
			net = value
		case "Ports":
			portMappings = parsePortMappings(value)
		case "Run":
			commands = append(commands, value)
		case "Exec":
			execCommand = value
		default:
			log.Fatalf("Unknown key: %s", key)
		}
	}
	fmt.Println(baseImage)
	fmt.Println(mem)
	fmt.Println(cpu)
	fmt.Println(net)
	fmt.Println(commands)
	fmt.Println(execCommand)
	fmt.Println(portMappings)
}

func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func parsePortMappings(value string) []jsonStruct.PortMapping {
	portMappings := []jsonStruct.PortMapping{}
	mappings := strings.Split(value, ",")
	for _, mapping := range mappings {
		mapping = strings.TrimSpace(mapping)
		if mapping == "" {
			continue
		}

		parts := strings.SplitN(mapping, ":", 2)
		if len(parts) != 2 {
			log.Fatalf("Invalid port mapping: %s", mapping)
		}

		hostPort, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			log.Fatalf("Invalid host port: %v", err)
		}

		containerPort, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			log.Fatalf("Invalid container port: %v", err)
		}

		portMappings = append(portMappings, jsonStruct.PortMapping{
			HostPort:      hostPort,
			ContainerPort: containerPort,
		})
	}

	return portMappings
}

//Base: ubuntu20
//Mem: 1G
//Cpu: 0,5,1
//Net: bridge
//Ports: 4444:1928,7777:2048
//Run: apt update
//Run: apt update
//Run: apt-get -y update
//Exec: ./main
