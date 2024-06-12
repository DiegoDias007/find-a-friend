package types

type Org struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Whatsapp string `json:"whatsapp"`
	Password string `json:"password"`
}

type CreateOrg struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Whatsapp string `json:"whatsapp"`
	Password string `json:"password"`
}
