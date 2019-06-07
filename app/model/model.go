package model

type (
	// Login body of a login post
	Login struct {
		User     string `json:"User"`
		Passwort string `json:"Passwort"`
	}

	// Registrierung body of register post
	Registrierung struct {
		Name       string `json:"Name"`
		EMail      string `json:"EMail"`
		Passwort   string `json:"Passwort"`
		Akzeptiert bool   `json:"Akzeptiert"`
	}

	// UpdateProfil body of profil put
	UpdateProfil struct {
		EMail    string `json:"EMail"`
		Passwort string `json:"Passwort"`
		Neu      string `json:"Neu"`
	}

	// LernenErgebnis body of lern post
	LernenErgebnis struct {
		Index    int  `json:"Index"`
		Ergebnis bool `json:"Ergebnis"`
	}

	// Karteikasten body of edit post/put
	Karteikasten struct {
		Kategorie      string `json:"Kategorie"`
		Unterkategorie string `json:"Unterkategorie"`
		Titel          string `json:"Titel"`
		Beschreibung   string `json:"Beschreibung"`
		Public         bool   `json:"Public"`
	}
)
