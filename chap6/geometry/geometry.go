package geometry

import "math"

type Point struct{ X, Y float64 }

// 昔ながらの関数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Point 型のメソッド
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Path: Point を直線で結びつける道のり
type Path []Point

// path に沿って進んだ距離を返す
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
