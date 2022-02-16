package ports

import (
	"github.com/aleszilagyi/kafka-golang-api/domain/model"
	"github.com/aleszilagyi/kafka-golang-api/adapters/kafka"
)

type RepositoryPort interface {
	Insert(course model.Course) error
}

type CreateCoursePort interface{
	Execute(input kafka.CourseInputDto)
}
