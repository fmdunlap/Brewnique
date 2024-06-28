import { RecipeCategory } from "@/lib/types"
import { BACKEND_URL } from "./api"

export const DEFAULT_CATEGORY_ID = 1

export async function getCategories() {
    const response = await fetch(BACKEND_URL + '/v1/recipes/categories')
    if (!response.ok) {
        throw new Error('Failed to fetch recipe categories')
    }
    return response.json() as Promise<RecipeCategory[]>
}

export async function getSubcategories(parentCategoryId: number) {
    const response = await fetch(BACKEND_URL + '/v1/recipes/categories/' + parentCategoryId)
    if (!response.ok) {
        throw new Error('Failed to fetch subcategories for category: ' + parentCategoryId)
    }
    return response.json() as Promise<RecipeCategory[]>
}