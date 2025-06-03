package template

type RelationConfig struct {
	Model      string    `json:"model"`
	Type       *string   `json:"type"`
	ForeignKey *string   `json:"foreignKey"`
	References *string   `json:"references"`
	Tags       *[]string `json:"tags"`
}

type ModelRelations map[string]RelationConfig

type ModelConfig struct {
	Relations ModelRelations `json:"relations"`
}

type Config struct {
	ModelConfig map[string]ModelConfig `json:"modelConfig"`
}
