package templates

type Person struct {
	Name       string     `json:"name"`
	Age        int        `json:"-"`
	Profession string     `json:"profession"`
	Address    Address    `json:"address"`
	CreatedAt  Formatdate `json:"created_at"`
	DeletedAt  Formatdate `json:"deleted_at,omitempty"`
}
