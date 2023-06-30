package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"

	"github.com/ecshreve/jepp/app/models"
	"github.com/ecshreve/jepp/graph"
	"github.com/ecshreve/jepp/graph/common"
)

// Category is the resolver for the category field.
func (r *clueResolver) Category(ctx context.Context, obj *models.Clue) (*models.Category, error) {
	context := common.GetContext(ctx)

	var category models.Category
	if err := context.Database.First(&category, obj.CategoryID).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

// Game is the resolver for the game field.
func (r *clueResolver) Game(ctx context.Context, obj *models.Clue) (*models.Game, error) {
	context := common.GetContext(ctx)

	var game models.Game
	if err := context.Database.First(&game, obj.GameID).Error; err != nil {
		return nil, err
	}

	return &game, nil
}

// Clue returns graph.ClueResolver implementation.
func (r *Resolver) Clue() graph.ClueResolver { return &clueResolver{r} }

type clueResolver struct{ *Resolver }