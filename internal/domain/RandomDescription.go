package domain

type RandomDescription struct {
	description string
	confidence  float32
}

func NewRandomDescription(description string, confidence float32) RandomDescription {
	return RandomDescription{description: description, confidence: confidence}
}
