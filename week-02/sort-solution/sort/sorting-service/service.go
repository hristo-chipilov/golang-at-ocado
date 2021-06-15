package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/bbsbb/go-at-ocado/sort/gen"
)

func newSortingService() gen.SortingRobotServer {
	return &sortingService{items: map[string]int{}}
}

type sortingService struct {
	items map[string]int
}

func (s *sortingService) LoadItems(ctx context.Context, request *gen.LoadItemsRequest) (*gen.LoadItemsResponse, error) {

	for _, item := range request.GetItems() {
		s.items[item.Code] += 1
	}

	return &gen.LoadItemsResponse{}, nil
}

func (s *sortingService) MoveItem(context.Context, *gen.MoveItemRequest) (*gen.MoveItemResponse, error) {
	// TODO: Implement this
	return nil, errors.New("not implemented")
}

func (s *sortingService) SelectItem(context.Context, *gen.SelectItemRequest) (*gen.SelectItemResponse, error) {
	if len(s.items) == 0 {
		fmt.Println("has no items")
		return nil, errors.New("no items to select from")
	}

	fmt.Println("has plenty of items")
	keys := make([]string, 0, len(s.items))
	for k := range s.items {
		keys = append(keys, k)
	}
	i := rand.Intn(len(keys))
	fmt.Println("selected item %s", keys[i])
	item := &gen.Item{}
	item.Code = keys[i]
	item.Label = "bbb"

	response := &gen.SelectItemResponse{}
	response.Item = item

	return response, nil
}
