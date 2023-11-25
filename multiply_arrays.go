package main

import (
	"fmt"
	"math"
)

func main() {
    number1 := [4]int{1, 0, 0, 0}
    number2 := [4]int{0, 0, 0, 1}

    numberAsInt := arrayToNumber(number1)

    fmt.Printf("\nNumber as int: %d\n", numberAsInt)
}

func arrayToNumber(number [4]int) int {
    numberAsInt := 0;

    for i := 0; i < len(number); i++ {
        a := number[i]
        place := len(number) - i - 1

        pow := int(math.Pow(10, float64(place)))
        numberAsInt += a * pow

        // fmt.Printf("a: %d, i: %d, numberAsInt: %d", a, i, numberAsInt)
        // fmt.Printf(" pow: %d\n", pow)
    }

    return numberAsInt
}
