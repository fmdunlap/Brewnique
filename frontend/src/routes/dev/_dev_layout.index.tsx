import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader } from '@/components/ui/card'
import { Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table'
import { createUser, deleteUserById, getComments, getRecipes, getUsers } from '@/lib/api'
import { Comment, NewUser, Recipe, User } from '@/lib/types'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'
import { useState } from 'react'
import { Separator } from "@/components/ui/separator.tsx";
import { useForm } from "@tanstack/react-form";
import { Input } from "@/components/ui/input.tsx";
import { Label } from '@/components/ui/label'
import { HoverCard, HoverCardContent, HoverCardTrigger } from '@/components/ui/hover-card'
import { DatabaseIcon } from 'lucide-react'

export const Route = createFileRoute('/dev/_dev_layout/')({
    component: Index,
})

type DevDataPanelProps = {
    title: string,
    children?: React.ReactNode
}

function DevDataPanel({ title, children }: DevDataPanelProps) {
    return (
        <Card >
            <CardHeader>
                <p className="text-xl">{title}</p>
            </CardHeader>
            <CardContent className="flex flex-col gap-4">
                {children}
            </CardContent>
        </Card>
    )
}

function UserCard() {
    const queryClient = useQueryClient()

    const users = useQuery({
        queryKey: ['users'], queryFn: getUsers
    })

    const deleteUser = useMutation({
        mutationFn: (userId: number) => {
            return deleteUserById(userId)
        },
        onSuccess: () => {
            queryClient.invalidateQueries({
                queryKey: ['users']
            })
        }
    })

    return (
        <DevDataPanel title="Users">
            <UserDataTable users={users.data} deleteUser={deleteUser.mutate} />
            <Separator />
            <p className="text-xl">New User</p>
            <Card className="w-2/5 mx-auto p-4">
                <NewUserForm />
            </Card>
        </DevDataPanel>
    )
}

function UserDataTable({ users, deleteUser }: { users?: User[], deleteUser: (userId: number) => void }) {
    return <Table>
        <TableCaption>Data about users, not including passwords of course</TableCaption>
        <TableHeader>
            <TableRow>
                <TableHead className="w-12">ID</TableHead>
                <TableHead>Username</TableHead>
                <TableHead>Email</TableHead>
                <TableHead>Created At</TableHead>
                <TableHead>Updated At</TableHead>
                <TableHead className="w-12"></TableHead>
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
                    <TableCell><Button variant={'destructive'} size={'sm'} onClick={() => deleteUser(user.id)}>Delete</Button></TableCell>
                </TableRow>
            ))}
        </TableBody>
    </Table>
}

function NewUserForm() {
    const queryClient = useQueryClient()

    const submitUser = useMutation({
        mutationFn: (newUser: NewUser) => {
            return createUser(newUser)
        },
        onSuccess: () => {
            queryClient.invalidateQueries({
                queryKey: ['users']
            })
            userForm.reset()
        }
    })

    const userForm = useForm({
        defaultValues: {
            username: '',
            email: ''
        },
        onSubmit: async (values) => {
            submitUser.mutate(values.value)
        }
    })

    return (
        <form onSubmit={(e) => {
            e.preventDefault()
            e.stopPropagation()
            userForm.handleSubmit()
        }}>
            <div className="flex flex-col gap-2">
                <div>
                    <userForm.Field name="username" children={(field) => {
                        return (
                            <>
                                <Label htmlFor={field.name}>Username</Label>
                                <Input
                                    id={field.name}
                                    name={field.name}
                                    value={field.state.value}
                                    onBlur={field.handleBlur}
                                    onChange={(e) => field.handleChange(e.target.value)}
                                />
                            </>
                        )
                    }} />
                </div>
                <div>
                    <userForm.Field name="email" children={(field) => {
                        return (
                            <>
                                <Label htmlFor={field.name}>Email</Label>
                                <Input
                                    id={field.name}
                                    name={field.name}
                                    value={field.state.value}
                                    onBlur={field.handleBlur}
                                    onChange={(e) => field.handleChange(e.target.value)}
                                />
                            </>
                        )
                    }} />
                </div>
                <userForm.Subscribe
                    selector={(state) => [state.canSubmit, state.isSubmitting]}
                    children={([canSubmit, isSubmitting]) => (
                        <div className="flex flex-row gap-2 pt-2">
                            <Button type="submit" disabled={!canSubmit}>
                                {isSubmitting ? '...' : 'Submit'}
                            </Button>
                            <Button type="reset" onClick={() => userForm.reset()}>
                                Reset
                            </Button>
                        </div>
                    )}
                />
            </div>
        </form>
    )


}

