// Declare package main
package main

//mporte function
import "fmt"

// Declaração do const da temperatura da água em Farenheint.
const temperatureF = 212.0

// Main function
func main() {

	tempF := temperatureF         //Variavel da temperatura de ebulição em Farenheint
	tempC := (tempF - 32) * 5 / 9 //Variavel da temperatura de ebulição em Celsius
	tempF = 2

	fmt.Println("A temeperatura de ebulição da água em °F", tempF)
	fmt.Println("A temeperatura de ebulição da água em °C", tempC)

}
