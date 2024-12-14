package main

import "fmt"

type Person struct {
	Name       string
	Age        uint8
	Occupation string
}

func (p Person) String() string {
	return fmt.Sprintf("I'm %s, I'm %d and I'm %s", p.Name, p.Age, p.Occupation)
}

type Category struct {
	Name   string
	Father *Category
}

func (c Category) HasParent() bool {
	return c.Father != nil
}

func (c *Category) SetParent(parent *Category) {
	c.Father = parent
}

func main() {
	davi := Person{
		Name:       "Davi Silva",
		Occupation: "Software Engineer",
		Age:        39,
	}

	joelma := Person{"Joelma", 36, "Wife"}

	fmt.Println(davi)
	fmt.Println(joelma)

	category := Category{Name: "Category 001"}

	if !category.HasParent() {
		fmt.Println("Category avoid")
	}

	category.SetParent(&Category{Name: "SubCategory"})

	fmt.Println(category)

}
