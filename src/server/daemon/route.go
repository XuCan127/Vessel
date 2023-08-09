package daemon

import (
	"Vessel/src/common/jsonStruct"
	"Vessel/src/server/constant"
	"encoding/json"
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
	// 从URL路径中获取名称参数
	//vars := mux.Vars(r)
	//name := vars["name"]
	//path := vars["path"]

	// 假设添加图片成功
	success := true
	msg := "Image added successfully."

	// 创建响应对象
	response := jsonStruct.ImageBaseAddResponse{
		Success: success,
		Msg:     msg,
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
