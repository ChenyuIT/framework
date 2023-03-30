package console

import (
	"testing"

	"github.com/stretchr/testify/assert"

	consolemocks "github.com/chenyuIT/framework/contracts/console/mocks"
	"github.com/chenyuIT/framework/support/file"
)

func TestJobMakeCommand(t *testing.T) {
	jobMakeCommand := &JobMakeCommand{}
	mockContext := &consolemocks.Context{}
	mockContext.On("Argument", 0).Return("").Once()
	err := jobMakeCommand.Handle(mockContext)
	assert.EqualError(t, err, "Not enough arguments (missing: name) ")

	mockContext.On("Argument", 0).Return("GoravelJob").Once()
	err = jobMakeCommand.Handle(mockContext)
	assert.Nil(t, err)
	assert.True(t, file.Exists("app/jobs/goravel_job.go"))
	assert.True(t, file.Remove("app"))
}
