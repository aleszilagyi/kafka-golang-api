package ports

import (
	"github.com/aleszilagyi/kafka-golang-api/adapters/kafka"
	"github.com/aleszilagyi/kafka-golang-api/domain/model"
)

type RepositoryPort interface {
	Insert(course model.Course) error
}

type CreateCoursePort interface {
	Execute(input kafka.CourseInputDto) (kafka.CourseOutputDto, error)
}
