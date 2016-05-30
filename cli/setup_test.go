package cli

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CliSuite struct {
	suite.Suite
}

func (s *CliSuite) SetupTest() {
}

func TestAllTests(t *testing.T) {
	suite.Run(t, new(CliSuite))
}
