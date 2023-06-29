package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"encoding/base64"
	"fmt"
	"math"

	"github.com/ecshreve/jepp/app/models"
	"github.com/ecshreve/jepp/graph/common"
	"github.com/ecshreve/jepp/graph/model"
)

// Clues is the resolver for the clues field.
func (r *categoryResolver) Clues(ctx context.Context, obj *models.Category) ([]*models.Clue, error) {
	context := common.GetContext(ctx)

	var clues []*models.Clue
	if err := context.Database.Find(&clues, &models.Clue{CategoryID: obj.ID}).Error; err != nil {
		return nil, err
	}

	return clues, nil
}

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

// Season is the resolver for the season field.
func (r *gameResolver) Season(ctx context.Context, obj *models.Game) (*model.Season, error) {
	context := common.GetContext(ctx)

	var season model.Season
	if err := context.Database.First(&season, obj.SeasonID).Error; err != nil {
		return nil, err
	}

	return &season, nil
}

// AirDate is the resolver for the airDate field.
func (r *gameResolver) AirDate(ctx context.Context, obj *models.Game) (string, error) {
	context := common.GetContext(ctx)

	var game models.Game
	if err := context.Database.First(&game, obj.ID).Error; err != nil {
		return "", err
	}

	return game.AirDate.Format("2006-01-02"), nil
}

// TapeDate is the resolver for the tapeDate field.
func (r *gameResolver) TapeDate(ctx context.Context, obj *models.Game) (string, error) {
	context := common.GetContext(ctx)

	var game models.Game
	if err := context.Database.First(&game, obj.ID).Error; err != nil {
		return "", err
	}

	return game.TapeDate.Format("2006-01-02"), nil
}

// CluesConnection is the resolver for the cluesConnection field.
func (r *gameResolver) CluesConnection(ctx context.Context, obj *models.Game, first *int64, after *string) (*model.CluesConnection, error) {
	context := common.GetContext(ctx)

	// The cursor is base64 encoded by convention, so we need to decode it first
	var decodedCursor string
	if after != nil {
		b, err := base64.StdEncoding.DecodeString(*after)
		if err != nil {
			return nil, err
		}
		decodedCursor = string(b)
	}

	if decodedCursor == "" {
		decodedCursor = "0"
	}

	// Here we could query the DB to get data, e.g.
	var edges []*model.CluesEdge
	hasNextPage := false

	var clues []*models.Clue
	if err := context.Database.Limit(1000).Order("id asc").Where("game_id = ? AND id > ?", obj.ID, decodedCursor).Find(&clues).Error; err != nil {
		return nil, err
	}

	bound := int(math.Min(float64(int(*first)), float64(len(clues))))
	for i := 0; i < bound; i++ {

		edges = append(edges, &model.CluesEdge{
			Cursor: base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", clues[i].ID))),
			Node:   clues[i],
		})
	}

	if len(edges) == int(*first) && len(edges) < len(clues) {
		hasNextPage = true
	}

	pageInfo := model.PageInfo{
		StartCursor: base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", edges[0].Node.ID))),
		EndCursor:   base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", edges[len(edges)-1].Node.ID))),
		HasNextPage: &hasNextPage,
	}

	cc := model.CluesConnection{
		Edges:    edges,
		PageInfo: &pageInfo,
	}

	return &cc, nil
}

// Clue is the resolver for the clue field.
func (r *queryResolver) Clue(ctx context.Context, clueID string) (*models.Clue, error) {
	context := common.GetContext(ctx)

	var clue models.Clue
	if err := context.Database.First(&clue, clueID).Error; err != nil {
		return nil, err
	}

	return &clue, nil
}

// Clues is the resolver for the clues field.
func (r *queryResolver) Clues(ctx context.Context) ([]*models.Clue, error) {
	context := common.GetContext(ctx)

	var clues []*models.Clue
	if err := context.Database.Find(&clues).Error; err != nil {
		return nil, err
	}

	return clues, nil
}

// Category is the resolver for the category field.
func (r *queryResolver) Category(ctx context.Context, categoryID string) (*models.Category, error) {
	context := common.GetContext(ctx)

	var category models.Category
	if err := context.Database.First(&category, categoryID).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*models.Category, error) {
	context := common.GetContext(ctx)

	var categories []*models.Category
	if err := context.Database.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

// Season is the resolver for the season field.
func (r *queryResolver) Season(ctx context.Context, seasonID string) (*model.Season, error) {
	context := common.GetContext(ctx)

	var season model.Season
	if err := context.Database.First(&season, seasonID).Error; err != nil {
		return nil, err
	}

	return &season, nil
}

// Seasons is the resolver for the seasons field.
func (r *queryResolver) Seasons(ctx context.Context) ([]*model.Season, error) {
	context := common.GetContext(ctx)

	var seasons []*model.Season
	if err := context.Database.Find(&seasons).Error; err != nil {
		return nil, err
	}

	return seasons, nil
}

// Game is the resolver for the game field.
func (r *queryResolver) Game(ctx context.Context, gameID string) (*models.Game, error) {
	context := common.GetContext(ctx)

	var game models.Game
	if err := context.Database.First(&game, gameID).Error; err != nil {
		return nil, err
	}

	return &game, nil
}

// Games is the resolver for the games field.
func (r *queryResolver) Games(ctx context.Context) ([]*models.Game, error) {
	context := common.GetContext(ctx)

	var games []*models.Game
	if err := context.Database.Find(&games).Error; err != nil {
		return nil, err
	}

	return games, nil
}

// Category returns CategoryResolver implementation.
func (r *Resolver) Category() CategoryResolver { return &categoryResolver{r} }

// Clue returns ClueResolver implementation.
func (r *Resolver) Clue() ClueResolver { return &clueResolver{r} }

// Game returns GameResolver implementation.
func (r *Resolver) Game() GameResolver { return &gameResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type categoryResolver struct{ *Resolver }
type clueResolver struct{ *Resolver }
type gameResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
