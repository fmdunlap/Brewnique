package data

import (
	"fmt"
	"log"
	"testing"
	"time"
)

type TestRecipeProvider struct {
	recipes       map[int64]*Recipe
	recipeRatings map[int64]map[int64]*RecipeRating
	nextID        int64
}

func NewTestRecipeProvider() *TestRecipeProvider {
	return &TestRecipeProvider{
		recipes:       make(map[int64]*Recipe),
		recipeRatings: make(map[int64]map[int64]*RecipeRating),
		nextID:        1,
	}
}

func (p *TestRecipeProvider) PutRecipe(recipe *Recipe) (*Recipe, error) {
	// Model the author_id and name uniqueness index
	for _, existingRecipe := range p.recipes {
		if existingRecipe.AuthorId == recipe.AuthorId && existingRecipe.Name == recipe.Name {
			return nil, fmt.Errorf("recipe with name %s and authorId %d already exists", recipe.Name, recipe.AuthorId)
		}
	}

	recipe.Id = p.nextID
	recipe.CreatedAt = time.Now()
	recipe.UpdatedAt = time.Now()
	p.recipes[recipe.Id] = recipe
	p.nextID++
	return recipe, nil
}

func (p *TestRecipeProvider) GetRecipe(id int64) (*Recipe, error) {
	recipe, ok := p.recipes[id]
	if !ok {
		return nil, fmt.Errorf("recipe with ID %d not found", id)
	}
	return recipe, nil
}

