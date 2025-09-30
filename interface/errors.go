package interfacego

import (
	"fmt"
	"math"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func getError() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func ErrorsDemo() {
	if err := getError(); err != nil {
		fmt.Println(err) // 自动调用 err.Error()
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeSqrt(f)
	}
	return math.Sqrt(f), nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func ErrorsSqrt() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
