// Package weather provides information about the weather.
package weather

// CurrentCondition is used to set information about the current condition of the weather.
var CurrentCondition string

// CurrentLocation is used to set the current location of where to forecast the weather from.
var CurrentLocation string

// Forecast provides information about the weather of a give city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
