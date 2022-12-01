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
