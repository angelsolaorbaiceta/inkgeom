# inkgeom

Module for 2D and 3D geometry.

## 2D Geometry

The basis of two dimensional geometry are points and vectors.
To create a point use the `MakePoint` function:

```go
var (
  pA = g2d.MakePoint(4.0, 5.0)
  pB = g2d.MakePoint(12.0, 45.0)
)
```

To compute the distance between two points:

```go
distance := pA.distanceTo(pB)
```

To create a vector between two points (`pA -> pB`):

```go
vec := pA.VectorTo(pB)
```

Both vectors and versos can be created using the `MakeVector` and `MakeVersor` functions:

```go
var (
  vector = g2d.MakeVector(1, 2) // { 1, 2 }
  versor = g2d.MakeVersor(1, 2) // { 1 / sqrt(5), 2 / sqrt(5) }
)
