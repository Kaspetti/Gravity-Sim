package objects

import (
	"image/color"
	"math"

	gsMath "github.com/Kaspetti/Gravity-Sim/internal/math"
	"github.com/hajimehoshi/ebiten/v2"
)


const (
    G = 6.674
)


type Object struct {
    Id          int
	Mass        float64
	Radius      float64
	Position    gsMath.Vector2
    Velocity    gsMath.Vector2 
    Delta       gsMath.Vector2

    Color       color.RGBA
}


func (obj Object) Draw(screen *ebiten.Image) {
    width, height := int(obj.Radius*2), int(obj.Radius*2)

    img := ebiten.NewImage(width, height)

    pixels := make([]byte, width*height*4)
    x, y := 0.0, 0.0

    for i := 0; i < len(pixels); i+=4 {
        dSqr := math.Pow(x-obj.Radius, 2) + math.Pow(y-obj.Radius, 2)
        if dSqr <= obj.Radius*obj.Radius {
            pixels[i] = obj.Color.R
            pixels[i+1] = obj.Color.G
            pixels[i+2] = obj.Color.B
            pixels[i+3] = obj.Color.A
        } else {
            pixels[i] = 0
            pixels[i+1] = 0
            pixels[i+2] = 0
            pixels[i+3] = 0
        }

        x += 1
        if x >= float64(width) {
            x = 0
            y += 1
        }
    }

    img.WritePixels(pixels)

    options := ebiten.DrawImageOptions{}
    options.GeoM.Translate(obj.Position.X-obj.Radius, obj.Position.Y-obj.Radius)

    screen.DrawImage(img, &options)
}


func (obj1 *Object) Update(objects []Object) {
    obj1.Delta = gsMath.Vector2{X: 0, Y: 0}

    for _, obj2 := range objects {
        if obj1.Id == obj2.Id {
            continue
        }

        dist := obj1.Position.Dist(obj2.Position)

        f := G * ((obj1.Mass * obj2.Mass) / math.Pow(dist, 2))
        dir := obj2.Position.Sub(obj1.Position).Normalize()
        fVec := dir.Mult(f)

        obj1.Delta = obj1.Delta.Add(fVec.Mult(1 / obj1.Mass))
    }

    obj1.Velocity = obj1.Velocity.Add(obj1.Delta)
    obj1.Position = obj1.Position.Add(obj1.Velocity)
}
