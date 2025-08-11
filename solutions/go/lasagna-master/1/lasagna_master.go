package lasagna

func PreparationTime(layers []string, averagePrepTime int) int {
	if averagePrepTime == 0 {
		averagePrepTime = 2
	}
	return len(layers) * averagePrepTime
}

func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "sauce":
			sauce += 0.2
		case "noodles":
			noodles += 50
		default:

		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsRecipe []string, myRecipe []string) {
	myRecipe[len(myRecipe)-1] = friendsRecipe[len(friendsRecipe)-1]
}

func ScaleRecipe(amountForTwoPortions []float64, portions int) []float64 {
	amountForRequiredPortions := make([]float64, len(amountForTwoPortions))
	copy(amountForRequiredPortions, amountForTwoPortions)
	for i := 0; i < len(amountForRequiredPortions); i++ {
		amountForRequiredPortions[i] = amountForRequiredPortions[i] * float64(portions) / 2
	}

	return amountForRequiredPortions

}
