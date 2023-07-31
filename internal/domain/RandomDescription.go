package domain

type RandomDescription struct {
	Description string  `json:"description"`
	Confidence  float32 `json:"confidence"`
}

func NewRandomDescription(description string, confidence float32) RandomDescription {
	return RandomDescription{Description: description, Confidence: confidence}
}
