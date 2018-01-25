package model

// declare new struct
type gopher struct {
	name string
	age int
	isAdult bool
}

type horse struct {
	name string
	weight int
}

// g is the reciever of this function
// jump is the name of the function
func (g gopher) Jump() string {
	if g.age < 65 {
		return g.name + " can jump hight\n"
	}
	return g.name + " can still jump\n"
}

func (h horse) Jump() string {
	if h.weight < 2500 {
		return h.name + " is not heavy, can jump hight\n"
	}
	return h.name + " can not still jump\n"
}

// interface provide a way to specific behavior
// convention for name one-method interface in Go, is to use method name plus an -er
type jumper interface {
	Jump() string
}

// in order to be accessed from outside packages, identifiers must be exported by adopting uppercase name convention for the first letter
func GetList() []jumper {
	phil := &gopher{name: "Phil", age: 30}
	validateAge(phil)
	noodles := &gopher{name: "Noodles", age: 90}
	barbaro := &horse{name: "Barbaro", weight: 2000}

	list := []jumper{phil, noodles, barbaro}
	return list
}

// * indicates a pointer to to the original gopher
func validateAge(g *gopher) {
	g.isAdult = g.age >= 21
}
