package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

func createOutputFile(filename string) (*os.File, error) {
	// Ensure "dist" folder exists
	if err := os.MkdirAll("dist", os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create dist folder: %w", err)
	}

	// Create the output file
	filePath := filepath.Join("dist", filename)
	f, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file %s: %w", filePath, err)
	}

	return f, nil
}
