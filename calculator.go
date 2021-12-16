package main

import (
    "fmt"
)

func main() {

  var number1, number2 float64

  var operator string

  fmt.Print("First number: ")
  fmt.Scanln(&number1)

  fmt.Print("Second number: ")
  fmt.Scanln(&number2)

  fmt.Print("Operator ( + - * / ): ")
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
        fmt.Println("division by zero is impossible")
      } else {
      fmt.Print(number1 / number2)
      }
    default:
      fmt.Println("invalid operator")

  }
}

