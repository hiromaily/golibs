package times

import (
	"fmt"
	"time"
)

type Timer struct {
	Start time.Time
	End   time.Time
}

//format
const (
	// 日付のフォーマット
	FORMAT_A string = "1/2"
	FORMAT_B string = "【1/2】"
	FORMAT_C string = "【1月2日】"

	// 日付のフォーマット(+曜日)
	FORMAT_A_WEEK string = "1/2(%s)"
	FORMAT_B_WEEK string = "【1/2(%s)】"
	FORMAT_C_WEEK string = "【1月2日(%s)】"
)

// day of week
var JAPANESE_WEEKDAYS = []string{"日", "月", "火", "水", "木", "金", "土"}

//Formatter Date
func GetFormatDate(strDate string, addWeek bool) string{
	//2016-05-13 16:52:49
	//To time object
	t, _ := time.Parse("2006-01-02 15:04:05", strDate)

	//t.Month()
	//t.Day()

	//Format
	var baseFormat string
	if addWeek{
		baseFormat = fmt.Sprintf(FORMAT_A_WEEK, JAPANESE_WEEKDAYS[t.Weekday()])
	}else{
		baseFormat = FORMAT_A
	}

	return t.Format(baseFormat)
}

//Formatter Time
func GetFormatTime(strTime string) string {
	t, _ := time.Parse("2006-01-02 15:04:05", strTime)

	//t.Hour()
	//t.Minute()
	//t.Second()

	return t.Format("15:04")
}

//Timer
func NewTimer() *time.Time{
	return &Timer{Start:time.Now(), End:nil}
}

func (self *Timer) EndTime(){
	self.End = time.Now()
	//elapsed time
	fmt.Println(self.End.Sub(self.Start))
}
