package main

import "fmt"

func printMulti(n int, c string) {
	for i := 0; i < n; i++ {
		fmt.Print(c)
	}
	//fmt.Println()
}
func printStars(n int) {
	printMulti(n, "*")
}

func q1(n int) {
	for i := 1; i <= n; i++ {
		printStars(i)
	}
}

func q2(n int) {
	for i := 1; i <= n; i++ {
		printStars(i)
	}
	for i := n - 1; i > 0; i-- {
		printStars(i)
	}
}

/*
func printStarsAtCenter(numOfStars int, row int, img [][]string) {
	col := len(img)
	var ws int = (col - numOfStars) / 2
	//fmt.Println(ws)
	//printMulti(ws, "_")
	for i := 0; i < numOfStars; i++ {
	}
	printStars(numOfStars)

}
*/

/*
func q3(img [][]string) {
	n := len(img)
	row := 0
	for i := 1; i <= n; i += 2 {
		img[row]
	}
	for i := n - 2; i > 0; i -= 2 {
		printStarsAtCenter(i, n)
		fmt.Println()
	}
}
*/

func q4(n int) {
	var img = make([][]string, n)
	for i := 0; i < n; i++ {
		img[i] = make([]string, n)
		for j := 0; j < n; j++ {
			img[i][j] = "_"
		}
	}
	printImg(img)
	drawDiamond(img)
	printImg(img)

	fillNonRecursive("$", 3, 0, img)
	printImg(img)
}

func drawDiamond(img [][]string) {
	n := len(img)
	numOfStars := 1
	reached := false
	for i := 0; i < n; i++ {
		for j := 0; j < numOfStars; j++ {
			img[i][(n-numOfStars)/2+j] = "*"
		}
		if numOfStars == n {
			reached = true
		}
		if reached {
			numOfStars -= 2
		} else {
			numOfStars += 2
		}
	}
}

func printImg(img [][]string) {
	for i := 0; i < len(img); i++ {
		for j := 0; j < len(img[i]); j++ {
			fmt.Print(img[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func fill(color string, x int, y int, img [][]string) {
	old := img[x][y]
	fillImpl(old, color, x, y, img)
}
func fillImpl(old string, color string, x int, y int, img [][]string) {
	n := len(img)
	if x < n && x >= 0 && y < n && y >= 0 && img[x][y] == old {
		img[x][y] = color
		fillImpl(old, color, x+1, y, img)
		fillImpl(old, color, x-1, y, img)
		fillImpl(old, color, x, y+1, img)
		fillImpl(old, color, x, y-1, img)
	}
}

func fillNonRecursive(color string, x int, y int, img [][]string) {
	old := img[x][y]
	fillNonRecursiveImpl(old, color, x, y, img)
}

type Point struct {
	x int
	y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func fillNonRecursiveImpl(old string, color string,
	x int, y int, img [][]string) {

	n := len(img)
	todo := make([]Point, 0, 10)
	todo = append(todo, Point{x, y})
	for len(todo) > 0 {
		//printSlice(todo)
		last := todo[len(todo)-1]
		//fmt.Printf("(%d, %d)\n", last.x, last.y)
		todo = todo[0 : len(todo)-1]
		if last.x >= 0 && last.x < n && last.y >= 0 && last.y < n &&
			img[last.x][last.y] == old {
			img[last.x][last.y] = color
			todo = append(todo, Point{last.x, last.y + 1})
			todo = append(todo, Point{last.x, last.y - 1})
			todo = append(todo, Point{last.x + 1, last.y})
			todo = append(todo, Point{last.x - 1, last.y})
		}
	}
}

func printSlice(slice []Point) {
	fmt.Print("[")
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i])
		fmt.Print("; ")
	}
	fmt.Println("]")
}

func main() {
	//fmt.Println("hi")
	//q1(5)
	//q2(5)
	//q3(11)
	q4(17)
}
