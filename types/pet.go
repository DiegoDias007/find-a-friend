package types

type Pet struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	City    string  `json:"city"`
	Species string  `json:"species"`
	Breed   string  `json:"breed"`
	Height  float64 `json:"height"`
	Weight  float64 `json:"weight"`
	OrgId   int     `json:"org_id"`
}

type CreatePet struct {
	Name    string  `json:"name"`
	City    string  `json:"city"`
	Species string  `json:"species"`
	Breed   string  `json:"breed"`
	Height  float64 `json:"height"`
	Weight  float64 `json:"weight"`
	OrgId   int     `json:"org_id"`
}
