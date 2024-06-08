package types

type Org struct {
	Id int
	Name string
	Address string
	Whatsapp string
	Password string
}

type CreateOrg struct {
	Name string
	Address string
	Whatsapp string
	Password string
}