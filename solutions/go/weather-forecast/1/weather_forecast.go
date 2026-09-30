// Package weather does the thing and stuff.
package weather

var (
    // CurrentCondition does what it says on the tin.
	CurrentCondition string
    // CurrentLocation, likewise, does what it says on the tin.
	CurrentLocation  string
)

// Forecast does the stuff and whatnot.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
