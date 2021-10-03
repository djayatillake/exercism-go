package lasagna

// PreparationTime calculates the time to prepare the lasagna based on layers and avg prep time
func PreparationTime(layers []string, avg_time_per_layer int) int {
	if avg_time_per_layer <= 0 {
		return len(layers) * 2
	}
	return len(layers) * avg_time_per_layer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, v := range layers {
		switch v {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return
}

// AddSecretIngredient adds secret ingredient from friends list to yours
func AddSecretIngredient(friendsList, myList []string) []string {
	return append(myList, friendsList[len(friendsList)-1])
}

// ScaleRecipe scales the amounts needed for two portions to the number wanted
func ScaleRecipe(amounts []float64, portions int) []float64 {
	var scaler float64
	scaler = float64(portions) / 2
	newamounts := []float64{}
	for _, v := range amounts {
		newamounts = append(newamounts, v*scaler)
	}
	return newamounts
}
