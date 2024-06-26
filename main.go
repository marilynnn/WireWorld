package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const sWidth = 900
const sHeight = 600

var scale int = 15
var InitField Field
var stopped = false
var counter = 0

// enum состояния клеток
// Empty - стандартная пустая клетка
// Conductor - клетка проводник
// Head - голова электрона
// Tail - хвост электрона

const (
	empty     = 0
	head      = 1
	tail      = 2
	conductor = 3
)

// координата клетки
type Cell struct {
	x int
	y int
}

// соседи клетки
var CellNeighbors = []Cell{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, 1}, {1, 1}, {1, 0},
	{1, -1}, {0, -1},
}

// поле
type Field map[Cell]int

// размер поля
var boundary Cell = Cell{sWidth / scale, sHeight / scale}

// подсчет количества соседних клеток типа "head"
func CountHeads(field Field) func(Cell) int {
	return func(cell Cell) int {
		hcount := 0
		for _, neighbours := range CellNeighbors {
			currNeighb := Cell{cell.x + neighbours.x,
				cell.y + neighbours.y}
			if field[currNeighb] == head {
				hcount++
			}
		}
		return hcount
	}
}

// вычисление состояния клетки в соотв с правилами автомата
func NewState(field Field) func(Cell) int {
	heads := CountHeads(field)
	return func(cell Cell) int {
		switch field[cell] {
		case empty:
			return empty
		case head:
			return tail
		case tail:
			return conductor
		case conductor:
			if heads(cell) == 1 || heads(cell) == 2 {
				return head
			}
			return conductor
		default:
			return field[cell]
		}
	}
}

func DrawCell(cell int) int {
	switch cell {
	case empty:
		return conductor
	case conductor:
		return head
	default:
		return cell
	}
}

func DrawTimer(x int, y int) {
	var templ = []Cell{
		{x, y}, {x + 1, y + 1}, {x + 2, y + 1}, {x + 3, y + 1}, {x + 4, y + 1},
		{x + 5, y}, {x + 4, y - 1}, {x + 3, y - 1}, {x + 2, y - 1}, {x + 1, y - 1}}
	for _, coord := range templ {
		InitField[coord] = conductor
	}
}

func DrawDiod(x int, y int) {
	var templ = []Cell{
		{x, y}, {x + 1, y}, {x + 1, y + 1}, {x + 1, y - 1},
		{x + 3, y}, {x + 2, y - 1}, {x + 2, y + 1}}
	for _, coord := range templ {
		InitField[coord] = conductor
	}
}

func DrawOr(x int, y int) {
	var templ = []Cell{
		{x, y}, {x, y + 2}, {x, y - 2},
		{x + 1, y}, {x + 1, y - 1}, {x + 1, y + 1},
		{x + 2, y}}
	for _, coord := range templ {
		InitField[coord] = conductor
	}
}

func DrawXor(x int, y int) {
	var templ = []Cell{
		{x, y}, {x, y - 2}, {x, y - 3}, {x, y - 4}, {x, y - 6},
		{x + 1, y - 1}, {x + 1, y - 2}, {x + 1, y - 4}, {x + 1, y - 5},
		{x + 2, y - 2}, {x + 2, y - 4},
		{x + 3, y - 2}, {x + 3, y - 3}, {x + 3, y - 4}}
	for _, coord := range templ {
		InitField[coord] = conductor
	}
}

func DrawNot(x int, y int) {
	var templ = []Cell{
		{x, y}, {x + 1, y}, {x + 2, y - 1}, {x + 2, y + 2}, {x + 2, y + 3},
		{x + 3, y + 1}, {x + 3, y - 1}, {x + 3, y + 3},
		{x + 4, y}, {x + 4, y + 1}, {x + 4, y + 2}, {x + 4, y + 3},
		{x + 5, y + 1}, {x + 5, y + 2}, {x + 6, y}}
	for _, coord := range templ {
		InitField[coord] = conductor
	}
}

// обновление поля
func FieldUpdate(field Field) Field {
	newField := make(Field)
	newCell := NewState(field)
	for x := 0; x < boundary.x; x++ {
		for y := 0; y < boundary.y; y++ {
			currCell := Cell{x, y}
			newField[currCell] = newCell(currCell)
		}
	}
	return newField
}

type Game struct{}

