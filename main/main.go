// https://www.hackerrank.com/challenges/game-of-two-stacks/problem

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func twoStacks(x int, a []int, b []int) int {
	var count int
	var sum int

	// Get the count from the first array
	var idxA, idxB = 0, 0
	for idxA < len(a) {
		var n1 = a[idxA]

		if sum+n1 <= x {
			sum += n1
			count++
			idxA++
		} else {
			break
		}
	}

	idxA--

	// Now, let's put second array into consideration
	var maxCount = count
	for idxB = 0; idxB < len(b); idxB++ {
		var n2 = b[idxB]

		// Let's try to add the value to sum with n2
		sum += n2
		count++

		for sum > x {
			// Adding new element yields to sum more than maximum value
			// Let's do some more logic. Reduce sum from the value of first array.
			if idxA < 0 {
				// No more elements can be removed from first array. Break it.
				break
			}

			var n1 = a[idxA]
			sum -= n1
			idxA--
			count--
		}

		if sum > x {
			// Even after reduction attempt, the sum is still more than the maximum value.
			// No more optimum solution.
			break
		} else if maxCount < count {
			maxCount = count
		}
	}

	return maxCount
}

func main() {
	//fmt.Println(twoStacks(10, []int{4, 2, 4, 6, 1}, []int{2, 1, 8, 5}))
	//fmt.Println(twoStacks(8, []int{2, 1, 1, 1, 1, 10}, []int{7, 8, 9}))
	//fmt.Println(twoStacks(20, []int{17, 1, 1, 1, 10, 11}, []int{8, 4, 8, 2, 5}))
	//fmt.Println(twoStacks(62, []int{7, 15, 12, 0, 5, 18, 17, 2, 10, 15, 4, 2, 9, 15, 13, 12, 16}, []int{12, 16, 6, 8, 16, 15, 18, 3, 11, 0, 17, 7, 6, 11, 14, 13, 15, 6, 18, 6, 16, 12, 16, 11, 16, 11}))

	var inputFiles = []string{"input01.txt", "input02.txt"}
	var outputFiles = []string{"output01.txt", "output02.txt"}

	for i := 0; i < len(inputFiles); i++ {
		var inputFile, _ = os.Open(inputFiles[i])
		defer inputFile.Close()

		var outputFile, _ = os.Open(outputFiles[i])
		defer outputFile.Close()

		var inputReader = bufio.NewReader(inputFile)
		var outputReader = bufio.NewReader(outputFile)

		var numOfCasesStr, _ = inputReader.ReadString('\n')
		var numOfCases, _ = strconv.Atoi(strings.TrimSpace(numOfCasesStr))

		for i := 0; i < numOfCases; i++ {
			var line, _ = inputReader.ReadString('\n')
			var data = strings.Split(strings.TrimSpace(line), " ")
			var aLength, _ = strconv.Atoi(data[0])
			var bLength, _ = strconv.Atoi(data[1])
			var max, _ = strconv.Atoi(data[2])

			// Read the first array data
			var firstArrayStr, _ = inputReader.ReadString('\n')
			var firstArray = strings.Split(strings.TrimSpace(firstArrayStr), " ")
			var firstArrayData = make([]int, aLength)

			for i := 0; i < aLength; i++ {
				firstArrayData[i], _ = strconv.Atoi(firstArray[i])
			}

			// Read the second array data
			var secondArrayStr, _ = inputReader.ReadString('\n')
			var secondArray = strings.Split(strings.TrimSpace(secondArrayStr), " ")
			var secondArrayData = make([]int, bLength)

			for i := 0; i < bLength; i++ {
				secondArrayData[i], _ = strconv.Atoi(secondArray[i])
			}

			// Read the expected output
			line, _ = outputReader.ReadString('\n')
			expected, _ := strconv.Atoi(strings.TrimSpace(line))

			var result = twoStacks(max, firstArrayData, secondArrayData)
			fmt.Printf("Case %d: ", i)

			if result == expected {
				fmt.Println("OK")
			} else {
				fmt.Printf("Failed: Expected %d, Actual %d\n", expected, result)
			}
		}
	}
}
