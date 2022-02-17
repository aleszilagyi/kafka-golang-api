package usecase

import (
	"github.com/aleszilagyi/kafka-golang-api/adapters/kafka"
	"github.com/aleszilagyi/kafka-golang-api/ports"
	"github.com/aleszilagyi/kafka-golang-api/domain/model"
	"github.com/google/uuid"
)

type CreateCourse struct {
	Repository ports.RepositoryPort
}

func (createCourse CreateCourse) Execute(input kafka.CourseInputDto) (kafka.CourseOutputDto, error) {

	id := uuid.New().String()
	course := model.NewCourse(id, input.Name, input.Description, input.Status)

	err := createCourse.Repository.Insert(course)
	if err != nil {
		return kafka.CourseOutputDto{}, err
	}

	output := kafka.NewCourseOutputDto(course.ID, course.Name, course.Description, string(course.Status))

	return output, nil
}
