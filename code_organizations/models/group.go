package models

// Group structure
type Group struct {
  Name    string
  Members []*Person
}

// AddMember adds a member into group
func (g *Group) AddMember(p *Person) {
  g.Members = append(g.Members, p)
}
