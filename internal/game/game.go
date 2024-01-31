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
    for i := range g.Objects {
        g.Objects[i].Update(g.Objects)
    }

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
            //{
            //    Id: 1,
            //    Mass: 50000,
            //    Radius: 250,
            //    Position: gsMath.Vector2{X: 6400, Y: 3600},
            //    Velocity: gsMath.Vector2{X: 0, Y: 0},
            //    Delta:  gsMath.Vector2{X: 0, Y: 0},

            //    Color: color.RGBA{255, 0, 0, 255},
            //},
            //{
            //    Id: 2,
            //    Mass: 5,
            //    Radius: 100,
            //    Position: gsMath.Vector2{X: 3000, Y: 2000},
            //    Velocity: gsMath.Vector2{X: 0, Y: 7},
            //    Delta:  gsMath.Vector2{X: 0, Y: 0},

            //    Color: color.RGBA{0, 255, 0, 255},
            //},
            {
                Id: 1,
                Mass: 10000,
                Radius: 200,
                Position: gsMath.Vector2{X: 7000, Y: 3600},
                Velocity: gsMath.Vector2{X: 0, Y: 3},
                Delta: gsMath.Vector2{X: 0, Y: 0},

                Color: color.RGBA{255, 0, 0, 255},
            },
            {
                Id: 2,
                Mass: 5000,
                Radius: 100,
                Position: gsMath.Vector2{X: 5000, Y: 3600},
                Velocity: gsMath.Vector2{X: 0, Y: -3},
                Delta: gsMath.Vector2{X: 0, Y: 0},

                Color: color.RGBA{255, 0, 0, 255},
            },
        },
    }

    ebiten.SetWindowSize(1280, 720)
    ebiten.SetWindowTitle("Gravity Simulator")
    ebiten.SetTPS(240)

    return ebiten.RunGame(&game)
}
