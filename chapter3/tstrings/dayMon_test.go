package tstrings

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestDayName test for mon of week name, to run
//   cd chapter3
//   go test -run TestDayName -v
func TestDayName(t *testing.T) {
	testBasenameOne := []struct {
		day      DayOfWeek
		expected string
	}{
		{day: Sunday, expected: "Sunday"},
		{day: Monday, expected: "Monday"},
		{day: Tuesday, expected: "Tuesday"},
		{day: Wednesday, expected: "Wednesday"},
		{day: Thursday, expected: "Thursday"},
		{day: Friday, expected: "Friday"},
		{day: Saturday, expected: "Saturday"},
		{day: 10, expected: "Invalid"},
		{day: 100, expected: "Invalid"},
		{day: -10, expected: "Invalid"},
		{day: -100, expected: "Invalid"},
	}
	for _, test := range testBasenameOne {
		title := fmt.Sprintf("Day:%d=>%s", test.day, test.expected)
		t.Run(title, func(t *testing.T) {
			assert.Equal(t, test.expected, test.day.DayName())
		})
	}
}

//   go test -run TestWeekend -v
func TestWeekend(t *testing.T) {
	testBasenameOne := []struct {
		day      DayOfWeek
		expected bool
	}{
		{day: Sunday, expected: true},
		{day: Monday, expected: false},
		{day: Tuesday, expected: false},
		{day: Wednesday, expected: false},
		{day: Thursday, expected: false},
		{day: Friday, expected: false},
		{day: Saturday, expected: true},
	}
	for _, test := range testBasenameOne {
		title := fmt.Sprintf("Day:%d=>%t", test.day, test.expected)
		t.Run(title, func(t *testing.T) {
			assert.Equal(t, test.expected, test.day.Weekend())
		})
	}
}

//   go test -run TestWeekday -v
func TestWeekday(t *testing.T) {
	testBasenameOne := []struct {
		day      DayOfWeek
		expected bool
	}{
		{day: Sunday, expected: false},
		{day: Monday, expected: true},
		{day: Tuesday, expected: true},
		{day: Wednesday, expected: true},
		{day: Thursday, expected: true},
		{day: Friday, expected: true},
		{day: Saturday, expected: false},
	}
	for _, test := range testBasenameOne {
		title := fmt.Sprintf("Day:%d=>%t", test.day, test.expected)
		t.Run(title, func(t *testing.T) {
			assert.Equal(t, test.expected, test.day.Weekday())
		})
	}
}

// TestDayName test for mon of week name, to run
//   cd chapter3
//   go test -run TestMonthName -v
func TestMonthName(t *testing.T) {
	testBasenameOne := []struct {
		mon      MonthOfYear
		expected string
	}{
		{mon: Jan, expected: "Jan"}, {mon: 0, expected: "Jan"},
		{mon: Feb, expected: "Feb"}, {mon: 1, expected: "Feb"},
		{mon: Mar, expected: "Mar"}, {mon: 2, expected: "Mar"},
		{mon: Apr, expected: "Apr"}, {mon: 3, expected: "Apr"},
		{mon: May, expected: "May"}, {mon: 4, expected: "May"},
		{mon: Jun, expected: "Jun"}, {mon: 5, expected: "Jun"},
		{mon: July, expected: "July"}, {mon: 6, expected: "July"},
		{mon: Aug, expected: "Aug"}, {mon: 7, expected: "Aug"},
		{mon: Sep, expected: "Sep"}, {mon: 8, expected: "Sep"},
		{mon: Oct, expected: "Oct"}, {mon: 9, expected: "Oct"},
		{mon: Nov, expected: "Nov"}, {mon: 10, expected: "Nov"},
		{mon: Dec, expected: "Dec"}, {mon: 11, expected: "Dec"},
		{mon: 100, expected: "Invalid"},
		{mon: -10, expected: "Invalid"},
		{mon: -100, expected: "Invalid"},
	}
	for _, test := range testBasenameOne {
		title := fmt.Sprintf("Day:%d=>%s", test.mon, test.expected)
		t.Run(title, func(t *testing.T) {
			assert.Equal(t, test.expected, test.mon.MonthName())
		})
	}
}
