package main

import (
	"math/rand"
	"time"
	"fmt"
	"os"
	"os/exec"
	"bufio"
)

const (
	W = 10
	H = 20
)

var (
	board [H][W]int
	shapes = [][][]int{
		{{1,1,1,1}},                // I
		{{1,1},{1,1}},              // O
		{{1,1,1},{0,1,0}},          // T
		{{1,1,0},{0,1,1}},          // S
		{{0,1,1},{1,1,0}},          // Z
		{{1,0,0},{1,1,1}},          // J
		{{0,0,1},{1,1,1}},          // L
	}
	curShape, curX, curY, curR int
	score int
)

func clearScreen() {
	exec.Command("clear").Run()
}

func draw() {
	clearScreen()
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			char := " "
			if board[y][x]>0 { char="X" }
			if onShape(x, y) { char="O" }
			fmt.Print(char)
		}
		fmt.Println()
	}
	fmt.Println("Score:", score)
}

func onShape(x, y int) bool {
	shape := shapes[curShape]
	for i, row := range shape {
		for j, cell := range row {
			if cell > 0 {
				if curY+i==y && curX+j==x { return true }
			}
		}
	}
	return false
}

func canMove(dx, dy, dr int) bool {
	shape := shapes[curShape]
	nr := (curR+dr)%4
	rot := rotate(shape, nr)
	for i, row := range rot {
		for j, cell := range row {
			if cell > 0 {
				x, y := curX+j+dx, curY+i+dy
				if x<0||x>=W||y<0||y>=H { return false }
				if board[y][x]>0 { return false }
			}
		}
	}
	return true
}

func rotate(shape [][]int, times int) [][]int {
	res := shape
	for t:=0;t<times;t++ {
		w,h := len(res[0]),len(res)
		rot := make([][]int,w)
		for i:=0;i<w;i++ { rot[i]=make([]int,h)}
		for i:=0;i<h;i++ {
			for j:=0;j<w;j++ {
				rot[j][h-i-1]=res[i][j]
			}
		}
		res=rot
	}
	return res
}

func land() {
	shape := rotate(shapes[curShape], curR)
	for i, row := range shape {
		for j, cell := range row {
			if cell>0 { board[curY+i][curX+j]=1 }
		}
	}
	for y:=H-1;y>=0;y-- {
		full:=true
		for x:=0;x<W;x++ { if board[y][x]==0 { full=false } }
		if full {
			for yy:=y;yy>0;yy-- { board[yy]=board[yy-1] }
			for x:=0;x<W;x++ { board[0][x]=0 }
			score+=10
			y++
		}
	}
}

func spawn() bool {
	curShape = rand.Intn(len(shapes))
	curX, curY, curR = W/2-1, 0, 0
	return canMove(0,0,0)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	spawn()
	draw()
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			c, _ := reader.ReadByte()
			switch c {
			case 'a':
				if canMove(-1,0,0) { curX-- }
			case 'd':
				if canMove(1,0,0) { curX++ }
			case 's':
				if canMove(0,1,0) { curY++ }
			case 'w':
				if canMove(0,0,1) { curR=(curR+1)%4 }
			}
			draw()
		}
	}()
	for {
		time.Sleep(500*time.Millisecond)
		if canMove(0,1,0) { curY++ }
		else {
			land()
			if !spawn() { fmt.Println("Game Over"); return }
		}
		draw()
	}
}