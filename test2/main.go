package test2

import (
	"errors"
	"fmt"
	"strconv"
)

// input = LLRR= output = 210122
// input = ==RLL output = 000210
// input = =LLRR output = 221012
// input = RRL=R output = 012001

func Run() {
	testCase := []string{
		"LLRR=",
		"==RLL",
		"=LLRR",
		"RRL=R",
		"RLL=L",
	}

	for _, t := range testCase {

		fmt.Println(t)
		result, err := encode(t)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(result)
	}

	for {
		var str string
		fmt.Print("enter string to encode: ")
		fmt.Scanf("%s", &str)
		if result, err := encode(str); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("result: ", result)
		}
	}
}

func encode(str string) (string, error) {
	result := []int{0}
	for i := 0; i < len(str); i++ {
		s := string(str[i])
		pre := len(result) - 1
		switch s {
		case "L":
			result = append(result, result[pre]-1)
			if result[pre]-1 < 0 {
				for i, _ := range result {
					result[i] += 1
				}
			}
		case "R":
			result = append(result, result[pre]+1)
		case "=":
			result = append(result, result[pre])
		default:
			return "", errors.New("please type only '=','R','L'")
		}
	}

	stringResult := ""
	for _, num := range result {
		stringResult += strconv.Itoa(num)
	}
	return stringResult, nil
}
