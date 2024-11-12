package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

// Define the test suite struct
type MathSuite struct {
	suite.Suite
}

// SetupTest method to run before each test
func (s *MathSuite) SetupTest() {
	// Setup code can be added here if needed
}

// TearDownTest method to run after each test
func (s *MathSuite) TearDownTest() {
	// Teardown code can be added here if needed
}

// TestAdd method to test the Add function
func (s *MathSuite) TestAdd() {
	result := Add(2, 3)
	s.Equal(5, result, "Expected 2 + 3 to equal 5")

	result = Add(-1, 1)
	s.Equal(0, result, "Expected -1 + 1 to equal 0")

	result = Add(0, 0)
	s.Equal(0, result, "Expected 0 + 0 to equal 0")

	result = Add(-2, 5)
	s.Equal(3, result, "Expected -2 + 5 to equal 3")
}

// TestMathSuite function to run the test suite
func TestMathSuite(t *testing.T) {
	suite.Run(t, &MathSuite{})
}
