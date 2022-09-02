package main

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world!!!"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix + name
}

// func main() {
// 	fmt.Println(Hello("Chris", "Spanish"))
// }