func (g *Game) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeySpace) && inpututil.KeyPressDuration(ebiten.KeySpace) == 1 {
		stopped = !stopped
	}

	if ebiten.IsKeyPressed(ebiten.KeyBackspace) && inpututil.KeyPressDuration(ebiten.KeyBackspace) == 1 {
		InitField = Field{}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) == 1 {

		UpdatedCell := Cell{}
		UpdatedCell.x, UpdatedCell.y = ebiten.CursorPosition()
		UpdatedCell.x = (UpdatedCell.x) / scale
		UpdatedCell.y = (UpdatedCell.y) / scale
		InitField[UpdatedCell] = DrawCell(InitField[UpdatedCell])

	}

	if ebiten.IsKeyPressed(ebiten.KeyT) && inpututil.KeyPressDuration(ebiten.KeyT) == 1 {

		UpdatedCell := Cell{}
		UpdatedCell.x, UpdatedCell.y = ebiten.CursorPosition()
		UpdatedCell.x = (UpdatedCell.x) / scale
		UpdatedCell.y = (UpdatedCell.y) / scale
		DrawTimer(UpdatedCell.x, UpdatedCell.y)
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) && inpututil.KeyPressDuration(ebiten.KeyD) == 1 {

		UpdatedCell := Cell{}
		UpdatedCell.x, UpdatedCell.y = ebiten.CursorPosition()
		UpdatedCell.x = (UpdatedCell.x) / scale
		UpdatedCell.y = (UpdatedCell.y) / scale
		DrawDiod(UpdatedCell.x, UpdatedCell.y)
	}

	if ebiten.IsKeyPressed(ebiten.KeyO) && inpututil.KeyPressDuration(ebiten.KeyO) == 1 {

		UpdatedCell := Cell{}
		UpdatedCell.x, UpdatedCell.y = ebiten.CursorPosition()
		UpdatedCell.x = (UpdatedCell.x) / scale
		UpdatedCell.y = (UpdatedCell.y) / scale
		DrawOr(UpdatedCell.x, UpdatedCell.y)
	}

	if ebiten.IsKeyPressed(ebiten.KeyX) && inpututil.KeyPressDuration(ebiten.KeyX) == 1 {

		UpdatedCell := Cell{}
		UpdatedCell.x, UpdatedCell.y = ebiten.CursorPosition()
		UpdatedCell.x = (UpdatedCell.x) / scale
		UpdatedCell.y = (UpdatedCell.y) / scale
		DrawXor(UpdatedCell.x, UpdatedCell.y)
	}

	if ebiten.IsKeyPressed(ebiten.KeyN) && inpututil.KeyPressDuration(ebiten.KeyN) == 1 {

		UpdatedCell := Cell{}
		UpdatedCell.x, UpdatedCell.y = ebiten.CursorPosition()
		UpdatedCell.x = (UpdatedCell.x) / scale
		UpdatedCell.y = (UpdatedCell.y) / scale
		DrawNot(UpdatedCell.x, UpdatedCell.y)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, 0, 0, sWidth, sHeight, color.Black)
	gridColor := &color.RGBA{255, 255, 255, 50}
	for y := 0.0; y < sHeight; y += float64(scale) {
		ebitenutil.DrawLine(screen, 0, y, sWidth, y, gridColor)
		//vector.StrokeLine(screen, 0.0, y, sWidth, sHeight, 5.0, gridColor)
	}
	for x := 0.0; x < sWidth; x += float64(scale) {
		ebitenutil.DrawLine(screen, x, 0, x, sWidth, gridColor)
		//vector.StrokeLine(screen, 0.0, y, sWidth, sHeight, 5.0, gridColor)
	}
	for x := 0; x < boundary.x; x++ {
		for y := 0; y < boundary.y; y++ {
			currCell := Cell{x, y}
			switch InitField[currCell] {
			case conductor:
				screen.Set(x*scale, y*scale, &color.RGBA{239, 254, 90, 250})
				ebitenutil.DrawRect(screen, float64(x*scale), float64(y*scale), float64(scale), float64(scale), &color.RGBA{239, 254, 90, 250})
			case head:
				screen.Set(x*scale, y*scale, &color.RGBA{0, 46, 255, 250})
				ebitenutil.DrawRect(screen, float64(x*scale), float64(y*scale), float64(scale), float64(scale), &color.RGBA{0, 46, 255, 250})
			case tail:
				screen.Set(x*scale, y*scale, &color.RGBA{55, 0, 0, 250})
				ebitenutil.DrawRect(screen, float64(x*scale), float64(y*scale), float64(scale), float64(scale), &color.RGBA{255, 0, 0, 250})
			default:
				break
			}
		}
	}

	if !stopped {
		counter += 1
		if counter%10 == 0 {
			InitField = FieldUpdate(InitField)
			counter = 0
		}
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return sWidth, sHeight
}

func main() {
	InitField = Field{}

	ebiten.SetWindowSize(sWidth, sHeight)
	ebiten.SetWindowTitle("WireWorld!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}
