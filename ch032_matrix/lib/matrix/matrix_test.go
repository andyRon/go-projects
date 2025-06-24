package matrix

import "testing"

func TestMatrix_GetShape(t *testing.T) {
	m := Matrix([][]float64{[]float64{1, 2, 3}, []float64{4, 5, 6}})
	t.Log("shape of m is:", m.GetShape())
}

func TestMatrix_Mul(t *testing.T) {
	m := Matrix([][]float64{[]float64{1, 2, 3}, []float64{4, 5, 6}})
	t.Log("m=", m)
	m.Mul(2)
	t.Log("m=", m)
}

func TestNewMatrix(t *testing.T) {
	m, _ := NewMatrix(3, 4)
	t.Log("m=", m)
}

func TestMatrix_Add(t *testing.T) {
	m1 := Matrix([][]float64{[]float64{1, 2, 3}, []float64{4, 5, 6}})
	m2 := Matrix([][]float64{[]float64{14, 12, 10}, []float64{16, 17, 19}})
	m3, _ := m1.Add(m2)
	t.Log("m3=", m3)
}

func TestMatrix_Minus(t *testing.T) {
	m1 := Matrix([][]float64{[]float64{1, 2, 3}, []float64{4, 5, 6}})
	m2 := Matrix([][]float64{[]float64{11, 12, 10}, []float64{16, 17, 19}})
	m3, _ := m1.Minus(m2)
	t.Log("m3=", m3)
}

func TestMatrix_MatMul(t *testing.T) {
	m1 := Matrix([][]float64{[]float64{1, 2, 3}, []float64{4, 5, 6}})
	m2 := Matrix([][]float64{[]float64{11, 12}, []float64{13, 14}, []float64{15, 16}})
	m3, _ := m1.MatMul(m2)
	t.Log("m3=", m3)
}

func TestMatrix_Transpose(t *testing.T) {
	m := Matrix([][]float64{[]float64{1, 2, 3}, []float64{4, 5, 6}})
	t.Log("m=", m)
	mt := m.Transpose()
	t.Log("mt=", mt)
}
