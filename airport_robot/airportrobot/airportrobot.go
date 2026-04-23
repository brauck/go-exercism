package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

type Italian struct {
	Lang     string
	Greeting string
}

func (i Italian) LanguageName() string {
	return i.Lang
}

/*
	 func (i Italian) Greet(name string) string {
		greet := fmt.Sprintf("I can speak %s: %s %s!", i.Lang, i.Greeting, name)
		return greet
	}
*/
func (i Italian) Greet(name string) string {
	greet := fmt.Sprintf("I can speak Italian: Ciao %s!", name)
	return greet
}

type Portuguese struct {
	Lang     string
	Greeting string
}

func (p Portuguese) LanguageName() string {
	return p.Lang
}

/*
	 func (p Portuguese) Greet(name string) string {
		greet := fmt.Sprintf("I can speak %s: %s %s!", p.Lang, p.Greeting, name)
		return greet
	}
*/
func (p Portuguese) Greet(name string) string {
	greet := fmt.Sprintf("I can speak Portuguese: Olá %s!", name)
	return greet
}

func SayHello(name string, greeter Greeter) string {
	return greeter.Greet(name)
}