func (p *TestRecipeProvider) ListRecipes() ([]*Recipe, error) {
	var recipes []*Recipe
	for _, recipe := range p.recipes {
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

func (p *TestRecipeProvider) ListRecipesByAuthorId(userId int64) ([]*Recipe, error) {
	var recipes []*Recipe
	for _, recipe := range p.recipes {
		if recipe.AuthorId == userId {
			recipes = append(recipes, recipe)
		}
	}
	return recipes, nil
}

func (p *TestRecipeProvider) UpdateRecipe(recipe *Recipe) (*Recipe, error) {
	existingRecipe, ok := p.recipes[recipe.Id]
	if !ok {
		return nil, fmt.Errorf("recipe with ID %d not found", recipe.Id)
	}
	if recipe.Name != "" {
		existingRecipe.Name = recipe.Name
	}
	if recipe.Ingredients != nil {
		existingRecipe.Ingredients = recipe.Ingredients
	}
	if recipe.Instructions != nil {
		existingRecipe.Instructions = recipe.Instructions
	}
	if recipe.Version != 0 {
		existingRecipe.Version = recipe.Version
	}

	p.recipes[recipe.Id].UpdatedAt = time.Now()
	p.recipes[recipe.Id].Version = recipe.Version
	p.recipes[recipe.Id].Name = recipe.Name
	p.recipes[recipe.Id].Ingredients = recipe.Ingredients
	p.recipes[recipe.Id].Instructions = recipe.Instructions

	return existingRecipe, nil
}

func (p *TestRecipeProvider) DeleteRecipe(id int64) error {
	if _, ok := p.recipes[id]; !ok {
		return fmt.Errorf("recipe with ID %d not found", id)
	}
	delete(p.recipes, id)
	return nil
}

func (p *TestRecipeProvider) GetRecipeRatings(recipeId int64) ([]*RecipeRating, error) {
	log.Printf("GetRecipeRatings(%d), %v", recipeId, p.recipeRatings)
	if _, err := p.GetRecipe(recipeId); err != nil {
		return make([]*RecipeRating, 0), fmt.Errorf("recipe with ID %d not found", recipeId)
	}

	if _, ok := p.recipeRatings[recipeId]; !ok {
		return make([]*RecipeRating, 0), nil
	}
	ratings := make([]*RecipeRating, 0)
	for _, recipeRating := range p.recipeRatings[recipeId] {
		ratings = append(ratings, recipeRating)
	}
	return ratings, nil
}

func (p *TestRecipeProvider) SetUserRecipeRating(recipeId int64, ratingVal int, userId int64) (*RecipeRating, error) {
	if _, ok := p.recipeRatings[recipeId]; !ok {
		p.recipeRatings[recipeId] = make(map[int64]*RecipeRating)
	}

	existingRating, ok := p.recipeRatings[recipeId][userId]
	if !ok {
		newRating := RecipeRating{
			Id:        p.nextID,
			RecipeId:  recipeId,
			UserId:    userId,
			Rating:    ratingVal,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		p.recipeRatings[recipeId][userId] = &newRating
		p.nextID++
		return &newRating, nil
	}
	existingRating.Rating = ratingVal
	existingRating.UpdatedAt = time.Now()
	p.recipeRatings[recipeId][userId] = existingRating

	return existingRating, nil

}

func (p *TestRecipeProvider) GetRecipesByUserId(userId int64) ([]*Recipe, error) {
	var recipes []*Recipe
	for _, recipe := range p.recipes {
		if recipe.AuthorId == userId {
			recipes = append(recipes, recipe)
		}
	}
	return recipes, nil
}

func (p *TestRecipeProvider) TearDown() {
	p.nextID = 1
	p.recipes = make(map[int64]*Recipe)
}

func TestRecipeService_CreateRecipe(t *testing.T) {
	provider := NewTestRecipeProvider()
	service := NewRecipeService(provider)

	type args struct {
		name         string
		authorId     int64
		ingredients  []string
		instructions []string
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
		preRun  func(t *testing.T, provider *TestRecipeProvider)
		expect  *Recipe
	}{
		{
			name: "create recipe",
			args: args{
				name:         "A recipe",
				authorId:     1,
				ingredients:  []string{"first ingredient", "second ingredient", "third ingredient"},
				instructions: []string{"first instruction", "second instruction", "third instruction"},
			},
			wantErr: false,
			expect: &Recipe{
				Id:       1,
				AuthorId: 1,
				Name:     "A recipe",
				Ingredients: []string{
					"first ingredient",
					"second ingredient",
					"third ingredient",
				},
				Instructions: []string{
					"first instruction",
					"second instruction",
					"third instruction",
				},
			},
		},
		{
			name: "create recipe with with same name and same authorId",
			args: args{
				name:         "A recipe",
				authorId:     1,
				ingredients:  []string{"first ingredient", "second ingredient", "third ingredient"},
				instructions: []string{"first instruction", "second instruction", "third instruction"},
			},
			wantErr: true,
			expect:  nil,
			preRun: func(t *testing.T, provider *TestRecipeProvider) {
				provider.PutRecipe(&Recipe{
					Name:         "A recipe",
					AuthorId:     1,
					Ingredients:  []string{"first ingredient", "second ingredient", "third ingredient"},
					Instructions: []string{"first instruction", "second instruction", "third instruction"},
				})
			},
		},
		{
			name: "create recipe with same name and different authorId",
			args: args{
				name:         "A recipe",
				authorId:     2,
				ingredients:  []string{"first ingredient", "second ingredient", "third ingredient"},
				instructions: []string{"first instruction", "second instruction", "third instruction"},
			},
			wantErr: false,
			expect: &Recipe{
				Id:       2,
				AuthorId: 2,
				Name:     "A recipe",
				Ingredients: []string{
					"first ingredient",
					"second ingredient",
					"third ingredient",
				},
				Instructions: []string{
					"first instruction",
					"second instruction",
					"third instruction",
				},
			},
			preRun: func(t *testing.T, provider *TestRecipeProvider) {
				provider.PutRecipe(&Recipe{
					Name:         "A recipe",
					AuthorId:     1,
					Ingredients:  []string{"first ingredient", "second ingredient", "third ingredient"},
					Instructions: []string{"first instruction", "second instruction", "third instruction"},
				})
			},
		},
		{
			name: "create recipe without authorId",
			args: args{
				name:         "A recipe",
				authorId:     0,
				ingredients:  []string{"first ingredient", "second ingredient", "third ingredient"},
				instructions: []string{"first instruction", "second instruction", "third instruction"},
			},
			wantErr: true,
			expect:  nil,
		},
		{
			name: "create recipe without ingredients",
			args: args{
				name:         "A recipe",
				authorId:     1,
				ingredients:  []string{},
				instructions: []string{"first instruction", "second instruction", "third instruction"},
			},
			wantErr: true,
			expect:  nil,
		},
		{
			name: "create recipe without instructions",
			args: args{
				name:         "A recipe",
				authorId:     1,
				ingredients:  []string{"first ingredient", "second ingredient", "third ingredient"},
				instructions: []string{},
			},
			wantErr: true,
			expect:  nil,
		},
		{
			name: "create recipe wihtout name",
			args: args{
				name:         "",
				authorId:     1,
				ingredients:  []string{"first ingredient", "second ingredient", "third ingredient"},
				instructions: []string{"first instruction", "second instruction", "third instruction"},
			},
			wantErr: true,
			expect:  nil,
		},
	}

	for _, tc := range testCases {
		if tc.preRun != nil {
			tc.preRun(t, provider)
		}
		t.Run(tc.name, func(t *testing.T) {
			recipe, err := service.CreateRecipe(tc.args.name, tc.args.authorId, tc.args.ingredients, tc.args.instructions)
			if tc.wantErr {
				if (err != nil) != tc.wantErr {
					t.Errorf("CreateRecipe() error = %v, wantErr %v", err, tc.wantErr)
					return
				}
				return
			}

			if tc.expect == nil {
				if recipe != nil {
					t.Errorf("CreateRecipe() = %v, want %v", recipe, tc.expect)
				}
				return
			}

			if recipe.Name != tc.expect.Name {
				t.Errorf("CreateRecipe() = %v, want %v", recipe, tc.expect)
			}
			if recipe.AuthorId != tc.expect.AuthorId {
				t.Errorf("CreateRecipe() = %v, want %v", recipe, tc.expect)
			}
			for i, ingredient := range tc.expect.Ingredients {
				if ingredient != tc.expect.Ingredients[i] {
					t.Errorf("CreateRecipe() = %v, want %v", recipe, tc.expect)
				}
			}
			for i, instruction := range tc.expect.Instructions {
				if instruction != tc.expect.Instructions[i] {
					t.Errorf("CreateRecipe() = %v, want %v", recipe, tc.expect)
				}
			}
		})
		provider.TearDown()
	}
}

func TestRecipeService_GetRecipeRatings(t *testing.T) {
	type args struct {
		recipeId int64
	}

	type expected struct {
		Ratings []*RecipeRating
		Rating  float64
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
		preRun  func(t *testing.T, provider *TestRecipeProvider) int64
		expect  expected
	}{
		{
			name: "get existing recipe ratings",
			args: args{
				recipeId: 1,
			},
			wantErr: false,
			preRun: func(t *testing.T, provider *TestRecipeProvider) int64 {
				rec, _ := provider.PutRecipe(&Recipe{
					Name:         "A recipe",
					AuthorId:     1,
					Ingredients:  []string{"first ingredient", "second ingredient", "third ingredient"},
					Instructions: []string{"first instruction", "second instruction", "third instruction"},
				})
				provider.SetUserRecipeRating(1, 1, 1)
				provider.SetUserRecipeRating(1, 2, 2)
				provider.SetUserRecipeRating(1, 3, 3)
				return rec.Id
			},
			expect: expected{
				Ratings: []*RecipeRating{
					{
						Id:       2,
						UserId:   1,
						RecipeId: 1,
						Rating:   1,
					},
					{
						Id:       3,
						UserId:   2,
						RecipeId: 1,
						Rating:   2,
					},
					{
						Id:       4,
						UserId:   3,
						RecipeId: 1,
						Rating:   3,
					},
				},
				Rating: 2,
			},
		},
		{
			name: "get non-existing recipe ratings",
			args: args{
				recipeId: 2,
			},
			wantErr: true,
			expect: expected{
				Ratings: nil,
				Rating:  0,
			},
		},
		{
			name: "get unrated recipe ratings",
			args: args{
				recipeId: 1,
			},
			wantErr: false,
			preRun: func(t *testing.T, provider *TestRecipeProvider) int64 {
				rec, _ := provider.PutRecipe(&Recipe{
					Name:         "A recipe",
					AuthorId:     1,
					Ingredients:  []string{"first ingredient", "second ingredient", "third ingredient"},
					Instructions: []string{"first instruction", "second instruction", "third instruction"},
				})
				return rec.Id
			},
			expect: expected{
				Ratings: nil,
				Rating:  0.0,
			},
		},
	}

	for _, tc := range testCases {
		provider := NewTestRecipeProvider()
		service := NewRecipeService(provider)

		recipeId := int64(0)

		if tc.preRun != nil {
			recipeId = tc.preRun(t, provider)
		}
		t.Run(tc.name, func(t *testing.T) {
			ratings, err := provider.GetRecipeRatings(recipeId)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetRecipeRatings() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			for _, expRating := range tc.expect.Ratings {
				found := false
				for _, rating := range ratings {
					if rating.Id == expRating.Id && rating.UserId == expRating.UserId && rating.RecipeId == expRating.RecipeId && rating.Rating == expRating.Rating {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("GetRecipeRatings() = %v, want %v", ratings, expRating)
				}
			}

			recipeRating, err := service.GetAverageRecipeRating(tc.args.recipeId)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetRecipeRating() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if recipeRating != tc.expect.Rating {
				t.Errorf("GetRecipeRating() = %v, want %v", recipeRating, tc.expect.Rating)
			}
		})
		provider.TearDown()
	}
}

func TestRecipeService_SetUserRecipeRating(t *testing.T) {
	type args struct {
		recipeId  int64
		ratingVal int
		userId    int64
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
		preRun  func(t *testing.T, provider *TestRecipeProvider) int64
		expect  float64
	}{
		{
			name: "set existing recipe rating",
			args: args{
				recipeId:  1,
				ratingVal: 5,
				userId:    1,
			},
			wantErr: false,
			preRun: func(t *testing.T, provider *TestRecipeProvider) int64 {
				rec, _ := provider.PutRecipe(&Recipe{
					Name:         "A recipe",
					AuthorId:     1,
					Ingredients:  []string{"first ingredient", "second ingredient", "third ingredient"},
					Instructions: []string{"first instruction", "second instruction", "third instruction"},
				})
				return rec.Id
			},
			expect: 5.0,
		},
		{
			name: "set existing recipe with multiple ratings",
			args: args{
				recipeId:  1,
				ratingVal: 5,
				userId:    1,
			},
			wantErr: false,
			preRun: func(t *testing.T, provider *TestRecipeProvider) int64 {
				rec, _ := provider.PutRecipe(&Recipe{
					Name:         "A recipe",
					AuthorId:     1,
					Ingredients:  []string{"first ingredient", "second ingredient", "third ingredient"},
					Instructions: []string{"first instruction", "second instruction", "third instruction"},
				})

				provider.SetUserRecipeRating(1, 1, 1) // < This should get overwritten
				provider.SetUserRecipeRating(1, 2, 2)
				provider.SetUserRecipeRating(1, 3, 3)

				return rec.Id
			},
			expect: 10.0 / 3.0,
		},
	}

	for _, tc := range testCases {
		provider := NewTestRecipeProvider()
		service := NewRecipeService(provider)

		recipeId := int64(0)

		if tc.preRun != nil {
			recipeId = tc.preRun(t, provider)
		}
		t.Run(tc.name, func(t *testing.T) {
			err := service.SetUserRecipeRating(recipeId, tc.args.ratingVal, tc.args.userId)
			if (err != nil) != tc.wantErr {
				t.Errorf("SetUserRecipeRating() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			rating, err := service.GetAverageRecipeRating(recipeId)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetRecipeRatings() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if rating != tc.expect {
				t.Errorf("SetUserRecipeRating() = %v, want %v", rating, tc.expect)
			}
		})
		provider.TearDown()
	}
}
