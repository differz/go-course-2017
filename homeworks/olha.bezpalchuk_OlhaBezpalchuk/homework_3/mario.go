package main

import (
    "fmt"
)

func BuildLine(material string, brickLength int, lineLength int) string {
    var piece, spaces string
    
    // use bricks
    for i := 0; i < brickLength; i++ {
        piece += material
    }

    // add spaces to the line
    for i := brickLength; i < lineLength; i++ {
        spaces += " "
    }
 
    return spaces + piece + "  " + piece + spaces
}

func BuildPyramid(material string, height int) {
    fmt.Printf("Height: %d\n", height)
    for i := 1; i <= height; i++ {
        fmt.Println(BuildLine(material, i, height))
    }
}

func main() {
    brick := "#"
    height := 5

    BuildPyramid(brick, height)
}
