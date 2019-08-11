package main

import "fmt"

const greetDefault = "Sup %s! We're back!"
const greetSpanish = "Hola %s! We're back!"
const greetFrench = "Bonjour %s! We're back!"

func Hello(name string, lang string) string {
	greet := greetDefault

	if name == "" {
		name = "Homie"
	}

	switch lang {
	case "French":
		greet = greetFrench

	case "Spanish":
		greet = greetSpanish
	}

	return fmt.Sprintf(greet, name)
}

func main() {
	fmt.Println(Hello("John", ""))
}
