// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
)

func (c *Clue) Category(ctx context.Context) (*Category, error) {
	result, err := c.Edges.CategoryOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryCategory().Only(ctx)
	}
	return result, err
}

func (c *Clue) Game(ctx context.Context) (*Game, error) {
	result, err := c.Edges.GameOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryGame().Only(ctx)
	}
	return result, err
}

func (ga *Game) Season(ctx context.Context) (*Season, error) {
	result, err := ga.Edges.SeasonOrErr()
	if IsNotLoaded(err) {
		result, err = ga.QuerySeason().Only(ctx)
	}
	return result, err
}