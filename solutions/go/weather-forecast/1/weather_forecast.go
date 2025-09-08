// Package weather contains functionality for evaluating weather conditions.
package weather

// CurrentCondition represents the current weather conditions.
var CurrentCondition string

// CurrentLocation represents the current location of weather reporting.
var CurrentLocation string

// Forecast sets and returns a rundown of the weather condition for the given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
