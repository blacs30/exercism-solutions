package chessboard

//import "fmt"

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.

/*
## 1. Given a Chessboard and a File, count how many squares are occupied

Implement the `CountInFile(board Chessboard, file string) int` function.
It should count the total number of occupied squares by ranging over a map. Return an integer.
Return a count of zero (`0`) if the given file cannot be found in the map.

```go
CountInFile(board, "A")
// => 3
```
*/
func CountInFile(cb Chessboard, file string) int {
	count := 0
	for _, val := range cb[file] {
		if val == true {
			count++
		}
	}
	return count
}

/*
## 2. Given a Chessboard and a Rank, count how many squares are occupied

Implement the `CountInRank(board Chessboard, rank int) int` function.
It should count the total number of occupied squares by ranging over the given rank. Return an integer.
Return a count of zero (`0`) if the given rank is not a valid one (not between `1` and `8`, inclusive).

```go
CountInRank(board, 2)
// => 1
```
*/
// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0
	for _, val := range cb {
		if rank > 0 && rank <= 8 && val[rank-1] == true {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for _ = range cb {
		for i := 1; i <= 8; i++ {
			count++
		}
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, val := range cb {
		for i := 0; i < 8; i++ {
			if val[i] == true {
				count++
			}
		}
	}
	return count
}
