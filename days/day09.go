package days

import "fmt"

var day09_sample = [][]int{
	{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
	{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
	{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
	{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
	{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
}
var day09_data = [][]int{
	{3, 5, 6, 6, 7, 8, 9, 5, 6, 7, 9, 5, 3, 2, 1, 2, 6, 7, 9, 8, 7, 5, 6, 8, 9, 9, 7, 6, 6, 5, 1, 0, 1, 3, 6, 7, 9, 3, 2, 9, 8, 7, 6, 4, 0, 4, 5, 6, 8, 9, 9, 9, 8, 8, 4, 5, 6, 8, 9, 1, 0, 2, 4, 9, 7, 9, 8, 6, 8, 9, 9, 2, 1, 9, 8, 9, 7, 8, 9, 9, 9, 0, 1, 3, 4, 5, 7, 8, 9, 2, 3, 5, 5, 7, 8, 9, 9, 7, 6, 7},
	{2, 6, 7, 5, 8, 9, 5, 4, 5, 6, 8, 9, 4, 3, 3, 4, 5, 9, 8, 7, 6, 4, 9, 9, 9, 8, 6, 5, 4, 3, 2, 1, 2, 4, 5, 6, 7, 9, 1, 9, 9, 9, 5, 3, 2, 3, 7, 8, 9, 4, 9, 8, 7, 6, 5, 7, 7, 9, 9, 9, 1, 9, 9, 8, 6, 5, 4, 5, 7, 8, 8, 9, 9, 7, 6, 6, 6, 7, 8, 9, 8, 1, 2, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 6, 7, 8, 9, 6, 5, 6},
	{1, 2, 3, 4, 9, 9, 6, 3, 4, 5, 7, 8, 9, 4, 4, 5, 6, 7, 9, 9, 8, 9, 8, 9, 9, 9, 7, 6, 5, 4, 3, 5, 3, 4, 6, 7, 8, 9, 9, 8, 9, 9, 9, 4, 6, 4, 5, 6, 9, 3, 2, 9, 8, 7, 6, 8, 9, 9, 9, 8, 9, 8, 9, 7, 6, 4, 3, 6, 5, 9, 7, 8, 9, 9, 5, 4, 5, 6, 9, 8, 7, 6, 4, 5, 6, 8, 9, 9, 1, 0, 1, 2, 3, 7, 8, 9, 1, 2, 3, 5},
	{0, 1, 2, 5, 6, 7, 9, 2, 3, 4, 5, 7, 9, 6, 5, 9, 7, 8, 9, 9, 9, 8, 7, 8, 9, 9, 8, 8, 7, 5, 4, 7, 5, 5, 8, 8, 9, 7, 6, 7, 8, 9, 8, 9, 6, 5, 6, 7, 8, 9, 1, 2, 9, 8, 7, 9, 9, 8, 9, 7, 6, 7, 9, 9, 5, 3, 2, 3, 4, 5, 6, 8, 9, 8, 2, 3, 6, 7, 8, 9, 8, 6, 5, 6, 7, 9, 8, 4, 3, 5, 2, 3, 4, 6, 9, 3, 2, 9, 9, 6},
	{2, 3, 3, 6, 7, 9, 8, 9, 4, 5, 6, 8, 9, 8, 6, 7, 8, 9, 4, 9, 8, 7, 6, 7, 9, 7, 9, 8, 7, 6, 8, 8, 8, 7, 9, 9, 8, 6, 5, 4, 9, 6, 7, 8, 9, 9, 8, 8, 9, 9, 2, 3, 4, 9, 8, 9, 8, 7, 6, 5, 4, 6, 7, 8, 9, 9, 0, 1, 6, 7, 8, 9, 9, 7, 6, 4, 7, 8, 9, 9, 9, 8, 9, 8, 8, 9, 7, 5, 4, 4, 3, 4, 5, 7, 8, 9, 9, 8, 8, 9},
	{3, 5, 7, 8, 9, 8, 7, 9, 9, 9, 7, 9, 5, 9, 7, 8, 9, 3, 2, 9, 7, 6, 5, 8, 9, 6, 5, 9, 8, 7, 9, 9, 9, 8, 9, 9, 9, 5, 4, 3, 4, 5, 6, 7, 9, 9, 9, 9, 9, 8, 9, 4, 9, 8, 9, 9, 9, 8, 5, 4, 3, 4, 5, 7, 7, 8, 9, 2, 7, 8, 9, 9, 9, 8, 7, 6, 7, 9, 9, 9, 8, 9, 6, 9, 9, 8, 7, 6, 5, 9, 5, 6, 6, 7, 9, 8, 9, 6, 7, 9},
	{4, 5, 6, 9, 8, 7, 6, 7, 8, 8, 9, 5, 4, 3, 9, 9, 9, 9, 9, 7, 6, 5, 4, 7, 9, 6, 4, 5, 9, 9, 0, 2, 3, 9, 9, 9, 8, 7, 5, 4, 8, 9, 9, 8, 9, 9, 9, 9, 8, 7, 8, 9, 8, 7, 8, 9, 8, 7, 6, 3, 2, 3, 4, 5, 6, 7, 8, 9, 8, 9, 9, 9, 8, 9, 9, 8, 9, 9, 9, 8, 7, 6, 5, 4, 3, 9, 9, 7, 9, 8, 9, 8, 7, 9, 8, 7, 6, 5, 6, 8},
	{5, 8, 7, 9, 8, 5, 4, 5, 6, 7, 8, 9, 4, 1, 2, 3, 9, 8, 9, 8, 9, 7, 5, 6, 8, 9, 2, 4, 9, 2, 1, 6, 7, 8, 9, 1, 9, 9, 9, 5, 7, 9, 9, 9, 9, 8, 9, 9, 9, 6, 9, 9, 8, 6, 9, 8, 6, 5, 4, 2, 1, 2, 3, 4, 5, 8, 9, 1, 9, 9, 9, 9, 7, 8, 2, 9, 8, 9, 9, 9, 9, 8, 7, 3, 2, 1, 0, 9, 8, 7, 8, 9, 9, 7, 6, 5, 4, 3, 4, 7},
	{6, 9, 9, 8, 6, 4, 2, 4, 5, 8, 9, 6, 3, 2, 3, 9, 8, 7, 6, 9, 9, 8, 7, 7, 9, 2, 1, 9, 7, 9, 3, 5, 9, 9, 9, 3, 9, 8, 7, 6, 7, 8, 9, 3, 9, 7, 6, 8, 9, 5, 6, 9, 6, 5, 4, 9, 9, 7, 6, 4, 2, 3, 4, 5, 6, 9, 1, 0, 1, 9, 8, 7, 5, 2, 1, 5, 6, 8, 9, 3, 4, 9, 8, 8, 3, 3, 1, 9, 7, 6, 8, 9, 9, 8, 7, 6, 5, 6, 5, 6},
	{8, 9, 9, 9, 8, 4, 3, 5, 7, 9, 6, 5, 4, 5, 6, 9, 7, 6, 5, 7, 8, 9, 8, 9, 9, 9, 9, 8, 6, 8, 9, 9, 7, 8, 8, 9, 9, 9, 9, 7, 8, 9, 9, 4, 9, 8, 5, 3, 4, 4, 5, 8, 9, 3, 2, 3, 9, 8, 5, 4, 3, 4, 5, 6, 7, 8, 9, 2, 2, 9, 7, 6, 4, 3, 0, 3, 4, 7, 8, 9, 5, 9, 7, 6, 5, 4, 9, 8, 7, 5, 7, 8, 8, 9, 8, 7, 6, 7, 6, 7},
	{9, 9, 9, 9, 9, 5, 4, 6, 7, 8, 9, 7, 5, 8, 7, 8, 9, 7, 4, 5, 6, 7, 9, 8, 9, 8, 7, 6, 5, 6, 8, 3, 5, 6, 7, 8, 9, 2, 4, 9, 9, 8, 8, 9, 8, 7, 6, 2, 1, 2, 6, 7, 8, 9, 1, 3, 9, 7, 6, 8, 6, 5, 6, 7, 8, 9, 8, 9, 9, 6, 5, 4, 3, 2, 1, 2, 3, 6, 7, 8, 9, 9, 8, 8, 9, 9, 8, 7, 6, 4, 5, 6, 7, 8, 9, 8, 8, 9, 7, 8},
	{9, 9, 8, 9, 9, 7, 5, 6, 7, 9, 9, 8, 6, 9, 8, 9, 8, 9, 3, 4, 5, 9, 8, 7, 8, 9, 7, 6, 4, 2, 1, 2, 4, 8, 9, 9, 3, 1, 4, 9, 8, 7, 7, 8, 9, 6, 5, 3, 2, 3, 5, 6, 9, 5, 3, 9, 9, 8, 7, 8, 9, 6, 8, 9, 9, 4, 6, 9, 8, 7, 7, 5, 5, 3, 2, 5, 4, 5, 6, 7, 8, 9, 9, 9, 8, 7, 6, 6, 4, 3, 6, 3, 5, 9, 9, 9, 9, 9, 9, 9},
	{8, 8, 7, 8, 9, 8, 9, 8, 9, 9, 9, 9, 8, 9, 9, 8, 7, 8, 2, 3, 9, 8, 9, 6, 7, 8, 9, 5, 3, 1, 0, 2, 3, 7, 7, 8, 9, 2, 3, 9, 9, 6, 6, 9, 8, 7, 6, 4, 6, 4, 5, 7, 9, 6, 9, 8, 6, 9, 8, 9, 8, 7, 9, 0, 1, 2, 3, 4, 9, 9, 8, 6, 9, 4, 5, 6, 5, 6, 7, 8, 9, 6, 9, 8, 7, 7, 5, 4, 3, 2, 1, 2, 6, 9, 8, 9, 8, 9, 4, 3},
	{7, 5, 6, 7, 8, 9, 2, 9, 9, 9, 8, 8, 9, 8, 9, 6, 5, 3, 1, 2, 9, 7, 6, 5, 6, 7, 8, 9, 4, 3, 1, 5, 4, 5, 6, 7, 8, 9, 9, 8, 7, 5, 5, 3, 9, 8, 7, 5, 8, 5, 6, 8, 9, 7, 9, 8, 5, 4, 9, 4, 9, 9, 3, 2, 3, 3, 4, 9, 3, 1, 9, 9, 8, 9, 7, 9, 8, 7, 8, 9, 4, 5, 7, 9, 6, 5, 4, 3, 2, 1, 0, 1, 9, 8, 7, 8, 7, 8, 9, 2},
	{6, 4, 5, 4, 6, 9, 3, 4, 9, 8, 7, 6, 6, 7, 8, 9, 3, 2, 0, 1, 9, 8, 9, 6, 7, 9, 9, 8, 6, 9, 8, 6, 5, 6, 8, 9, 9, 8, 7, 6, 5, 3, 4, 2, 3, 9, 8, 7, 8, 9, 7, 8, 9, 9, 8, 7, 6, 3, 2, 3, 8, 9, 9, 5, 4, 4, 9, 8, 9, 2, 3, 9, 7, 9, 9, 9, 9, 8, 9, 3, 2, 3, 9, 8, 7, 6, 9, 5, 4, 2, 3, 2, 3, 9, 6, 7, 6, 8, 9, 0},
	{4, 3, 4, 3, 9, 8, 9, 9, 8, 7, 6, 5, 5, 9, 9, 6, 4, 3, 1, 2, 3, 9, 9, 9, 8, 9, 4, 9, 8, 9, 8, 7, 6, 8, 9, 9, 8, 7, 6, 4, 3, 2, 1, 0, 1, 2, 9, 8, 9, 9, 8, 9, 6, 7, 9, 8, 9, 4, 3, 4, 6, 7, 8, 9, 5, 9, 8, 7, 8, 9, 9, 8, 6, 8, 9, 7, 9, 9, 3, 2, 1, 2, 3, 9, 8, 9, 8, 6, 7, 3, 4, 5, 9, 8, 5, 4, 5, 7, 8, 9},
	{4, 2, 1, 2, 4, 6, 7, 9, 8, 6, 4, 3, 4, 8, 9, 7, 5, 4, 5, 3, 5, 6, 9, 9, 9, 2, 3, 5, 9, 7, 9, 8, 7, 9, 8, 9, 9, 9, 6, 5, 4, 3, 2, 1, 2, 4, 9, 9, 4, 3, 9, 4, 5, 6, 7, 9, 8, 5, 4, 5, 6, 8, 9, 7, 6, 7, 9, 6, 5, 9, 8, 7, 5, 6, 5, 6, 8, 9, 9, 4, 2, 4, 9, 9, 9, 9, 9, 7, 8, 6, 5, 6, 9, 7, 4, 3, 4, 6, 7, 9},
	{4, 1, 0, 1, 3, 4, 7, 8, 9, 7, 3, 2, 3, 7, 8, 9, 6, 5, 7, 4, 6, 7, 8, 9, 2, 1, 2, 3, 4, 6, 7, 9, 8, 9, 7, 9, 9, 8, 7, 7, 6, 5, 6, 7, 3, 4, 7, 8, 9, 2, 1, 2, 3, 7, 8, 9, 7, 6, 5, 6, 9, 9, 9, 8, 7, 9, 6, 5, 4, 9, 8, 6, 4, 3, 4, 5, 6, 7, 8, 9, 4, 9, 8, 9, 8, 9, 9, 8, 9, 8, 6, 9, 8, 6, 5, 2, 3, 4, 5, 9},
	{3, 2, 1, 2, 3, 5, 6, 7, 8, 9, 5, 9, 5, 6, 7, 8, 9, 9, 8, 9, 7, 9, 9, 9, 9, 9, 9, 4, 9, 9, 9, 9, 9, 5, 6, 7, 8, 9, 9, 8, 8, 6, 7, 6, 4, 5, 6, 7, 8, 9, 3, 4, 9, 8, 9, 2, 9, 8, 6, 7, 8, 9, 9, 9, 9, 8, 7, 6, 5, 9, 7, 6, 4, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 5, 7, 9, 9, 9, 2, 9, 7, 9, 8, 7, 5, 3, 4, 6, 7, 8},
	{4, 5, 3, 4, 5, 9, 8, 8, 9, 9, 9, 8, 9, 7, 8, 9, 9, 9, 9, 5, 9, 7, 8, 9, 8, 7, 8, 9, 8, 7, 8, 9, 5, 4, 5, 7, 9, 8, 9, 9, 9, 9, 8, 7, 5, 6, 8, 8, 9, 8, 9, 9, 8, 9, 1, 0, 2, 9, 7, 8, 9, 9, 9, 5, 4, 9, 9, 7, 6, 9, 8, 7, 2, 1, 2, 5, 6, 7, 9, 9, 9, 9, 5, 4, 5, 7, 8, 9, 3, 9, 8, 9, 9, 7, 6, 4, 5, 6, 8, 9},
	{5, 6, 4, 5, 6, 7, 8, 9, 7, 8, 9, 7, 8, 9, 9, 9, 8, 9, 7, 4, 5, 6, 9, 8, 7, 6, 7, 8, 9, 6, 9, 6, 4, 3, 4, 9, 9, 7, 9, 8, 7, 8, 9, 8, 8, 7, 9, 9, 8, 7, 8, 9, 7, 8, 9, 9, 3, 9, 8, 9, 9, 9, 8, 9, 9, 9, 8, 9, 9, 8, 7, 6, 1, 0, 1, 6, 9, 8, 9, 9, 8, 7, 6, 7, 8, 9, 9, 5, 4, 9, 9, 9, 9, 8, 7, 5, 6, 9, 9, 7},
	{9, 8, 9, 6, 7, 8, 9, 6, 5, 4, 7, 6, 9, 9, 9, 9, 7, 8, 5, 3, 4, 9, 8, 7, 6, 5, 1, 3, 4, 5, 7, 9, 6, 5, 9, 8, 7, 6, 8, 9, 6, 7, 8, 9, 9, 8, 9, 9, 7, 6, 7, 5, 6, 8, 9, 8, 9, 9, 9, 9, 9, 9, 6, 9, 8, 7, 7, 8, 9, 9, 8, 4, 3, 1, 2, 7, 8, 9, 8, 8, 9, 8, 7, 8, 9, 3, 2, 9, 9, 8, 6, 8, 9, 9, 9, 8, 7, 8, 9, 6},
	{5, 9, 9, 8, 9, 9, 6, 5, 4, 3, 4, 5, 7, 7, 8, 9, 6, 5, 4, 2, 9, 8, 7, 6, 5, 4, 2, 5, 5, 7, 8, 9, 7, 6, 7, 9, 8, 8, 9, 6, 5, 6, 7, 8, 9, 9, 8, 7, 6, 5, 3, 4, 9, 9, 6, 7, 8, 9, 9, 9, 8, 7, 5, 3, 7, 6, 6, 7, 9, 7, 6, 5, 4, 2, 3, 8, 9, 8, 7, 6, 4, 9, 8, 9, 9, 9, 9, 8, 7, 9, 5, 7, 8, 9, 8, 9, 8, 9, 6, 5},
	{4, 6, 9, 9, 7, 6, 4, 5, 3, 2, 3, 4, 5, 6, 7, 8, 9, 7, 9, 1, 2, 9, 9, 8, 6, 6, 3, 4, 9, 8, 9, 4, 9, 8, 8, 9, 9, 9, 7, 5, 4, 8, 9, 9, 8, 6, 9, 8, 7, 3, 2, 9, 8, 9, 4, 3, 4, 5, 9, 8, 7, 6, 4, 2, 3, 4, 5, 9, 8, 9, 7, 6, 5, 6, 4, 9, 8, 9, 8, 4, 3, 4, 9, 9, 9, 8, 7, 6, 5, 3, 4, 6, 7, 9, 7, 9, 9, 2, 3, 4},
	{2, 9, 8, 7, 6, 4, 3, 2, 1, 0, 1, 5, 6, 7, 8, 9, 9, 9, 8, 9, 3, 4, 8, 9, 9, 8, 5, 6, 8, 9, 2, 3, 4, 9, 9, 8, 7, 6, 8, 4, 3, 4, 9, 8, 6, 5, 6, 9, 6, 4, 9, 8, 7, 8, 9, 2, 3, 4, 5, 9, 6, 5, 2, 1, 2, 3, 5, 6, 7, 8, 9, 9, 6, 7, 5, 6, 7, 8, 9, 6, 4, 6, 7, 9, 8, 7, 6, 5, 4, 2, 3, 2, 3, 5, 6, 7, 8, 9, 6, 5},
	{0, 1, 9, 8, 6, 5, 4, 3, 2, 1, 7, 8, 9, 8, 9, 8, 9, 9, 7, 8, 9, 6, 7, 9, 8, 7, 6, 8, 9, 5, 0, 2, 3, 4, 9, 9, 8, 5, 4, 3, 2, 9, 8, 7, 6, 4, 9, 8, 7, 9, 8, 7, 6, 7, 6, 9, 9, 5, 7, 8, 9, 4, 3, 2, 3, 4, 5, 9, 8, 9, 9, 8, 7, 8, 9, 7, 9, 9, 9, 8, 5, 9, 9, 8, 7, 6, 5, 4, 2, 1, 0, 1, 8, 9, 8, 9, 9, 8, 7, 6},
	{1, 2, 3, 9, 9, 6, 8, 6, 4, 5, 6, 7, 8, 9, 8, 7, 9, 8, 6, 8, 9, 7, 8, 9, 9, 9, 7, 8, 9, 3, 1, 2, 9, 9, 8, 9, 7, 6, 5, 4, 3, 9, 9, 5, 4, 3, 4, 9, 9, 8, 7, 6, 5, 4, 5, 7, 8, 9, 9, 9, 7, 6, 4, 6, 5, 5, 9, 8, 9, 1, 0, 9, 8, 9, 9, 8, 9, 6, 4, 9, 9, 8, 9, 9, 9, 7, 6, 6, 3, 2, 3, 5, 6, 8, 9, 9, 8, 9, 9, 8},
	{5, 3, 4, 9, 8, 9, 8, 7, 5, 7, 9, 8, 9, 9, 7, 6, 5, 4, 5, 7, 8, 9, 9, 3, 4, 5, 9, 9, 8, 9, 2, 9, 8, 9, 7, 8, 9, 7, 7, 5, 9, 8, 7, 6, 3, 2, 9, 8, 9, 7, 6, 5, 4, 3, 4, 5, 9, 9, 8, 9, 8, 9, 6, 8, 6, 9, 8, 7, 8, 9, 9, 8, 9, 7, 9, 9, 7, 6, 5, 9, 8, 7, 8, 9, 9, 8, 7, 8, 5, 4, 5, 6, 7, 9, 3, 5, 7, 9, 8, 9},
	{7, 4, 9, 8, 7, 8, 9, 8, 6, 8, 9, 9, 8, 9, 8, 8, 6, 3, 7, 9, 9, 6, 5, 4, 9, 9, 8, 9, 7, 8, 9, 8, 7, 8, 6, 7, 8, 9, 8, 9, 8, 9, 8, 3, 2, 1, 9, 7, 6, 6, 5, 3, 2, 2, 3, 4, 7, 6, 7, 8, 9, 9, 7, 9, 8, 9, 9, 6, 7, 8, 9, 7, 5, 6, 9, 9, 8, 7, 9, 9, 7, 6, 6, 7, 8, 9, 8, 9, 8, 5, 7, 7, 9, 0, 2, 3, 9, 8, 7, 8},
	{9, 9, 8, 7, 6, 6, 8, 9, 9, 9, 9, 8, 7, 6, 9, 4, 3, 2, 9, 6, 8, 9, 9, 9, 8, 8, 7, 9, 6, 7, 9, 7, 6, 5, 4, 5, 7, 8, 9, 9, 7, 6, 5, 4, 9, 3, 4, 9, 5, 4, 3, 1, 0, 1, 2, 3, 4, 5, 9, 9, 1, 9, 8, 9, 9, 9, 8, 5, 6, 7, 9, 9, 4, 9, 8, 9, 9, 9, 8, 1, 0, 4, 5, 8, 9, 0, 9, 9, 7, 6, 7, 9, 8, 9, 9, 9, 8, 7, 6, 5},
	{9, 8, 9, 8, 5, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 2, 1, 0, 4, 5, 7, 9, 8, 9, 7, 6, 5, 7, 5, 6, 7, 9, 3, 1, 2, 6, 6, 7, 8, 9, 8, 7, 9, 9, 8, 9, 9, 8, 7, 6, 5, 4, 2, 2, 3, 4, 5, 7, 8, 9, 2, 4, 9, 9, 8, 7, 6, 4, 5, 6, 7, 8, 9, 8, 7, 9, 8, 7, 6, 3, 2, 3, 6, 7, 9, 2, 3, 9, 8, 9, 9, 8, 7, 9, 8, 9, 9, 6, 5, 4},
	{8, 7, 6, 7, 3, 4, 7, 9, 9, 4, 9, 8, 7, 9, 6, 9, 2, 2, 3, 4, 5, 6, 7, 8, 9, 5, 4, 3, 4, 5, 9, 8, 9, 2, 3, 4, 5, 8, 9, 9, 9, 9, 8, 7, 7, 9, 8, 9, 8, 6, 5, 4, 3, 3, 6, 5, 6, 7, 9, 5, 4, 9, 3, 1, 9, 6, 4, 3, 4, 5, 6, 9, 8, 7, 6, 7, 9, 6, 5, 4, 3, 4, 5, 8, 9, 3, 4, 5, 9, 9, 9, 7, 6, 8, 6, 9, 8, 9, 2, 3},
	{7, 6, 4, 3, 2, 4, 5, 8, 9, 3, 2, 9, 9, 8, 9, 8, 9, 9, 9, 9, 6, 8, 8, 9, 7, 6, 5, 1, 6, 6, 9, 7, 8, 9, 4, 5, 6, 7, 8, 9, 9, 8, 7, 6, 6, 7, 7, 9, 9, 7, 6, 5, 5, 5, 9, 6, 9, 8, 9, 6, 9, 8, 9, 9, 9, 8, 6, 5, 5, 6, 7, 8, 9, 6, 5, 6, 9, 8, 7, 6, 7, 5, 6, 8, 9, 4, 8, 9, 9, 8, 7, 6, 5, 6, 5, 6, 7, 9, 1, 2},
	{9, 8, 4, 4, 1, 5, 6, 7, 8, 9, 3, 4, 9, 7, 6, 7, 9, 8, 7, 8, 9, 9, 9, 9, 8, 7, 6, 9, 8, 9, 8, 6, 7, 9, 6, 6, 7, 9, 9, 9, 8, 7, 6, 5, 4, 5, 6, 8, 7, 9, 8, 6, 7, 9, 8, 9, 8, 9, 9, 9, 9, 7, 6, 8, 9, 8, 7, 6, 8, 7, 8, 9, 6, 5, 4, 6, 7, 9, 8, 9, 8, 9, 7, 9, 7, 6, 7, 9, 8, 7, 6, 5, 4, 3, 4, 5, 8, 9, 0, 1},
	{9, 9, 2, 1, 0, 7, 8, 8, 9, 5, 4, 9, 7, 6, 5, 9, 8, 7, 6, 7, 7, 8, 9, 9, 9, 8, 9, 9, 9, 8, 7, 5, 6, 8, 9, 7, 8, 9, 1, 0, 9, 8, 7, 6, 3, 2, 4, 5, 6, 9, 9, 8, 9, 8, 7, 6, 7, 8, 9, 8, 9, 6, 5, 6, 7, 9, 8, 7, 9, 8, 9, 9, 8, 6, 3, 4, 5, 6, 9, 9, 9, 9, 8, 9, 8, 9, 8, 9, 9, 8, 7, 6, 5, 2, 3, 4, 5, 8, 9, 2},
	{8, 7, 6, 2, 1, 3, 5, 9, 9, 6, 5, 6, 9, 4, 4, 3, 1, 2, 4, 5, 6, 7, 8, 9, 9, 9, 8, 9, 8, 7, 6, 4, 6, 7, 8, 9, 9, 9, 9, 9, 7, 9, 8, 3, 2, 1, 2, 6, 9, 8, 9, 9, 9, 7, 6, 5, 6, 7, 8, 7, 8, 9, 4, 3, 8, 9, 9, 9, 9, 9, 7, 9, 8, 7, 6, 5, 6, 7, 8, 9, 3, 4, 9, 9, 9, 8, 9, 3, 2, 9, 8, 7, 9, 3, 4, 5, 6, 7, 8, 9},
	{9, 8, 5, 4, 3, 4, 5, 7, 8, 9, 9, 9, 8, 3, 2, 1, 0, 1, 2, 4, 5, 6, 9, 9, 8, 7, 6, 9, 7, 6, 5, 3, 2, 6, 7, 9, 8, 9, 8, 7, 6, 5, 4, 2, 1, 0, 1, 9, 8, 7, 8, 7, 8, 9, 8, 4, 3, 4, 5, 6, 9, 9, 4, 2, 9, 9, 6, 9, 8, 7, 6, 5, 9, 8, 8, 6, 7, 8, 9, 9, 2, 3, 5, 9, 8, 7, 8, 9, 1, 0, 9, 6, 5, 4, 5, 6, 7, 9, 9, 6},
	{8, 7, 6, 5, 6, 5, 6, 8, 9, 9, 8, 7, 6, 5, 3, 2, 1, 2, 6, 5, 6, 7, 8, 9, 9, 8, 4, 9, 9, 5, 4, 2, 0, 4, 5, 6, 7, 8, 9, 9, 9, 6, 5, 3, 2, 1, 9, 8, 7, 6, 5, 6, 7, 8, 9, 2, 1, 3, 4, 6, 7, 8, 9, 9, 7, 8, 5, 6, 9, 8, 5, 4, 3, 9, 9, 7, 8, 9, 6, 8, 9, 9, 9, 7, 6, 6, 7, 8, 9, 9, 8, 7, 9, 8, 7, 7, 8, 9, 3, 5},
	{9, 8, 7, 7, 8, 9, 7, 9, 8, 7, 9, 8, 7, 6, 7, 3, 2, 3, 4, 8, 7, 8, 9, 1, 2, 9, 9, 8, 7, 6, 5, 3, 2, 3, 4, 7, 8, 9, 8, 9, 8, 9, 6, 4, 3, 2, 3, 9, 8, 7, 6, 7, 8, 9, 6, 3, 2, 3, 4, 5, 6, 9, 9, 7, 6, 5, 4, 9, 8, 7, 6, 8, 5, 7, 9, 8, 9, 4, 5, 6, 9, 8, 9, 6, 4, 5, 6, 7, 9, 2, 9, 8, 9, 9, 8, 9, 9, 9, 2, 1},
	{6, 9, 9, 8, 9, 9, 8, 9, 7, 6, 8, 9, 8, 9, 9, 8, 5, 4, 9, 9, 9, 9, 8, 9, 9, 8, 9, 9, 9, 7, 6, 5, 3, 4, 8, 8, 9, 8, 7, 5, 7, 9, 6, 5, 6, 3, 5, 6, 9, 8, 7, 8, 9, 7, 5, 4, 3, 4, 5, 6, 7, 8, 9, 6, 5, 4, 3, 2, 9, 9, 8, 7, 6, 8, 9, 9, 4, 3, 4, 9, 9, 7, 8, 5, 3, 5, 9, 8, 9, 1, 9, 9, 1, 0, 9, 7, 9, 8, 9, 2},
	{4, 5, 9, 9, 6, 5, 9, 9, 8, 5, 4, 4, 9, 9, 8, 7, 6, 7, 8, 9, 4, 5, 7, 9, 8, 7, 8, 9, 9, 8, 7, 6, 4, 5, 6, 7, 8, 9, 6, 4, 8, 9, 7, 6, 7, 8, 9, 7, 8, 9, 8, 9, 8, 7, 6, 5, 4, 5, 6, 7, 8, 9, 7, 6, 4, 3, 1, 0, 1, 2, 9, 9, 7, 9, 6, 5, 4, 2, 9, 8, 7, 6, 2, 1, 2, 9, 8, 9, 9, 9, 8, 5, 3, 2, 3, 5, 6, 7, 8, 9},
	{6, 9, 7, 6, 5, 4, 5, 9, 7, 6, 2, 3, 2, 3, 9, 8, 7, 8, 9, 4, 3, 2, 9, 8, 7, 6, 7, 9, 8, 9, 8, 7, 6, 8, 7, 8, 9, 9, 8, 5, 6, 7, 9, 9, 8, 9, 9, 9, 9, 4, 9, 1, 9, 8, 7, 6, 5, 7, 7, 8, 9, 9, 9, 8, 6, 7, 9, 1, 9, 3, 9, 9, 9, 9, 8, 6, 2, 1, 0, 9, 9, 5, 3, 0, 4, 8, 7, 8, 9, 8, 7, 6, 5, 4, 8, 9, 8, 9, 9, 1},
	{9, 8, 9, 8, 6, 7, 9, 9, 8, 4, 1, 0, 1, 2, 3, 9, 8, 9, 4, 3, 2, 0, 9, 7, 6, 5, 4, 5, 7, 8, 9, 8, 7, 9, 9, 9, 1, 9, 7, 6, 8, 9, 9, 9, 9, 9, 7, 8, 9, 2, 1, 0, 1, 9, 9, 8, 7, 9, 9, 9, 7, 8, 9, 9, 7, 9, 8, 9, 7, 9, 8, 9, 3, 2, 9, 7, 3, 9, 9, 8, 7, 3, 2, 1, 3, 4, 6, 7, 8, 9, 8, 7, 8, 5, 7, 8, 9, 9, 9, 0},
	{9, 7, 8, 9, 7, 9, 8, 6, 4, 3, 2, 1, 2, 4, 4, 9, 9, 6, 5, 4, 4, 9, 8, 6, 5, 4, 3, 4, 5, 9, 9, 9, 8, 9, 2, 1, 0, 9, 8, 7, 9, 9, 9, 8, 9, 7, 6, 9, 8, 9, 3, 1, 2, 4, 9, 9, 9, 5, 6, 7, 6, 7, 8, 9, 9, 8, 7, 6, 6, 8, 7, 8, 9, 1, 9, 8, 9, 8, 9, 9, 6, 5, 3, 2, 5, 6, 7, 9, 9, 4, 9, 8, 9, 6, 8, 9, 9, 9, 8, 9},
	{8, 6, 9, 8, 9, 8, 7, 6, 5, 4, 3, 2, 3, 4, 9, 8, 9, 7, 9, 5, 6, 7, 9, 9, 7, 3, 2, 3, 4, 7, 8, 9, 9, 4, 3, 9, 1, 2, 9, 8, 9, 8, 7, 6, 5, 3, 5, 6, 7, 9, 9, 2, 3, 9, 8, 7, 6, 4, 5, 4, 5, 8, 9, 8, 9, 7, 6, 5, 5, 5, 6, 7, 8, 9, 4, 9, 9, 7, 9, 8, 7, 6, 5, 6, 6, 7, 8, 9, 9, 2, 3, 9, 8, 7, 9, 8, 9, 7, 6, 7},
	{7, 5, 7, 7, 8, 9, 8, 7, 6, 5, 4, 3, 4, 9, 8, 7, 6, 9, 8, 9, 8, 9, 9, 8, 6, 5, 1, 0, 5, 6, 7, 8, 9, 5, 9, 8, 9, 9, 9, 9, 9, 8, 9, 7, 4, 2, 4, 5, 5, 6, 8, 9, 9, 9, 9, 9, 7, 3, 1, 3, 6, 9, 5, 7, 9, 6, 5, 4, 3, 4, 5, 9, 9, 1, 2, 9, 8, 6, 4, 9, 8, 7, 8, 9, 8, 8, 9, 9, 8, 1, 2, 4, 9, 8, 9, 7, 7, 9, 4, 6},
	{4, 3, 4, 6, 5, 6, 9, 9, 8, 6, 5, 6, 5, 9, 9, 6, 5, 6, 7, 8, 9, 9, 9, 6, 5, 4, 3, 2, 4, 5, 6, 9, 9, 9, 9, 7, 9, 8, 9, 9, 9, 7, 6, 5, 4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 3, 4, 9, 2, 0, 1, 2, 3, 4, 9, 9, 9, 4, 3, 2, 5, 9, 8, 9, 2, 3, 9, 7, 5, 3, 4, 9, 8, 9, 9, 9, 9, 7, 5, 7, 0, 1, 4, 6, 9, 7, 6, 5, 4, 3, 5},
	{3, 2, 6, 5, 4, 5, 5, 6, 9, 7, 6, 7, 9, 8, 7, 5, 4, 5, 6, 8, 9, 9, 8, 7, 6, 5, 6, 3, 5, 6, 7, 8, 9, 8, 7, 6, 7, 7, 8, 9, 8, 7, 6, 4, 2, 0, 1, 2, 5, 7, 8, 9, 9, 3, 2, 9, 8, 7, 3, 2, 4, 5, 9, 8, 9, 8, 9, 2, 0, 2, 6, 7, 8, 9, 9, 8, 5, 4, 2, 3, 4, 9, 8, 9, 9, 9, 7, 4, 2, 1, 2, 3, 5, 9, 8, 8, 9, 5, 7, 6},
	{2, 1, 0, 1, 2, 3, 4, 9, 9, 8, 7, 9, 8, 7, 5, 4, 3, 4, 7, 6, 8, 9, 9, 8, 9, 6, 7, 5, 6, 7, 8, 9, 8, 7, 6, 4, 5, 6, 9, 8, 7, 6, 5, 4, 3, 1, 3, 4, 5, 6, 7, 8, 9, 2, 1, 9, 8, 6, 4, 3, 4, 9, 8, 7, 5, 7, 8, 9, 1, 4, 5, 6, 8, 9, 8, 7, 3, 2, 1, 2, 3, 5, 7, 9, 9, 8, 6, 5, 3, 4, 3, 4, 5, 6, 9, 9, 8, 9, 8, 8},
	{3, 2, 1, 2, 6, 4, 9, 8, 9, 9, 8, 9, 9, 6, 5, 4, 2, 3, 6, 5, 7, 8, 9, 9, 8, 7, 9, 9, 7, 8, 9, 7, 9, 8, 5, 3, 4, 5, 7, 9, 8, 7, 6, 5, 4, 2, 7, 6, 7, 8, 9, 9, 5, 1, 0, 3, 9, 7, 6, 7, 5, 9, 9, 6, 4, 6, 7, 9, 2, 5, 6, 7, 9, 9, 6, 5, 4, 3, 0, 1, 2, 4, 6, 7, 8, 9, 7, 6, 6, 5, 9, 7, 6, 7, 9, 8, 7, 8, 9, 9},
	{5, 4, 3, 4, 5, 9, 8, 7, 9, 9, 9, 8, 7, 6, 3, 2, 1, 2, 5, 4, 6, 7, 8, 9, 9, 8, 9, 9, 8, 9, 7, 6, 5, 4, 3, 2, 5, 6, 7, 8, 9, 9, 8, 6, 5, 9, 8, 7, 8, 9, 9, 6, 4, 2, 1, 2, 9, 8, 7, 8, 9, 8, 9, 4, 3, 4, 8, 8, 9, 7, 7, 8, 9, 8, 7, 6, 5, 7, 6, 4, 3, 4, 6, 7, 9, 9, 8, 9, 8, 9, 8, 9, 7, 9, 8, 7, 6, 7, 8, 9},
	{6, 5, 4, 9, 9, 8, 7, 6, 7, 9, 8, 7, 6, 5, 4, 5, 0, 1, 2, 3, 4, 5, 9, 9, 9, 9, 9, 8, 9, 8, 9, 8, 6, 5, 4, 5, 6, 7, 8, 9, 9, 9, 9, 7, 7, 8, 9, 8, 9, 9, 8, 7, 6, 9, 2, 3, 4, 9, 8, 9, 8, 7, 8, 9, 2, 5, 6, 7, 8, 9, 8, 9, 6, 9, 8, 7, 7, 9, 7, 6, 7, 8, 9, 9, 9, 8, 9, 9, 9, 8, 7, 8, 9, 9, 7, 6, 5, 4, 6, 6},
	{7, 6, 9, 8, 8, 9, 6, 5, 6, 7, 9, 8, 7, 6, 5, 2, 1, 2, 3, 6, 5, 6, 7, 8, 9, 9, 7, 7, 6, 7, 8, 9, 7, 6, 5, 6, 7, 8, 9, 8, 7, 8, 9, 9, 9, 9, 0, 9, 7, 8, 9, 8, 9, 8, 9, 6, 5, 6, 9, 8, 9, 6, 9, 8, 9, 6, 7, 8, 9, 7, 9, 9, 5, 4, 9, 8, 9, 4, 9, 8, 9, 9, 8, 7, 8, 7, 8, 9, 8, 7, 5, 9, 8, 7, 6, 5, 4, 3, 5, 5},
	{9, 9, 8, 7, 6, 2, 3, 4, 5, 6, 7, 9, 9, 7, 4, 3, 2, 4, 5, 7, 7, 8, 8, 9, 8, 8, 6, 7, 5, 4, 5, 9, 8, 7, 8, 7, 8, 9, 9, 8, 6, 7, 8, 7, 8, 9, 1, 5, 6, 9, 8, 9, 8, 7, 8, 9, 6, 9, 7, 6, 8, 4, 5, 6, 9, 9, 8, 9, 7, 5, 7, 8, 9, 3, 4, 9, 4, 2, 1, 9, 5, 6, 9, 6, 5, 6, 9, 7, 6, 5, 4, 5, 9, 9, 7, 7, 6, 2, 3, 4},
	{8, 9, 7, 6, 5, 1, 2, 3, 4, 6, 9, 8, 7, 6, 5, 4, 3, 5, 6, 9, 8, 9, 9, 8, 7, 6, 5, 4, 4, 3, 4, 6, 9, 9, 9, 8, 9, 6, 9, 7, 5, 4, 7, 6, 7, 9, 3, 4, 5, 9, 7, 6, 9, 6, 7, 8, 9, 9, 6, 5, 4, 3, 4, 5, 7, 8, 9, 6, 5, 4, 2, 4, 8, 9, 9, 6, 5, 3, 2, 3, 4, 9, 8, 7, 6, 7, 8, 9, 5, 2, 3, 4, 9, 8, 7, 6, 5, 4, 5, 5},
	{7, 6, 5, 4, 4, 0, 1, 4, 6, 7, 9, 9, 9, 7, 6, 7, 4, 6, 7, 8, 9, 3, 3, 9, 8, 7, 6, 3, 1, 2, 3, 4, 9, 8, 7, 9, 8, 7, 9, 8, 5, 3, 4, 5, 6, 8, 9, 5, 9, 8, 7, 5, 6, 5, 6, 9, 8, 7, 8, 6, 5, 2, 5, 7, 8, 9, 9, 7, 3, 2, 0, 5, 6, 7, 8, 9, 9, 4, 5, 6, 5, 6, 9, 8, 7, 8, 9, 4, 3, 1, 2, 3, 4, 9, 8, 7, 6, 8, 7, 8},
	{8, 7, 4, 3, 2, 1, 2, 3, 7, 8, 9, 9, 9, 8, 9, 9, 8, 7, 9, 9, 3, 2, 1, 0, 9, 8, 5, 4, 0, 1, 2, 9, 9, 8, 6, 7, 9, 9, 8, 7, 3, 2, 3, 5, 7, 8, 9, 9, 8, 7, 6, 4, 3, 4, 5, 9, 7, 6, 7, 7, 2, 1, 3, 6, 6, 7, 8, 9, 4, 3, 1, 6, 7, 8, 9, 9, 8, 9, 6, 7, 6, 7, 8, 9, 8, 9, 5, 4, 2, 0, 1, 9, 5, 6, 9, 8, 7, 8, 8, 9},
	{9, 8, 5, 4, 4, 3, 4, 4, 5, 6, 7, 9, 9, 9, 7, 5, 9, 8, 9, 7, 9, 9, 3, 4, 9, 7, 5, 4, 1, 2, 9, 8, 7, 6, 5, 9, 8, 7, 6, 5, 4, 3, 5, 7, 8, 9, 9, 9, 9, 8, 7, 9, 1, 0, 9, 8, 6, 5, 3, 2, 1, 0, 2, 4, 5, 9, 9, 9, 5, 5, 3, 8, 9, 9, 9, 8, 7, 9, 7, 8, 9, 8, 9, 4, 9, 9, 6, 9, 3, 1, 9, 8, 9, 9, 9, 9, 8, 9, 9, 9},
	{8, 7, 6, 6, 7, 8, 9, 5, 7, 8, 9, 9, 9, 6, 5, 4, 3, 9, 4, 6, 9, 8, 9, 5, 9, 8, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 9, 9, 8, 9, 5, 4, 6, 9, 9, 9, 8, 9, 9, 9, 9, 8, 9, 1, 3, 9, 7, 5, 4, 5, 3, 2, 4, 5, 6, 8, 9, 8, 6, 7, 8, 9, 9, 9, 8, 9, 6, 8, 9, 9, 9, 9, 4, 3, 9, 8, 7, 8, 4, 9, 8, 6, 7, 8, 9, 9, 9, 6, 7, 8},
	{9, 8, 9, 8, 9, 9, 7, 6, 7, 9, 9, 9, 8, 9, 7, 9, 2, 1, 3, 9, 9, 7, 8, 9, 9, 6, 5, 4, 3, 4, 9, 6, 5, 4, 3, 2, 3, 4, 9, 8, 7, 5, 7, 8, 9, 8, 7, 8, 9, 9, 9, 7, 9, 2, 9, 9, 8, 9, 9, 6, 4, 3, 6, 6, 7, 9, 5, 9, 7, 8, 9, 3, 9, 8, 7, 6, 5, 7, 9, 9, 8, 9, 9, 2, 3, 9, 8, 9, 9, 8, 7, 5, 6, 9, 8, 9, 7, 5, 3, 2},
	{9, 9, 8, 9, 6, 9, 8, 7, 8, 9, 9, 8, 7, 6, 5, 4, 3, 3, 9, 8, 9, 6, 7, 9, 8, 7, 6, 5, 7, 9, 8, 7, 6, 6, 4, 3, 7, 5, 6, 9, 9, 6, 8, 9, 3, 5, 6, 7, 9, 9, 7, 6, 7, 9, 8, 9, 9, 9, 8, 7, 8, 4, 7, 9, 8, 9, 4, 5, 9, 9, 1, 2, 3, 9, 8, 9, 4, 6, 8, 7, 6, 7, 8, 9, 4, 6, 9, 8, 6, 7, 5, 4, 5, 8, 7, 8, 9, 4, 2, 1},
	{9, 8, 7, 6, 5, 4, 9, 8, 9, 2, 3, 9, 8, 7, 6, 9, 4, 9, 8, 7, 6, 5, 6, 8, 9, 9, 8, 6, 7, 8, 9, 9, 7, 6, 5, 4, 5, 8, 7, 9, 8, 7, 9, 5, 4, 6, 9, 8, 9, 8, 7, 5, 7, 8, 7, 8, 9, 9, 9, 9, 6, 5, 8, 9, 9, 6, 5, 7, 8, 9, 0, 9, 4, 9, 9, 7, 3, 9, 8, 6, 5, 6, 7, 8, 9, 9, 8, 6, 4, 3, 6, 3, 5, 8, 6, 7, 9, 5, 1, 0},
	{9, 9, 9, 5, 4, 3, 4, 9, 2, 1, 2, 3, 9, 8, 9, 8, 9, 9, 8, 6, 5, 4, 7, 8, 7, 8, 9, 8, 9, 9, 7, 6, 9, 8, 9, 7, 6, 7, 8, 9, 9, 8, 9, 6, 5, 9, 8, 9, 9, 7, 7, 4, 6, 5, 6, 8, 9, 9, 9, 8, 7, 6, 9, 7, 9, 7, 6, 7, 9, 9, 9, 8, 9, 8, 7, 5, 2, 3, 4, 4, 4, 5, 6, 9, 8, 9, 5, 4, 3, 2, 1, 2, 4, 6, 5, 7, 8, 9, 2, 1},
	{9, 8, 7, 6, 5, 4, 5, 6, 9, 0, 1, 2, 3, 9, 8, 7, 9, 8, 7, 5, 4, 3, 4, 5, 6, 7, 9, 9, 4, 3, 4, 5, 6, 9, 9, 8, 9, 8, 9, 0, 1, 9, 9, 8, 9, 8, 7, 9, 8, 6, 5, 3, 2, 4, 5, 9, 9, 8, 6, 9, 8, 9, 7, 6, 8, 9, 7, 9, 9, 9, 8, 7, 7, 9, 9, 3, 1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 6, 2, 1, 0, 1, 2, 3, 4, 6, 7, 8, 9, 2},
	{0, 9, 9, 8, 7, 6, 6, 7, 8, 9, 4, 3, 9, 8, 9, 6, 9, 8, 7, 6, 5, 2, 3, 4, 5, 6, 7, 9, 3, 2, 3, 6, 9, 8, 9, 9, 9, 9, 7, 9, 9, 8, 8, 9, 8, 7, 6, 4, 9, 8, 7, 4, 0, 6, 7, 8, 9, 3, 5, 6, 9, 7, 6, 5, 7, 9, 8, 9, 9, 8, 9, 6, 6, 7, 8, 9, 2, 4, 5, 6, 4, 5, 7, 9, 8, 9, 5, 4, 3, 2, 1, 2, 3, 4, 6, 7, 9, 9, 9, 3},
	{1, 2, 3, 9, 9, 7, 7, 8, 9, 6, 5, 9, 8, 7, 6, 5, 7, 9, 8, 7, 6, 1, 4, 5, 7, 9, 8, 9, 2, 1, 2, 9, 9, 7, 8, 9, 7, 6, 5, 8, 7, 6, 7, 9, 8, 6, 5, 3, 4, 9, 6, 4, 2, 4, 9, 9, 3, 2, 4, 6, 9, 9, 5, 4, 3, 1, 9, 9, 8, 7, 8, 4, 5, 9, 7, 8, 9, 6, 6, 9, 5, 6, 7, 8, 9, 7, 6, 5, 5, 4, 3, 3, 7, 6, 8, 9, 6, 7, 8, 9},
	{3, 3, 9, 9, 7, 9, 8, 9, 9, 9, 9, 7, 7, 5, 5, 4, 5, 6, 9, 8, 9, 2, 5, 6, 7, 8, 9, 4, 1, 0, 9, 8, 8, 6, 9, 9, 7, 6, 4, 3, 4, 5, 7, 9, 5, 4, 3, 2, 3, 9, 7, 5, 9, 5, 7, 8, 9, 1, 3, 9, 8, 7, 6, 5, 4, 3, 4, 5, 9, 6, 5, 3, 4, 7, 6, 7, 8, 9, 7, 8, 8, 7, 9, 9, 0, 9, 8, 7, 7, 5, 4, 4, 5, 7, 8, 9, 5, 9, 9, 2},
	{4, 9, 8, 7, 6, 6, 9, 4, 9, 8, 7, 6, 5, 4, 3, 2, 4, 9, 8, 7, 7, 3, 4, 5, 9, 9, 4, 3, 2, 9, 8, 7, 6, 5, 7, 8, 9, 4, 3, 2, 1, 2, 9, 8, 6, 5, 4, 3, 4, 5, 9, 9, 8, 7, 8, 9, 1, 0, 1, 2, 9, 9, 7, 6, 5, 4, 5, 6, 9, 9, 7, 2, 1, 4, 5, 6, 9, 9, 8, 9, 9, 9, 3, 2, 1, 5, 9, 9, 8, 6, 5, 5, 9, 8, 9, 9, 9, 8, 9, 3},
	{5, 7, 9, 6, 4, 4, 2, 3, 4, 9, 8, 7, 4, 3, 2, 1, 9, 8, 7, 6, 5, 4, 7, 6, 7, 8, 9, 4, 9, 9, 8, 6, 5, 4, 5, 9, 8, 9, 4, 3, 0, 1, 9, 8, 7, 6, 5, 4, 5, 6, 8, 9, 9, 8, 9, 3, 2, 1, 2, 3, 4, 9, 8, 9, 6, 8, 6, 9, 8, 7, 6, 3, 2, 3, 4, 5, 6, 7, 9, 7, 6, 5, 4, 5, 2, 4, 6, 9, 8, 7, 6, 6, 7, 9, 9, 9, 9, 7, 8, 9},
	{9, 9, 8, 7, 3, 2, 1, 2, 9, 8, 9, 8, 5, 4, 3, 0, 1, 9, 9, 8, 7, 5, 6, 7, 8, 9, 9, 9, 8, 7, 6, 5, 4, 3, 4, 9, 7, 9, 9, 2, 1, 2, 4, 9, 8, 9, 6, 5, 6, 7, 9, 9, 9, 9, 7, 5, 3, 2, 3, 4, 6, 8, 9, 8, 7, 8, 7, 9, 9, 6, 5, 4, 3, 4, 5, 7, 7, 9, 9, 9, 9, 6, 5, 6, 9, 5, 6, 9, 9, 8, 7, 8, 9, 7, 9, 8, 7, 6, 7, 8},
	{8, 9, 9, 8, 5, 3, 2, 3, 9, 7, 6, 5, 4, 3, 2, 1, 2, 3, 4, 9, 8, 9, 7, 8, 9, 5, 8, 9, 9, 8, 7, 8, 3, 2, 9, 8, 6, 8, 8, 9, 4, 3, 4, 5, 9, 8, 7, 7, 9, 8, 9, 9, 8, 9, 8, 6, 5, 3, 4, 5, 7, 9, 3, 9, 9, 9, 8, 9, 8, 7, 6, 6, 8, 5, 6, 8, 8, 9, 9, 7, 8, 9, 9, 9, 8, 9, 7, 8, 9, 9, 8, 9, 4, 6, 9, 8, 6, 5, 8, 3},
	{6, 9, 8, 7, 5, 4, 5, 4, 5, 9, 8, 6, 5, 4, 3, 9, 3, 5, 6, 9, 9, 9, 8, 9, 2, 3, 7, 8, 9, 9, 8, 5, 4, 3, 9, 6, 5, 6, 7, 8, 9, 4, 5, 9, 8, 9, 9, 8, 9, 9, 5, 9, 7, 8, 9, 8, 6, 4, 5, 7, 9, 3, 2, 1, 0, 1, 9, 1, 9, 8, 7, 8, 9, 6, 8, 9, 9, 4, 7, 6, 7, 9, 8, 7, 6, 7, 9, 9, 9, 9, 9, 4, 3, 9, 8, 7, 5, 4, 3, 2},
	{4, 3, 9, 8, 7, 5, 6, 7, 6, 7, 9, 8, 7, 8, 9, 8, 9, 6, 7, 8, 9, 9, 9, 4, 3, 4, 6, 9, 9, 9, 8, 7, 9, 9, 8, 6, 4, 6, 9, 9, 6, 5, 9, 8, 7, 5, 6, 9, 7, 9, 4, 3, 6, 7, 9, 8, 7, 5, 6, 8, 9, 4, 3, 2, 3, 2, 9, 0, 9, 9, 8, 9, 8, 7, 9, 5, 4, 3, 3, 5, 9, 7, 9, 9, 7, 8, 9, 9, 9, 8, 8, 9, 2, 1, 9, 9, 4, 3, 2, 1},
	{9, 2, 3, 9, 7, 6, 7, 8, 9, 8, 9, 9, 8, 9, 8, 7, 8, 9, 8, 9, 9, 8, 7, 6, 4, 5, 9, 8, 7, 9, 9, 9, 8, 6, 5, 4, 3, 7, 8, 9, 7, 6, 9, 8, 6, 4, 7, 5, 6, 8, 9, 2, 4, 6, 7, 9, 8, 8, 7, 8, 9, 5, 5, 3, 4, 5, 8, 9, 8, 7, 9, 6, 9, 8, 9, 3, 2, 1, 2, 3, 4, 5, 6, 9, 8, 9, 9, 8, 7, 7, 7, 8, 9, 9, 8, 7, 6, 4, 3, 4},
	{8, 9, 5, 9, 8, 7, 9, 9, 4, 9, 7, 9, 9, 8, 9, 6, 9, 9, 9, 8, 7, 9, 8, 7, 6, 9, 8, 7, 6, 9, 7, 8, 9, 7, 6, 7, 4, 6, 7, 8, 9, 9, 8, 7, 5, 2, 2, 4, 5, 7, 9, 1, 3, 4, 5, 7, 9, 9, 8, 9, 9, 9, 6, 4, 5, 6, 7, 9, 9, 6, 5, 4, 6, 9, 5, 4, 3, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9, 9, 6, 5, 6, 7, 8, 9, 9, 8, 9, 6, 4, 5},
	{7, 8, 9, 9, 9, 8, 9, 2, 3, 4, 6, 9, 8, 7, 8, 5, 8, 7, 8, 9, 6, 5, 9, 9, 8, 9, 9, 6, 5, 4, 6, 9, 9, 8, 8, 6, 5, 6, 7, 9, 2, 1, 9, 5, 4, 1, 0, 2, 4, 6, 8, 9, 8, 5, 9, 8, 9, 2, 9, 9, 9, 8, 9, 5, 6, 7, 9, 7, 5, 4, 3, 2, 3, 5, 9, 8, 4, 3, 5, 6, 9, 7, 9, 9, 9, 8, 9, 8, 7, 4, 3, 2, 3, 4, 6, 9, 8, 7, 5, 7},
	{6, 9, 8, 8, 9, 9, 5, 3, 4, 5, 9, 8, 7, 6, 5, 4, 6, 6, 7, 8, 9, 4, 3, 1, 9, 9, 9, 8, 4, 3, 4, 8, 6, 9, 9, 8, 6, 7, 8, 9, 9, 9, 7, 4, 3, 2, 1, 2, 3, 4, 9, 9, 7, 6, 7, 8, 9, 1, 2, 9, 8, 7, 8, 9, 7, 9, 8, 7, 6, 5, 4, 3, 4, 9, 8, 7, 6, 4, 6, 7, 8, 8, 9, 4, 8, 7, 9, 9, 8, 6, 5, 4, 5, 9, 7, 8, 9, 8, 7, 8},
	{5, 5, 6, 7, 9, 7, 6, 9, 8, 6, 8, 9, 6, 5, 4, 3, 4, 5, 6, 7, 9, 6, 1, 0, 1, 9, 8, 9, 3, 2, 2, 6, 4, 6, 9, 8, 7, 9, 9, 9, 8, 7, 6, 5, 4, 3, 2, 3, 4, 5, 7, 8, 9, 7, 9, 9, 1, 0, 9, 8, 9, 6, 9, 8, 9, 9, 9, 8, 7, 6, 5, 4, 5, 6, 9, 8, 7, 5, 7, 8, 9, 9, 2, 3, 5, 6, 8, 9, 8, 7, 9, 8, 7, 8, 9, 9, 9, 9, 8, 9},
	{3, 4, 7, 8, 9, 8, 7, 8, 9, 7, 9, 8, 9, 5, 5, 2, 3, 6, 7, 9, 9, 7, 4, 2, 9, 8, 7, 8, 9, 1, 0, 2, 3, 5, 6, 9, 8, 9, 8, 9, 9, 8, 9, 6, 5, 4, 3, 4, 5, 6, 8, 9, 9, 8, 9, 4, 2, 9, 8, 7, 8, 5, 6, 7, 8, 9, 9, 9, 8, 9, 8, 5, 6, 7, 8, 9, 8, 9, 8, 9, 1, 0, 1, 4, 9, 7, 9, 3, 9, 9, 9, 9, 8, 9, 9, 8, 9, 9, 9, 4},
	{1, 5, 6, 7, 8, 9, 8, 9, 3, 9, 8, 7, 8, 4, 2, 1, 0, 5, 8, 7, 8, 9, 6, 9, 8, 6, 6, 7, 9, 2, 1, 2, 3, 4, 5, 6, 9, 8, 7, 5, 2, 9, 8, 7, 6, 5, 4, 6, 8, 7, 8, 9, 9, 9, 6, 5, 6, 7, 9, 6, 8, 3, 4, 5, 6, 8, 9, 9, 9, 8, 7, 6, 7, 9, 9, 8, 9, 9, 9, 3, 2, 1, 9, 9, 8, 9, 3, 2, 3, 4, 9, 9, 9, 5, 6, 7, 8, 9, 4, 3},
	{2, 3, 4, 9, 8, 9, 9, 0, 2, 4, 9, 6, 5, 4, 3, 4, 1, 4, 5, 6, 7, 8, 9, 9, 9, 5, 4, 9, 8, 9, 2, 3, 4, 5, 7, 8, 9, 9, 4, 3, 1, 2, 9, 9, 7, 6, 5, 7, 8, 8, 9, 8, 9, 8, 7, 6, 9, 9, 8, 5, 4, 2, 5, 4, 6, 7, 8, 9, 9, 9, 8, 7, 9, 8, 7, 7, 8, 9, 8, 9, 3, 9, 8, 9, 7, 8, 9, 0, 1, 9, 8, 7, 8, 9, 7, 8, 9, 5, 3, 2},
	{3, 4, 5, 6, 7, 8, 9, 1, 3, 9, 8, 7, 7, 5, 4, 3, 2, 3, 4, 5, 6, 7, 9, 8, 7, 6, 2, 3, 7, 8, 9, 9, 5, 9, 8, 9, 8, 7, 5, 4, 4, 3, 4, 9, 8, 7, 6, 8, 9, 9, 7, 6, 4, 9, 8, 9, 8, 7, 6, 4, 3, 1, 2, 3, 5, 6, 7, 8, 9, 9, 9, 9, 8, 7, 6, 5, 6, 8, 7, 9, 9, 8, 7, 8, 6, 7, 9, 1, 9, 8, 7, 6, 7, 9, 8, 9, 4, 3, 2, 1},
	{4, 5, 6, 9, 9, 9, 9, 2, 4, 5, 9, 9, 9, 8, 7, 5, 6, 5, 5, 7, 8, 9, 8, 7, 6, 4, 3, 4, 5, 6, 7, 8, 9, 9, 9, 7, 9, 8, 7, 6, 7, 9, 5, 6, 9, 8, 9, 9, 0, 9, 8, 7, 6, 8, 9, 9, 9, 9, 7, 8, 4, 3, 4, 7, 8, 8, 9, 9, 7, 8, 9, 8, 7, 6, 4, 4, 6, 7, 6, 7, 8, 9, 6, 7, 5, 6, 7, 9, 8, 9, 6, 5, 7, 8, 9, 4, 3, 2, 1, 0},
	{5, 6, 7, 8, 9, 6, 8, 9, 9, 6, 7, 8, 9, 9, 8, 7, 7, 6, 7, 8, 9, 3, 9, 8, 7, 5, 4, 7, 6, 9, 8, 9, 9, 7, 6, 6, 7, 9, 8, 7, 8, 9, 6, 9, 9, 9, 8, 9, 1, 2, 9, 8, 7, 9, 9, 8, 9, 9, 8, 9, 5, 4, 5, 6, 7, 8, 9, 7, 6, 7, 6, 9, 6, 5, 3, 2, 3, 4, 5, 8, 9, 6, 5, 6, 3, 4, 9, 8, 7, 6, 5, 4, 5, 8, 9, 6, 4, 4, 2, 3},
	{6, 8, 8, 9, 3, 4, 5, 9, 8, 9, 8, 9, 2, 3, 9, 9, 8, 8, 9, 9, 2, 1, 2, 9, 7, 6, 5, 8, 9, 9, 9, 8, 8, 9, 5, 4, 3, 4, 9, 8, 9, 9, 9, 8, 8, 9, 7, 8, 9, 4, 5, 9, 8, 9, 8, 7, 6, 7, 9, 7, 6, 5, 7, 7, 9, 9, 7, 6, 5, 4, 5, 9, 5, 4, 2, 1, 4, 7, 6, 7, 8, 9, 4, 1, 2, 3, 4, 9, 8, 7, 6, 6, 6, 7, 8, 9, 6, 5, 6, 9},
	{7, 8, 9, 1, 2, 3, 9, 8, 7, 9, 9, 0, 1, 4, 5, 9, 9, 9, 8, 9, 1, 0, 1, 9, 8, 7, 8, 9, 9, 9, 8, 7, 6, 5, 4, 3, 2, 1, 2, 9, 9, 9, 8, 7, 6, 7, 6, 7, 8, 9, 6, 7, 9, 8, 7, 6, 5, 6, 9, 8, 8, 7, 8, 9, 9, 8, 6, 5, 4, 3, 9, 8, 7, 5, 3, 3, 4, 6, 9, 9, 9, 4, 2, 0, 1, 2, 3, 4, 9, 8, 8, 7, 7, 8, 9, 9, 9, 6, 7, 8},
	{8, 9, 9, 2, 3, 9, 8, 7, 6, 7, 8, 9, 9, 9, 9, 8, 7, 6, 7, 8, 9, 1, 2, 3, 9, 9, 9, 8, 7, 6, 9, 8, 7, 6, 8, 2, 1, 0, 1, 2, 9, 8, 7, 6, 5, 4, 5, 6, 9, 9, 8, 9, 9, 9, 3, 3, 4, 7, 8, 9, 9, 8, 9, 8, 9, 9, 8, 6, 6, 4, 5, 9, 6, 5, 4, 5, 6, 7, 8, 9, 5, 4, 3, 5, 4, 3, 4, 5, 9, 9, 9, 9, 8, 9, 9, 9, 8, 9, 8, 9},
	{9, 9, 8, 9, 4, 9, 9, 9, 5, 6, 7, 9, 8, 7, 8, 7, 6, 5, 6, 7, 8, 9, 3, 9, 9, 7, 9, 7, 6, 5, 4, 9, 8, 7, 7, 3, 2, 2, 2, 3, 4, 9, 8, 7, 5, 3, 8, 7, 8, 8, 9, 9, 9, 8, 2, 1, 5, 6, 7, 8, 9, 9, 8, 7, 8, 9, 9, 7, 7, 9, 9, 9, 7, 6, 8, 7, 8, 9, 9, 7, 6, 5, 9, 9, 8, 6, 5, 6, 7, 8, 9, 4, 9, 9, 9, 8, 7, 8, 9, 9},
	{9, 8, 7, 9, 9, 8, 9, 8, 7, 7, 9, 9, 8, 6, 5, 4, 5, 4, 5, 9, 9, 9, 9, 8, 7, 5, 8, 9, 5, 4, 3, 5, 9, 7, 6, 5, 4, 3, 3, 4, 9, 8, 9, 7, 6, 2, 6, 5, 7, 7, 8, 9, 8, 7, 6, 2, 3, 4, 9, 9, 9, 8, 7, 6, 7, 8, 9, 8, 9, 8, 7, 9, 8, 9, 9, 8, 9, 3, 9, 8, 7, 9, 8, 7, 9, 9, 9, 7, 8, 9, 2, 3, 4, 9, 8, 7, 6, 7, 9, 8},
	{9, 7, 6, 7, 6, 7, 8, 9, 8, 9, 8, 7, 6, 5, 4, 3, 2, 3, 6, 6, 7, 8, 9, 9, 8, 6, 7, 9, 4, 2, 1, 9, 9, 9, 7, 7, 6, 4, 5, 9, 8, 7, 5, 3, 2, 1, 0, 4, 5, 6, 7, 8, 9, 8, 5, 4, 5, 6, 7, 8, 9, 7, 6, 5, 6, 8, 9, 9, 8, 7, 6, 5, 9, 9, 9, 9, 1, 2, 3, 9, 8, 9, 7, 6, 8, 7, 8, 9, 9, 4, 3, 5, 8, 9, 8, 6, 5, 4, 5, 7},
	{9, 8, 5, 4, 5, 8, 9, 9, 9, 6, 9, 9, 8, 7, 5, 2, 1, 3, 4, 5, 6, 7, 8, 9, 9, 7, 9, 8, 9, 3, 9, 8, 9, 8, 9, 9, 7, 5, 7, 9, 9, 8, 6, 4, 3, 2, 1, 2, 3, 5, 6, 7, 9, 9, 6, 5, 6, 7, 9, 9, 8, 7, 5, 4, 5, 7, 9, 8, 9, 8, 7, 6, 7, 8, 9, 9, 9, 3, 4, 5, 9, 8, 7, 5, 7, 6, 7, 8, 9, 9, 8, 6, 7, 8, 9, 7, 4, 3, 4, 5},
	{9, 9, 4, 3, 4, 9, 9, 8, 7, 5, 6, 8, 9, 8, 7, 5, 2, 3, 8, 9, 8, 9, 9, 8, 9, 9, 8, 7, 8, 9, 8, 7, 8, 7, 9, 8, 9, 7, 8, 9, 9, 9, 6, 5, 4, 3, 4, 5, 6, 9, 9, 8, 9, 9, 8, 6, 7, 9, 9, 9, 6, 5, 4, 3, 4, 6, 6, 7, 8, 9, 8, 7, 8, 9, 7, 9, 8, 9, 9, 9, 9, 7, 6, 4, 4, 5, 6, 7, 8, 9, 8, 7, 8, 9, 9, 9, 6, 6, 5, 9},
	{9, 8, 7, 5, 6, 7, 9, 6, 5, 4, 5, 7, 9, 9, 5, 4, 3, 4, 6, 7, 9, 8, 7, 7, 6, 8, 9, 6, 9, 8, 9, 6, 5, 5, 8, 7, 8, 9, 9, 9, 9, 9, 9, 6, 5, 8, 5, 6, 7, 8, 9, 9, 9, 8, 9, 9, 8, 9, 9, 8, 9, 4, 3, 2, 3, 4, 5, 9, 9, 4, 9, 9, 9, 5, 6, 9, 7, 8, 7, 8, 9, 5, 4, 3, 2, 4, 5, 6, 7, 8, 9, 8, 9, 8, 7, 8, 9, 7, 9, 8},
	{6, 9, 8, 6, 9, 9, 8, 9, 6, 3, 2, 9, 8, 7, 6, 5, 4, 9, 7, 9, 8, 7, 5, 4, 5, 9, 4, 5, 6, 7, 8, 9, 3, 3, 3, 5, 6, 9, 8, 9, 8, 9, 8, 7, 6, 7, 6, 8, 8, 9, 9, 9, 8, 7, 9, 8, 9, 9, 8, 7, 6, 5, 6, 1, 2, 3, 5, 8, 9, 3, 2, 1, 2, 4, 9, 8, 6, 5, 6, 7, 9, 9, 5, 0, 1, 2, 3, 4, 5, 6, 8, 9, 8, 7, 6, 7, 9, 9, 8, 7},
	{5, 6, 9, 9, 7, 7, 6, 7, 9, 2, 1, 2, 9, 8, 9, 7, 6, 7, 9, 9, 8, 6, 5, 3, 4, 1, 2, 4, 5, 6, 9, 3, 2, 1, 2, 3, 9, 8, 7, 9, 7, 8, 9, 9, 7, 9, 7, 9, 9, 9, 9, 8, 9, 6, 5, 7, 8, 9, 8, 6, 4, 3, 1, 0, 1, 2, 6, 7, 8, 9, 4, 2, 3, 9, 8, 7, 5, 4, 5, 6, 9, 8, 9, 1, 2, 3, 5, 6, 8, 9, 9, 9, 7, 6, 5, 6, 7, 8, 9, 6},
	{6, 7, 9, 8, 6, 4, 5, 9, 8, 9, 0, 1, 2, 9, 9, 8, 7, 9, 8, 7, 6, 5, 4, 2, 1, 0, 1, 2, 7, 8, 9, 4, 3, 2, 3, 9, 8, 7, 6, 5, 6, 7, 8, 9, 9, 4, 9, 2, 1, 9, 8, 7, 6, 5, 4, 6, 7, 8, 9, 8, 6, 4, 3, 2, 6, 4, 5, 6, 7, 8, 9, 9, 9, 8, 9, 5, 4, 3, 5, 9, 8, 7, 8, 9, 4, 4, 6, 7, 8, 9, 9, 9, 8, 8, 6, 7, 8, 9, 4, 5},
	{7, 8, 9, 4, 5, 3, 4, 5, 6, 8, 9, 9, 9, 1, 2, 9, 8, 9, 9, 8, 7, 6, 5, 3, 2, 1, 2, 7, 6, 8, 9, 9, 4, 3, 4, 9, 8, 7, 5, 4, 5, 6, 7, 8, 9, 3, 2, 1, 0, 1, 9, 9, 8, 7, 5, 7, 8, 9, 9, 8, 7, 8, 4, 4, 5, 7, 6, 7, 8, 9, 9, 8, 8, 7, 5, 7, 3, 2, 9, 8, 7, 5, 7, 8, 9, 5, 7, 8, 9, 6, 7, 8, 9, 9, 9, 8, 9, 2, 3, 4},
	{8, 9, 8, 3, 2, 1, 3, 3, 4, 5, 6, 7, 8, 9, 3, 4, 9, 8, 7, 9, 8, 7, 8, 4, 4, 2, 3, 4, 5, 6, 7, 8, 9, 5, 9, 7, 6, 6, 2, 3, 4, 5, 6, 8, 9, 5, 3, 3, 2, 3, 4, 9, 8, 7, 6, 7, 9, 7, 8, 9, 8, 6, 5, 5, 6, 8, 7, 8, 9, 9, 8, 7, 6, 5, 3, 2, 0, 1, 2, 9, 6, 4, 6, 7, 8, 9, 9, 9, 3, 5, 6, 9, 7, 6, 5, 9, 0, 1, 2, 3},
	{9, 8, 7, 6, 5, 0, 1, 2, 3, 5, 7, 9, 9, 9, 9, 9, 7, 6, 5, 6, 9, 8, 9, 7, 7, 3, 4, 6, 7, 8, 8, 9, 9, 9, 8, 7, 5, 4, 1, 2, 3, 4, 6, 8, 9, 6, 5, 4, 6, 4, 5, 6, 9, 9, 7, 8, 9, 6, 9, 1, 9, 7, 9, 8, 7, 9, 8, 9, 5, 4, 9, 8, 7, 6, 5, 3, 3, 2, 9, 8, 5, 3, 5, 6, 8, 9, 2, 0, 2, 3, 9, 8, 7, 6, 3, 2, 1, 2, 3, 4},
	{8, 7, 6, 5, 4, 3, 2, 4, 5, 6, 7, 8, 9, 9, 8, 7, 5, 4, 3, 5, 6, 9, 7, 6, 5, 4, 5, 7, 8, 9, 9, 9, 9, 8, 7, 6, 5, 3, 2, 3, 4, 6, 7, 9, 8, 7, 6, 5, 6, 8, 9, 7, 9, 9, 8, 9, 8, 7, 9, 0, 9, 8, 9, 9, 8, 9, 9, 8, 6, 5, 6, 9, 9, 8, 7, 6, 5, 9, 8, 7, 6, 4, 6, 8, 9, 3, 2, 1, 2, 9, 8, 7, 6, 5, 4, 3, 2, 3, 4, 5},
}

type hight_map struct {
	size       point
	heights    [][]int
	low_map    [][]int // 0 is unknown, -1 low, 1 not low
	low_points []point
	risk_val   int
}

func create_height_map(heights [][]int) (hm hight_map) {

	x_size := len(heights)
	y_size := len(heights[0])

	hm = hight_map{
		point{x_size, y_size},
		heights,
		make([][]int, x_size),
		[]point{},
		0,
	}

	for i := range hm.low_map {
		hm.low_map[i] = make([]int, y_size)
	}

	return hm
}

func (hm *hight_map) get_val(x int, y int) int {
	return hm.heights[x][y]
}

// check if point is lower than V
func (hm *hight_map) check_side_points(x int, y int) (is_low bool) {
	is_low = false
	v := hm.get_val(x, y)
	if x != 0 {
		if hm.get_val(x-1, y) <= v {
			return false
		}
		hm.low_map[x-1][y] = 1
	}

	if y != 0 {
		if hm.get_val(x, y-1) <= v {
			return false
		}
		hm.low_map[x][y-1] = 1
	}

	if x != hm.size.x-1 {
		if hm.get_val(x+1, y) <= v {
			return false
		}
		hm.low_map[x+1][y] = 1

	}

	if y != hm.size.y-1 {
		if hm.get_val(x, y+1) <= v {
			return false
		}
		hm.low_map[x][y+1] = 1
	}

	return true
}

func (hm *hight_map) scan() {
	for x, r := range hm.heights {
		for y, v := range r {
			if hm.low_map[x][y] == 0 {
				if hm.check_side_points(x, y) {
					fmt.Println(x, y, v)
					hm.low_points = append(hm.low_points, point{x, y})
					hm.low_map[x][y] = -1
					hm.risk_val += v + 1
				} else {
					hm.low_map[x][y] = 1
				}
			}
		}
	}
}

func Day09A() int {

	hm := create_height_map(day09_data)
	hm.scan()

	return hm.risk_val
}

func Day09B() int {

	return 0
}
