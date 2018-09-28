package main

import "fmt"

type Cat struct {
	name string
	scientificName string
	category string
}

func New(name, scintificName, category string) Cat {
	return Cat{
		name: name,
		scientificName: scintificName,
		category: category,
	}
}

func (cat *Cat) SetName(name string) {
	cat.name = name
}

func (cat Cat) SetNameOfCopy(name string) {
	cat.name = name
}

func (cat Cat) Name() string {
	return cat.name
}

func (cat Cat) ScientificName() string {
	return cat.scientificName
}

func (cat Cat) Category() string {
	return cat.category
}

func(cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %s)", cat.scientificName, cat.category, cat.name)
}

func main() {
	cat := New("little pig", "American Shorthair", "cat")
	cat.SetName("monster")
	fmt.Printf("The cat: %s\n", cat)

	cat.SetNameOfCopy("Little pig")
	fmt.Printf("The cat: %s\n", cat)

	type Pet interface {
		SetName(name string)
		Name() string
		Category() string
		ScientificName() string
	}
	_, ok := interface{}(cat).(Pet)
	fmt.Printf("Cat implements interface Pet: %v\n", ok)
	_, ok = interface{}(&cat).(Pet)
	fmt.Printf("*Cat implements interface Pet: %v\n", ok)

}