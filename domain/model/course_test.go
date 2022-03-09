package model_test

import (
	"testing"

	"github.com/aleszilagyi/kafka-golang-api/domain/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var uuidMock = uuid.New().String()
var courseMock = model.Course{
	ID:          uuidMock,
	Name:        "course ok",
	Description: "description ok",
	Status:      "READY",
}

func TestNewCourse_Ready(t *testing.T) {
	course := model.NewCourse(uuidMock, "course ok", "description ok", "READY")

	require.EqualValues(t, course, courseMock)
}

func TestReturnStatusEnum_Undefined(t *testing.T) {
	course := model.NewCourse(uuidMock, "course ok", "description ok", "teste")

	require.EqualValues(t, course.Status, "UNDEFINED")
}

func TestReturnStatusEnum_InProgress(t *testing.T) {
	course := model.NewCourse(uuidMock, "course ok", "description ok", "IN_PROGRESS")

	require.EqualValues(t, course.Status, "IN_PROGRESS")
}