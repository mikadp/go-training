package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

const aConst int = 64

func main() {

	var aString string = "This is test"
	fmt.Println(aString)
	fmt.Printf("The variable's type is %T\n", aString)

	var anInteger int = 50
	fmt.Println(anInteger)
	fmt.Printf("the variable's type is %T\n", anInteger)

	fmt.Println(aConst)

	//User input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("you entered:", input)

	//Conver string inputs to other types
	fmt.Print("Enter number: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value of number:", aFloat)
	}

	//Math
	i1, i2, i3 := 12, 34, 65
	intSum := i1 + i2 + i3
	fmt.Println("integer sum: ", intSum)

	f1, f2, f3 := 3.45, 5.67, 4.77
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum: ", floatSum)

	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("rounded Float sum: ", floatSum)

	//Time
	n := time.Now()
	fmt.Println("time is:", n)
	fmt.Println(n.Format(time.ANSIC))

}
