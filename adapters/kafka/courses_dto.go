package kafka

type CourseInputDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type CourseOutputDto struct {
	ID          string `json:"name"`
	Name        string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func NewCourseOutputDto(id, name, description, status string) CourseOutputDto {
	return CourseOutputDto{
		ID:          id,
		Name:        name,
		Description: description,
		Status:      status,
	}
}

func NewCourseInputDto() CourseInputDto {
	return CourseInputDto{}
}