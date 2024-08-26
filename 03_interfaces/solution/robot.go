package main

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

type Language struct {
	name     string
	greeting string
}

func NewLanguage(name, greeting string) *Language {
	return &Language{name: name, greeting: greeting}
}

func (l Language) LanguageName() string {
	return l.name
}

func (l Language) Greet(name string) string {
	return fmt.Sprintf("%v %v!", l.greeting, name)
}

func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %v: %v", g.LanguageName(), g.Greet(name))
}

func main() {
	italian := NewLanguage("Italian", "Ciao")
	french := NewLanguage("French", "Bonjour")
	fmt.Println(SayHello("Bob", italian))
	fmt.Println(SayHello("Bobette", french))
}
