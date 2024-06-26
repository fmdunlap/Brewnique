package service

import (
	"fmt"

	"brewnique.fdunlap.com/internal/data"
)

type RecipeProvider interface {
	SetUserRecipeRating(recipeId int64, ratingVal int, userId int64) (*data.RecipeRating, error)
	DeleteRecipe(id int64) error
	GetRecipe(id int64) (*data.Recipe, error)
	GetRecipesByUserId(userId int64) ([]*data.Recipe, error)
	GetRecipeRatings(recipeId int64) ([]*data.RecipeRating, error)
	ListCategories() ([]*data.RecipeCategory, error)
	ListSubcategories(categoryId int64) ([]*data.RecipeCategory, error)
	ListRecipes() ([]*data.Recipe, error)
	ListRecipesByAuthorId(userId int64) ([]*data.Recipe, error)
	PutRecipe(recipe *data.Recipe) (*data.Recipe, error)
	UpdateRecipe(recipe *data.Recipe) (*data.Recipe, error)
	GetAttributes() ([]*data.Attribute, error)
	GetAttributeValues(attributeId int64) ([]*data.AttributeValue, error)
	GetTags() ([]*data.Tag, error)
	GetRecipeTags(recipeId int64) ([]*data.RecipeTag, error)
	PutRecipeTags(recipeId int64, tags []*data.RecipeTag) error
}

type RecipeService struct {
	recipeProvider RecipeProvider
}

func NewRecipeService(recipeProvider RecipeProvider) *RecipeService {
	return &RecipeService{
		recipeProvider: recipeProvider,
	}
}

func (s *RecipeService) CreateRecipe(name string, authorId int64, ingredients []string, instructions []string) (*data.Recipe, error) {
	if name == "" {
		return nil, fmt.Errorf("name is not set")
	}
	if authorId == 0 {
		return nil, fmt.Errorf("authorId is not set")
	}
	if len(ingredients) == 0 {
		return nil, fmt.Errorf("ingredients is not set")
	}
	if len(instructions) == 0 {
		return nil, fmt.Errorf("instructions is not set")
	}

	return s.recipeProvider.PutRecipe(&data.Recipe{
		Name:         name,
		Ingredients:  ingredients,
		Instructions: instructions,
		AuthorId:     authorId,
	})
}

func (s *RecipeService) GetRecipe(id int64) (*data.Recipe, error) {
	return s.recipeProvider.GetRecipe(id)
}

func (s *RecipeService) GetUserRecipes(userId int64) ([]*data.Recipe, error) {
	return s.recipeProvider.ListRecipesByAuthorId(userId)
}

func (s *RecipeService) ListRecipes() ([]*data.Recipe, error) {
	return s.recipeProvider.ListRecipes()
}

func (s *RecipeService) ListRecipesByAuthorId(userId int64) ([]*data.Recipe, error) {
	return s.recipeProvider.ListRecipesByAuthorId(userId)
}

func (s *RecipeService) UpdateRecipe(recipe *data.Recipe) (*data.Recipe, error) {
	return s.recipeProvider.UpdateRecipe(recipe)
}

func (s *RecipeService) DeleteRecipe(id int64) error {
	return s.recipeProvider.DeleteRecipe(id)
}

func (s *RecipeService) GetAverageRecipeRating(recipeId int64) (float64 /* rating */, error) {
	ratings, err := s.recipeProvider.GetRecipeRatings(recipeId)
	if err != nil {
		return 0, err
	}

	if len(ratings) == 0 {
		return 0, nil
	}

	ratingSum := 0
	for _, rating := range ratings {
		ratingSum += rating.Rating
	}

	return float64(ratingSum) / float64(len(ratings)), nil
}

func (s *RecipeService) SetUserRecipeRating(recipeId int64, ratingVal int, userId int64) error {
	_, err := s.recipeProvider.SetUserRecipeRating(recipeId, ratingVal, userId)
	return err
}

func (s *RecipeService) ListCategories() ([]*data.RecipeCategory, error) {
	return s.recipeProvider.ListCategories()
}

func (s *RecipeService) ListSubcategories(categoryId int64) ([]*data.RecipeCategory, error) {
	return s.recipeProvider.ListSubcategories(categoryId)
}

func (s *RecipeService) GetAttributes() ([]*data.Attribute, error) {
	attributes, err := s.recipeProvider.GetAttributes()
	if err != nil {
		return nil, err
	}

	for _, attribute := range attributes {
		attribute.Values = make([]data.AttributeValue, 0)
		attributeVals, err := s.recipeProvider.GetAttributeValues(attribute.Id)
		if err != nil {
			return nil, err
		}
		for _, attributeVal := range attributeVals {
			attribute.Values = append(attribute.Values, *attributeVal)
		}
	}

	return attributes, nil
}

func (s *RecipeService) GetAttributeValues(attributeId int64) ([]*data.AttributeValue, error) {
	return s.recipeProvider.GetAttributeValues(attributeId)
}

func (s *RecipeService) GetTags() ([]*data.Tag, error) {
	return s.recipeProvider.GetTags()
}
