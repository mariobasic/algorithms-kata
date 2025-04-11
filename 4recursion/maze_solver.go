package _recursion

type Path struct {
	path []Point
}

func (p *Path) push(val Point) {
	p.path = append(p.path, val)
}

func (p *Path) pop() Point {
	if len(p.path) == 0 {
		return Point{}
	}
	last := p.path[len(p.path)-1]
	p.path = p.path[:len(p.path)-1]
	return last
}

type Point struct {
	x, y int
}

// func walk(maze []string, wall rune, curr Point, end Point, seen [][]bool, path *data_structure.Stack[Point]) bool {
func walk(maze []string, wall rune, curr Point, end Point, seen [][]bool, path *Path) bool {
	// base case
	// 1. off the map
	if curr.x < 0 || curr.x >= len(maze[0]) || curr.y < 0 || curr.y >= len(maze) {
		return false
	}
	// 2. on a wall?
	if maze[curr.y][curr.x] == uint8(wall) {
		return false
	}
	// 3. at the end?
	if curr.x == end.x && curr.y == end.y {
		//*path = append(*path, curr)
		path.push(curr)
		return true
	}

	if seen[curr.y][curr.x] {
		return false
	}
	// 3 steps in each recursion
	// 1. pre
	// 2. recurse
	// 3. post

	// pre
	seen[curr.y][curr.x] = true
	//*path = append(*path, curr)
	path.push(curr)

	// recurse
	for i := 0; i < len(getDir()); i++ {
		a := getDir()[i]
		x, y := a[0], a[1]
		if walk(maze, wall, Point{x: curr.x + x, y: curr.y + y}, end, seen, path) {
			return true
		}

	}
	//post
	//*path = (*path)[:len(*path)-1]
	path.pop()

	return false
}

func solve(maze []string, wall rune, start Point, end Point) []Point {
	var seen [][]bool
	//path := data_structure.NewStack[Point]()
	var path Path

	for i := 0; i < len(maze); i++ {
		seen = append(seen, make([]bool, len(maze[0])))
	}

	walk(maze, wall, start, end, seen, &path)

	return path.path
}

func getDir() [4][2]int {
	return [4][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
}
