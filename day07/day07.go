package day07

import (
	"strconv"
	"strings"
)

func Step01(terminalOutput string) int {
	tree := parseTerminalOutput(terminalOutput)
	return sumDirsWithTotalSizeOf(tree, 100000)
}

type dirEntry struct {
	isDirectory bool
	size        int
	parent      *dirEntry
	children    map[string]dirEntry
}

func parseTerminalOutput(output string) dirEntry {
	root := dirEntry{isDirectory: true, children: make(map[string]dirEntry)}
	cwd := root

	var first string
	for len(output) > 0 {
		first, output = splitFirst(output, "\n$")
		cwd = processOutput(first, root, cwd)
	}

	return root
}

func processOutput(output string, root, cwd dirEntry) dirEntry {
	command, commandOutput := splitFirst(output, "\n")
	switch command {
	case "$ cd /":
		return getRoot(cwd)
	case "$ cd ..":
		return *cwd.parent
	case "$ ls":
		processDirListing(cwd, commandOutput)
		return cwd
	default:
		// case "$ cd <dirname>":
		name := output[5:]
		return cwd.children[name]
	}
}

func processDirListing(cwd dirEntry, output string) {
	lines := strings.Split(output, "\n")
	for _, l := range lines {
		switch l[:3] {
		case "dir":
			cwd.children[l[4:]] = dirEntry{
				isDirectory: true,
				children:    make(map[string]dirEntry),
				parent:      &cwd,
			}
		default:
			fields := strings.Fields(l)
			cwd.children[fields[1]] = dirEntry{size: toInt(fields[0])}
		}
	}
}

func sumDirsWithTotalSizeOf(root dirEntry, maxDirSize int) int {
	sum := 0
	dirs := findDirsWithTotalSizeOf(root, maxDirSize)
	for _, d := range dirs {
		sum += calculateTotalSize(d)
	}

	return sum
}

func findDirsWithTotalSizeOf(dir dirEntry, maxDirSize int) []dirEntry {
	dirs := make([]dirEntry, 0)
	allDirs := getDescendantDirectories(dir)
	for _, d := range allDirs {
		if calculateTotalSize(d) <= maxDirSize {
			dirs = append(dirs, d)
		}
	}

	return dirs
}

func getRoot(dir dirEntry) dirEntry {
	for dir.parent != nil {
		dir = *dir.parent
	}

	return dir
}

func getDescendantDirectories(root dirEntry) []dirEntry {
	result := make([]dirEntry, 0)
	stack := make([]dirEntry, 0)

	stack = append(stack, root)

	for len(stack) > 0 {
		dir := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, dir)
		stack = append(stack, getChildDirectories(dir.children)...)
	}

	return result
}

func getChildDirectories(children map[string]dirEntry) []dirEntry {
	result := make([]dirEntry, 0)
	for _, c := range children {
		if c.isDirectory {
			result = append(result, c)
		}
	}

	return result
}

func calculateTotalSize(dir dirEntry) int {
	size := 0
	for _, d := range getDescendantDirectories(dir) {
		size += getSize(d)
	}

	return size
}

func getSize(dir dirEntry) int {
	size := 0
	for _, c := range dir.children {
		if !c.isDirectory {
			size += c.size
		}
	}

	return size
}

func splitFirst(s string, sep string) (string, string) {
	index := strings.Index(s, sep)
	if index == -1 {
		return s, ""
	} else {
		return s[:index], s[index+1:]
	}
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
