export type User = {
    id: number,
    username: string,
    email: string,
    created_at: string,
    updated_at: string
}

export type NewUser = {
    username: string,
    email: string
}

export type Recipe = {
    id: number,
    name: string,
    author: User,
    created_at?: Date,
    updated_at?: Date,
    ingredients: string[],
    instructions: string[]
    category: string,
    subcategory: string,
    attributes: RecipeAttribute[],
    tags: string[]
}

export type NewRecipe = {
    name: string,
    author_id: number,
    ingredients: string[],
    instructions: string[],
    category_id: number,
    subcategory_id: number,
    attribute_ids: number[],
    tag_ids: number[]
}

export type RecipeCategory = {
    id: number,
    name: string,
    parent_id: number | null
}

export type RecipeAttribute = {
    name: string,
    type: string,
    value: string
}

export type RecipeAttributeOptions = {
    id: number,
    name: string,
    type: string,
    values: RecipeAttributeOptionValue[]
}

export type RecipeAttributeOptionValue = {
    id: number,
    value: string
}

export type RecipeTag = {
    id: number,
    recipe_id: number,
    tag_id: number,
    name: string
}

export type Comment = {
    id: number,
    content: string,
    author: User,
    recipe_id: number,
    children: Comment[],
    created_at: string,
    updated_at: string
}