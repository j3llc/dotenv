package dotenv

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Computes the absolute path to the .env file based on the path provided.
// I there are no errors it returns the absolute path otherwise the error
func absDotEnv(path string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return abs, nil
}

// Check if the file exists and returns an exists boolean and the file size.
func fileExists(abs string) (exists bool, size int64) {
	info, err := os.Stat(abs)
	if err != nil {
		return false, 0
	}
	size = info.Size()
	return true, size
}

// Reads and parsed the .env file. If there is an error it returns that error
func readAndParse(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if itDoes := startsWith(line, "#"); itDoes {
			continue
		}
		wg.Add(1)
		go parseLine(ctx, line, &wg)
	}
	wg.Wait()
	return nil
}

// Parses line by expected format of k=v
func parseLine(ctx context.Context, line string, wg *sync.WaitGroup) {
	defer wg.Done()
	var err error
	var key string
	var value string
	var kv []string
	cleanup := func() {
		if err != nil {
			if key != "" {
				err := os.Unsetenv(key)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	defer cleanup()
loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		default:
			//todo clean this to work beeter in identifying special character lines
			if line == "\n" || line == "\r" || line == "\r\n" || line == "\t" {
				err = LPR_INVALID_LINE
				break loop
			}
			kv = strings.Split(line, "=")
			if len(kv) < 2 {
				err = LPR_SHORT
				break loop
			} else if len(kv) > 2 {
				err = LPR_LONG
				break loop
			}
			key = kv[0]
			value = kv[1]
			if strings.Contains(value, "'") {
				strings.ReplaceAll(value, "'", "")
			}
			if strings.Contains(value, "\"") {
				strings.ReplaceAll(value, "\"", "")
			}
			err = os.Setenv(key, value)
			break loop
		}
	}
}

// Check if the source string starts with the pattern provided.
func startsWith(source, pattern string) bool {
	if len(pattern) > len(source) {
		return false
	}
	source = strings.ToLower(source)
	pattern = strings.ToLower(pattern)
	return source[:len(pattern)] == pattern
}
