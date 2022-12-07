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
	root := dirEntry{isDirectory: true}
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
		// TODO: Implement actual cd /
		return cwd
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
			cwd.children[l[4:]] = dirEntry{isDirectory: true}
		default:
			fields := strings.Fields(l)
			cwd.children[fields[1]] = dirEntry{size: toInt(fields[0])}
		}
	}
}

func sumDirsWithTotalSizeOf(root dirEntry, maxDirSize int) int {
	panic("unimplemented")
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
