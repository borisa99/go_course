package main

type bot interface {
	getGreeding() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeding(eb)
	printGreeding(sb)
}

func printGreeding(b bot) {
	println(b.getGreeding())
}
func (englishBot) getGreeding() string {
	return "Hello!"
}
func (spanishBot) getGreeding() string {
	return "Hola!"
}
