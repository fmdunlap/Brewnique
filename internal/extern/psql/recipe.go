package psql

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"brewnique.fdunlap.com/internal/data"
	"github.com/lib/pq"
)

type RecipeDbRow struct {
	Id            int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	AuthorId      int64
	Name          string
	Ingredients   pq.StringArray
	Instructions  pq.StringArray
	CategoryId    int64
	SubcategoryId int64
	Version       int
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

const recipeQuery = `
SELECT
	recipes.id,
	recipes.created_at,
	recipes.updated_at,
	recipes.name,
	recipes.ingredients,
	recipes.instructions,
	recipes.version,
	recipes.author_id,
	recipes.category_id,
	recipes.subcategory_id,
	c.name AS category_name,
	sc.name AS subcategory_name,
	(
		SELECT json_build_object(
			'id', u.id,
			'username', u.username,
			'email', u.email,
			'created_at', u.created_at,
			'updated_at', u.updated_at
		)
		FROM users AS u
		WHERE u.id = recipes.author_id
	) AS author,
	(
		SELECT json_agg(json_build_object(
			'type', a.type,
			'name', a.name,
			'value', av.value
		))
		FROM recipe_attributes AS ra
		INNER JOIN attributes AS a ON a.id = ra.attribute_value_id
		INNER JOIN attribute_values AS av ON av.id = ra.attribute_value_id
		WHERE ra.recipe_id = recipes.id
	) AS attribute_values,
	(
		SELECT
			json_agg(json_build_object(
				'id', rt.id,
				'recipe_id', recipes.id,
				'tag_id', t.id,
				'name', t.name
			))
		FROM recipe_tags AS rt
		INNER JOIN tags AS t ON t.id = rt.tag_id
		WHERE rt.recipe_id = recipes.id
	) AS tags
FROM recipes
LEFT JOIN categories AS c ON c.id = recipes.category_id
LEFT JOIN categories AS sc ON sc.id = recipes.subcategory_id
`

func convertRecipeFieldsToRecipe(recipeRow RecipeDbRow, categoryName, subcategoryName, userValueJSON, attributeValuesJSON, tagsJSON sql.NullString) (*data.Recipe, error) {
	recipeAttributes := make([]data.AttributeDTO, 0)
	if attributeValuesJSON.Valid {
		err := json.Unmarshal([]byte(attributeValuesJSON.String), &recipeAttributes)
		if err != nil {
			return nil, err
		}
	}

	recipeTags := make([]data.RecipeTag, 0)
	if tagsJSON.Valid {
		err := json.Unmarshal([]byte(tagsJSON.String), &recipeTags)
		if err != nil {
			return nil, err
		}
	}
	tags := make([]data.Tag, 0)
	for _, tag := range recipeTags {
		tags = append(tags, data.Tag{
			Id:   tag.Id,
			Name: tag.Name,
		})
	}

	author := data.User{}
	if userValueJSON.Valid {
		err := json.Unmarshal([]byte(userValueJSON.String), &author)
		if err != nil {
			return nil, err
		}
	}

	recipe := data.Recipe{
		Id:            recipeRow.Id,
		CreatedAt:     recipeRow.CreatedAt,
		UpdatedAt:     recipeRow.UpdatedAt,
		AuthorId:      recipeRow.AuthorId,
		Author:        author,
		Name:          recipeRow.Name,
		Ingredients:   recipeRow.Ingredients,
		Instructions:  recipeRow.Instructions,
		Category:      categoryName.String,
		CategoryId:    recipeRow.CategoryId,
		Subcategory:   subcategoryName.String,
		SubcategoryId: recipeRow.SubcategoryId,
		Version:       recipeRow.Version,
		Attributes:    recipeAttributes,
		Tags:          tags,
	}

	return &recipe, nil
}

func convertRecipeQueryRowResult(row *sql.Row) (*data.Recipe, error) {
	recipeRow := RecipeDbRow{}
	var categoryName, subcategoryName, userValueJSON, attributeValuesJSON, tagsJSON sql.NullString
	err := row.Scan(
		&recipeRow.Id,
		&recipeRow.CreatedAt,
		&recipeRow.UpdatedAt,
		&recipeRow.Name,
		(*pq.StringArray)(&recipeRow.Ingredients),
		(*pq.StringArray)(&recipeRow.Instructions),
		&recipeRow.Version,
		&recipeRow.AuthorId,
		&recipeRow.CategoryId,
		&recipeRow.SubcategoryId,
		&categoryName,
		&subcategoryName,
		&userValueJSON,
		&attributeValuesJSON,
		&tagsJSON,
	)

	if err != nil {
		return nil, err
	}

	return convertRecipeFieldsToRecipe(recipeRow, categoryName, subcategoryName, userValueJSON, attributeValuesJSON, tagsJSON)
}

func converRecipeQueryResult(rows *sql.Rows) ([]*data.Recipe, error) {
	recipes := make([]*data.Recipe, 0)
	for rows.Next() {
		recipeRow := RecipeDbRow{}
		var categoryName, subcategoryName, userValueJSON, attributeValuesJSON, tagsJSON sql.NullString
		err := rows.Scan(
			&recipeRow.Id,
			&recipeRow.CreatedAt,
			&recipeRow.UpdatedAt,
			&recipeRow.Name,
			(*pq.StringArray)(&recipeRow.Ingredients),
			(*pq.StringArray)(&recipeRow.Instructions),
			&recipeRow.Version,
			&recipeRow.AuthorId,
			&recipeRow.CategoryId,
			&recipeRow.SubcategoryId,
			&categoryName,
			&subcategoryName,
			&userValueJSON,
			&attributeValuesJSON,
			&tagsJSON,
		)

		if err != nil {
			return nil, err
		}

		recipe, err := convertRecipeFieldsToRecipe(recipeRow, categoryName, subcategoryName, userValueJSON, attributeValuesJSON, tagsJSON)
		if err != nil {
			return nil, err
		}

		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (p PostgresProvider) GetRecipe(id int64) (*data.Recipe, error) {
	// Includes a subquery to get the recipe's attributes
	res := p.db.QueryRow(recipeQuery+"WHERE recipes.id = $1", id)
	return convertRecipeQueryRowResult(res)
}

func (p PostgresProvider) ListRecipes() ([]*data.Recipe, error) {
	// Select all recipes using a join to get author info, category, subcategory, attriubtes, and tags
	rows, err := p.db.Query(recipeQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return converRecipeQueryResult(rows)
}

func (p PostgresProvider) ListRecipesByAuthorId(userId int64) ([]*data.Recipe, error) {
	rows, err := p.db.Query(recipeQuery+"WHERE recipes.author_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return converRecipeQueryResult(rows)
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
