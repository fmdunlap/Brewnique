package psql

import (
	"brewnique.fdunlap.com/internal/data"
	"github.com/lib/pq"
	"log"
	"time"
)

type RecipeDbRow struct {
	Id           int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	AuthorId     int64
	Name         string
	Ingredients  pq.StringArray
	Instructions pq.StringArray
	Version      int
}

func (r *RecipeDbRow) ToRecipe() data.Recipe {
	return data.Recipe{
		Id:           r.Id,
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
		AuthorId:     r.AuthorId,
		Name:         r.Name,
		Ingredients:  r.Ingredients,
		Instructions: r.Instructions,
		Version:      r.Version,
	}
}

type RecipeRatingDbRow struct {
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	UserId    int64
	RecipeId  int64
	Rating    int
}

func (r *RecipeRatingDbRow) ToRecipeRating() data.RecipeRating {
	return data.RecipeRating{
		Id:       r.Id,
		UserId:   r.UserId,
		RecipeId: r.RecipeId,
		Rating:   r.Rating,
	}
}

type RecipeTagDbRow struct {
	Id       int64
	RecipeId int64
	TagId    int64
}

type RecipeTagDbRowWithName struct {
	Id       int64
	RecipeId int64
	TagId    int64
	Name     string
}

func (r *RecipeTagDbRowWithName) ToRecipeTag() data.RecipeTag {
	return data.RecipeTag{
		Id:       r.Id,
		RecipeId: r.RecipeId,
		TagId:    r.TagId,
		Name:     r.Name,
	}
}

func (p PostgresProvider) GetRecipe(id int64) (*data.Recipe, error) {
	res := p.db.QueryRow("SELECT id, created_at, updated_at, name, ingredients, instructions, version FROM recipes WHERE id = $1", id)

	recipeRow := RecipeDbRow{}
	err := res.Scan(&recipeRow)
	if err != nil {
		return nil, err
	}

	recipe := recipeRow.ToRecipe()

	return &recipe, nil
}

