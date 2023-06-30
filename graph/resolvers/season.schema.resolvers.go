package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"time"

	"github.com/ecshreve/jepp/app/models"
	"github.com/ecshreve/jepp/graph"
	"github.com/ecshreve/jepp/graph/common"
)

// StartDate is the resolver for the startDate field.
func (r *seasonResolver) StartDate(ctx context.Context, obj *models.Season) (string, error) {
	context := common.GetContext(ctx)

	var season models.Season
	if err := context.Database.First(&season, obj.ID).Error; err != nil {
		return "", err
	}

	return season.StartDate.Format(time.RFC3339), nil
}

// EndDate is the resolver for the endDate field.
func (r *seasonResolver) EndDate(ctx context.Context, obj *models.Season) (string, error) {
	context := common.GetContext(ctx)

	var season models.Season
	if err := context.Database.First(&season, obj.ID).Error; err != nil {
		return "", err
	}

	return season.EndDate.Format(time.RFC3339), nil
}

// Season returns graph.SeasonResolver implementation.
func (r *Resolver) Season() graph.SeasonResolver { return &seasonResolver{r} }

type seasonResolver struct{ *Resolver }
