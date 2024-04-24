package internal

import (
    "fmt"
)


var firePalette = []string {
        " ",
        ".",
        ",",
        "|",
        "?",
        "{",
        "}",
        "$",
        "#",
        "M",
	}


func CreateFire(width, height int) []int {
    pixelsArray := make([]int, width*height)

    for j := 0; j < width*height; j++ {
        pixelsArray[j] = 0
    }

    return pixelsArray
}

func UpdateFire(pixelsArray []int, width int, height int) []int {
    for j := width*height - 1; j >= 0; j-- {
        UpdatePixel(pixelsArray, j, width, height)
    }

    return pixelsArray
}

func RenderFire(pixelsArray []int, width, height int) {
    for j := 0; j < height; j++ {
        for k := 0; k < width; k++ {
            fmt.Printf("%s ", firePalette[pixelsArray[j*width + k]])
        }
        fmt.Println()
    }
}
