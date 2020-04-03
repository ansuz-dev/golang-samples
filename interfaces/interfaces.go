package main

import (
  "fmt"
)

// Vehicle structure
type Vehicle interface {
  speed() int
}

// Car structure
type Car struct {
  Speed  int
  Weight int
}

func (c Car) speed() int {
  return c.Speed
}

// Bike structure
type Bike struct {
  Speed int
}

func (b Bike) speed() int {
  return b.Speed
}

// Dog structure
type Dog struct {
  Speed int
}

func (d Dog) run() int {
  return d.Speed
}

func getSpeed(v Vehicle) int {
  return v.speed()
}

func getSpeedX(v interface{}) int {
  switch v.(type) {
  case Car, Bike:
    return v.(Vehicle).speed()
  case Dog:
    return v.(Dog).run()
  default:
    return 0
  }
}

func main() {
  car := Car{Speed: 100}
  bike := Bike{Speed: 50}
  dog := Dog{Speed: 10}

  speed := getSpeed(car)
  fmt.Println("Car's speed :", speed)

  speed = getSpeed(bike)
  fmt.Println("Bike's speed :", speed)

  // But, can't call getSpeed with dog
  // speed = getSpeed(dog)
  // fmt.Println("Dog's speed :", speed)

  // Instead, call getSpeedX
  speed = getSpeedX(dog)
  fmt.Println("Dog's speed :", speed)
  fmt.Println()

  // Empty interface
  var x interface{}
  x = 5
  fmt.Println(x == x)

  // x = []int{1, 2, 3}
  // fmt.Println(x == x)

  // Type assertion
  var v interface{} = Car{Speed: 60}
  var ve Vehicle = v.(Vehicle)
  fmt.Println("ve :", ve)
  var car2 Car = v.(Car)
  fmt.Println("car2 :", car2)
  // var bike2 Bike = v.(Bike)
  // fmt.Println("bike2 :", bike2)

  // dog2 := x.(Dog)
  // dog2.Speed = 15
  // fmt.Println(dog2)

  var y int8
  y = int8(x.(int))
  fmt.Println("y: ", y)
}
