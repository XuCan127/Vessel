package jsonStruct

type ImageBase struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CreatedTime string `json:"create time"` //创建时间
}

type ImageBaseAddResponse struct {
	Success   bool      `json:"success"`
	Msg       string    `json:"msg"`
	ImageBase ImageBase `json:"image base"`
}

type ImageBaseRemoveResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

type ImageBaseListResponse struct {
	Success   bool        `json:"success"`
	Msg       string      `json:"msg"`
	ImageBase []ImageBase `json:"image bases"`
}
