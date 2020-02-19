package utils

import (
	"fmt"
	"testing"
)

/*
go test -v -count=1 ./utils -run TestGetTrueRandIntn
*/
func TestGetTrueRandIntn(t *testing.T) {
	fmt.Printf("the random int is: %+v\n", GetTrueRandIntn(10))
}
