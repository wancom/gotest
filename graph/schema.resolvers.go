package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"
)

// Send is the resolver for the send field.
func (r *mutationResolver) Send(ctx context.Context, text string) (string, error) {
	err := r.Redis.Publish(text)
	if err != nil {
		return "", err
	}
	return text, nil
}

// Message is the resolver for the message field.
func (r *queryResolver) Message(ctx context.Context) ([]string, error) {
	panic(fmt.Errorf("not implemented: Message - message"))
}

// Submsg is the resolver for the submsg field.
func (r *subscriptionResolver) Submsg(ctx context.Context) (<-chan string, error) {
	return r.Redis.SubscribeCh(ctx), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
