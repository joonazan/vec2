package vec2

import "math"

type Matrix [6]float64

var Identity Matrix = Matrix{
	1, 0, 0,
	0, 1, 0,
}

/*
Kaikki seuraavat
palauttavat matriisin, joka ...
*/

// kääntää vastapäivään
func Rotation(a float64) Matrix {
	sini := math.Sin(a)
	kosini := math.Cos(a)
	return Matrix{
		kosini, -sini, 0,
		sini, kosini, 0,
	}
}

// liikuttaa
func Translation(v Vector) Matrix {
	return Matrix{
		1, 0, v.X,
		0, 1, v.Y,
	}
}

func Scale(x, y float64) Matrix {
	return Matrix{
		x, 0, 0,
		0, y, 0,
	}
}

// multiplies m with v as if v was (x, y, 1)
func (m Matrix) Transform(v Vector) Vector {
	return Vector{
		m[0]*v.X + m[1]*v.Y + m[2],
		m[3]*v.X + m[4]*v.Y + m[5],
	}
}

// multiplies m with v as if v was (x, y, 0)
func (m Matrix) TransformDirection(v Vector) Vector {
	return Vector{
		m[0]*v.X + m[1]*v.Y,
		m[3]*v.X + m[4]*v.Y,
	}
}

// skalaari kertaa matriisi
func (m Matrix) Mulf(r float64) Matrix {
	for i := range m {
		m[i] *= r
	}
	return m
}

// matriisitulo
func (m Matrix) Mul(n Matrix) Matrix {
	return Matrix{
		m[0]*n[0] + m[1]*n[3], m[0]*n[1] + m[1]*n[4], m[0]*n[2] + m[1]*n[5] + m[2],
		m[3]*n[0] + m[4]*n[3], m[3]*n[1] + m[4]*n[4], m[3]*n[2] + m[4]*n[5] + m[5],
	}
}

// käänteismatriisi
func (m Matrix) Inverse() Matrix {
	return Matrix{
		m[4], -m[1], m[1]*m[5] - m[2]*m[4],
		-m[3], m[0], m[2]*m[3] - m[0]*m[5],
	}.Mulf(1 / (m[0]*m[4] - m[1]*m[3]))
}

// muunna 32-bittiseksi
func (m Matrix) To32() [6]float32 {
	n := [6]float32{}
	for i, f64 := range m {
		n[i] = float32(f64)
	}
	return n
}
