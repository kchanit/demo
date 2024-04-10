package util

import (
	"time"

	"github.com/centraldigital/cfw-core-lib/pkg/constant/datetimeconst"
	"github.com/centraldigital/cfw-core-lib/pkg/model/jsonmodel"
	"github.com/centraldigital/cfw-core-lib/pkg/util/typeconvertutil"
)

const (
	DATETIME_WITHZONE_FORMAT = "2006-01-02T15:04:05-07:00"
)

// warning: this ignore parsing error
func DateStringToDateMinTime(ds string) jsonmodel.DateMinTime {
	t, _ := time.Parse(datetimeconst.DATE_FORMAT, ds)
	return jsonmodel.DateMinTime(t)
}

func DateStringToDateMaxTime(ds string) jsonmodel.DateMaxTime {
	t, _ := time.Parse(datetimeconst.DATE_FORMAT, ds)
	return jsonmodel.DateMaxTime(t)
}

func DatetimeToDatetimeNoTString(t time.Time) string {
	return t.Format(datetimeconst.DATETIME_NO_T)
}

func DatetimeToDateString(t time.Time) string {
	return t.Format(datetimeconst.DATE_FORMAT)
}

func DateStringToDatetime(ds string) time.Time {
	t, _ := time.Parse(datetimeconst.DATE_FORMAT, ds)
	return t
}

func DatetimeStringToDatetime(s string) string {
	return s[:10]
}

// 0 if they are same date,
// -1 if the first date is before
// and 1 if the first date is after
func CompareDateWithDatetime(dateString string, datetime time.Time) (res *int, err error) {
	dMinString := jsonmodel.DateMinTime(datetime).String()
	if dateString == dMinString {
		res = typeconvertutil.ToPtr(0)
		return
	}

	if dateString < dMinString {
		res = typeconvertutil.ToPtr(-1)
		return
	}

	res = typeconvertutil.ToPtr(1)
	return
}

// this roughly calculate by string comparing d1 > d2
func IsDateStringAfter(d1, d2 string) bool {
	return d1 > d2
}

// this roughly calculate by string comparing d1 < d2
func IsDateStringBefore(d1, d2 string) bool {
	return d1 < d2
}

// this roughly calculate by string comparing d1 >= d2
func IsDateStringAfterEqual(d1, d2 string) bool {
	return d1 >= d2
}

// this roughly calculate by string comparing d1 <= d2
func IsDateStringBeforeEqual(d1, d2 string) bool {
	return d1 <= d2
}

// compare as unix
func IsDatetimeAfter(d1, d2 time.Time) bool {
	return d1.Unix() > d2.Unix()
}
