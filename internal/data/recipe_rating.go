package data

import "time"

type RecipeRating struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    int64     `json:"user_id"`
	RecipeId  int64     `json:"recipe_id"`
	Rating    int       `json:"rating"`
}
