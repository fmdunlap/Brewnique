import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader } from '@/components/ui/card'
import { Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table'
import { getComments, getRecipes, getUsers } from '@/lib/api'
import { Comment, Recipe, User } from '@/lib/types'
import { useQuery } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'
import { useState } from 'react'

export const Route = createFileRoute('/dev/_layout/')({
    component: Index,
})

type DevDataPanelProps = {
    title: string,
    children?: React.ReactNode
}

function DevDataPanel({ title, children }: DevDataPanelProps) {
    return (
        <Card>
            <CardHeader>
                <p className="text-xl">{title}</p>
            </CardHeader>
            <CardContent>
                {children}
            </CardContent>
        </Card>
    )
}

function UserDataTable({ users }: { users?: User[] }) {
    return <Table>
        <TableCaption>Data about users, not including passwords of course</TableCaption>
        <TableHeader>
            <TableRow>
                <TableHead className="w-12">ID</TableHead>
                <TableHead>Username</TableHead>
                <TableHead>Email</TableHead>
                <TableHead>Created At</TableHead>
                <TableHead>Updated At</TableHead>
            </TableRow>
        </TableHeader>
        <TableBody>
            {users && users.map(user => (
                <TableRow key={user.id}>
                    <TableCell>{user.id}</TableCell>
                    <TableCell>{user.username}</TableCell>
                    <TableCell>{user.email}</TableCell>
                    <TableCell>{user.created_at}</TableCell>
                    <TableCell>{user.updated_at}</TableCell>
                </TableRow>
            ))}
        </TableBody>
    </Table>
}

function RecipeView({ recipe }: { recipe: Recipe }) {
    return (
        <Card>
            <CardHeader>
                <p className="text-xl">{recipe.name}</p>
            </CardHeader>
            <CardContent>
                <p>Author: {recipe.author.username} - <a href={`mailto:${recipe.author.email}`}>{recipe.author.email}</a></p>
                <p>Ingredients: {recipe.ingredients.join(', ')}</p>
                <p>Instructions: {recipe.instructions.join(', ')}</p>
                <p>Category: {recipe.category.name}</p>
                <p>Subcategory: {recipe.subcategory.name}</p>
                <p>Attributes: </p>
                <ul className="pl-4 list-disc list-inside">
                    {recipe.attributes.map(attribute => (
                        <li key={attribute.name}>{attribute.name}: {attribute.value}</li>
                    ))}
                </ul>
                <p>Tags: </p>
                <ul className="pl-4 list-disc list-inside">
                    {recipe.tags.map(tag => (
                        <li key={tag.id}>{tag.name}</li>
                    ))}
                </ul>
            </CardContent>
        </Card>
    )
}

function RecipeDataTable({ recipes }: { recipes?: Recipe[] }) {
    const [selectedRecipe, setSelectedRecipe] = useState<Recipe | null>(null)

    const handleRecipeClick = (recipe: Recipe) => {
        setSelectedRecipe(recipe)
    }

    return (
        <>
            <Table>
                <TableCaption>Data about recipes, not including ingredients or instructions</TableCaption>
                <TableHeader>
                    <TableRow>
                        <TableHead className="w-12">ID</TableHead>
                        <TableHead>Name</TableHead>
                        <TableHead>Author</TableHead>
                        <TableHead>Created At</TableHead>
                        <TableHead>Updated At</TableHead>
                        <TableHead className="w-12">View</TableHead>
                    </TableRow>
                </TableHeader>
                <TableBody>
                    {recipes && recipes.map(recipe => (
                        <TableRow key={recipe.id}>
                            <TableCell>{recipe.id}</TableCell>
                            <TableCell>{recipe.name}</TableCell>
                            <TableCell>{recipe.author_id}</TableCell>
                            <TableCell>{recipe.created_at}</TableCell>
                            <TableCell>{recipe.updated_at}</TableCell>
                            <TableCell><Button variant={'outline'} onClick={() => handleRecipeClick(recipe)}>View</Button></TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
            {selectedRecipe && <RecipeView recipe={selectedRecipe} />}
        </>
    )
}

function CommentComponent({ comment, indentation = 0 }: { comment: Comment, indentation?: number }) {
    return (
        <div>
            <div className="flex flex-row gap-2">
                {indentation > 0 && <pre>{'  '.repeat(indentation)}</pre>}
                <p>{comment.content}</p>
                <p>{comment.author.username}</p>
                <p>{comment.recipe_id}</p>
            </div>
            {comment.children.map(child => <CommentComponent key={child.id} comment={child} indentation={indentation + 1} />)}
        </div>
    )
}

function Index() {
    const users = useQuery({
        queryKey: ['users'], queryFn: getUsers
    })

    const recipes = useQuery({
        queryKey: ['recipes'], queryFn: getRecipes
    })

    const comments = useQuery({
        queryKey: ['comments'], queryFn: async () => {
            const comments = await getComments(2)
            console.log(comments)
            return comments
        }
    })

    return (
        <div className='flex flex-col gap-4'>
            <DevDataPanel title="Users">
                <UserDataTable users={users.data} />
            </DevDataPanel>
            <DevDataPanel title="Recipes">
                <RecipeDataTable recipes={recipes.data} />
            </DevDataPanel>
            <DevDataPanel title="Comments">
                {comments.data?.map(comment => (
                    <CommentComponent key={comment.id} comment={comment} />
                ))}
            </DevDataPanel>
        </div>
    )
}