// Package xsdt
// Do not Edit. This stuff it's been automatically generated.
// Generated at 2022-04-05 07:17:08.59722 +0200 CEST m=+0.095667126
package xsdt

import (
	"fmt"
	"time"
)

func (me Boolean) IsValid(optional bool) bool {
	return true
}

func (me Boolean) IsZero() bool {
	return true
}

func (me AnyType) IsValid(optional bool) bool {
	return true
}

func (me AnyType) IsZero() bool {
	return true
}

func (me Decimal) IsValid(optional bool) bool {
	return true
}

func (me Decimal) IsZero() bool {
	return true
}

func ToDecimal(f interface{}) Decimal {
	return Decimal(fmt.Sprintf("%v", f))
}

func (me String) IsValid(optional bool) bool {
	return true
}

func (me String) IsZero() bool {
	return true
}

func (me Base64Binary) IsValid(optional bool) bool {
	return true
}

func (me Base64Binary) IsZero() bool {
	return true
}

func (me Date) IsValid(optional bool) bool {
	return true
}

func (me Date) IsZero() bool {
	return me == "" || me == "0001-01-01"
}

func ToDateFromTime(tm time.Time) Date {
	return Date(tm.Format("2006-01-02"))
}

func (me DateTime) IsValid(optional bool) bool {

	if me.IsZero() {
		if optional {
			return true
		}
		return false
	}

	return true
}

func (me DateTime) IsZero() bool {
	return me == "" || me == "0001-01-01T00:00:00Z"
}

func ToDateTimeFromTime(tm time.Time) DateTime {
	return DateTime(tm.Format(time.RFC3339))
}
