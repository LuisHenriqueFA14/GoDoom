package main

import (
    "github.com/LuisHenriqueFA14/GoDoom/internal"
    "fmt"
    "time"
)

const (
    WIDTH = 30
    HEIGHT = 30
)

func main() {
    pixels := internal.CreateFire(WIDTH, HEIGHT)

    for {
        fmt.Print("\033[H\033[2J")
        pixels = internal.UpdateFire(pixels, WIDTH, HEIGHT)
        internal.RenderFire(pixels, WIDTH, HEIGHT)
        time.Sleep(100 * time.Millisecond)
    }
}
