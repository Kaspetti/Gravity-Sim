package game

import (
	"image/color"

	gsMath "github.com/Kaspetti/Gravity-Sim/internal/math"
	"github.com/Kaspetti/Gravity-Sim/internal/objects"
	"github.com/hajimehoshi/ebiten/v2"
)


const (
    WORLDWIDTH  = 12800
    WORLDHEIGHT = 7200
)


type Game struct {
    Objects []objects.Object
}


func (g *Game) Update() error {
    return nil
}


func (g *Game) Draw(screen *ebiten.Image) {
    for _, obj := range g.Objects {
        obj.Draw(screen)
    }
}


func (g *Game) Layout(screenWidth, screenHeight int) (viewportWidth, viewportHeight int) {
    return WORLDWIDTH, WORLDHEIGHT
}


func StartGame() error {
    game := Game{
        Objects: []objects.Object{
            {
                Mass: 1,
                Radius: 250,
                Position: gsMath.Vector2{X: 500, Y: 500},
                Velocity: gsMath.Vector2{X: 0, Y: 0},

                Color: color.RGBA{255, 0, 0, 255},
            },
        },
    }

    ebiten.SetWindowSize(1280, 720)
    ebiten.SetWindowTitle("Gravity Simulator")

    return ebiten.RunGame(&game)
}
