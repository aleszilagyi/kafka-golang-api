package model

type StatusEnum string

const (
	UNAVAILABLE StatusEnum = "UNAVAILABLE"
	DEVELOPMENT StatusEnum = "DEVELOPMENT"
	ON_HOLD     StatusEnum = "ON_HOLD"
	OUTDATED    StatusEnum = "OUTDATED"
	READY       StatusEnum = "READY"
	UP          StatusEnum = "UP"
)

type Course struct {
	ID          string
	Name        string
	Description string
	Status      StatusEnum
}

func ReturnStatusEnum(status string) StatusEnum {
	switch status {
	case "UP":
		return UP

	case "READY":
		return READY

	case "OUTDATED":
		return OUTDATED

	case "ON_HOLD":
		return ON_HOLD

	case "DEVELOPMENT":
		return DEVELOPMENT

	default:
		return UNAVAILABLE
	}
}

func NewCourse(id, name, description, status string) Course {
	statusCourse := ReturnStatusEnum(status)

	return Course{
		ID:          id,
		Name:        name,
		Description: description,
		Status:      statusCourse,
	}
}
