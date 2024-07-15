package service

import (
	"brewnique.fdunlap.com/internal/data"
	"fmt"
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
	ListAttributeDefinitions() ([]*data.AttributeDefinition, error)
	GetAttributeValues(attributeId int64) ([]*data.AttributeValue, error)
	ListTags() ([]*data.Tag, error)
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

type CreateRecipeParams struct {
	Name          string   `json:"name"`
	AuthorId      int64    `json:"author_id"`
	Ingredients   []string `json:"ingredients"`
	Instructions  []string `json:"instructions"`
	CategoryId    int64    `json:"category"`
	SubcategoryId int64    `json:"subcategory"`
	AttributeIds  []int64  `json:"attributes"`
	TagIds        []int64  `json:"tags"`
}

func (p *CreateRecipeParams) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("name is not set")
	}
	if p.AuthorId == 0 {
		return fmt.Errorf("authorId is not set")
	}
	if len(p.Ingredients) == 0 {
		return fmt.Errorf("ingredients is not set")
	}
	if len(p.Instructions) == 0 {
		return fmt.Errorf("instructions is not set")
	}

	return nil
}

func (s *RecipeService) getAttributesFromValueIds(attributeIds []int64) ([]data.AttributeDTO, error) {
	attributes := make([]data.AttributeDTO, 0)
	attributeDefinitions, err := s.recipeProvider.ListAttributeDefinitions()
	if err != nil {
		return nil, err
	}
	for _, attributeDefinition := range attributeDefinitions {
		for _, attributeValue := range attributeIds {
			if attributeDefinition.Id == attributeValue {
				attributes = append(attributes, data.AttributeDTO{
					Name:  attributeDefinition.Name,
					Type:  attributeDefinition.Type,
					Value: attributeDefinition.Values[attributeValue].Value,
				})
			}
		}
	}

	return attributes, nil
}

func (s *RecipeService) getTagsFromTagIds(tagIds []int64) ([]data.Tag, error) {
	tagDefinitions, err := s.recipeProvider.ListTags()
	if err != nil {
		return nil, err
	}
	tagMap := make(map[int64]*data.Tag)
	for _, tagDefinition := range tagDefinitions {
		tagMap[tagDefinition.Id] = tagDefinition
	}
	tags := make([]data.Tag, 0)
	for _, tagId := range tagIds {
		tag, ok := tagMap[tagId]
		if !ok {
			return nil, fmt.Errorf("tag with id %d not found", tagId)
		}
		tags = append(tags, *tag)
	}

	return tags, nil
}

func (s *RecipeService) CreateRecipe(params CreateRecipeParams) (*data.Recipe, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	attributes, err := s.getAttributesFromValueIds(params.AttributeIds)
	if err != nil {
		return nil, err
	}

	tags, err := s.getTagsFromTagIds(params.TagIds)
	if err != nil {
		return nil, err
	}

	return s.recipeProvider.PutRecipe(&data.Recipe{
		AuthorId:      params.AuthorId,
		Name:          params.Name,
		Ingredients:   params.Ingredients,
		Instructions:  params.Instructions,
		CategoryId:    params.CategoryId,
		SubcategoryId: params.SubcategoryId,
		Attributes:    attributes,
		Tags:          tags,
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

func (s *RecipeService) GetAttributes() ([]*data.AttributeDefinition, error) {
	attributes, err := s.recipeProvider.ListAttributeDefinitions()
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
	return s.recipeProvider.ListTags()
}
