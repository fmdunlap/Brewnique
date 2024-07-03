import { BACKEND_URL } from "./api";

import { NewRecipe, Recipe } from "../types";


export async function getRecipes() {
    const response = await fetch(BACKEND_URL + '/v1/recipes')
    if (!response.ok) {
        throw new Error('Failed to fetch recipes')
    }
    return response.json() as Promise<Recipe[]>
}

export async function createRecipe(recipe: NewRecipe) {
    const response = await fetch(BACKEND_URL + '/v1/recipes', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(recipe)
    })
    if (!response.ok) {
        throw new Error('Failed to create recipe')
    }
    return response.json() as Promise<Recipe>

}