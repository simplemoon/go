package main

import (
	"gorm.io/gen"
)

type User struct {
	Name string
	Age  int
}

type Querier interface {
	// SELECT * FROM @@table WHERE id=@id
	GetByID(id int) (gen.T, error)

	// GetUsersByRole query data by roles and return it as *slice of pointer*
	//   (The below blank line is required to comment for the generated method)
	//
	// SELECT * FROM @@table WHERE role IN @rolesName
	GetByRoles(rolesName ...string) ([]*gen.T, error)

	// InsertValue insert value
	//
	// INSERT INTO @@table (name, age) VALUES (@name, @age)
	InsertValue(name string, age int) error
}

func main() {

	g := gen.NewGenerator(gen.Config{})

	g.ApplyInterface(func(Querier) {

	}, User{}, g.GenerateModel("employee"))

	g.Execute()
}
