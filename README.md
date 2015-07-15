#vec2, a 2d vector math library for Golang
This library provides 2d vectors and transformation matrices.

For vectors there are in-place operations as methods and copying operations as functions. There are multiple convenient matrix constructors. Matrices can be multiplied with each other and can be used to transform vectors.

The special thing about this library is that it has 2d transform matrices that leave out the third row (because it is always 0 0 1) to save space and time. For that reason there are two methods for transforming a vector with a matrix: `Transform` and `TransformDirection`. The latter one leaves out the translation, which is just like multiplying a 3x3 matrix with (x, y, 0).

`v.Cross(v2)` calculates the 3d cross product of two 2d vectors and returns the resulting z as a scalar. `v.Crossf(r)` takes a scalar that is the length of a z-axis aligned vector.
