package interfaces

import "fmt"

type Walker interface {
	Walking()
	Running()
}

type Bio interface {
	Bio()
}

type Flyer interface {
	Flying()
}

type Animal struct {
	Breed    string
	Sound    string
	Nature string
	CanFly   bool
	CanSwim  bool
	CanWalk  bool
}

type Dog struct {
	Animal
	Name string
	Legs int
}

type Cat struct {
	Animal
	Name string
	Legs int
}

type Eagle struct {
	Animal
	Wings int
	Legs  int
}

/*
	Declaring functions
*/

func (d Dog) Walking() {
	fmt.Printf("Walking...")
}

func (c Cat) Walking() {
	fmt.Printf("Walking...")
}

func (d Dog) Running() {
	fmt.Printf("Running on %v legs...", d.Legs)
}

func (c Cat) Running() {
	fmt.Printf("Running on %v legs...", c.Legs)
}

func (e Eagle) Flying() {
	fmt.Printf("Can fly with %v wings.", e.Wings)
}

func (e Eagle) Walking() {
	fmt.Printf("Can walk on their %v legs...", e.Legs)
}

func (d Dog) Bio() {
	fmt.Println("Name:", d.Name)
	fmt.Println("Breed:", d.Breed)
	fmt.Println("Sound:", d.Sound)
}
