package fastlyservice

type FastlyServiceModel struct {
	Comment    string `json:"comment"`
	CustomerID string `json:"customer_id"`
	ID         string `json:"id"`
	Name       string `json:"name"`
	Versions   []struct {
		Active           interface{} `json:"active"`
		Comment          string      `json:"comment"`
		CreatedAt        string      `json:"created_at"`
		DeletedAt        interface{} `json:"deleted_at"`
		Deployed         interface{} `json:"deployed"`
		InheritServiceID interface{} `json:"inherit_service_id"`
		Locked           string      `json:"locked"`
		Number           string      `json:"number"`
		Service          string      `json:"service"`
		ServiceID        string      `json:"service_id"`
		Staging          interface{} `json:"staging"`
		Testing          interface{} `json:"testing"`
		UpdatedAt        string      `json:"updated_at"`
	} `json:"versions"`
}
