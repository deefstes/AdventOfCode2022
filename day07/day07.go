package day07

import (
	"strconv"
	"strings"

	"github.com/deefstes/AdventOfCode2022/helpers"
)

type Directory struct {
	Name        string
	Files       []File
	Directories []Directory
}

type File struct {
	Name string
	Size int64
}

func ParseDir(input *[]string) Directory {
	var retval Directory
	elements := strings.Split((*input)[0], " ")
	if !strings.HasPrefix((*input)[0], "$ cd ") {
		return retval
	}

	retval.Name = elements[2]
	*input = (*input)[1:]

	for {
		if len(*input) == 0 {
			return retval
		}
		if (*input)[0] == "$ ls" {
			retval.Files = append(retval.Files, ListFiles(input)...)
		} else if (*input)[0] == "$ cd .." {
			*input = (*input)[1:]
			return retval
		} else if strings.HasPrefix((*input)[0], "$ cd ") {
			retval.Directories = append(retval.Directories, ParseDir(input))
		}

		//
		// if len(*input) == 0 {
		// 	return retval
		// }
	}
}

func ListFiles(input *[]string) []File {
	var retval []File
	if (*input)[0] != "$ ls" {
		return retval
	}

	*input = (*input)[1:]

	for {
		if len(*input) == 0 {
			return retval
		}
		elements := strings.Split((*input)[0], " ")
		if elements[0] == "dir" {
			// ignore it, we'll deal with it when "$ cd x" comes around
			*input = (*input)[1:]
		} else if helpers.IsNumber(elements[0]) {
			size, _ := strconv.ParseInt(elements[0], 10, 64)
			retval = append(retval, File{Size: size, Name: elements[1]})
			*input = (*input)[1:]
		} else if (*input)[0] == "$ cd .." {
			//*input = (*input)[1:]
			return retval
		} else if strings.HasPrefix((*input)[0], "$ cd ") {
			return retval
		}
	}
}

func (d *Directory) CalcDirSize() int64 {
	var total int64
	for _, f := range d.Files {
		total = total + f.Size
	}
	for _, sd := range d.Directories {
		total = total + sd.CalcDirSize()
	}

	return total
}

func (d *Directory) AllDirSizes() []File {
	var retval []File
	retval = append(retval, File{Name: d.Name, Size: d.CalcDirSize()})
	for _, sd := range d.Directories {
		retval = append(retval, sd.AllDirSizes()...)
	}

	return retval
}
