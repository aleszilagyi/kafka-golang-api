package repository

import (
	"github.com/aleszilagyi/kafka-golang-api/domain/model"
	"database/sql"
)

type CourseMySQLRepository struct {
	Db *sql.DB
}

func (courseRepository CourseMySQLRepository) Insert(course model.Course) error {
	statement, err := courseRepository.Db.Prepare(`Insert into courses(id, name, description, status) values(?,?,?,?)`)
	if err != nil {
		return err
	}

	_, err = statement.Exec(
		course.ID,
		course.Name,
		course.Description,
		course.Status,
	)

	if err != nil {
		return err
	}

	return nil
}
