package config

type Defaults struct {
	BasePath string
}

// Init Set defaults for each variable needed across the project
func Get() Defaults {
	return Defaults{
		// TODO might be a better way ?
		// Customize to your needs
		BasePath: "w/perso/cellular-automata-go/",
	}
}
