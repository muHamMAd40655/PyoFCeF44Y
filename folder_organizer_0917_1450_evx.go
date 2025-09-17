// 代码生成时间: 2025-09-17 14:50:14
package main

import (
    "os"
    "path/filepath"
    "log"
    "strings"
    "sort"
    "fmt"
)

// FolderOrganizer represents the main structure for organizing folders
type FolderOrganizer struct {
    RootPath string
}

// NewFolderOrganizer creates a new FolderOrganizer instance
func NewFolderOrganizer(rootPath string) *FolderOrganizer {
    return &FolderOrganizer{RootPath: rootPath}
}

// OrganizeFolders goes through all the directories starting from the root path
// and sorts the files alphabetically within each folder.
func (f *FolderOrganizer) OrganizeFolders() error {
    err := filepath.WalkDir(f.RootPath, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if d.IsDir() {
            return f.sortFilesInDirectory(path)
        }
        return nil
    })
    return err
}

// sortFilesInDirectory sorts the files in the given directory alphabetically
func (f *FolderOrganizer) sortFilesInDirectory(dirPath string) error {
    files, err := os.ReadDir(dirPath)
    if err != nil {
        return err
    }

    var fileNames []string
    for _, file := range files {
        if !file.IsDir() {
            fileNames = append(fileNames, file.Name())
        }
    }

    sort.Strings(fileNames) // Sort file names alphabetically
    for _, name := range fileNames {
        oldPath := filepath.Join(dirPath, name)
        newPath := filepath.Join(dirPath, strings.ToLower(name)) // Convert to lowercase for consistency
        if oldPath != newPath { // Rename only if the path has changed
            err := os.Rename(oldPath, newPath)
            if err != nil {
                return err
            }
        }
    }
    return nil
}

func main() {
    rootPath := "/path/to/your/directory" // Replace with the actual root path
    organizer := NewFolderOrganizer(rootPath)
    err := organizer.OrganizeFolders()
    if err != nil {
        log.Fatalf("Error organizing folders: %v
", err)
    } else {
        fmt.Println("Folders organized successfully.")
    }
}
