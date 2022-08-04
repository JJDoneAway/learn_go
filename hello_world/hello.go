package main

import (
	"fmt"
	"strings"
)

const ES = "ES"
const FR = "FR"
const DE = "DE"
const EN_Hello = "Hello"
const ES_Hello = "Hola"
const FR_Hello = "Bonjour"
const DE_Hello = "Hallo"

func main() {
	fmt.Println(Hello("Johannes", ""))
}

func Hello(name string, language string) string {
	name = strings.TrimSpace(name)

	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s %s!", helloWord(language), name)

}

func helloWord(language string) (hello string) {
	language = strings.ToUpper(strings.TrimSpace(language))

	switch language {
	case ES:
		hello = ES_Hello
	case FR:
		hello = FR_Hello
	case DE:
		hello = DE_Hello
	default:
		hello = EN_Hello
	}
	return
}
