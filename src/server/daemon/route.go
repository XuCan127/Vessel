package daemon

import (
	"Vessel/src/common/jsonStruct"
	"Vessel/src/server/constant"
	"Vessel/src/server/imageBaseCenter"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (daemon *Daemon) VersionHandler(w http.ResponseWriter, r *http.Request) {
	response := jsonStruct.VersionResponse{
		Version: constant.Version,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (daemon *Daemon) ImageBaseAddHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	path := vars["path"]
	success := true
	msg := ""
	imageBase := jsonStruct.ImageBase{}

	isExist, err := imageBaseCenter.IsExist(name)
	if isExist {
		success = false
		msg = "Found exist imageBase."
		goto rtn
	}
	if err != nil {
		success = false
		msg = err.Error()
		goto rtn
	}

	imageBase, err = imageBaseCenter.AddImageBase(name, path)

rtn:
	// 创建响应对象
	response := jsonStruct.ImageBaseAddResponse{
		Success:   success,
		Msg:       msg,
		ImageBase: imageBase,
	}

	// 将响应对象转换为 JSON 格式
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 设置响应头部
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (daemon *Daemon) ImageBaseRemoveHandler(w http.ResponseWriter, r *http.Request) {

}
func (daemon *Daemon) ImageBaseListHandler(w http.ResponseWriter, r *http.Request) {

}

func (daemon *Daemon) ContainerLaunchHandler(w http.ResponseWriter, r *http.Request) {

}

func (daemon *Daemon) ContainerPSHandler(w http.ResponseWriter, r *http.Request) {

}
