package math


type Vector2 struct {
    X   float64
    Y   float64
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
