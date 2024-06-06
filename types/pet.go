package types

type Pet struct {
	Id      int
	Name    string
	City    string
	Species string
	Breed   string
	Height  float64
	Weight  float64
	Org_id  int
}

type CreatePet struct {
	Name    string
	City    string
	Species string
	Breed   string
	Height  float64
	Weight  float64
	Org_id  int
}
