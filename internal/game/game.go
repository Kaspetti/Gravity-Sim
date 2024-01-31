package game

import "github.com/hajimehoshi/ebiten/v2"


const (
    WORLDWIDTH  = 10000
    WORLDHEIGHT = 10000
)


type Game struct {
}


func (g *Game) Update() error {
    return nil
}


func (g *Game) Draw(screen *ebiten.Image) {

}


func (g *Game) Layout(screenWidth, screenHeight int) (viewportWidth, viewportHeight int) {
    return WORLDWIDTH, WORLDHEIGHT
}


func StartGame() error {
    game := Game{}

    ebiten.SetWindowSize(1280, 720)
    ebiten.SetWindowTitle("Gravity Simulator")

    return ebiten.RunGame(&game)
}
