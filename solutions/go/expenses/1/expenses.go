package expenses

import (
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

/*
Implement the generic `Filter` function to filter records according to a criteria given by a function.
This filter function accepts a collection of records and a predicate function and returns only the records in the collection that satisfy the predicate.

```go
records := []Record{
  {Day: 1, Amount: 15, Category: "groceries"},
  {Day: 11, Amount: 300, Category: "utility-bills"},
  {Day: 12, Amount: 28, Category: "groceries"},
}

// Day1Records only returns true for records that are from day 1
func Day1Records(r Record) bool {
  return r.Day == 1
}

Filter(records, Day1Records)
// =>
// [
//   {Day: 1, Amount: 15, Category: "groceries"}
// ]
```
*/
// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var newRecords []Record
	for _, entry := range in {
		if predicate != nil && predicate(entry) {
			newRecords = append(newRecords, entry)
		}
	}
	return newRecords
}

/*
Implement the `ByDaysPeriod` function that will help Bob create such filters.
This function accepts a `DaysPeriod` and returns function that takes a record and tells whether the record is in the period of time specified by the `DaysPeriod` given as argument.

```go
records := []Record{
  {Day: 1, Amount: 15, Category: "groceries"},
  {Day: 11, Amount: 300, Category: "utility-bills"},
  {Day: 12, Amount: 28, Category: "groceries"},
  {Day: 26, Amount: 300, Category: "university"},
  {Day: 28, Amount: 1300, Category: "rent"},
}

period := DaysPeriod{From: 1, To: 15}

Filter(records, ByDaysPeriod(period))
// =>
// [
//   {Day: 1, Amount: 15, Category: "groceries"},
//   {Day: 11, Amount: 300, Category: "utility-bills"},
//   {Day: 12, Amount: 28, Category: "groceries"},
// ]
```
*/

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	start, end := p.From, p.To

	return func(in Record) bool {
		if in.Day >= start && in.Day <= end {
			return true
		} else {
			return false
		}
	}
}

/*
Implement the `ByCategory` function that will help Bob create such filters.
This function accepts a category and returns a function that takes a record and tells whether the category of this record is the same as the category given as the argument.

```go
records := []Record{
  {Day: 1, Amount: 15, Category: "groceries"},
  {Day: 11, Amount: 300, Category: "utility-bills"},
  {Day: 12, Amount: 28, Category: "groceries"},
  {Day: 28, Amount: 1300, Category: "rent"},
}

Filter(records, ByCategory("groceries"))
// =>
// [
//   {Day: 1, Amount: 15, Category: "groceries"},
//   {Day: 12, Amount: 28, Category: "groceries"},
// ]
```
*/
// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	category := c
	return func(in Record) bool {
		if in.Category == category {
			return true
		} else {
			return false
		}
	}
}

/*

Implement the `TotalByPeriod` function to return a sum of expenses in the days period.

```go
records := []Record{
  {Day: 15, Amount: 16, Category: "entertainment"},
  {Day: 32, Amount: 20, Category: "groceries"},
  {Day: 40, Amount: 30, Category: "entertainment"}
}

p1 := DaysPeriod{From: 1, To: 30}
p2 := DaysPeriod{From: 31, To: 60}

TotalByPeriod(records, p1)
// => 16

TotalByPeriod(records, p2)
// => 50
*/
// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	records := Filter(in, ByDaysPeriod(p))
	var sum float64
	for _, record := range records {
		sum += float64(record.Amount)
	}
	return sum
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	records := Filter(in, ByCategory(c))
	if len(records) < 1 {
		return 0, fmt.Errorf("unknown category")
	}
	return TotalByPeriod(records, p), nil
}
