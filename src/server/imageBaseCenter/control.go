package imageBaseCenter

import (
	"Vessel/src/common/jsonStruct"
	"Vessel/src/common/term"
	"fmt"
	"time"
)

import "sync"

var mutex sync.Mutex

func init() {

}

func IsExist(name string) (bool, error) {
	mutex.Lock()
	defer mutex.Unlock()

	imageFile, err := readDatabase()
	if err != nil {
		return false, fmt.Errorf("Failed to get ImageBaseFile: %v", err)
	}

	// 遍历ImageBase列表，查找具有特定名称的ImageBase对象
	for _, image := range imageFile {
		if image.Name == name {
			return true, nil // 返回找到的ImageBase对象
		}
	}

	return false, nil
}

func AddImageBase(name, path string) (jsonStruct.ImageBase, error) {
	mutex.Lock()
	defer mutex.Unlock()

	newBase := jsonStruct.ImageBase{Name: name, CreatedTime: time.Now().Format(term.CreateTimeFmt)}
	oldImageFile, err := readDatabase()
	if err != nil {
		return jsonStruct.ImageBase{}, fmt.Errorf("Failed to get ImageBaseFile: %v", err)
	}
	oldImageFile = append(oldImageFile, newBase)

	writeDatabase(oldImageFile)
	return newBase, nil
}
