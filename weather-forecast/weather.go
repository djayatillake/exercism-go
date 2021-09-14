//Package COMMENT
package weather

var (
//these are two string variables
	CurrentCondition string
//this is the second string variable
	CurrentLocation  string
)

//Log returns the current condition for a location
func Log(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
