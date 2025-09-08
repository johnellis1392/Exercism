package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100.0)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(productionRate) * (successRate / 100.0) / 60.0)
}

const SINGLE_CAR_COST uint = 10_000
const BATCH_CAR_COST uint = 95_000

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var batch_total uint = BATCH_CAR_COST * (uint(carsCount) / 10)
    var incidental_total uint = SINGLE_CAR_COST * (uint(carsCount) % 10)
    return batch_total + incidental_total
}
