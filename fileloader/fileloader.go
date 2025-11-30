package fileloader

import( 
    "os"
    "path/filepath"
)

func Load(path string) ([]byte, error) {
    absolutePath, err := filepath.Abs(path)

    if err != nil {
        return nil, err
    }

    fileContent, err := os.ReadFile(absolutePath)

    if err != nil {
        return nil, err
    }

    return fileContent, nil
}
