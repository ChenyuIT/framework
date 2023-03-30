package crypt

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/chenyuIT/framework/facades"
	"github.com/chenyuIT/framework/testing/mock"
)

type ApplicationTestSuite struct {
	suite.Suite
}

func TestApplicationTestSuite(t *testing.T) {
	mockConfig := mock.Config()
	mockConfig.On("GetString", "app.key").Return("11111111111111111111111111111111").Once()

	facades.Crypt = NewApplication()
	suite.Run(t, new(ApplicationTestSuite))
	mockConfig.AssertExpectations(t)
}

func (s *ApplicationTestSuite) SetupTest() {

}

func (s *ApplicationTestSuite) TestEncryptString() {
	encryptString, err := facades.Crypt.EncryptString("Goravel")
	s.NoError(err)
	s.NotEmpty(encryptString)
}

func (s *ApplicationTestSuite) TestDecryptString() {
	payload, err := facades.Crypt.EncryptString("Goravel")
	s.NoError(err)
	s.NotEmpty(payload)

	value, err := facades.Crypt.DecryptString(payload)
	s.NoError(err)
	s.Equal("Goravel", value)

	_, err = facades.Crypt.DecryptString("Goravel")
	s.Error(err)
}
