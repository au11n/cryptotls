package fips140

const (
	indicatorUnset uint8 = iota
	indicatorFalse
	indicatorTrue
)

func ResetServiceIndicator()  {}
func ServiceIndicator() bool  { return false }
func RecordApproved()         {}
func RecordNonApproved()      {}
