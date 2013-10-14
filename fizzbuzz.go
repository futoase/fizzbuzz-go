package main

import "strconv"

func FizzBuzz(i int) string{
  if i % 3 == 0 && i % 5 == 0 {
    return "FizzBuzz"
  } else if i % 3 == 0 {
    return "Fizz"
  } else if i % 5 == 0 {
    return "Buzz"
  }
  return strconv.Itoa(i)
}
