package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"

	graph1 "github.com/ecshreve/jepp/graph/generated"
	"github.com/ecshreve/jepp/graph/model"
)

// Season is the resolver for the season field.
func (r *gameResolver) Season(ctx context.Context, obj *model.Game) (*model.Season, error) {
	var season model.Season
	if err := r.DB.First(&season, obj.SeasonID).Error; err != nil {
		return nil, err
	}

	return &season, nil
}

// Clues is the resolver for the clues field.
func (r *gameResolver) Clues(ctx context.Context, obj *model.Game) ([]*model.Clue, error) {
	var clues []*model.Clue
	if err := r.DB.Where("game_id = ?", obj.ID).Find(&clues).Error; err != nil {
		return nil, err
	}

	return clues, nil
}

// Game returns graph1.GameResolver implementation.
func (r *Resolver) Game() graph1.GameResolver { return &gameResolver{r} }

type gameResolver struct{ *Resolver }
