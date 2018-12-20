package main

// This works
// import "github.com/chrismwendt/Foo"

// This doesn't: main.go:7:8: code in directory /Users/chrismwendt/go/src/github.com/chrismwendt/foo expects import "github.com/chrismwendt/Foo"
import "github.com/chrismwendt/foo"

var foo Foo.FooType

func main() {
}
