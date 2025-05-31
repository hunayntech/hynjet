package metadata

type Relation struct {
	Key        string  `json:"key"`
	Model      string  `json:"model"`
	Type       *string `json:"type"`
	ForeignKey *string `json:"foreignKey"`
	References *string `json:"references"`
}
