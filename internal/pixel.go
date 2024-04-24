package internal

import (
    "math/rand"
)

const MAX_PIXEL_VALUE = 9

func UpdatePixel(pixelsArray []int, j int, width int, height int) {
    decay := rand.Intn(2)

    if width*height > j + width {
        pixelsArray[j] = (pixelsArray[j + width]) - decay

        if pixelsArray[j] < 0 {
            pixelsArray[j] = 0
        }
    } else {
        pixelsArray[j] = MAX_PIXEL_VALUE
    }
}
