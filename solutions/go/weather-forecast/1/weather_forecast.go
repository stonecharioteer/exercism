/* Package weather: program that can forecast the current weather condition of various cities in Goblinocus.*/
package weather

// CurrentCondition: Current weather condition.
var CurrentCondition string

// CurrentLocation: Current location.
var CurrentLocation string

/* Forecast the weather using the current location and condition.
 */
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
