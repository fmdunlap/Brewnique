import { RecipeAttributeOptions } from "../types"
import { BACKEND_URL } from "./api"

export async function getAttributes() {
    const response = await fetch(BACKEND_URL + '/v1/recipes/attributes')
    if (!response.ok) {
        throw new Error('Failed to fetch attributes')
    }
    return response.json() as Promise<RecipeAttributeOptions[]>
}