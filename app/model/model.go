package model

type (
	// Login body of a login post
	Login struct {
		User     string `json:"User"`
		Passwort string `json:"Passwort"`
	}
)
