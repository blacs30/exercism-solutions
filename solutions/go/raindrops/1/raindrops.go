package raindrops

/*
- is divisible by 3, add "Pling" to the result.
- is divisible by 5, add "Plang" to the result.
- is divisible by 7, add "Plong" to the result.
- **is not** divisible by 3, 5, or 7, the result should be the number as a string.
*/
import "fmt"

func Convert(number int) string {
	res := ""
	if number%3 == 0 {
		res = fmt.Sprintf("Pling")
	}
	if number%5 == 0 {
		res += fmt.Sprintf("Plang")
	}
	if number%7 == 0 {
		res += fmt.Sprintf("Plong")
	}
	if res == "" {
		res = fmt.Sprintf("%d", number)
	}
	return res
}
