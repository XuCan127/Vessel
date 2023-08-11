package imageBaseCenter

import (
	"Vessel/src/common/jsonStruct"
	"Vessel/src/server/constant"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func readDatabase() ([]jsonStruct.ImageBase, error) {
	fd, err := os.OpenFile(constant.ImagesList, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("open file " + constant.ImagesList + " fail: " + err.Error())
	}
	defer fd.Close()

	var images []jsonStruct.ImageBase

	// 逐行读取文件内容
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		// 解析JSON数据到ImageBase对象
		var image jsonStruct.ImageBase
		err := json.Unmarshal([]byte(line), &image)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON data: %v", err)
		}

		images = append(images, image)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	return images, nil
}

// 将数据库内容写入文件
func writeDatabase(images []jsonStruct.ImageBase) error {
	fd, err := os.OpenFile(constant.ImagesList, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("open file " + constant.ImagesList + " fail: " + err.Error())
	}
	defer fd.Close()
	// 清空文件内容
	err = fd.Truncate(0)
	if err != nil {
		return fmt.Errorf("failed to truncate file: %v", err)
	}

	// 将每个ImageBase对象转换为JSON字符串，并写入文件
	writer := bufio.NewWriter(fd)
	for _, image := range images {
		jsonData, err := json.Marshal(image)
		if err != nil {
			return fmt.Errorf("failed to marshal image to JSON: %v", err)
		}

		_, err = writer.WriteString(string(jsonData) + "\n")
		if err != nil {
			return fmt.Errorf("failed to write to file: %v", err)
		}
	}

	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("failed to flush writer: %v", err)
	}

	return nil
}
