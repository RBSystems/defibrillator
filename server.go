package main

import "github.com/byuoitav/defibrillator/packages/pulse"

func main() {
	err := pulse.Check()
	if err != nil {
		panic(err)
	}
}
