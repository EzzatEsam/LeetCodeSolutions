package main
import (
	"fmt"
)
func generateParenthesis(n int) []string {
    return  GetCombinations(n, 0)
}


func GetCombinations(stackLeft , stackRight int) []string {
	combs := []string{}

	if stackRight == 1 && stackLeft == 0 {
		return []string{")"}
	}
	if stackLeft > 0 {
		combsIfLeft := GetCombinations(stackLeft-1, stackRight +1)
		for _ , comb := range combsIfLeft {
			combs = append(combs, "(" + comb)
		}
	}

	if stackRight > 0   {
		combsIfRight := GetCombinations(stackLeft, stackRight-1)
		for _ , comb := range combsIfRight {
			combs = append(combs, ")" + comb)
		}
	}
	return combs
}

func main()  {
	tests := []int{1,2,3}
	// print the results of test
	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(generateParenthesis(test))
	}
}
