package main

import "github.com/Kaspetti/Gravity-Sim/internal/game"


func main() {
    if err := game.StartGame(); err != nil {
        panic(err)
    }
}
