package main

import "fmt"

func main() {
	celsius := 36.6
	fahrenheit := celsius*9/5 + 32

	fmt.Printf("%.1f°C = %.1f°F\n", celsius, fahrenheit)
}
