package main

import "fmt"

const (
	englishHelloPrefix  = "Hello, "
	spanishHelloPrefix  = "Hola, "
	frenchHelloPrefix   = "Bonjour, "
	arabicHelloPrefix   = "مرحبا, "
	filipinoHelloPrefix = "Kamusta, "
)

func Hello(name, language string) string {
	switch language {
	case "Spanish":
		if name == "" {
			name = "Mundo"
		}
		return spanishHelloPrefix + name + "!"
	case "French":
		if name == "" {
			name = "Monde"
		}
		return frenchHelloPrefix + name + "!"
	case "Arabic":
		if name == "" {
			name = "العالم"
		}
		return arabicHelloPrefix + name + "!"
	case "Filipino":
		if name == "" {
			name = "Mundo"
		}
		return filipinoHelloPrefix + name + "!"
	default:
		if name == "" {
			name = "World"
		}
		return englishHelloPrefix + name + "!"
	}
}

func main() {
	fmt.Println(Hello("World", "English"))
}
