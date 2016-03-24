package main

import "fmt"
import "github.com/SirWojtek/GoShips/objects"

func main() {
	fmt.Println("Hello World!")
	gameState := objects.NewGameState()
	gameState.Paint()
}
