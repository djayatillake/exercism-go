package lasagna

// TODO: define the 'OvenTime()' function
func OvenTime() int {
    return 40
}

// TODO: define the 'RemainingOvenTime()' function
func RemainingOvenTime(actual_minutes int) int {
    return OvenTime() - actual_minutes
}

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers int) int {
    return layers * 2
}

// TODO: define the 'ElapsedTime()' function
func ElapsedTime(layers int, actual_minutes int) int {
    return PreparationTime(layers) + actual_minutes
}