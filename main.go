package main

import "fmt"

func main() {
	chip()
}

func chip() {
	robot := `        ▛▀▀▀▀▀▀▀▀▀▀▀▜
        ▌  ◉     ◉  ▐
        ▌     ◡     ▐
        ▙▄▄▄▄▄▄▄▄▄▄▄▟
`
	fmt.Print(robot)
}
