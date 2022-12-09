package day07

type dirEntry struct {
	isDirectory bool
	size        int
	parent      *dirEntry
	children    map[string]dirEntry
}

func (dir dirEntry) root() dirEntry {
	for dir.parent != nil {
		dir = *dir.parent
	}

	return dir
}

func (dir dirEntry) directories() []dirEntry {
	result := make([]dirEntry, 0)
	for _, c := range dir.children {
		if c.isDirectory {
			result = append(result, c)
		}
	}

	return result
}

func (dir dirEntry) files() []dirEntry {
	result := make([]dirEntry, 0)
	for _, c := range dir.children {
		if !c.isDirectory {
			result = append(result, c)
		}
	}

	return result
}

func (root dirEntry) allDirectories() []dirEntry {
	result := make([]dirEntry, 0)
	stack := newStack()

	stack.push(root)

	for !stack.isEmpty() {
		dir := stack.pop()

		result = append(result, dir)
		stack.push(dir.directories()...)
	}

	return result
}

func (root dirEntry) query(predicate func(d dirEntry) bool) []dirEntry {
	filteredDirs := make([]dirEntry, 0)
	for _, d := range root.allDirectories() {
		if predicate(d) {
			filteredDirs = append(filteredDirs, d)
		}
	}

	return filteredDirs
}

func (dir dirEntry) getSize() int {
	size := 0
	for _, c := range dir.files() {
		size += c.size
	}

	return size
}

func (dir dirEntry) totalSize() int {
	size := 0
	for _, d := range dir.allDirectories() {
		size += d.getSize()
	}

	return size
}
