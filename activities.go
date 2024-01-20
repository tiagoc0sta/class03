package main

import (
	"fmt"
	"strings"
)

func main() {
	activity02()
	activity03()
	activity04()
	activity05()
	activity06()
	activity07()
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
	fmt.Printf("///////////// END OF EXERCISE 03 /////////////////\n\n")
}

func activity04() {
/*4. Complex Number Electrical Engineering:

In an electrical engineering context, create a program where you define two complex impedance values (e.g., Z1 = 3 + 4i ohms and Z2 = 5 - 2i ohms). Calculate and display the addition of these two impedances.
*/
	var Z1 complex64 = 3 + 4i
	var Z2 complex64 = 5 - 2i

	addition := Z1 + Z2
	totalOhms := real(addition)

	fmt.Printf("The addition between two impedances is: Z1 + Z2 = %v\n", addition)
	fmt.Printf("The real part of this addition is = %v\n", totalOhms)
	fmt.Printf("///////////// END OF EXERCISE 04 /////////////////\n\n")
}

func activity05() {
	/*5. Booleans Weather Forecast:

Develop a program to suggest whether to carry an umbrella. Use two boolean variables: isRaining and chanceOfRain. Set conditions such that if it's raining or the chance of rain is more than 50%, the program suggests taking an umbrella. Display the suggestion.*/
	var isRaining bool = false
	var chanceOfRain float64 = 60.0

	if isRaining || chanceOfRain > 50.0 {
		fmt.Println("Take the umbrella with you")
	} else {
		fmt.Println("Do not take the umbrella with you")
	}
	fmt.Printf("///////////// END OF EXERCISE 05 /////////////////\n\n")

}

func activity06() {
	/*6. Strings Travel Blog Post:

Write a program for a travel blog post. Include a string variable for the username and another for the blog post content. Count the number of characters in the post, and check if it contains the hashtag "#travel". Display the username, post content, character count, and whether the hashtag is present.
*/
	userName := "Tiago"
	postContent :=  "lacus sed turpis tincidunt id aliquet risus feugiat in ante"
	checkIfContains := "#travel"
	charCount := len(postContent)

	if strings.Contains(postContent, checkIfContains){
		fmt.Printf("The sentence: %s\nWas wrote by %s, has %d characters and contains %s", postContent, userName, charCount, checkIfContains)
	}else {
		fmt.Printf("The sentence: %s\nWas wrote by %s, has %d characters and does not contain %s", postContent, userName,charCount, checkIfContains)
	}
	fmt.Printf("///////////// END OF EXERCISE 06 /////////////////\n\n")

}

func activity07() {
	/*7. Constants Astronomy:

Create a program with constants representing the Earth's orbital period around the sun (365.25 days), the gravitational constant (6.67430 × 10^-11 m^3 kg^-1 s^-2), and the average surface temperature of the Earth (14 degrees Celsius). Display these constants with appropriate descriptions.
*/ 
	const earthOrbitalAroundSun float64 = 365.25
	const gravitationalConstant float64 = 6.67430e-11 // m^3 kg^-1 s^-2
	const averageSurfaceTemp float64 = 14.0 // degree Celsius

	fmt.Printf("The Earth's orbital period around the sun is %.2f days\n", earthOrbitalAroundSun)
	fmt.Printf("Gravitational Constant is %.2e m^3 kg^-1 s^-2\n", gravitationalConstant)
	fmt.Printf("The average surface temperature of the Earth is %.2f degrees celsius\n", averageSurfaceTemp)
	fmt.Printf("///////////// END OF EXERCISE 07 /////////////////\n\n")

}