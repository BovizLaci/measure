package measure

import (
	"fmt"
	"time"
)

func MeasureGoFuncSpeed(f func(parameters ...any) ([]any, error), funcParams ...any) error {
	fmt.Println()
	start := time.Now()
	results, err := f(funcParams...)
	elapsed := time.Since(start)
	if err != nil {
		return err
	}
	fmt.Println("Results:", results)
	fmt.Printf("Duration: %v ms\n", elapsed.Milliseconds())
	fmt.Printf("Duration: %v s\n", elapsed.Seconds())
	return nil
}


func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ReverseSlice(s []any) ([]any) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
