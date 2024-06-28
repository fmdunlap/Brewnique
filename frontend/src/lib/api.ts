import { Comment, NewUser, Recipe, User } from "./types"

const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

export async function getUsers() {
    const response = await fetch(BACKEND_URL + '/v1/users')
    if (!response.ok) {
        throw new Error('Failed to fetch users')
    }
    return response.json() as Promise<User[]>
}

export async function createUser(newUser: NewUser) {
    const response = await fetch(BACKEND_URL + '/v1/users', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(newUser)
    })
    if (!response.ok) {
        throw new Error('Failed to create user')
    }
    return response.json() as Promise<User>
}

export async function deleteUserById(userId: number) {
    const response = await fetch(BACKEND_URL + '/v1/users/' + userId, {
        method: 'DELETE'
    })
    if (!response.ok) {
        throw new Error('Failed to delete user')
    }
    return response.json() as Promise<User>
}

export async function getRecipes() {
    const response = await fetch(BACKEND_URL + '/v1/recipes')
    if (!response.ok) {
        throw new Error('Failed to fetch recipes')
    }
    return response.json() as Promise<Recipe[]>
}

export async function getComments(recipeId: number) {
    const response = await fetch(BACKEND_URL + '/v1/comments/recipe/' + recipeId)
    if (!response.ok) {
        throw new Error('Failed to fetch comments')
    }
    return response.json() as Promise<Comment[]>
}