package imageBaseCenter

import (
	"Vessel/src/common/jsonStruct"
	"fmt"
)

func IsExist(name string) (bool, error) {
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

func AddImageBase(newBase jsonStruct.ImageBase) error {
	oldImageFile, err := readDatabase()
	if err != nil {
		return fmt.Errorf("Failed to get ImageBaseFile: %v", err)
	}
	oldImageFile = append(oldImageFile, newBase)

	return writeDatabase(oldImageFile)
}

func ListImageBases() ([]jsonStruct.ImageBase, error) {
	return readDatabase()
}
