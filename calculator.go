package main

import (
    "fmt"
)

const errorMessageOperator = "invalid operator"
const errorMessageDivision = "division by zero is impossible"


func main() {

  var number1, number2 float64

  var operator string

  fmt.Print("First number: ")
  fmt.Scanln(&number1)

  fmt.Print("Second number: ")
  fmt.Scanln(&number2)

  fmt.Print("Enter ( + - * / ): ")
  fmt.Scanln(&operator)

  switch operator {
    case "+" :
      fmt.Print(number1 + number2)

    case "-" :
      fmt.Print(number1 - number2)

    case "*" :
      fmt.Print(number1 * number2)

    case "/" :
      if number2 == 0.0 {
        fmt.Println(errorMessageDivision)
      } else {
      fmt.Print(number1 / number2)
      }

    default:
      fmt.Println(errorMessageOperator)

  }
}

