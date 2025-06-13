package handler

// 响应数据格式类型

// ResponseStatus 当http服务器无法正确返回结果时，返回ResponseStatus
type ResponseStatus struct {
	//使用结构体标签tag，当对Status字段进行json或xml序列化时，字段名称为status
	Status string `json:"status" xml:"status"`
}

// STLFileList 返回stl文件列表的结构体类型
type STLFileList struct {
	STLList []string `json:"stl_list" xml:"stl_list"`
}

type STLFileName struct {
	Name string `json:"name" xml:"name"`
}
