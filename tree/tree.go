package tree

import (
	"fmt"
	"github.com/ajstarks/svgo"
	"math/rand"
	"os"
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

	for i, c := range t.children {
		rootHeight := t.length
		if i != 0 && t.length > 10 {
			rootHeight = rand.Intn(t.length)
		}

		canvas.Gtransform(fmt.Sprintf("translate(%d,%d)", 0, rootHeight))
		c.draw(canvas)
		canvas.Gend()
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

func New() Tree {
	var c []tree
	for i := 0; i < 5; i++ {
		t := tree{
			length:   80,
			rotation: -60 + rand.Intn(120),
			children: []tree{},
		}
		c = append(c, t)
	}

	return &tree{
		children: c,
		length:   100,
		rotation: -10 + rand.Intn(20),
	}
}
