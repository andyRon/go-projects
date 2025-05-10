package stl

// TriangleFace STL文件中的三角面元
type TriangleFace struct {
	N [3]float32 `json:"n" bson:"n"`
	A [3]float32 `json:"a" bson:"n"`
	B [3]float32 `json:"b" bson:"n"`
	C [3]float32 `json:"c" bson:"n"`
}

// 识别的STL文件是二进制格式
// 二进制STL文件用固定的字节数来给出三角面元的几何信息。
// 文件起始的80个字节是文件头，用于存贮文件名；
// 紧接着用 4 个字节的整数来描述模型的三角面元个数，
// 后面逐个给出每个三角面元的几何信息。每个三角面元占用固定的50个字节，依次是:
// 3个4字节浮点数(三角面元的法线矢量)
// 3个4字节浮点数(三角面元第1个顶点的坐标)
// 3个4字节浮点数(三角面元第2个顶点的坐标)
// 3个4字节浮点数(三角面元第3个顶点的坐标)
// 三角面片的最后2个字节用来描述三角面元的属性信息。
// 一个完整二进制STL文件的大小为三角面元数乘以 50再加上84个字节。

type ModelSTL struct {
	Name              string         `json:"name" bson:"name"`
	FaceNum           int32          `json:"face_num" bson:"face_num"`
	TriangleFaceArray []TriangleFace `json:"triangle_face_array" bson:"triangle_face_array"`
}
