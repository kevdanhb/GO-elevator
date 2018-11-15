package main

import (
  "fmt"
  "math"
  "time"
)

var currentfloor float64
var traveltime float64
var elevatorfloor float64 = 0
var floordif float64
var i float64
var gotofloor float64

func main() {
  fmt.Println("What floor are you currently on?")
  fmt.Scan(&currentfloor)
  //for each floor it takes 1 seconds to travel between each one
  floordif = math.Abs(currentfloor - elevatorfloor)
  traveltime = floordif * 1
  fmt.Println("The elevator will get to you in", traveltime,"seconds.")
  if elevatorfloor > currentfloor {
    for floordif > 0 {
      time.Sleep(1 * time.Second)
      fmt.Println(elevatorfloor)
      elevatorfloor--
      floordif--
    }
  } else {
    for floordif > 0 {
      time.Sleep(1 * time.Second)
      fmt.Println(elevatorfloor)
      elevatorfloor++
      floordif--
    }
  }
    SelectFloor(0, 7)
}
  func SelectFloor(min float64, max float64) {
    fmt.Println("Elevator is here. Please come in and select a floor 0-7.")
    fmt.Scan(&gotofloor)
    if (gotofloor < min || gotofloor > max || gotofloor == elevatorfloor) {
      fmt.Println("Invalid floor please select 0-7")
      time.Sleep(1 * time.Second)
      SelectFloor(0, 7)
    }
  floordif = math.Abs(elevatorfloor - gotofloor)
  if elevatorfloor > gotofloor {
    for floordif > 0 {
      fmt.Println(elevatorfloor)
      elevatorfloor--
      time.Sleep(1 * time.Second)
      floordif--
    }
  } else {
    for floordif > 0 {
      fmt.Println(elevatorfloor)
      elevatorfloor++
      time.Sleep(1 * time.Second)
      floordif--
    }
  }
  fmt.Println("You have arrived at floor", elevatorfloor,"Now GTFO!!")
  main()
}
