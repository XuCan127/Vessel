package impl

import (
	"Vessel/src/common/term"
	"fmt"
	"log"
	"os"
)

func main() {
	// 创建基础层和顶层目录
	baseLayer := "base-layer"
	topLayer := "top-layer"
	err := os.Mkdir(baseLayer, 0755)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(topLayer, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// 挂载OverlayFS
	mergedLayer := "merged-layer"
	overlayWorkDir := "overlay-work"
	err = os.Mkdir(mergedLayer, 0755)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(overlayWorkDir, 0755)
	if err != nil {
		log.Fatal(err)
	}
	err = mountOverlayFS(baseLayer, topLayer, mergedLayer, overlayWorkDir)
	if err != nil {
		log.Fatal(err)
	}

}

// 使用mount系统调用挂载OverlayFS
func mountOverlayFS(lowerDir, upperDir, mergedDir, workDir string) error {
	return term.Mount("overlay", mergedDir, "overlay", 0,
		fmt.Sprintf("lowerdir=%s,upperdir=%s,workdir=%s", lowerDir, upperDir, workDir))
}
