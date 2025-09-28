package foundations

import "fmt"

// Structs are defined with `struct` keyword.
type City struct {
	// List attributes and types.
	Name       string
	Population int

	// Structs are types. So ...
	GPSLocation struct {
		Latitude  float64
		Longitude float64
	}
}

// We can add methods to structs.
// Actually, ... we add them to defined types.
func (c City) IsBig() bool {
	return c.Population >= 1_000_00
}

// We need to have a pointer to it if we want to modify an object in its functions.
func (c *City) MakeBigger() {
	c.Population += c.Population / 10
}

// We can also make "static" functions.
// Still needs to be called on an object.
func (City) IsRentAffordable() bool {
	return false
}

// There are no constructors, destructors, or other special functions.
// Normal pattern:
// - type definition
// - "constructor" function: `func New<Type>( ... ) Type { ... }`
//   May return just the type, or pointer, or errors, whatever is useful.
// - member functions

// Need a destructor?
// Add a `close` function and use it with `defer`.

func structs() {
	// We can also define structs (and other types) within functions.
	type Country struct {
		Name   string
		Cities []City
	}

	// Even for just one variable.
	continent := struct {
		Countries []Country
	}{
		Countries: nil,
	}

	// Structs are value types. If we create a new object, all attributes have
	// their type's zero-value.
	city := City{}
	fmt.Println(
		city.Name,                  // ""
		city.Population,            // 0
		city.GPSLocation.Latitude,  // 0.0
		city.GPSLocation.Longitude, // 0.0
	)

	// We can initialize values within the `{}`.
	berlin := City{
		Name:       "Berlin",
		Population: 3_685_265,

		// Not naming types can become very verbose later.
		GPSLocation: struct {
			Latitude  float64
			Longitude float64
		}{52.540431727242385, 13.412172631133856},
	}

	// Assigning to another objects means (shallow) copying.
	city = berlin
	city.Name = "Neu-Berlin"
	fmt.Println(berlin.Name, city.Name) // Berlin Neu-Berlin
}
