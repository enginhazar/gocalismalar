package pakets

import "fmt"

func Yaz(bakiye *float64) {
	*bakiye -= 20
	fmt.Printf("Bakiye :%.2f\n", *bakiye)
}
