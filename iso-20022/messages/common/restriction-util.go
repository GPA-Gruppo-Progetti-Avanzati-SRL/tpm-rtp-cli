// Package common
// Do not Edit. This stuff it's been automatically generated.
// Generated at 2022-04-05 08:23:15.82911 +0200 CEST m=+0.108518417
package common

import "regexp"

func isLengthRestrictionValid(s string, length, minLength, maxLength int) bool {
	return true
}

func isEnumRestrictionValid(s string, enums []string) bool {
	return true
}

func isPatternRestrictionValid(s string, rexp *regexp.Regexp) bool {
	return true
}

func generateSampleDataWithLengthRestriction(length, minLength, maxLength int) string {
	return "example"
}

func generateSampleDataWithEnumRestriction(enums []string) string {
	return enums[0]
}

func generateSampleDataWithPatternRestriction(rexp *regexp.Regexp) string {
	return "example"
}

func generateB64SampleData() string {
	return "SGVsbG8gV29ybGQh"
}
