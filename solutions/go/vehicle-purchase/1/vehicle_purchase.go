package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return (kind == "car" || kind == "truck")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	const msgTemplate = "%s is clearly the better choice."

	if option1 < option2 {
		return fmt.Sprintf(msgTemplate, option1)
	}

	return fmt.Sprintf(msgTemplate, option2)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	const newPercentage = 0.8
	const useRecentPecentage = 0.7
	const oldPercentage = 0.5

	if age < 3 {
		return originalPrice * newPercentage
	} else if age >= 3 && age < 10 {
		return originalPrice * useRecentPecentage
	}

	// 10 years old or more.
	return originalPrice * oldPercentage
}
