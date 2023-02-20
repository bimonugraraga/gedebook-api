package utils

func CalculateOffset(limit int, page int) int {
	pageToCalculate := page - 1
	return pageToCalculate * limit
}
