package main

import (
	"fmt"
	"syscall/js"
	"time"

	zxcvbn "github.com/nbutton23/zxcvbn-go"
	zxcvbnfork "github.com/trustelem/zxcvbn"
)

func main() {
	fmt.Println("hello from Golang using Web Assembly!")

	js.Global().Set("passwordStrength", zxcvbnWrapper())
	js.Global().Set("passwordStrengthFork", zxcvbnForkWrapper())

	<-make(chan bool)
}

func zxcvbnWrapper() js.Func {
	jsFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		start := time.Now()
		if len(args) != 1 {
			return "Invalid number of arguments passed"
		}
		password := args[0].String()
		score := zxcvbn.PasswordStrength(password, nil)
		elapsed := time.Since(start)
		fmt.Printf("zxcvbn password evaluation took %s\n", elapsed.String())
		return map[string]interface{}{
			"password":           score.Password,
			"calc_time":          score.CalcTime,
			"crack_time":         score.CrackTime,
			"crack_time_display": score.CrackTimeDisplay,
			"entropy":            score.Entropy,
			"score":              score.Score,
		}
	})
	return jsFunc
}

func zxcvbnForkWrapper() js.Func {
	jsFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		start := time.Now()
		if len(args) != 1 {
			return "Invalid number of arguments passed"
		}
		password := args[0].String()
		score := zxcvbnfork.PasswordStrength(password, nil)
		elapsed := time.Since(start)
		fmt.Printf("zxcvbnFork password evaluation took %s\n", elapsed.String())
		return map[string]interface{}{
			"password":  password,
			"calc_time": score.CalcTime,
			"score":     score.Score,
		}
	})
	return jsFunc
}
