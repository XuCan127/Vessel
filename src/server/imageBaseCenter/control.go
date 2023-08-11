package imageBaseCenter

import (
	"Vessel/src/common/jsonStruct"
	"fmt"
	"sync"
)

var controlMutex sync.Mutex

func IsNameExist(name string) (bool, error) {
	controlMutex.Lock()
	defer controlMutex.Unlock()
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

func IsIdExist(Id string) (bool, error) {
	controlMutex.Lock()
	defer controlMutex.Unlock()
	imageFile, err := readDatabase()
	if err != nil {
		return false, fmt.Errorf("Failed to get ImageBaseFile: %v", err)
	}

	// 遍历ImageBase列表，查找具有特定名称的ImageBase对象
	for _, image := range imageFile {
		if image.ID == Id {
			return true, nil // 返回找到的ImageBase对象
		}
	}

	return false, nil
}

func AddImageBase(newBase jsonStruct.ImageBase) error {
	controlMutex.Lock()
	defer controlMutex.Unlock()
	oldImageFile, err := readDatabase()
	if err != nil {
		return fmt.Errorf("Failed to get ImageBaseFile: %v", err)
	}
	oldImageFile = append(oldImageFile, newBase)

	return writeDatabase(oldImageFile)
}

func RemoveImageBase(Id string) error {
	controlMutex.Lock()
	defer controlMutex.Unlock()
	oldImageFile, err := readDatabase()
	var newImageBases []jsonStruct.ImageBase
	if err != nil {
		return fmt.Errorf("Failed to get ImageBaseFile: %v", err)
	}
	for _, base := range oldImageFile {
		if base.ID == Id {
			continue
		}
		newImageBases = append(newImageBases, base)
	}

	return writeDatabase(newImageBases)
}

func ListImageBases() ([]jsonStruct.ImageBase, error) {
	controlMutex.Lock()
	defer controlMutex.Unlock()
	return readDatabase()
}
