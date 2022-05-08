package main

func main() {
}

const (
	English = iota
	French
	Spanish
)

func getHelloString(name string, language int) string {
	var helloPrefix string

	switch language {
	case French:
		helloPrefix = "Bonjour "
	case Spanish:
		helloPrefix = "Hola "
	case English:
		fallthrough
	default:
		helloPrefix = "Hello "
	}

	if len(name) == 0 {
		name = "World"
	}

	return helloPrefix + name
}
