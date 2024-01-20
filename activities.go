package main

import "fmt"

func main() {
	activity02()
	activity03()

}

func activity02 () {
/*2. Integer Library Fine

Create a program to calculate library fines. Assume a daily fine of $1 for overdue books. Define two integer variables: overdueDays (e.g., set to 7 days) and dailyFine. Calculate the total fine and display it.
*/
	var overdueDays int = 7
	var dailyFine int = 1

	var totalFine int = overdueDays * dailyFine

	fmt.Printf("The Overdue days are %.d\nThe daily Fine is $%.2f\nTotal fine is $%.2f\n", overdueDays, float64(dailyFine),float64(totalFine))
	fmt.Printf("///////////// END OF EXERCISE 02 /////////////////\n\n")
}

func activity03() {
/*3. Floating-Point Gas Station:

Write a program that calculates the cost of fuel for a road trip. Include variables for the price per gallon of gasoline (e.g., $3.75), the number of gallons needed (e.g., 15.5 gallons), and a discount percentage for loyalty members (e.g., 5%). Calculate the total cost with and without the discount, and display both.
*/

pricePerGallonOfGasoline := 3.75
numberOfGallonsNeeded := 15.5
discountPercentage := 0.05

totalCostWithoutDiscount := pricePerGallonOfGasoline * numberOfGallonsNeeded
totalCostWithDiscount := totalCostWithoutDiscount*(1-discountPercentage)

fmt.Printf("Total cost without discount = $%.2f\nTotal cost with discount = $%.2f\n", totalCostWithoutDiscount, totalCostWithDiscount )
fmt.Printf("///////////// END OF EXERCISE 03 /////////////////\n")

}