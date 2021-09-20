package config_test

import (
	"testing"

	"github.com/go-mongo/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConfigSuite struct {
	suite.Suite
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, &ConfigSuite{})
}

func (suite *ConfigSuite) SetupTest() {
}

func (suite *ConfigSuite) TestGetLocation() {
	message := config.Initialize("../").GetLocalization("en", "en.message.succeed", "product")

	assert.Equal(suite.T(), message, "success retrieve data product")
}
