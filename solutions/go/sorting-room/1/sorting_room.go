package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
/*
Jen wants numbers to return strings like "This is the number 2.0" (including one digit after the decimal):

fmt.Println(DescribeNumber(-12.345))
// Output: This is the number -12.3
*/
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// // DescribeNumberBox should return a string describing the NumberBox.
/*

Jen wants number boxes to return strings like `"This is a box containing the number 2.0"` (again, including one digit after the decimal):

```go
fmt.Println(DescribeNumberBox(numberBoxContaining{12}))
// Output: This is a box containing the number 12.0
```
*/
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
/*
Jen needs a helper function to extract the number from a `FancyNumberBox`.
If the `FancyNumberBox` is a `FancyNumber`, extract its value and convert it from a `string` to an `int`.
Any other type of `FancyNumberBox` should return 0.

```go
fmt.Println(ExtractFancyNumber(FancyNumber{"10"}))
// Output: 10
fmt.Println(ExtractFancyNumber(AnotherFancyNumber{"4"}))
// Output: 0
```
*/
func ExtractFancyNumber(fnb FancyNumberBox) int {
	if val, ok := fnb.(FancyNumber); ok {
		value, err := strconv.Atoi(val.Value())
		if err != nil {
			return 0
		}
		return value
	}
	return 0
}

// // DescribeFancyNumberBox should return a string describing the FancyNumberBox.
/*

If the `FancyNumberBox` is a `FancyNumber`, Jen wants strings saying `"This is a fancy box containing the number 4.0"`.
Any other type of `FancyNumberBox` should say `"This is a fancy box containing the number 0.0"`.

```go
fmt.Println(DescribeFancyNumberBox(FancyNumber{"10"}))
// Output: This is a fancy box containing the number 10.0
fmt.Println(DescribeFancyNumberBox(AnotherFancyNumber{"4"}))
// Output: This is a fancy box containing the number 0.0
```

NOTE: we should use the `ExtractFancyNumber` function!
*/
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	fn := float64(ExtractFancyNumber(fnb))
	if fn == 0 {
		return fmt.Sprintf("This is a fancy box containing the number %.1f", fn)
	} else {
		return fmt.Sprintf("This is a fancy box containing the number %.1f", fn)

	}
}

// // DescribeAnything should return a string describing whatever it contains.
/*

This is the main function Jen needs which takes any input (the empty interface means any value at all: `interface{}`).
`DescribeAnything` should delegate to the other functions based on the type of the value passed in.
More specifically:

- `int` and `float64` should both delegate to `DescribeNumber`
- `NumberBox` should delegate to `DescribeNumberBox`
- `FancyNumberBox` should delegate to `DescribeFancyNumberBox`
- anything else should result in `"Return to sender"`

```go
fmt.Println(DescribeAnything(numberBoxContaining{12.345}))
// Output: This is a box containing the number 12.3
fmt.Println(DescribeAnything("some string"))
// Output: Return to sender
```
*/
func DescribeAnything(i interface{}) string {
	switch v := i.(type) {
	case int:
		return DescribeNumber(float64(v))
	case float64:
		return DescribeNumber(v)
	case NumberBox:
		return DescribeNumberBox(v)
	case FancyNumberBox:
		return DescribeFancyNumberBox(v)
	default:
		return "Return to sender"
	}
}
