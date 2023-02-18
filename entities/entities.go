package entities

type Person struct {
	ID        uint   `gorm:"PrimaryKey"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

type Product struct {
	ID   uint   `grom:"Primarykey"`
	Name string `json:"name"`
	Cost uint   `json:"cost"`
}
