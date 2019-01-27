package tree

import (
	"fmt"
	"github.com/ajstarks/svgo"
	"math"
	"math/rand"
	"os"
)

const (
	leafFactor = 0.2
)

type Tree interface {
	Draw(x, y int, filename string) error
}

type tree struct {
	children []tree
	length   int
	rotation int
}

func (t *tree) draw(canvas *svg.SVG) {
	canvas.Gtransform(fmt.Sprintf("rotate(%d)", t.rotation))
	canvas.Line(0, 0, 0, t.length, "fill:none;stroke:black")

	if len(t.children) > 0 {
		for i, c := range t.children {
			rootHeight := t.length
			if i != 0 && t.length > 10 {
				rootHeight = rand.Intn(t.length)
			}

			canvas.Gtransform(fmt.Sprintf("translate(%d,%d)", 0, rootHeight))
			c.draw(canvas)
			canvas.Gend()
		}
	} else {
		var x,y []int
		leafSize := int(float64(t.length) * leafFactor)
		for i := 0; i<3; i++ {
			x = append(x, rand.Intn(leafSize * 2) - leafSize)
			y = append(y, rand.Intn(leafSize * 2) - leafSize)
		}
		canvas.Polygon(x, y, "fill:white;stroke:black")
	}
	canvas.Gend()
}

func (t *tree) Draw(x, y int, filename string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	canvas := svg.New(f)
	canvas.Start(x, y)
	canvas.TranslateRotate(x/2, y, 180)
	t.draw(canvas)
	canvas.Gend()
	canvas.End()
	return nil
}

func New(childCount, depth, length, angleRange int) Tree {
	var c []tree
	if depth > 0 {
		childLength := int(0.8 * float64(length))
		grandchildrenCount := int(math.Max(1, float64(childCount)-1))
		for i := 0; i < childCount; i++ {
			t := New(grandchildrenCount, depth-1, childLength, 60).(*tree)
			c = append(c, *t)
		}
	}

	return &tree{
		children: c,
		length:   length,
		rotation: rand.Intn(2*angleRange) - angleRange,
	}
}
