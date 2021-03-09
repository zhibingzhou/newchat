package global

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

type GVA_MODEL struct {
	ID        int        `json:"id" gorm:"primarykey"`
	CreatedAt time.Time  `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"comment:更新时间"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"comment:删除时间"`
}

type TimeStamp time.Time

func (ts TimeStamp) MarshalJSON() ([]byte, error) {
	origin := time.Time(ts)
	return []byte(strconv.FormatInt(origin.UnixNano()/1000000, 10)), nil
}

func (ts *TimeStamp) ToTime() time.Time {
	return time.Time(*ts)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format.
func (ts *TimeStamp) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	millis, err := strconv.ParseInt(string(data), 10, 64)

	*ts = TimeStamp(time.Unix(0, millis*int64(time.Millisecond)))
	return err
}

func (ts TimeStamp) ToString() string {
	return ts.ToTime().Format("2006-01-02 15:04:05")
}

func (ts TimeStamp) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(ts)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

// Scan valueof time.Time
func (ts *TimeStamp) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*ts = TimeStamp(value)
		return nil
	}
	//i, err = strconv.ParseInt(sc, 10, 64)

	return fmt.Errorf("can not convert %v to timestamp", v)
}
