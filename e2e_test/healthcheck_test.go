package e2e_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type EndToEndTest struct {
	suite.Suite
}

func TestEndToEndSuite(t *testing.T) {
	suite.Run(t, new(EndToEndTest))
}

func (s *EndToEndTest) TestPlaceholder() {
	s.True(true)
}
