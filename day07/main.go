package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type File struct {
	size int
	name string
}

type Directory struct {
	prev_dir    *Directory
	files       []File
	directories []*Directory
	name        string
}

func (directory *Directory) Cd(name string) *Directory {
	if name == ".." {
		return directory.prev_dir
	}
	for _, dir := range directory.directories {
		if dir.name == name {
			return dir
		}
	}
	return nil
}

func (directory *Directory) Mkdir(name string) {
	directory.directories = append(directory.directories, &Directory{name: name, prev_dir: directory})
}

func (directory *Directory) Touch(name string, size int) {
	directory.files = append(directory.files, File{name: name, size: size})
}

func (directory *Directory) Size() int {
	size := 0
	for _, file := range directory.files {
		size += file.size
	}
	for _, dir := range directory.directories {
		size += dir.Size()
	}
	return size
}

func (directory *Directory) CollectDirs() []*Directory {
	dirs := make([]*Directory, 0)
	dirs = append(dirs, directory)
	for _, dir := range directory.directories {
		collected := dir.CollectDirs()
		for _, d := range collected {
			dirs = append(dirs, d)
		}
	}
	return dirs
}

func main() {
	file_name := os.Args[1]
	file, err := os.Open(file_name)

	if err != nil {
		fmt.Println("Error opening file!")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var current_directory *Directory

	ls := false
	dirs_to_ls := make([]string, 0)

	for scanner.Scan() {
		text := strings.Fields(scanner.Text())
		if text[0] == "$" {
			if ls {
				ls = false
				create_dirs(dirs_to_ls, current_directory)
				dirs_to_ls = make([]string, 0)
			}
			switch text[1] {
			case "cd":
				if current_directory == nil {
					current_directory = &Directory{name: text[2]}
				} else {
					current_directory = current_directory.Cd(text[2])
				}
			case "ls":
				ls = true
			}
		} else if ls {
			dirs_to_ls = append(dirs_to_ls, scanner.Text())
		}
	}

	if ls {
		create_dirs(dirs_to_ls, current_directory)
	}

	for {
		if current_directory.name == "/" {
			break
		}
		current_directory = current_directory.prev_dir
	}

	root_size := current_directory.Size()

	result := 0
	collected_dirs := current_directory.CollectDirs()
	for _, dir := range collected_dirs {
		size := dir.Size()
		if size <= 100000 {
			result += size
		}
	} 

	fmt.Println("All directories below 100000 have size of", result)

	sort.Slice(collected_dirs, func(i, j int) bool {
		return collected_dirs[i].Size() < collected_dirs[j].Size()
	})

	for _, dir := range collected_dirs {
		size := dir.Size()
		if 40000000 > root_size - size  {
			fmt.Println("You can delete directory", dir.name, "with size of", size)
			return
		}
	} 

	file.Close()
}

func create_dirs(dirs_to_ls []string, current_directory *Directory) {
	for _, entry := range dirs_to_ls {
		params := strings.Fields(entry)
		if params[0] == "dir" {
			current_directory.Mkdir(params[1])
		} else {
			size, _ := strconv.Atoi(params[0])
			current_directory.Touch(params[1], size)
		}
	}
}
