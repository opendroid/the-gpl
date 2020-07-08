package tstrings

// DayOfWeekName supports these methods
type DayOfWeekName interface {
	DayName() string
	Weekend() bool
	Weekday() bool
}

// DayName returns name of day to a corresponding day number
func (d DayOfWeek) DayName() string {
	day := [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	if d >= Sunday && d <= Saturday {
		return day[d]
	}
	return "Invalid"
}

// Weekend returns true if a day falls on weekend
func (d DayOfWeek) Weekend() bool {
	switch d {
	case Saturday, Sunday:
		return true
	default:
		return false
	}
}

// Weekday returns true if a day falls on M-F
func (d DayOfWeek) Weekday() bool {
	return !d.Weekend()
}

// MonthOfYearName supports name of month
type MonthOfYearName interface {
	MonthName() string
}

// MonthName tells name of the month corresponding to a month number
func (m MonthOfYear) MonthName() string {
	mon := [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun",
		"July", "Aug", "Sep", "Oct", "Nov", "Dec"}
	if m >= Jan && m <= Dec {
		return mon[m]
	}
	return "Invalid"
}
