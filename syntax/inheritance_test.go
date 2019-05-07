package syntax

import "testing"

func TestPrintPerson(t *testing.T) {
	var (
		s Student
		d Doctor
	)
	PrintPerson(s)
	PrintPerson(d)
}
