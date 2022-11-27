package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Reader ready")
	fmt.Println("Enter the denominations separated by comma ','")
	denoms := readNextLineAndConvertThemToIntegers(",")
	fmt.Println("Enter the amout that you want to make")
	amount := readNextInt()

	// Time to make the amount from the denominations
	result := makeAmount(denoms, amount)

	if result != nil {
		fmt.Println("\nResult\n ")
		for i := len(result) - 1; i >= 0; i-- {
			denom := result[i]
			fmt.Println(strconv.Itoa(denom.denomination), "\tx\t", strconv.Itoa(denom.count))
		}
	} else {
		fmt.Println("\nCould not find a combination\n ")
	}
}

// This structure is to hold the result.
// ie. How many currencies of each denomination will be
// needed to make the amount
type Denom struct {
	denomination int
	count int
}

// A slice type of the Denom structure for convenience
type DenomList []Denom

// This function evaluates the denominations and find the combination
// of denominations that are needed to make the amount.
// Returns a DenomList if the denominations are sufficient to make the amount.
// Else returns nil
func makeAmount(denominations []int, amount int) DenomList {
	// Sort the denominations for ease of processing
	sort.Ints(denominations)

	// Remove all the denominations that are bigger than the amount since they
	// can never make the amount
	size := len(denominations)
	for ; size > 0; size-- {
		if (denominations[size - 1] <= amount) {
			break
		}
	}
	if size > 0 {
		denominations = denominations[0 : size]
	} else {
		// All the denominations were greater than the amount. Hence the amount
		// can't be made from the denominations
		return nil
	}

	// So, there are denominations in the set that is less than or equal to the
	// amount. Now we need to traverse through them to find the combination
	// with the least number of currencies to make the amount

	// Initialize the denomList slice to be passed around in the recursive calls
	denomList := make(DenomList, 0, size)
	for i := 0; i < size; i++ {
		denomList = append(denomList, Denom{
			denomination: denominations[i],
			count: 0,
		})
	}

	// Now, invoke the makeAmount0 function on each denomination in the decreasing
	// order of the value. The idea is to check whether the amount can be made if
	// the specified denomination is the biggest denomination in the set. See the
	// implementation of makeAmount0 to make more sense of the logic.
	for i := size - 1; i >= 0; i-- {
		// Initialize the denomination counts to 0
		for j := size - 1; j >= 0; j-- {
			denomList[j].count = 0
		}
		result := makeAmount0(denominations, i, amount, denomList)
		// nil result means the amount cannot be reached with the specified denomination
		// being the biggest one in the set. On the next iteration, the loop will
		// invoke makeAmount0 with the next biggest denomination
		if result != nil {
			return result
		}
	}
	return nil
}

// This is a recursive function which runs on the following logic
// 	1. Adds the specified denomination to the existing denomination sets to see if it can make the amount
//		
// 		1.1 if it can make the amount, then return the result
//		
//		1.2 if it cannot reach the amount by this addition, then do a recursive call to do one more addition of the same denomination
//		
//		1.3 if the addition exeeds the amount, then try adding the next biggest denomination by doing another recursive call
func makeAmount0(denominations []int, position int, amount int, denomList DenomList) DenomList {
	if position < 0 {
		// Position 0 means the recursion has ran out of elements. Returning nil
		return nil
	}
	denom := denominations[position]
	newAmount := amount - denom

	if newAmount == 0 {
		// We got an exact match. Return the result
		denomObj := &denomList[position]
		denomObj.count = denomObj.count + 1
		return denomList
 	} else if newAmount < 0{
		// Adding this denomination is too much. Lets try with a smaller denomination
		return makeAmount0(denominations, position - 1, amount, denomList)
	} else {
		// newAmount is still greater than 0. Lets try adding the same denomination again
		denomObj := &denomList[position]
		denomObj.count = denomObj.count + 1
		return makeAmount0(denominations, position, newAmount, denomList)
	}
}

var scanner *bufio.Scanner
// Lazy initialization of scanner
func getScanner() *bufio.Scanner {
	if scanner == nil {
		scanner = bufio.NewScanner(os.Stdin)
	}
	return scanner
} 

// Read the next line from the console
func readNextLine() string {
	sc := getScanner()
	sc.Scan()
	return strings.TrimSpace(sc.Text())
}

// Read the next integer from the console
func readNextInt() int {
	t := readNextLine()
	i, _ := strconv.Atoi(t)
	return i
}

// Read the next line from the console, split them by the delimiter and convert them to numbers
func readNextLineAndConvertThemToIntegers(delimiter string) []int {
	line := readNextLine()
	splits := strings.Split(line, delimiter)
	size := len(splits)
	output := make([]int, 0, size)
	for i := 0; i < size; i++ {
		str := splits[i]
		if num, err := strconv.Atoi(str); err == nil {
			output = append(output, num)
		}
	}
	return output
}