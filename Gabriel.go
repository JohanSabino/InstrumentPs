package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

type uribita struct {
	persona
	lpm bool
}

func main() {
	p1 := persona{
		Nombre:   "Gabriel",
		Apellido: "Sabino",
		Edad:     19,
	}
	p2 := persona{
		Nombre:   "Alejandra",
		Apellido: "Daza",
		Edad:     26,
	}
	p3 := persona{
		Nombre:   "Andrés",
		Apellido: "Sabino",
		Edad:     28,
	}

	personas := []persona{p1,
		p2,
		p3,
	}
	fmt.Println(personas)

	bs, err := json.Marshal(personas)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)
	fmt.Println(string(bs))
}
