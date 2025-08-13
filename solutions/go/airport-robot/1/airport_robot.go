package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
/*## 1. Create the abstract greeting functionality

You will not write the code for the different languages yourself so you need to structure your code for the robot so that other developers can easily add more languages later.

As a first step, define an interface `Greeter` with two methods.

- `LanguageName` which returns the name of the language (a `string`) that the robot is supposed to greet the visitor in.
- `Greet` which accepts a visitor's name (a `string`) and returns a `string` with the greeting message in a specific language.

Next, implement a function `SayHello` that accepts the name of the visitor and anything that implements the `Greeter` interface as arguments and returns the desired greeting string.
For example, imagine a German `Greeter` implementation for which `LanguageName` returns `"German"` and `Greet` returns `"Hallo {name}!"`:

```go
SayHello("Dietrich", germanGreeter)
// => "I can speak German: Hallo Dietrich!"
```
*/
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(visitor string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s!", greeter.LanguageName(), greeter.Greet(visitor))
}

type Italian struct {
	name string
}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s", name)
}

type Portuguese struct {
	name string
}

func (i Portuguese) LanguageName() string {
	return "Portuguese"
}

func (i Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s", name)
}
