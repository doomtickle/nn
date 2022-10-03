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
			a[i][j] = 1
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
			m.Arr[i][j] += float64(rand.RandInt(10))
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
      for k:=0; k<a.Cols;k++ {
        sum += a.Arr[i][k] * b.Arr[k][j]
      }
      res.Arr[i][j] = sum
		}
	}

  return res
}

func (m *Matrix) Transpose() *Matrix {
  res := NewMatrix(m.Cols, m.Rows)

  for i:=0;i<m.Rows;i++{
    for j:=0;j<m.Cols; j++ {
      res.Arr[j][i] = m.Arr[i][j]
    }
  }

  return res
}
