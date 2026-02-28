package intermediate

import (
	"fmt"
	"time"
)

func main() {

	// Cirrent local time
	// fmt.Println(time.Now())

	// // Specific time
	// specificTime := time.Date(2026, time.July, 30, 12, 0, 0, 0, time.UTC)
	// fmt.Println("Specific Time:", specificTime)

	// // Parse time
	// parsedTime, _ := time.Parse("2006-01-02", "2026-02-27") // Mon Jan 2 2006 15:04:05 MST 2006
	// parsedTime1, _ := time.Parse("06-01-02", "26-02-27")    // Mon Jan 2 2006 15:04:05 MST 2006
	// // parsedTime2, _ := time.Parse("6-1-2", "26-2-27")        // Mon Jan 2 2006 15:04:05 MST 2006
	// parsedTime2, _ := time.Parse("06-1-2", "26-2-27")             // Mon Jan 2 2006 15:04:05 MST 2006 /// tiene que tener el mismo formato
	// parsedTime3, _ := time.Parse("06-1-2 15-04", "26-2-27 18-47") // Mon Jan 2 2006 15:04:05 MST 2006 /// tiene que tener el mismo formato

	// fmt.Println(parsedTime)
	// fmt.Println(parsedTime1)
	// fmt.Println(parsedTime2)
	// fmt.Println(parsedTime3)

	// // Formatting Time
	t := time.Now()
	// fmt.Println("Formating Time", t.Format("06-01-02 15-15-05"))

	// oneDayLater := t.Add(time.Hour * 24)
	// fmt.Println(oneDayLater.Weekday())

	// fmt.Println("Rounded Time:", t.Round(time.Hour))
	// loc, _ := time.LoadLocation("America/Santiago")
	// t = time.Date(2026, time.February, 27, 10, 05, 40, 00, time.UTC)

	// // Convert this to the specific time zone
	// tLocal := t.In(loc)

	// // Perform Rounding
	// roundendTime := t.Round(time.Hour)
	// roundedTimeLocal := roundendTime.In(loc)

	// fmt.Println("Original Time (UTC)", t)
	// fmt.Println("Original Time (Local)", tLocal)
	// fmt.Println("Rounded Time (UTC)", roundendTime)
	// fmt.Println("Rounded Time (Local)", roundedTimeLocal)

	fmt.Println("Truncated Time", t.Truncate(time.Hour))

	loc, _ := time.LoadLocation("America/New_York")
	tInNY := time.Now().In(loc)
	fmt.Println("Time In New York:", tInNY)

	t1 := time.Date(2026, time.February, 27, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2026, time.February, 27, 18, 0, 0, 0, time.UTC)
	duration := t2.Sub(t1)
	fmt.Println("Duration:", duration)

	// Compare times
	fmt.Println("t2 is after t1?", t2.After(t1))

}
