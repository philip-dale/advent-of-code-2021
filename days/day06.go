package days

var day06_sample = []int{3, 4, 3, 1, 2}
var day06_data = []int{2, 1, 1, 4, 4, 1, 3, 4, 2, 4, 2, 1, 1, 4, 3, 5, 1, 1, 5, 1, 1, 5, 4, 5, 4, 1, 5, 1, 3, 1, 4, 2, 3, 2, 1, 2, 5, 5, 2, 3, 1, 2, 3, 3, 1, 4, 3, 1, 1, 1, 1, 5, 2, 1, 1, 1, 5, 3, 3, 2, 1, 4, 1, 1, 1, 3, 1, 1, 5, 5, 1, 4, 4, 4, 4, 5, 1, 5, 1, 1, 5, 5, 2, 2, 5, 4, 1, 5, 4, 1, 4, 1, 1, 1, 1, 5, 3, 2, 4, 1, 1, 1, 4, 4, 1, 2, 1, 1, 5, 2, 1, 1, 1, 4, 4, 4, 4, 3, 3, 1, 1, 5, 1, 5, 2, 1, 4, 1, 2, 4, 4, 4, 4, 2, 2, 2, 4, 4, 4, 2, 1, 5, 5, 2, 1, 1, 1, 4, 4, 1, 4, 2, 3, 3, 3, 3, 3, 5, 4, 1, 5, 1, 4, 5, 5, 1, 1, 1, 4, 1, 2, 4, 4, 1, 2, 3, 3, 3, 3, 5, 1, 4, 2, 5, 5, 2, 1, 1, 1, 1, 3, 3, 1, 1, 2, 3, 2, 5, 4, 2, 1, 1, 2, 2, 2, 1, 3, 1, 5, 4, 1, 1, 5, 3, 3, 2, 2, 3, 1, 1, 1, 1, 2, 4, 2, 2, 5, 1, 2, 4, 2, 1, 1, 3, 2, 5, 5, 3, 1, 3, 3, 1, 4, 1, 1, 5, 5, 1, 5, 4, 1, 1, 1, 1, 2, 3, 3, 1, 2, 3, 1, 5, 1, 3, 1, 1, 3, 1, 1, 1, 1, 1, 1, 5, 1, 1, 5, 5, 2, 1, 1, 5, 2, 4, 5, 5, 1, 1, 5, 1, 5, 5, 1, 1, 3, 3, 1, 1, 3, 1}

type population_info struct {
	timer int
	count int
}

type population_groups struct {
	groups     []population_info
	spawn_time int
	child_time int
}

func create_groups(data []int) population_groups {
	pg := population_groups{
		[]population_info{{1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}},
		7,
		2,
	}

	for _, v := range data {
		for i, g := range pg.groups {
			if g.timer == v {
				pg.groups[i].count++
			}
		}
	}
	return pg
}

func (pg *population_groups) time_step() {
	for i, g := range pg.groups {
		if g.timer == 0 {
			pg.groups = append(pg.groups, population_info{pg.spawn_time + pg.child_time - 1, pg.groups[i].count})
			pg.groups[i].timer = pg.spawn_time - 1
		} else {
			pg.groups[i].timer--
		}

	}
}

func (pg *population_groups) totals() int {
	sum := 0
	for _, g := range pg.groups {
		sum += g.count
	}
	return sum
}

func (pg *population_groups) merge() {
	for {
		merge_found := false
		i_dest := 0
		i_dupe := 0

		for i1, g1 := range pg.groups {
			for i2, g2 := range pg.groups {
				if g1.timer == g2.timer && i1 != i2 {
					merge_found = true
					i_dest = i1
					i_dupe = i2
					break
				}
			}
			if merge_found {
				break
			}
		}

		if merge_found {
			pg.groups[i_dest].count += pg.groups[i_dupe].count
			pg.groups = append(pg.groups[:i_dupe], pg.groups[i_dupe+1:]...)
		} else {
			break
		}
	}
}

func Day06A() int {

	pg := create_groups(day06_data)
	days := 80
	for d := 0; d < days; d++ {
		pg.time_step()
		pg.merge()
	}
	return pg.totals()
}

func Day06B() int {

	pg := create_groups(day06_data)
	days := 256
	for d := 0; d < days; d++ {
		pg.time_step()
		pg.merge()
	}
	return pg.totals()
}
