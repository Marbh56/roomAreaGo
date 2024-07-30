package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	converstionFactor := 0.09290304
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What is the length of the room in feet? ")
	length, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print("Couldn't get length")
		return
	}
	length = strings.TrimSpace(length)
	lengthInt, err := strconv.Atoi(length)
	if err != nil {
		fmt.Print("Couldn't convert length")
		return
	}

	fmt.Print("What is the width of the room in feet? ")
	width, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print("Couldn't get width")
		return
	}
	width = strings.TrimSpace(width)
	widthInt, err := strconv.Atoi(width)
	if err != nil {
		fmt.Print("Couldn't convert width")
		return
	}

	areaImperial := lengthInt * widthInt
	areaMetric := float64(areaImperial) * converstionFactor

	fmt.Printf("You entered dimensions of %v feet by %v feet.", length, width)
	fmt.Printf("The area is\n")
	fmt.Printf("%v square feet\n", areaImperial)
	fmt.Printf("%.3f square meters\n", areaMetric)
}
