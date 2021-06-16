package entities

type WithHeld struct {
	Scope        string   `json:"scope"`
	CountryCodes []string `json:"country_codes"`
}

type Includes struct {
	Users []User `json:"users,omitempty"`
}

type Error struct {
	Detail       string `json:"detail"`
	Title        string `json:"title"`
	ResourceType string `json:"resource_type,omitempty"`
	Parameter    string `json:"parameter,omitempty"`
	Value        string `json:"value,omitempty"`
	Type         string `json:"type,omitempty"`
}

type BaseData struct {
	Includes *Includes `json:"includes,omitempty"`
	Meta     *Meta     `json:"meta,omitempty"`
	Error    []*Error  `json:"error,omitempty"`
}

type UserResp struct {
	Data *User `json:"data,omitempty"`
	*BaseData
}

type UsersResp struct {
	Data []*User `json:"data,omitempty"`
	*BaseData
}
