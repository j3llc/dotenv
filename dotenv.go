// Package dotenv is a zero-dependency package that loads environment variables from an .env
// file into the executable process. Storing configuration in the environment separate from code.
// It is based on The Twelve-Factor App methodology.
package dotenv

import (
	"fmt"
)

// Loads .env file from present working directory. If there is an error it returns that error
// otherwise it returns nil. This functions uses [LoadPath] passing a default path using
// the prsent working directory and assuming there is a .env file in it
func Load() error {
	path := "./.env"
	if err := LoadPath(path); err != nil {
		return err
	}
	return nil
}

// Loads .env file from specified path. If there is an error it returns that error
func LoadPath(path string) error {
	path, err := absDotEnv(path)
	if err != nil {
		return err
	}
	exists, size := fileExists(path)
	if !exists {
		return fmt.Errorf("%s not found", path)
	}
	// if file is biggern than 2mg
	if size > (1 << 20) {
		return fmt.Errorf("File size is too big. Max is 20mg file size is %d", size)
	}
	if err = readAndParse(path); err != nil {
		return err
	}
	return nil
}
