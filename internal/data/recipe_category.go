package data

type RecipeCategory struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	ParentId *int64 `json:"parent_id"`
}
