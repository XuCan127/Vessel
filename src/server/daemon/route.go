package daemon

import (
	"Vessel/src/common/jsonStruct"
	"Vessel/src/common/term"
	"Vessel/src/server/constant"
	"Vessel/src/server/imageBaseCenter"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func (daemon *Daemon) VersionHandler(w http.ResponseWriter, r *http.Request) {
	sendResponse(jsonStruct.VersionResponse{
		Version: constant.Version,
	}, w)
}

func (daemon *Daemon) ImageBaseAddHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	argName := vars["name"]
	argPath := vars["path"]

	imageBaseID := uuid.New().String()
	unzipPath := filepath.Join(constant.ImagesDir, imageBaseID)

	// 创建响应对象
	response := jsonStruct.ImageBaseAddResponse{
		Success:   true,
		Msg:       "",
		ImageBase: jsonStruct.ImageBase{},
	}

	isExist, err := imageBaseCenter.IsNameExist(argName)
	if isExist {
		response.Success = false
		response.Msg = "Found exist imageBase."
		goto rtn
	}
	if err != nil {
		response.Success = false
		response.Msg = err.Error()
		goto rtn
	}

	err = term.Unzip(argPath, unzipPath)
	if err != nil {
		response.Success = false
		response.Msg = err.Error()
		goto rtn
	}

	response.ImageBase.Name = argName
	response.ImageBase.CreatedTime = time.Now().Format(term.CreateTimeFmt)
	response.ImageBase.ID = imageBaseID

	err = imageBaseCenter.AddImageBase(response.ImageBase)
	if err != nil {
		response.Success = false
		response.Msg = err.Error()
		os.RemoveAll(unzipPath)
	}

rtn:
	sendResponse(response, w)
}

func (daemon *Daemon) ImageBaseRemoveHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	argImageId := vars["imageId"]

	unzipPath := filepath.Join(constant.ImagesDir, argImageId)

	// 创建响应对象
	response := jsonStruct.ImageBaseRemoveResponse{
		Success: true,
		Msg:     "",
	}
	isExist, err := imageBaseCenter.IsIdExist(argImageId)
	if !isExist {
		response.Success = false
		response.Msg = "Not Found exist imageBase."
		goto rtn
	}
	if err != nil {
		response.Success = false
		response.Msg = err.Error()
		goto rtn
	}
	err = imageBaseCenter.RemoveImageBase(argImageId)
	if err != nil {
		response.Success = false
		response.Msg = err.Error()
		goto rtn
	}
	os.RemoveAll(unzipPath)

rtn:
	sendResponse(response, w)
}
func (daemon *Daemon) ImageBaseListHandler(w http.ResponseWriter, r *http.Request) {
	response := jsonStruct.ImageBaseListResponse{
		Success:    true,
		Msg:        "",
		ImageBases: nil,
	}

	bases, err := imageBaseCenter.ListImageBases()
	if err != nil {
		response.Success = false
		response.Msg = err.Error()
		goto rtn
	}
	response.ImageBases = bases
rtn:
	sendResponse(response, w)
}

func (daemon *Daemon) ContainerLaunchHandler(w http.ResponseWriter, r *http.Request) {

}

func (daemon *Daemon) ContainerPSHandler(w http.ResponseWriter, r *http.Request) {

}

func sendResponse(responseJson any, w http.ResponseWriter) {
	// 将响应对象转换为 JSON 格式
	jsonResponse, err := json.Marshal(responseJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 设置响应头部
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
