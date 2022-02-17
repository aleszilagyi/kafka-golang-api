package model

type statusEnum string

const (
	Undefined  statusEnum = "UNDEFINED"
	InProgress            = "IN_PROGRESS"
	Ready                 = "READY"
)

type Course struct {
	ID          string
	Name        string
	Description string
	Status      statusEnum
}

func returnStatusEnum(status string) statusEnum {
	switch status {
	case "IN_PROGRESS":
		return InProgress

	case "READY":
		return Ready

	default:
		return Undefined
	}
}

func NewCourse(id, name, description, status string) Course {
	statusCourse := returnStatusEnum(status)

	return Course{
		ID:          id,
		Name:        name,
		Description: description,
		Status:      statusCourse,
	}
}
