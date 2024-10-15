package main

import "fmt"

func main() {
	fmt.Println(squareIsWhite("a2"))
}

func squareIsWhite(coordinates string) bool {

	return (coordinates[0]+coordinates[1])%2 == 1

}
