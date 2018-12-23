package main

import (
	"github.com/kdhageman/terraingen/tree"
	"github.com/rs/zerolog/log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	t := tree.New()
	if err := t.Draw(256, 256, "out/test.svg"); err != nil {
		log.Debug().Msgf("Failed to draw: %s", err)
	}
}
