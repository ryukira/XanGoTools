package conv

import (
	"log"
	"testing"
)

func TestStringNumberToNumber(t *testing.T) {
	a, b := StringNumberToNumber("0000050000")
	log.Println(a, b)
	log.Println("asdas")
}