function RecipeCard({ setCommentRecipeId }: { setCommentRecipeId: (recipeId: number) => void }) {
    const recipes = useQuery({
        queryKey: ['recipes'], queryFn: getRecipes
    })

    return (
        <DevDataPanel title="Recipes">
            <RecipeDataTable recipes={recipes.data} onShowRecipeComments={setCommentRecipeId} />
        </DevDataPanel>
    )
}

function RecipeDetailHoverCard({ recipe }: { recipe: Recipe }) {
    return (
        <HoverCard>
            <HoverCardTrigger>
                <div className="hover:bg-slate-300 hover:text-slate-700 rounded-md text-sm w-8 h-8 flex flex-col items-center justify-center">
                    <DatabaseIcon className="h-6 w-6 mx-auto my-auto" />
                </div>
            </HoverCardTrigger>
            <HoverCardContent className="w-96">
                <RecipeView recipe={recipe} />
            </HoverCardContent>
        </HoverCard>
    )
}

function RecipeView({ recipe }: { recipe: Recipe }) {
    return (
        <div className="flex flex-col gap-1">
            <p className="text-md">{recipe.name}</p>
            <div className="flex flex-col gap-1 text-sm">
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
            </div>
        </div>
    )
}

function RecipeDataTable({ recipes, onShowRecipeComments }: { recipes?: Recipe[], onShowRecipeComments: (recipeId: number) => void }) {
    return (
        <Table>
            <TableCaption>Data about recipes, not including ingredients or instructions</TableCaption>
            <TableHeader>
                <TableRow>
                    <TableHead className="w-10" />
                    <TableHead className="w-12">ID</TableHead>
                    <TableHead>Name</TableHead>
                    <TableHead>Author</TableHead>
                    <TableHead>Created At</TableHead>
                    <TableHead>Updated At</TableHead>
                    <TableHead className="w-12"></TableHead>
                    <TableHead className="w-12" />
                </TableRow>
            </TableHeader>
            <TableBody>
                {recipes && recipes.map(recipe => (
                    <TableRow key={recipe.id}>
                        <TableCell><RecipeDetailHoverCard recipe={recipe} /></TableCell>
                        <TableCell>{recipe.id}</TableCell>
                        <TableCell>{recipe.name}</TableCell>
                        <TableCell>{recipe.author.username}</TableCell>
                        <TableCell>{recipe.created_at}</TableCell>
                        <TableCell>{recipe.updated_at}</TableCell>
                        <TableCell><Button variant={'outline'} size={'sm'} onClick={() => onShowRecipeComments(recipe.id)}>Comments</Button></TableCell>
                        <TableCell><Button variant={'destructive'} size={'sm'} onClick={() => onShowRecipeComments(recipe.id)}>Delete</Button></TableCell>
                    </TableRow>
                ))}
            </TableBody>
        </Table>
    )
}

function CommentCard({ commentRecipeId }: { commentRecipeId: number | null }) {
    const comments = useQuery({
        queryKey: ['comments', commentRecipeId], queryFn: async () => {
            if (commentRecipeId === null) {
                return []
            }
            const comments = await getComments(commentRecipeId)
            return comments
        }
    })

    return (
        <DevDataPanel title="Comments">
            {comments.data?.map(comment => (
                <CommentView key={comment.id} comment={comment} />
            ))}
        </DevDataPanel>
    )
}

function CommentView({ comment, indentation = 0 }: { comment: Comment, indentation?: number }) {
    return (
        <div>
            <div className="flex flex-row gap-2">
                {indentation > 0 && <pre>{'  '.repeat(indentation)}</pre>}
                <p>{comment.content}</p>
                <p>{comment.author.username}</p>
                <p>{comment.recipe_id}</p>
            </div>
            {comment.children.map(child => <CommentView key={child.id} comment={child} indentation={indentation + 1} />)}
        </div>
    )
}

function Index() {
    const [commentRecipeId, setCommentRecipeId] = useState<number | null>(null)

    return (
        <div className='flex flex-col gap-4'>
            <UserCard />
            <RecipeCard setCommentRecipeId={setCommentRecipeId} />
            <CommentCard commentRecipeId={commentRecipeId} />
        </div>
    )
}