func (p PostgresProvider) ListRecipes() ([]*data.Recipe, error) {
	rows, err := p.db.Query("SELECT id, author_id, created_at, updated_at, name, ingredients, instructions, version FROM recipes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes []*data.Recipe
	for rows.Next() {
		var recipeRow RecipeDbRow
		err = rows.Scan(&recipeRow.Id, &recipeRow.AuthorId, &recipeRow.CreatedAt, &recipeRow.UpdatedAt, &recipeRow.Name, (*pq.StringArray)(&recipeRow.Ingredients), (*pq.StringArray)(&recipeRow.Instructions), &recipeRow.Version)
		if err != nil {
			return nil, err
		}
		recipe := recipeRow.ToRecipe()
		recipes = append(recipes, &recipe)
	}

	return recipes, nil
}

func (p PostgresProvider) ListRecipesByAuthorId(userId int64) ([]*data.Recipe, error) {
	rows, err := p.db.Query("SELECT id, created_at, updated_at, name, ingredients, instructions, version FROM recipes WHERE author_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes []*data.Recipe
	for rows.Next() {
		var recipeRow RecipeDbRow
		err = rows.Scan(
			&recipeRow.Id,
			&recipeRow.CreatedAt,
			&recipeRow.UpdatedAt,
			&recipeRow.Name,
			&recipeRow.Ingredients,
			&recipeRow.Instructions,
			&recipeRow.Version,
		)
		if err != nil {
			return nil, err
		}

		recipe := recipeRow.ToRecipe()
		recipes = append(recipes, &recipe)
	}

	return recipes, nil
}

func (p PostgresProvider) PutRecipe(recipe *data.Recipe) (*data.Recipe, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	log.Printf("Creating recipe with name %s, ingredients %v, instructions %v", recipe.Name, recipe.Ingredients, recipe.Instructions)

	var insertedRecipe RecipeDbRow
	err = tx.QueryRow(`
        INSERT INTO recipes (name, author_id, ingredients, instructions, version)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at, updated_at, author_id, name, ingredients, instructions, version
    `, recipe.Name, recipe.AuthorId, pq.Array(recipe.Ingredients), pq.Array(recipe.Instructions), recipe.Version).Scan(
		&insertedRecipe.Id,
		&insertedRecipe.CreatedAt,
		&insertedRecipe.UpdatedAt,
		&insertedRecipe.AuthorId,
		&insertedRecipe.Name,
		&insertedRecipe.Ingredients,
		&insertedRecipe.Instructions,
		&insertedRecipe.Version,
	)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	recipeResult := insertedRecipe.ToRecipe()

	return &recipeResult, nil
}

func (p PostgresProvider) UpdateRecipe(recipe *data.Recipe) (*data.Recipe, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var updatedRecipe RecipeDbRow
	err = tx.QueryRow(`
        UPDATE recipes
        SET name = $1, ingredients = $2, instructions = $3, version = $4, updated_at = NOW()
        WHERE id = $5
        RETURNING id, created_at, updated_at, author_id, name, ingredients, instructions, version
    `, recipe.Name, pq.Array(recipe.Ingredients), pq.Array(recipe.Instructions), recipe.Version, recipe.Id).Scan(
		&updatedRecipe.Id,
		&updatedRecipe.CreatedAt,
		&updatedRecipe.UpdatedAt,
		&updatedRecipe.AuthorId,
		&updatedRecipe.Name,
		&updatedRecipe.Ingredients,
		&updatedRecipe.Instructions,
		&updatedRecipe.Version,
	)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	recipeResult := updatedRecipe.ToRecipe()
	return &recipeResult, nil
}

func (p PostgresProvider) DeleteRecipe(id int64) error {
	_, err := p.db.Exec("DELETE FROM recipes WHERE id = $1", id)
	return err
}

func (p PostgresProvider) GetRecipesByUserId(userId int64) ([]*data.Recipe, error) {
	rows, err := p.db.Query("SELECT id, created_at, updated_at, author_id, name, ingredients, instructions, version FROM recipes WHERE author_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	recipes := []*data.Recipe{}
	for rows.Next() {
		var recipeRow RecipeDbRow
		err = rows.Scan(
			&recipeRow.Id,
			&recipeRow.CreatedAt,
			&recipeRow.UpdatedAt,
			&recipeRow.AuthorId,
			&recipeRow.Name,
			&recipeRow.Ingredients,
			&recipeRow.Instructions,
			&recipeRow.Version,
		)
		if err != nil {
			return nil, err
		}

		recipe := recipeRow.ToRecipe()
		recipes = append(recipes, &recipe)
	}

	return recipes, nil
}

func (p PostgresProvider) GetRecipeRatings(recipeId int64) ([]*data.RecipeRating, error) {
	ratings := []*data.RecipeRating{}
	rows, err := p.db.Query("SELECT id, created_at, updated_at, user_id, recipe_id, rating FROM recipe_ratings WHERE recipe_id = $1", recipeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ratingRow RecipeRatingDbRow
		err = rows.Scan(
			&ratingRow.Id,
			&ratingRow.CreatedAt,
			&ratingRow.UpdatedAt,
			&ratingRow.UserId,
			&ratingRow.RecipeId,
			&ratingRow.Rating,
		)
		if err != nil {
			return nil, err
		}

		rating := ratingRow.ToRecipeRating()
		ratings = append(ratings, &rating)
	}

	return ratings, nil
}

// Sets the rating for a user on a recipe. Each user can only rate a recipe once.
func (p PostgresProvider) SetUserRecipeRating(recipeId int64, ratingVal int, userId int64) (*data.RecipeRating, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Check for an existing rating
	// If the rating already exists, and the ratingVal is the same, do nothing
	// If the rating already exists, and the ratingVal is different, update the rating
	// If the rating does not exist, create a new rating

	var existingRating RecipeRatingDbRow
	err = tx.QueryRow("SELECT id, created_at, updated_at, user_id, recipe_id, rating FROM recipe_ratings WHERE recipe_id = $1 AND user_id = $2", recipeId, userId).Scan(
		&existingRating.Id,
		&existingRating.CreatedAt,
		&existingRating.UpdatedAt,
		&existingRating.UserId,
		&existingRating.RecipeId,
		&existingRating.Rating,
	)

	// Rating exists, check if ratingVal is the same. If not, or update it
	if err == nil {
		if existingRating.Rating == ratingVal {
			rating := existingRating.ToRecipeRating()
			return &rating, nil
		}

		var updatedRating RecipeRatingDbRow
		updatedRating.RecipeId = recipeId
		updatedRating.UserId = userId
		updatedRating.Rating = ratingVal
		updatedRating.UpdatedAt = time.Now()

		err = tx.QueryRow("UPDATE recipe_ratings SET ratingVal = $1, updated_at = NOW() WHERE recipe_id = $2 AND user_id = $3", ratingVal, recipeId, userId).Scan(
			&updatedRating.Id,
			&updatedRating.CreatedAt,
			&updatedRating.UpdatedAt,
			&updatedRating.UserId,
			&updatedRating.RecipeId,
			&updatedRating.Rating,
		)
		if err != nil {
			return nil, err
		}

		rating := updatedRating.ToRecipeRating()
		return &rating, nil
	}

	insertedRating := RecipeRatingDbRow{}
	// Rating does not exist, create it
	err = tx.QueryRow("INSERT INTO recipe_ratings (recipe_id, user_id, ratingVal) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at, recipe_id, user_id, ratingVal",
		recipeId,
		userId,
		ratingVal,
	).Scan(
		&insertedRating.Id,
		&insertedRating.CreatedAt,
		&insertedRating.UpdatedAt,
		&insertedRating.RecipeId,
		&insertedRating.UserId,
		&insertedRating.Rating,
	)

	rating := insertedRating.ToRecipeRating()
	return &rating, nil
}

func (p PostgresProvider) GetRecipeTags(recipeId int64) ([]*data.RecipeTag, error) {
	// Join the recipe_tags table with the tags table to get the tag name
	rows, err := p.db.Query("SELECT id, recipe_id, tag_id, name FROM recipe_tags JOIN tags ON recipe_tags.tag_id = tags.id WHERE recipe_id = $1", recipeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipeTags []*data.RecipeTag
	for rows.Next() {
		var recipeTagRow RecipeTagDbRowWithName
		err = rows.Scan(
			&recipeTagRow.Id,
			&recipeTagRow.RecipeId,
			&recipeTagRow.TagId,
			&recipeTagRow.Name,
		)
		if err != nil {
			return nil, err
		}
		recipeTag := recipeTagRow.ToRecipeTag()
		recipeTags = append(recipeTags, &recipeTag)
	}

	return recipeTags, nil
}

func (p PostgresProvider) PutRecipeTags(recipeId int64, tags []*data.RecipeTag) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Clear out any existing tags for this recipe
	_, err = tx.Exec("DELETE FROM recipe_tags WHERE recipe_id = $1", recipeId)
	if err != nil {
		return err
	}

	// Insert the new tags
	for _, tag := range tags {
		insertedRecipeTag := RecipeTagDbRow{
			RecipeId: recipeId,
			TagId:    tag.TagId,
		}
		err = tx.QueryRow("INSERT INTO recipe_tags (recipe_id, tag_id) VALUES ($1, $2) RETURNING id, recipe_id, tag_id",
			recipeId,
			tag.TagId,
		).Scan(
			&insertedRecipeTag.Id,
			&insertedRecipeTag.RecipeId,
			&insertedRecipeTag.TagId,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
