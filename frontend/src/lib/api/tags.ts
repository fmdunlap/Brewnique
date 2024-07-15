import { RecipeTag } from "../types"
import { BACKEND_URL } from "./api"

export async function getTags() {
    const response = await fetch(BACKEND_URL + '/v1/recipes/tags')
    if (!response.ok) {
        throw new Error('Failed to fetch tags')
    }
    return response.json() as Promise<RecipeTag[]>
}