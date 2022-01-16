package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/fashionscape/fsbot-api/graph/generated"
	"github.com/fashionscape/fsbot-api/graph/model"
)

func (r *mutationResolver) CreateOutfit(ctx context.Context, input model.InputOutfit) (*model.Outfit, error) {
	// Send to internal function. Instantiate an outfit, add created/updated timestamps, and store it.
	fmt.Println("This is where an outfit will be created")
	// Create handlers
	return nil, nil
}

func (r *mutationResolver) UpdateOutfit(ctx context.Context, input model.InputOutfit) (*model.Outfit, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetOutfit(ctx context.Context, outfitID string) (*model.Outfit, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetRandomOutfit(ctx context.Context) (*model.Outfit, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
