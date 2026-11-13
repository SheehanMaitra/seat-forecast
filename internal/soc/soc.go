// Package soc fetches and parses Rutgers Schedule of Classes data.
package soc

// Section is one registrable section of a course.
type Section struct {
	Index  string // the 5-digit registration index
	Course string // course title or code
	Open   bool   // true when seats are available
}
