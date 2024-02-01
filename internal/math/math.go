package math

import "math"

type Vector2 struct {
	X float64
	Y float64
}

func (v1 Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{
		v1.X + v2.X,
		v1.Y + v2.Y,
	}
}

func (v1 Vector2) Sub(v2 Vector2) Vector2 {
	return Vector2{
		v1.X - v2.X,
		v1.Y - v2.Y,
	}
}

func (v Vector2) Mult(s float64) Vector2 {
	return Vector2{
		v.X * s,
		v.Y * s,
	}
}

func (v1 Vector2) Dist(v2 Vector2) float64 {
	return math.Sqrt(math.Pow(v1.X - v2.X, 2) + math.Pow(v1.Y - v2.Y, 2))
}


func (v Vector2) Norm() Vector2 {
    return v.Mult(1/v.Len())
}


func (v Vector2) Len() float64 {
    return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}
