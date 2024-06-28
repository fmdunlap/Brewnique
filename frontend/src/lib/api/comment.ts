import { BACKEND_URL } from "./api";

import { Comment } from "../types";


export async function getComments(recipeId: number) {
    const response = await fetch(BACKEND_URL + '/v1/comments/recipe/' + recipeId)
    if (!response.ok) {
        throw new Error('Failed to fetch comments')
    }
    return response.json() as Promise<Comment[]>
}