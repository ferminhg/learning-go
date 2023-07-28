package domain

type DescriptionGenerator interface {
	Run(title string) ([]RandomDescription, error)
}
