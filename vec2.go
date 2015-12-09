package vec2

import "math"

type Vector struct {
	X, Y float64
}

func (self *Vector) Add(other Vector) {
	self.X += other.X
	self.Y += other.Y
}

func (self *Vector) Sub(other Vector) {
	self.X -= other.X
	self.Y -= other.Y
}

func (self *Vector) Mul(other float64) {
	self.X *= other
	self.Y *= other
}

func (self Vector) Dot(other Vector) float64 {
	return self.X*other.X + self.Y*other.Y
}

func (self Vector) Cross(other Vector) float64 {
	return self.X*other.Y - self.Y*other.X
}

func (self Vector) Crossf(other float64) Vector {
	return Vector{-self.Y * other, self.X * other}
}

func (self Vector) LengthSquared() float64 {
	return self.X*self.X + self.Y*self.Y
}

func (self Vector) Length() float64 {
	return math.Sqrt(self.LengthSquared())
}

func (self *Vector) Normalize() {
	self.Mul(1 / self.Length())
}

func (self Vector) Normalized() Vector {
	return Mul(self, 1/self.Length())
}

func (v Vector) Plus(v2 Vector) Vector {
	return Add(v, v2)
}

func (v Vector) Minus(v2 Vector) Vector {
	return Sub(v, v2)
}

func (v Vector) Times(r float64) Vector {
	return Mul(v, r)
}

func Add(v, u Vector) Vector {
	return Vector{v.X + u.X, v.Y + u.Y}
}

func Sub(v, u Vector) Vector {
	return Vector{v.X - u.X, v.Y - u.Y}
}

func Mul(v Vector, r float64) Vector {
	return Vector{v.X * r, v.Y * r}
}
