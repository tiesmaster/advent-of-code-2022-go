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

	for len(output) > 0 {
		index := strings.Index(output[1:], "$")
		if index == -1 {
			index = len(output)
		}

		cwd = processOutput(output[:index], root, cwd)
		output = output[index:]

	}

	return root
}

func processOutput(output string, root, cwd dirEntry) dirEntry {
	index := strings.Index(output, "\n")
	command := output[:index]
	switch command {
	case "$ cd /":
		// TODO: Implement actual cd /
		return cwd
	case "$ cd ..":
		return *cwd.parent
	case "$ ls":
		processDirListing(cwd, output[index:])
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

func splitFirst(s string, sep string) (string, string)

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
