package firsttest

import (
	"errors"
	"fmt"
)

func Hello(name, language string) (string, error) {
	if name == "" {
		name = "World"
	}

	prefix := ""

	switch language {
	case "english":
		prefix = "Hello"
	case "spanish":
		prefix = "Hola"
	case "german":
		prefix = "Hallo"
	default:
		return "", errors.New("need to provide supported language")
	}
	return prefix + ", " + name, nil
}

func main() {
	fmt.Println(Hello("John", "english"))
}
