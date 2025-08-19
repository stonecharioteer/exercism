package thefarm
import (
    "errors"
    "fmt"
    )


func DivideFood(fodderCalculator FodderCalculator, numberOfCows int) (float64, error) {
    // return the amount of food per cow
    fodderAmount, errorFodder:= fodderCalculator.FodderAmount(numberOfCows)
    if errorFodder != nil {
        return 0, errorFodder
    }
    fatteningFactor, errorFatteningFactor := fodderCalculator.FatteningFactor()
    if errorFatteningFactor != nil {
        return 0, errorFatteningFactor
    }
    return fodderAmount*fatteningFactor/float64(numberOfCows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, numberOfCows int) (float64, error) {
    if numberOfCows > 0 {
        return DivideFood(fodderCalculator, numberOfCows)
    }
    return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
    NumberOfCows int
    Message string
}

func (e *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.NumberOfCows, e.Message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numberOfCows int) error {
    if numberOfCows < 0 {
        return &InvalidCowsError {
            NumberOfCows: numberOfCows,
            Message: "there are no negative cows",
        }
    } else if numberOfCows == 0 {
        return &InvalidCowsError {
            NumberOfCows: numberOfCows,
            Message: "no cows don't need food",
        }
    }
    return nil
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
