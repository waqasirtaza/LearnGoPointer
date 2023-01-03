package main

import "fmt"
const (
	Active = true
	Inactive = false
)
type SecurityTag bool

type Item struct {
	name string
	tag SecurityTag
}

func activate(tag *SecurityTag) {
	*tag = Active
}
func deactivate(tag *SecurityTag) {
	*tag = Inactive
}
func checkout(item []Item) {
	for i :=0; i < len(item); i++ {
		deactivate(&item[i].tag)
	}
}
func main() {
	shirts := Item{"Shirts",Active}
	pants := Item{"Pants",Active}
	purse := Item{"Purse",Active}
	watch := Item{"Watch",Active}
	items := []Item{shirts,pants,purse,watch}
	fmt.Println(items)
	deactivate(&items[0].tag)
	fmt.Println(items)
	checkout(items)
	fmt.Println(items)

}