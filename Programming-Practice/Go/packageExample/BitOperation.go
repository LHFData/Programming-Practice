package packageExample

import "fmt"

func BitAND(a uint, b uint) uint {
	return a & b
}
func BitOR(a uint, b uint) uint {
	return a | b
}
func BitXOR(a uint, b uint) uint {
	return a ^ b
}
func BitLeftMove(a uint, count int) uint {
	return a << count
}
func BitRightMove(a uint, count int) uint {
	return a >> count
}
func BitAndNot(a uint, b uint) uint {
	return a &^ b
}
func BitOutput(a uint) {
	fmt.Printf("%08b\n", a)
}
