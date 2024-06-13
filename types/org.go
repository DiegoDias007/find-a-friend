package types

type Org struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Whatsapp string `json:"whatsapp"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateOrg struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Whatsapp string `json:"whatsapp"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOrg struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
