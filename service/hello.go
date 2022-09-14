package service

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

const (
	english = "English"
	spanish = "Spanish"
	french  = "French"
)

func Hello(name string, language string) string {
	prefix := getHelloPrefixByLanguage(language)
	return prefix + name
}

func getHelloPrefixByLanguage(language string) (prefix string) {
	switch language {
	case english:
		prefix = englishHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
