package rao

type SendApiResp struct {
	Code int         `json:"code"`
	Em   string      `json:"em"`
	Et   string      `json:"et"`
	Data interface{} `json:"data"`
}

type FileUploadBase64Req struct {
	FileString string `json:"file_string"`
	PathDir    string `json:"path_dir"`
	FileName   string `json:"file_name"`
}

type FileUploadBase64Resp struct {
	Path string `json:"path"`
}
