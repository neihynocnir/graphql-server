package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/neihynocnir/graphql-server/graph/generated"
	"github.com/neihynocnir/graphql-server/graph/model"
)

func (r *queryResolver) FindCategories(ctx context.Context) (*model.Categories, error) {

	baseURL := os.Getenv("BASE_URL")
	token := os.Getenv("TOKEN")

	resp, err := http.Get(baseURL + `?token=` + token)
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	err = json.Unmarshal(body, &r.categories)
	if err != nil {
		print(err)
	}

	return r.categories, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
