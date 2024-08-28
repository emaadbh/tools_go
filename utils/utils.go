package utils

import "fmt"

func ConstValue() string {
	return DashGenerator(25) + " %v RECORD " + DashGenerator(25) + "\n"
}

func StarGenerator(number int) string {
	var value string

	for i := 0; i < number; i++ {
		value += "*"
	}

	return value

}
func DashGenerator(number int) string {
	var value string

	for i := 0; i < number; i++ {
		value += "-"
	}

	return value

}

func EmadGenerator() string {

	return ` 
.------------------------------------------------------.
|      __  __        ___     ____  __    __  __    ___ |
|     (  \/  )      (   \   (_  _)/  \  /  \(  )  / __)|
|      )    (        ) ) )    )( ( () )( () ))(__ \__ \|
|     (_/\/\_)      (___/    (__) \__/  \__/(____)(___/|
'-----------------------V0.0.1-------------------------'
`

}

func Exit() bool {
	var index string
	fmt.Print(" To use again, press any key and press enter - ")
	fmt.Println("exit type: q , exit or 0 ")

	_, _ = fmt.Scan(&index)
	if index == "0" || index == "exit" || index == "q" {
		return false
	} else {
		return true
	}
}
