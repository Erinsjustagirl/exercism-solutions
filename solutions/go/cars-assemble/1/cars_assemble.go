package cars

// i had to phone chatgpt for the last function because i was too tired and confused to do it. soggy :c - Erin

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var successPercentage = float64(successRate) / 100;
    return float64(productionRate) * successPercentage;
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60.0);
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var groupsOfTen = carsCount / 10;
	var carsLeftover = carsCount % 10;

	var costOfGroups = groupsOfTen * 95000;
	var costOfLeftover = carsLeftover * 10000;

	return uint(costOfGroups + costOfLeftover);
}
