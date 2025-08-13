package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
/*
Implement a function `TotalBirdCount` that accepts a slice of `int`s that contains the bird count per day.
It should return the total number of birds that you counted.
*/
func TotalBirdCount(birdsPerDay []int) int {
	totalBirds := 0
	for i := 0; i < len(birdsPerDay); i++ {
		totalBirds += birdsPerDay[i]
	}
	return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
/*
Implement a function `BirdsInWeek` that accepts a slice of bird counts per day and a week number.

It returns the total number of birds that you counted in that specific week.
You can assume weeks are always tracked completely.
*/
func BirdsInWeek(birdsPerDay []int, week int) int {
	totalBirdsInWeek := 0
	startDay := week*7 - 7
	for i := startDay; i < startDay+7; i++ {
		totalBirdsInWeek += birdsPerDay[i]
	}
	return totalBirdsInWeek
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
/*

You figured out that this bird always spent every second day in your garden.

You do not know exactly where it was in between those days but definitely not in your garden.

Your bird watcher intuition also tells you that the bird was in your garden on the first day that you tracked in your list.

Given this new information, write a function `FixBirdCountLog` that takes a slice of birds counted per day as an argument and returns the slice after correcting the counting mistake.
*/
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] = birdsPerDay[i] + 1
	}
	return birdsPerDay
}
