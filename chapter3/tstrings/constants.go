package tstrings

const (
	commaCharSpacing = 3 // put comma after 3 characters
	t100B            = 100
	t1K              = 1024
	t10K             = 10240
)

const (
	bit0 = 1 << iota
	bit1
	bit2
	bit3
	bit4
	bit5
	bit6
	bit7
)

// Declare KB, Kilobyte, Megabyte, Giga, Tera, Peta,
// Untyped constants have 256 bit precision
const (
	_ = 1 << (10 * iota)
	KB
	MB
	GB
	TB // Terra 1 << 32
	PB // Peta
	EB // Exa Byte Google has 15 Exabyte of data
	ZB // Zetta Byte  1 << 64
	YB // Yotta Byte
	BB // Bronto byte - https://whatsabyte.com/
	GO // Geop byte
)

// DayOfWeek Declare days of week
type DayOfWeek int8

const (
	// Sunday first day of week
	Sunday DayOfWeek = iota
	// Monday second day of week
	Monday
	// Tuesday third day of week
	Tuesday
	// Wednesday forth day of week
	Wednesday
	// Thursday fifth day of week
	Thursday
	// Friday day 6
	Friday
	// Saturday last day
	Saturday
)

// MonthOfYear assign a number to a month,  Jan = 0
type MonthOfYear int8

const (
	// Jan month 0
	Jan MonthOfYear = iota
	// Feb month 1
	Feb
	// Mar month 2
	Mar
	// Apr month 4
	Apr
	// May month 5
	May
	// Jun month 6
	Jun
	// July month 7
	July
	// Aug month 8
	Aug
	// Sep month 9
	Sep
	// Oct month 10
	Oct
	// Nov month 11
	Nov
	// Dec merry Christmas
	Dec
)
