package main

import "fmt"

/** Mario problem

A program for printing two half-pyramids made of special symbol (for example, #) of a defined height with two spaces
between them. For example, for height 4 it prints this:

   #  #
  ##  ##
 ###  ###
####  ####

The height is defined as a variable in the program source code.
The console output of the program prints the value of the height and half-pyramids.
 */
func main() {

	singlePyramidSize := 4
	pyramidsSeparatorLength := 2
	pictureWidth := singlePyramidSize + pyramidsSeparatorLength + singlePyramidSize
	pictureHeight := singlePyramidSize

	fmt.Println("Pyramid size: ", singlePyramidSize)

	for verticalIndex := 0; verticalIndex < pictureHeight; verticalIndex++ {
		startNewLine()
		for horizontalIndex := 0; horizontalIndex < pictureWidth; horizontalIndex++ {
			printAppropriateSymbol(horizontalIndex, singlePyramidSize, verticalIndex, pyramidsSeparatorLength)
		}
	}
}
func startNewLine() (int, error) {
	return fmt.Println()
}
func printAppropriateSymbol(horizontalIndex int, singlePyramidSize int, verticalIndex int, pyramidsSeparatorLength int) {
	isFirstPyramidArea := horizontalIndex <= singlePyramidSize-1
	isSeparatorArea := horizontalIndex > singlePyramidSize-1 && horizontalIndex <= singlePyramidSize-1+pyramidsSeparatorLength
	isSecondPyramidArea := horizontalIndex > singlePyramidSize-1+pyramidsSeparatorLength

	if isFirstPyramidArea {
		printSymbolForFirstPyramid(horizontalIndex, singlePyramidSize, verticalIndex)
	} else if isSeparatorArea {
		printSeparatorSymbol()
	} else if isSecondPyramidArea {
		printSecondPyramidSymbol(horizontalIndex, pyramidsSeparatorLength, singlePyramidSize, verticalIndex)
	}
}
func printSecondPyramidSymbol(horizontalIndex int, pyramidsSeparatorLength int, singlePyramidSize int, verticalIndex int) {
	indexInSinglePyramid := horizontalIndex - pyramidsSeparatorLength - singlePyramidSize
	idBeforeDiagonal := indexInSinglePyramid <= (verticalIndex)
	if idBeforeDiagonal {
		fmt.Print("#")
	} else {
		fmt.Print(" ")
	}
}
func printSeparatorSymbol() {
	fmt.Print(" ")
}
func printSymbolForFirstPyramid(horizontalIndex int, singlePyramidSize int, verticalIndex int) {
	isBehindDiagonal := horizontalIndex >= (singlePyramidSize - 1 - (verticalIndex))
	if isBehindDiagonal {
		fmt.Print("#")
	} else {
		fmt.Print(" ")
	}
}
