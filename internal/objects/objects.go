package objects

import (
	"image/color"
	"math"

	gsMath "github.com/Kaspetti/Gravity-Sim/internal/math"
	"github.com/hajimehoshi/ebiten/v2"
)


type Object struct {
	Mass        float64
	Radius      float64
	Position    gsMath.Vector2
    Velocity    gsMath.Vector2 

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
