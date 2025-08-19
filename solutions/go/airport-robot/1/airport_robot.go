package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
import "fmt"
type Greeter interface {
    LanguageName() string
    Greet(visitorName string) string
}

func SayHello(visitorName string, greeter Greeter) string {
    return greeter.Greet(visitorName)
}

type Italian struct {}

func (i Italian) LanguageName() string {
    return "Italian"
}

func (i Italian) Greet(visitorName string) string {
    return fmt.Sprintf("I can speak Italian: Ciao %s!", visitorName)
}

type Portuguese struct {}

func (p Portuguese) LanguageName() string {
    return "Portuguese"
}

func (p Portuguese) Greet(visitorName string) string {
    return fmt.Sprintf("I can speak Portuguese: Olá %s!", visitorName)
}