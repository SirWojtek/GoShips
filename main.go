package main

import "fmt"
import "./objects"

func main() {
	fmt.Println("Hello World!")
	ship := objects.NewShip(objects.Rect{10, 10, 10, 10})
	ship.GetDamage(10)
}
