package templates

type Address struct {
	Addressline1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zipcode"`
	Country      string `json:"country,omitempty"`
}
