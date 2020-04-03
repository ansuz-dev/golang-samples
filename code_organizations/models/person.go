package models

// Person structure
type Person struct {
  Name    string
  Age     int
  Address string
}

// GrowUp increases Age by 1
func (p *Person) GrowUp() {
  p.Age++
}
