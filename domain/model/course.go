package model

type Course struct {
	ID          string
	Name        string
	Description string
	Status      string
}

func NewCourse(id string, name string, description string, status string) Course {
	return Course{
		ID:          id,
		Name:        name,
		Description: description,
		Status:      status,
	}
}
