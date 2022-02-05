package days

type octo_data struct {
	x_size      int
	y_size      int
	energy      [][]int
	has_flashed [][]bool
}

var day11_sample = [][]int{
	{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
	{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
	{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
	{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
	{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
	{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
	{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
	{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
	{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
	{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
}
var day11_data = [][]int{
	{2, 2, 3, 8, 5, 1, 8, 6, 1, 4},
	{4, 5, 5, 2, 3, 8, 8, 5, 5, 3},
	{2, 5, 6, 2, 1, 2, 1, 1, 4, 3},
	{2, 6, 6, 6, 6, 8, 5, 3, 3, 7},
	{7, 5, 7, 5, 5, 1, 8, 7, 8, 4},
	{3, 5, 7, 2, 5, 3, 4, 8, 7, 1},
	{8, 4, 1, 1, 7, 1, 8, 2, 8, 3},
	{7, 7, 4, 2, 6, 6, 8, 3, 8, 5},
	{1, 2, 3, 5, 1, 3, 3, 2, 3, 1},
	{2, 5, 4, 6, 1, 6, 5, 3, 4, 5},
}

func create_octo_data(data [][]int) octo_data {
	od := octo_data{
		len(data),
		len(data[0]),
		data,
		make([][]bool, len(data)),
	}
	for i := range od.has_flashed {
		od.has_flashed[i] = make([]bool, od.y_size)
	}

	return od
}

func (od *octo_data) clear_and_increment() {
	for i, r := range od.energy {
		for j := range r {
			if od.has_flashed[i][j] == true {
				od.has_flashed[i][j] = false
				od.energy[i][j] = 1
			} else {
				od.energy[i][j] += 1
			}
		}
	}
}

func (od *octo_data) flash_cell(x int, y int, inc bool) int {
	flash_count := 0
	if x < 0 || x >= od.x_size || y < 0 || y >= od.y_size {
		return flash_count
	}

	if inc {
		od.energy[x][y] += 1
	}

	if od.has_flashed[x][y] == false && od.energy[x][y] > 9 {
		od.has_flashed[x][y] = true
		flash_count++
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				if i == 0 && j == 0 {
					// done this already
				} else {
					flash_count += od.flash_cell(x+i, y+j, true)
				}
			}
		}
	}
	return flash_count
}

func (od *octo_data) flash_all() int {
	flash_count := 0
	for i, r := range od.energy {
		for j := range r {
			flash_count += od.flash_cell(i, j, false)
		}
	}
	return flash_count
}

func Day11A() int {
	steps := 100
	od := create_octo_data(day11_data)
	flash_count := 0
	for i := 0; i < steps; i++ {
		od.clear_and_increment()
		flash_count += od.flash_all()
	}
	return flash_count
}

func Day11B() int {
	od := create_octo_data(day11_data)
	all_flash := false
	steps := 0
	for !all_flash {
		od.clear_and_increment()
		all_flash = od.flash_all() == 100
		steps++
	}
	return steps
}
