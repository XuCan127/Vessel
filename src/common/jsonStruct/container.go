package jsonStruct

type PortMapping struct {
	HostPort      int `json:"dst"`
	ContainerPort int `json:"src"`
}

type Container struct {
	Pid         string        `json:"pid"`
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	CreatedTime string        `json:"createTime"`
	Volume      string        `json:"volume"`
	PortMaps    []PortMapping `json:"port maps"`
}

type ContainerLaunchResponse struct {
	Success    bool      `json:"success"`
	Msg        string    `json:"msg"`
	Containers Container `json:"container"`
}

type ContainerPSResponse struct {
	Success    bool        `json:"success"`
	Msg        string      `json:"msg"`
	Containers []Container `json:"containers"`
}
