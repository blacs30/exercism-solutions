package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

/*
```go
celsiusUnit := Celsius
fahrenheitUnit := Fahrenheit

celsiusUnit.String()
// => °C
fahrenheitUnit.String()
// => °F
fmt.Sprint(celsiusUnit)
// Output: °C
```
*/
// Add a String method to the TemperatureUnit type

func (tu TemperatureUnit) String() string {
	unit := []string{"°C", "°F"}
	return unit[tu]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (t Temperature) String() string {
	return fmt.Sprintf("%d %s", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

/*
```go
mphUnit := MilesPerHour
mphUnit.String()
// => mph
fmt.Sprint(mphUnit)
// Output: mph

kmhUnit := KmPerHour
kmhUnit.String()
// => km/h
fmt.Sprint(kmhUnit)
// Output: km/h
```
*/
// Add a String method to SpeedUnit
func (su SpeedUnit) String() string {
	unit := []string{"km/h", "mph"}
	return unit[su]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

/*

```go
windSpeedNow := Speed{
    magnitude: 18,
    unit: KmPerHour,
}
windSpeedNow.String(windSpeedNow)
// => 18 km/h
fmt.Sprintf(windSpeedNow)
// Output: 18 km/h

windSpeedYesterday := Speed{
    magnitude: 22,
    unit: MilesPerHour,
}
windSpeedYesterday.String(windSpeedYesterday)
// => 22 mph
fmt.Sprint(windSpeedYesterday)
// Output: 22 mph
```
*/
// Add a String method to Speed
func (sp Speed) String() string {
	return fmt.Sprintf("%d %s", sp.magnitude, sp.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

/*
sfData.String()
// => San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity
fmt.Sprint(sfData)
// Output: San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity
```
*/
// Add a String method to MeteorologyData

func (md MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity", md.location, md.temperature, md.windDirection, md.windSpeed, md.humidity)
}
