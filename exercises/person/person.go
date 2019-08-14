package person

// Mamal ...
type Mamal interface {
	Gender() string
}

// Animal ...
type Animal struct {
	sex string
}

// Person ...
type Person struct {
	sex string
}

// Gender ...
func (p Person) Gender() string {
	return p.sex
}

// Gender ...
func (p Animal) Gender() string {
	return p.sex
}
