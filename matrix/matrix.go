package matrix

import (
	"log"

	"github.com/doomtickle/nn/rand"
)

type Matrix struct {
	Rows, Cols int
	Arr        [][]float64
}

func NewMatrix(rows, cols int) *Matrix {
	var (
		m Matrix
	)

	m.Rows = rows
	m.Cols = cols
	a := make([][]float64, m.Rows)
	for i := 0; i < m.Rows; i++ {
		a[i] = make([]float64, m.Cols)
		for j := 0; j < m.Cols; j++ {
			a[i][j] = 0
		}
	}
	m.Arr = a

	return &m
}

func (m *Matrix) Scale(n float64) {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			m.Arr[i][j] *= n
		}
	}
}

func (m *Matrix) Randomize() {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			m.Arr[i][j] = rand.RandFloat(-1, 1)
		}
	}
}

func (m *Matrix) Add(n float64) {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			m.Arr[i][j] += n
		}
	}
}

func (m *Matrix) AddMatrix(n *Matrix) {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			m.Arr[i][j] += n.Arr[i][j]
		}
	}
}

func (m *Matrix) SubtractMatrix(n *Matrix) {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			m.Arr[i][j] -= n.Arr[i][j]
		}
	}
}

func Subtract(a, b *Matrix) *Matrix {
  res := NewMatrix(a.Rows, a.Cols)
	for i := 0; i < a.Rows; i++ {
		for j := 0; j < a.Cols; j++ {
			res.Arr[i][j] = a.Arr[i][j] - b.Arr[i][j]
		}
	}

  return res
}

func (m *Matrix) MultMatrixByElement(n *Matrix) {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			m.Arr[i][j] *= n.Arr[i][j]
		}
	}
}

func (m *Matrix) MultMatrix(n *Matrix) *Matrix {
	if m.Cols != n.Rows {
		log.Fatal("Cols of a must match rows of B.")
	}
	a := m
	b := n
	res := NewMatrix(a.Rows, b.Cols)

	for i := 0; i < res.Rows; i++ {
		for j := 0; j < res.Cols; j++ {
			sum := float64(0)
			for k := 0; k < a.Cols; k++ {
				sum += a.Arr[i][k] * b.Arr[k][j]
			}
			res.Arr[i][j] = sum
		}
	}

	return res
}

func Transpose(m *Matrix) *Matrix {
	res := NewMatrix(m.Cols, m.Rows)

	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			res.Arr[j][i] = m.Arr[i][j]
		}
	}

	return res
}

type MapFunc func(float64) float64

func StaticMap(m *Matrix, fn MapFunc) *Matrix {
  result := NewMatrix(m.Rows, m.Cols)
  a := make([][]float64, m.Rows)
	for i := 0; i < m.Rows; i++ {
    a[i] = make([]float64, m.Cols)
		for j := 0; j < m.Cols; j++ {
			a[i][j] = fn(m.Arr[i][j])
		}
	}

  result.Arr = a

  return result
}

func (m *Matrix) Map(fn MapFunc) {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			m.Arr[i][j] = fn(m.Arr[i][j])
		}
	}
}

func FromArray(arr []float64) *Matrix {
	m := NewMatrix(len(arr), 1)
	for i := 0; i < len(arr); i++ {
		m.Arr[i][0] = arr[i]
	}

	return m
}

func (m *Matrix) ToArray() []float64 {
  var a []float64
  for i:=0;i<m.Rows;i++{
    for j:=0;j<m.Cols;j++ {
      a = append(a, m.Arr[i][j])
    }
  }

  return a
}

func Multiply(a, b *Matrix) *Matrix {
	if a.Cols != b.Rows {
		log.Fatal("Cols of a must match rows of B.")
	}
	res := NewMatrix(a.Rows, b.Cols)

	for i := 0; i < res.Rows; i++ {
		for j := 0; j < res.Cols; j++ {
			sum := float64(0)
			for k := 0; k < a.Cols; k++ {
				sum += a.Arr[i][k] * b.Arr[k][j]
			}
			res.Arr[i][j] = sum
		}
	}

	return res
}
