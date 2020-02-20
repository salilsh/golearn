package main

import (
	"fmt"
)

type vehicle struct {
	doors string
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	civic := sedan{
		vehicle: vehicle{
			doors: "wide",
			color: "white",
		},
		luxury: true,
	}

	hummer := truck{
		vehicle: vehicle{
			doors: "broad",
			color: "camo",
		},
		fourWheel: true,
	}

	fmt.Println(civic.doors, civic.color, civic.luxury)
	fmt.Println(hummer.doors, hummer.color, hummer.fourWheel)

}
