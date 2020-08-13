package main

import "fmt"

type bot interface {
	getGreeting() string
	//getBotVersion() float64
	//respondToUser(user) string
}

type englishBot struct {}
type spanishBot struct {}

func main()  {
	eb := englishBot{}
	sb := spanishBot{}

	//printGreetingEnglish(eb)
	//printGreetingSpanish(sb)

	printGreeting(eb)
	printGreeting(sb)
}

//func printGreetingEnglish(eb englishBot){
//	fmt.Println(eb.getGreeting())
//}
//
//func printGreetingSpanish(sb spanishBot){
//	fmt.Println(sb.getGreeting())
//}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}


