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
    author_id: number,
    author: User,
    created_at: string,
    updated_at: string
    ingredients: string[],
    instructions: string[]
    category: RecipeCategory,
    subcategory: RecipeCategory,
    attributes: RecipeAttribute[],
    tags: RecipeTag[]
}

export type RecipeCategory = {
    id: number,
    name: string,
    parent_id: number | null
}

export type RecipeAttribute = {
    id: number,
    name: string,
    type: string,
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