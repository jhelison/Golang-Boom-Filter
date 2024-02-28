package main_test

import (
	main "bloom"
	"testing"

	"github.com/stretchr/testify/suite"
)

// BloomTestSuite is a basic bloom test suite
type BloomTestSuite struct {
	suite.Suite
}

// TestBloomTestSuite runs all tests under the bloom test suite
func TestBloomTestSuite(t *testing.T) {
	s := new(BloomTestSuite)
	suite.Run(t, s)
}

func (suite *BloomTestSuite) TestBloom() {
	testCases := []struct {
		name     string
		inputs   []string
		testWord string
		outcome  bool
	}{
		{
			name:     "Pass, maybe in my set",
			inputs:   []string{"test", "this is a test", "more test"},
			testWord: "test",
			outcome:  true,
		},
		{
			name:     "Not pass, not in set",
			inputs:   []string{"test", "this is a test", "more test asd"},
			testWord: "test2123",
			outcome:  false,
		},
		{
			name:     "Pass, maybe on my set 1",
			inputs:   []string{"test", "this is a test", "more test asd"},
			testWord: "this is a test",
			outcome:  true,
		},
		{
			name:     "Pass, maybe on my set 2",
			inputs:   []string{"test", "this is a test", "more test asd"},
			testWord: "more test asd",
			outcome:  true,
		},
		{
			name:     "Pass, maybe on my set 3",
			inputs:   []string{"test", "this is a test", "more test asd"},
			testWord: "test",
			outcome:  true,
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			bloom := main.NewBloom(
				20,
				5,
				main.FNVHashStrategy{},
			)
			for _, input := range tc.inputs {
				bloom.Add([]byte(input))
			}

			res := bloom.Check([]byte(tc.testWord))

			if tc.outcome {
				suite.Require().True(res)
			} else {
				suite.Require().False(res)
			}
		})
	}
}
