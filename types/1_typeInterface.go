package types

import (
	"fmt"
)

// Animal interface defines the behavior that all animals must implement
// In Go, interfaces are implicitly implemented - no need to explicitly declare implementation
type Animal interface {
	MakeSound() string
	GetName() string
}

// Dog struct represents a dog with its properties
type Dog struct {
	Name string
	Age  int
}

// GetName implements the Animal interface for Dog
func (d Dog) GetName() string {
	return d.Name
}

// MakeSound implements the Animal interface for Dog
func (d Dog) MakeSound() string {
	return "Woof"
}

// Cat struct represents a cat with its properties
type Cat struct {
	Name string
	Age  int
}

// GetName implements the Animal interface for Cat
func (c Cat) GetName() string {
	return c.Name
}

// MakeSound implements the Animal interface for Cat
func (c Cat) MakeSound() string {
	return "Meow"
}

/** Interfaces in Go
Key concepts:
- Interfaces define a set of method signatures
- Types implicitly implement interfaces by implementing all required methods
- No explicit declaration of interface implementation is needed
- Interfaces enable polymorphism and code reuse
- Multiple types can implement the same interface
*/
func TypeInterface() {
	fmt.Println(" ")
	fmt.Println("TYPES: Interface")

	// Create instances of Dog and Cat
	dog := Dog{Name: "Buddy", Age: 3}
	cat := Cat{Name: "Whiskers", Age: 2}

	// Both Dog and Cat can be passed to PrintAnimal because they implement the Animal interface
	PrintAnimal(dog)
	PrintAnimal(cat)
}

// PrintAnimal accepts any type that implements the Animal interface
// This demonstrates polymorphism in Go
func PrintAnimal(animal Animal) {
	fmt.Println("Animal:", animal.GetName(), "Sound:", animal.MakeSound())
}
