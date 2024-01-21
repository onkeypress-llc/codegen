package cgfs_test

import (
	"fmt"

	"github.com/onkeypress-llc/codegen/cg/cgi"
)

func isFSInterface(fs cgi.FSInterface) bool {
	return true
}

type fileExistsCase struct {
	filename string
	expected bool
}

func existsCase(filename string, expected bool) *fileExistsCase {
	return &fileExistsCase{filename: filename, expected: expected}
}

func testExistsCase(fs cgi.FSInterface, c *fileExistsCase) error {
	if exists := fs.Exists(c.filename); exists != c.expected {
		return fmt.Errorf("File %s existence incorrect %t", c.filename, exists)
	}
	return nil
}

func testExistsCases(fs cgi.FSInterface, cases ...*fileExistsCase) error {
	for i := range cases {
		c := cases[i]
		err := testExistsCase(fs, c)
		if err != nil {
			return err
		}
	}
	return nil
}

type fileReadCase struct {
	filename      string
	expectedError bool
	expectedValue string
}

func readCase(filename, expectedValue string, expectedError bool) *fileReadCase {
	return &fileReadCase{filename: filename, expectedValue: expectedValue, expectedError: expectedError}
}

func testReadCase(fs cgi.FSInterface, c *fileReadCase) error {
	if value, err := fs.Read(c.filename); value != c.expectedValue || (err != nil) != c.expectedError {
		return fmt.Errorf("Read file %s resulted in unexpected result [%s, %s]", c.filename, value, err)
	}
	return nil
}

func testReadCases(fs cgi.FSInterface, cases ...*fileReadCase) error {
	for i := range cases {
		c := cases[i]
		if err := testReadCase(fs, c); err != nil {
			return err
		}
	}
	return nil
}
