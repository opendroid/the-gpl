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
	Sunday    DayOfWeek = iota // Sunday first day of week
	Monday                     // Monday second day of week
	Tuesday                    // Tuesday third day of week
	Wednesday                  // Wednesday forth day of week
	Thursday                   // Thursday fifth day of week
	Friday                     // Friday is day 6
	Saturday                   // Saturday last day
)

// MonthOfYear assign a number to a month,  Jan = 0
type MonthOfYear int8

const (
	Jan  MonthOfYear = iota // Jan month 0
	Feb                     // Feb month 1
	Mar                     // Mar month 2
	Apr                     // Apr month 4
	May                     // May month 5
	Jun                     // Jun month 6
	July                    // July month 7
	Aug                     // Aug month 8
	Sep                     // Sep month 9
	Oct                     // Oct month 10
	Nov                     // Nov month 11
	Dec                     // Dec merry Christmas
)
