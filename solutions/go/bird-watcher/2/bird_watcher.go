package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
	for i := 0; i < len(birdsPerDay); i++ {
		sum += birdsPerDay[i]
	}
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	sum := 0
	startDay := (week - 1) * 7
	if startDay > len(birdsPerDay) {
		return 0
	}
	endDay := startDay + 7
	if endDay > len(birdsPerDay) {
		endDay = len(birdsPerDay) - 1
	}
	for i := startDay; i < endDay; i++ {
		sum += birdsPerDay[i]
	}
	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	var value int
	for i := 0; i < len(birdsPerDay); i++ {
		if i == 0 || i%2 == 0 {
			value = birdsPerDay[i] + 1
		} else {
			value = birdsPerDay[i]
		}
		birdsPerDay = append(birdsPerDay, value)

	}
	return birdsPerDay

}
