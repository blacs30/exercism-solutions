// Package weather reports the weather forecast
// for the current location.
package weather

// CurrentCondition represents the current weather conditions in form of a string.
var CurrentCondition string

// CurrentLocation represents the location for the weather forecast in form of a string.
var CurrentLocation string

// Forecast takes a city and conditation string
// and returns a description of the weather as string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
