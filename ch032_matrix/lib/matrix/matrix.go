package matrix

import "errors"

type Matrix [][]float64 // 矩阵

// GetShape 定义获取矩阵形状的方法
func (m1 Matrix) GetShape() (s [2]int) {
	s[0] = len(m1)
	s[1] = len(m1[0])
	return
}

// Mul 定义矩阵与标量c相乘的运算
func (m1 *Matrix) Mul(c float64) {
	row, col := m1.GetShape()[0], m1.GetShape()[1]
	mTemp := make([][]float64, row, row)
	//将m1矩阵复制到临时变量复制给临时变量mTemp
	copy(mTemp, *m1)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			mTemp[i][j] *= c
		}
	}
	//将临时变量mTemp矩阵复制回给m1
	copy(*m1, mTemp)
}

// NewMatrix 矩阵的初始化方法，r为矩阵的行数，c为矩阵的列数
func NewMatrix(r, c int) (mat Matrix, err error) {
	//行数或列数小于等于零，返回err
	if r <= 0 || c <= 0 {
		err = errors.New("rows and columns of the matrix must >0")
		return
	}
	//初始化mat的行
	mat = make([][]float64, r, r)
	//此处不使用for range是因为我们要改变遍历元素的值
	for i := 0; i < r; i++ {
		mat[i] = make([]float64, c, c)
	}
	return
}

// Add 矩阵的加法
func (m1 Matrix) Add(m2 Matrix) (m3 Matrix, err error) {
	if m1.GetShape() != m2.GetShape() {
		err = errors.New("the shape of the two matrix are not equal")
		return
	}
	//获取矩阵的形状
	r, c := m1.GetShape()[0], m1.GetShape()[1]
	m3, _ = NewMatrix(r, c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			m3[i][j] = m1[i][j] + m2[i][j]
		}
	}
	return
}

// Minus 矩阵的减法
func (m1 Matrix) Minus(m2 Matrix) (m3 Matrix, err error) {
	if m1.GetShape() != m2.GetShape() {
		err = errors.New("the shape of the two matrix are not equal")
		return
	}
	r, c := m1.GetShape()[0], m1.GetShape()[1]
	m3, _ = NewMatrix(r, c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			m3[i][j] = m1[i][j] - m2[i][j]
		}
	}
	return
}

// MatMul 两个矩阵相乘
func (m1 Matrix) MatMul(m2 Matrix) (m3 Matrix, err error) {
	r1, c1 := m1.GetShape()[0], m1.GetShape()[1]
	r2, c2 := m2.GetShape()[0], m2.GetShape()[1]
	if c1 != r2 {
		err = errors.New("the shape of the two matrix are not match")
		return
	}
	//初始化计算结果m3
	m3, _ = NewMatrix(r1, c2)
	for i := 0; i < r1; i++ {
		for j := 0; j < c2; j++ {
			//取出m1的第i行
			v1 := RowVector(m1[i])
			v2 := make([]float64, r2, r2)
			for k := 0; k < r2; k++ {
				//将m2的第j列值依次放入v2中
				v2[k] = m2[k][j]
			}
			//利用行向量乘法计算m3的矩阵元素
			m3[i][j], _ = v1.Dot(v2)
		}
	}
	return
}

// Transpose 定义矩阵的转置运算，本质上就是将矩阵元素的索引值i和j进行互换
func (m1 Matrix) Transpose() (m2 Matrix) {
	r, c := m1.GetShape()[0], m1.GetShape()[1]
	m2, _ = NewMatrix(c, r)
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			m2[i][j] = m1[j][i]
		}
	}
	return
}
