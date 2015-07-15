package vec2

import (
	"fmt"
	"testing"
)

func TestInverse(t *testing.T) {
	mat := Rotation(1.4).Mul(Translation(Vector{1235, 623}))
	inverse := mat.Inverse()

	fmt.Println(mat)
	fmt.Println(inverse)
	fmt.Println(mat.Mul(inverse))
}
