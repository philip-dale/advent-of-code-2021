package days

import (
	"fmt"
	"strings"
)

type cave struct {
	label       string
	connections []string
	is_small    bool
	visited     bool
}

type cave_map struct {
	caves           map[string]cave
	two_visit       string
	two_visit_count int
	paths           [][]string
}

const start_str = "start"
const end_str = "end"

var day12_sample1 = [][]string{
	{"start", "A"},
	{"start", "b"},
	{"A", "c"},
	{"A", "b"},
	{"b", "d"},
	{"A", "end"},
	{"b", "end"},
}

var day12_sample2 = [][]string{
	{"dc", "end"},
	{"HN", "start"},
	{"start", "kj"},
	{"dc", "start"},
	{"dc", "HN"},
	{"LN", "dc"},
	{"HN", "end"},
	{"kj", "sa"},
	{"kj", "HN"},
	{"kj", "dc"},
}

var day12_sample3 = [][]string{
	{"fs", "end"},
	{"he", "DX"},
	{"fs", "he"},
	{"start", "DX"},
	{"pj", "DX"},
	{"end", "zg"},
	{"zg", "sl"},
	{"zg", "pj"},
	{"pj", "he"},
	{"RW", "he"},
	{"fs", "DX"},
	{"pj", "RW"},
	{"zg", "RW"},
	{"start", "pj"},
	{"he", "WI"},
	{"zg", "he"},
	{"pj", "fs"},
	{"start", "RW"},
}

var day12_data = [][]string{
	{"bm", "XY"},
	{"ol", "JS"},
	{"bm", "im"},
	{"RD", "ol"},
	{"bm", "QI"},
	{"JS", "ja"},
	{"im", "gq"},
	{"end", "im"},
	{"ja", "ol"},
	{"JS", "gq"},
	{"bm", "AF"},
	{"RD", "start"},
	{"RD", "ja"},
	{"start", "ol"},
	{"cj", "bm"},
	{"start", "JS"},
	{"AF", "ol"},
	{"end", "QI"},
	{"QI", "gq"},
	{"ja", "gq"},
	{"end", "AF"},
	{"im", "QI"},
	{"bm", "gq"},
	{"ja", "QI"},
	{"gq", "RD"},
}

func (c cave) String() string {
	str := c.label + " - " + fmt.Sprint(c.is_small) + " - ["
	for _, c := range c.connections {
		str += c + ", "
	}
	str += "]"
	return str
}

func create_mapping(data [][]string) cave_map {
	cm := cave_map{
		make(map[string]cave),
		"",
		0,
		[][]string{},
	}

	for _, c := range data {
		entry0, exists0 := cm.caves[c[0]]
		if !exists0 {
			entry0.label = c[0]
			entry0.is_small = (c[0] == strings.ToLower(c[0]))
			entry0.visited = false
		}

		entry1, exists1 := cm.caves[c[1]]
		if !exists1 {
			entry1.label = c[1]
			entry1.is_small = (c[1] == strings.ToLower(c[1]))
			entry1.visited = false
		}

		entry1.connections = append(entry1.connections, entry0.label)
		entry0.connections = append(entry0.connections, entry1.label)
		cm.caves[c[0]] = entry0
		cm.caves[c[1]] = entry1
	}
	return cm
}

func (cm cave_map) path_count() int {
	return len(cm.paths)
}

func (cm *cave_map) set_visited(current cave, state bool) {
	current.visited = state
	cm.caves[current.label] = current
}

func (cm *cave_map) add_path(path []string) {
	for _, p := range cm.paths {
		match := true
		if len(p) != len(path) {
			continue
		} else {
			for i, v := range p {
				if path[i] != v {
					match = false
					break
				}
			}
		}
		if match {
			return
		}
	}

	path_tmp := make([]string, len(path))
	copy(path_tmp, path)
	cm.paths = append(cm.paths, path_tmp)
}

func (cm *cave_map) scan_step(current_str string, path []string) {
	current := cm.caves[current_str]
	if current.is_small {
		if current.visited {
			return
		}
		if current_str == cm.two_visit {
			if cm.two_visit_count == 1 {
				cm.set_visited(current, true)
			}
			cm.two_visit_count++
		} else {
			cm.set_visited(current, true)
		}
	}

	path = append(path, current.label)

	for _, c := range current.connections {
		if c == end_str {
			cm.add_path(path)
		} else {
			cm.scan_step(cm.caves[c].label, path)
		}
	}
	cm.set_visited(current, false)
	if current_str == cm.two_visit {
		cm.two_visit_count--
	}
	return
}

func (cm *cave_map) scan() {
	start := cm.caves[start_str]
	for _, c := range start.connections {
		cm.set_visited(start, true)

		path := []string{start_str}
		cm.scan_step(cm.caves[c].label, path)
	}
	return
}

func (cm cave_map) get_small() (small_caves []string) {
	for _, v := range cm.caves {
		if v.is_small {
			if v.label != start_str && v.label != end_str {
				small_caves = append(small_caves, v.label)
			}
		}
	}
	return small_caves
}

func Day12A() int {
	cm := create_mapping(day12_data)
	cm.scan()
	return cm.path_count()
}

func Day12B() int {
	cm := create_mapping(day12_data)
	small_caves := cm.get_small()

	for _, c := range small_caves {
		cm.two_visit = c
		cm.scan()
	}

	return cm.path_count()
}
