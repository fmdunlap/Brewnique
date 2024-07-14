import { NewRecipe, Recipe, RecipeCategory } from "@/lib/types"
import { Button } from "@/components/ui/button.tsx"
import { Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table.tsx"
import { HoverCard, HoverCardContent, HoverCardTrigger } from "@/components/ui/hover-card"
import { DatabaseIcon } from "lucide-react"
import { ColumnDef } from "@tanstack/react-table"
import { DataTable } from "../ui/data-table"
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "../ui/accordion"
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
import { createRecipe } from "@/lib/api/recipe"
import { useForm } from "@tanstack/react-form"
import { User } from "@/lib/types";
import { Label } from "@radix-ui/react-label"
import { Input } from "../ui/input"
import { Card, CardContent } from "../ui/card"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "../ui/select"
import { getUsers } from "@/lib/api/user"
import { getCategories, getSubcategories } from "@/lib/api/category"

export function RecipeDataAccordion({ recipes, onShowRecipeComments, categories }: { recipes?: Recipe[], onShowRecipeComments: (recipeId: number) => void, categories?: RecipeCategory[] }) {
    return (
        <Accordion type="multiple" className="w-11/12 mx-auto">
            <AccordionItem value="recipe-data-table">
                <AccordionTrigger>Recipes</AccordionTrigger>
                <AccordionContent>
                    <RecipeDataTable recipes={recipes} onShowRecipeComments={onShowRecipeComments} />
                </AccordionContent>
            </AccordionItem>
            <AccordionItem value="recipe-category-table">
                <AccordionTrigger>Recipe Categories</AccordionTrigger>
                <AccordionContent className="w-4/5 mx-auto">
                    {categories != undefined &&
                        <DataTable columns={CategoryColumns} data={categories} />
                    }
                </AccordionContent>
            </AccordionItem>
            <AccordionItem value="new-recipe">
                <AccordionTrigger>New Recipe</AccordionTrigger>
                <AccordionContent>
                    <Card>
                        <CardContent className="p-6">
                            <NewRecipeForm />
                        </CardContent>
                    </Card>
                </AccordionContent>
            </AccordionItem>
        </Accordion>
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
                <p>Author: {recipe.author.username} - <a href={`mailto:${recipe.author.email}`}>{recipe.author.email}</a> - {recipe.author_id ?? "NO_ID"}</p>
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

const CategoryColumns: ColumnDef<RecipeCategory>[] = [
    {
        accessorKey: 'id',
        header: () => <div>ID</div>,
    },
    {
        accessorKey: 'name',
        header: () => <div className="w-full">Name</div>
    },
    {
        accessorKey: 'parent_id',
        header: () => <div className="w-1/5">Parent ID</div>,
        cell: ({ row }) => {
            const id = row.original.parent_id ? row.original.parent_id : "PRIMARY";

            return <div>{id}</div>
        }
    }
]

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
                        <TableCell>{recipe.created_at != null ? recipe.created_at.toString() : ""}</TableCell>
                        <TableCell>{recipe.updated_at != null ? recipe.updated_at.toString() : ""}</TableCell>
                        <TableCell><Button variant={'outline'} size={'sm'} onClick={() => onShowRecipeComments(recipe.id)}>Comments</Button></TableCell>
                        <TableCell><Button variant={'destructive'} size={'sm'} onClick={() => onShowRecipeComments(recipe.id)}>Delete</Button></TableCell>
                    </TableRow>
                ))}
            </TableBody>
        </Table>
    )
}

function NewRecipeForm() {
    const queryClient = useQueryClient()

    const submitRecipe = useMutation({
        mutationFn: async (newRecipe: NewRecipe) => {
            // Fake resolve chain to mock server response
            await new Promise(resolve => setTimeout(resolve, 1000))
            console.log(newRecipe)
        },
        onSuccess: () => {
            queryClient.invalidateQueries({
                queryKey: ['recipes']
            })
        }
    })

    const recipeForm = useForm({
        defaultValues: {
            name: '',
            author_id: 1,
            ingredients: [""],
            instructions: [""],
            category_id: 1,
            subcategory_id: '7', // Points at "American Ale"
            attributes: [0],
            tags: [0]
        },
        onSubmit: async (values) => {
            await submitRecipe.mutateAsync({
                name: values.value.name,
                author_id: values.value.author_id,
                ingredients: values.value.ingredients,
                instructions: values.value.instructions,
                category_id: values.value.category_id,
                subcategory_id: parseInt(values.value.subcategory_id),
                attributes: values.value.attributes,
                tags: values.value.tags
            })
        }
    })

    const userOptions = useQuery({
        queryKey: ['users'], queryFn: getUsers
    })
    const categoryOptions = useQuery({
        queryKey: ['recipeCategories'], queryFn: () => {
            return getCategories()
        }
    })
    const subcategoryOptions = useQuery({
        queryKey: ['subcategories', recipeForm.state.values.category_id],
        queryFn: async () => {
            const subcategories = await getSubcategories(recipeForm.state.values.category_id)
            return subcategories
        }
    })

    return (
        <form onSubmit={(e) => {
            e.preventDefault()
            e.stopPropagation()
            recipeForm.handleSubmit()
        }}>
            <div className="flex flex-col gap-2">
                <recipeForm.Field name="name" children={(field) => {
                    return (
                        <>
                            <Label htmlFor={field.name}>Name</Label>
                            <Input
                                id={field.name}
                                name={field.name}
                                value={field.state.value}
                                onBlur={field.handleBlur}
                                onChange={(e) => { field.handleChange(e.target.value) }}
                            />
                        </>
                    )
                }} />
                <recipeForm.Field name="author_id" children={(field) => {
                    return (
                        <>
                            <Label htmlFor={field.name}>Author</Label>
                            <Select onValueChange={(v) => field.handleChange(parseInt(v))} defaultValue={field.state.value.toString()}>
                                <SelectTrigger>
                                    <SelectValue placeholder="Select a verified email to display" />
                                </SelectTrigger>
                                <SelectContent>
                                    {userOptions.data?.map(user => (
                                        <SelectItem key={user.id} value={user.id.toString()}>{user.username} - {user.email}</SelectItem>
                                    ))}
                                </SelectContent>
                            </Select>
                        </>
                    )
                }} />
                <recipeForm.Field name="ingredients" mode="array">
                    {(field) => {
                        return (
                        <div className="mx-auto flex flex-col gap-2 w-5/6">
                            {field.state.value.map((_, i) => {
                            return (
                                <recipeForm.Field key={i} name={`ingredients[${i}]`}>
                                {(subField) => {
                                    return (
                                    <div className="flex flex-row gap-2">
                                        <Input
                                            value={subField.state.value}
                                            onChange={(e) =>
                                            subField.handleChange(e.target.value)
                                            }
                                        />
                                    </div>
                                    )
                                }}
                                </recipeForm.Field>
                            )
                            })}
                            <Button
                            onClick={() => field.pushValue("")}
                            type="button"
                            >
                            Add Ingredient
                            </Button>
                        </div>
                        )
                    }}
                    </recipeForm.Field>
                <recipeForm.Field name="category_id" children={(field) => {
                    return (
                        <>
                            <Label htmlFor={field.name}>Category</Label>
                            <Select
                                onValueChange={(v) => {
                                    field.handleChange(parseInt(v))
                                    queryClient.invalidateQueries({
                                        queryKey: ['subcategories']
                                    })
                            }}
                            defaultValue={field.state.value.toString()}>
                                <SelectTrigger>
                                    <SelectValue placeholder="Select a category to display" />
                                </SelectTrigger>
                                <SelectContent>
                                    {categoryOptions.data?.map(category => (
                                        <SelectItem key={category.id} value={category.id.toString()}>{category.name}</SelectItem>
                                    ))}
                                </SelectContent>
                            </Select>
                        </>
                    )
                }} />
                <recipeForm.Field name="subcategory_id" children={(field) => {
                    return (
                        <>
                            <Label htmlFor={field.name}>Subcategory</Label>
                            <Select onValueChange={(v) => {
                                field.handleChange(v);
                            }} defaultValue={field.state.value}>
                                <SelectTrigger>
                                    <SelectValue placeholder="Select a subcategory to display" />
                                </SelectTrigger>
                                <SelectContent>
                                    {subcategoryOptions.data?.map(category => (
                                        <SelectItem key={category.id} value={category.id.toString()}>{category.name}</SelectItem>
                                    ))}
                                </SelectContent>
                            </Select>
                        </>
                    )
                }} />
                <recipeForm.Subscribe
                    selector={(state) => [state.canSubmit, state.isSubmitting]}
                    children={([canSubmit, isSubmitting]) => (
                        <Button type="submit" disabled={!canSubmit}>
                            {isSubmitting ? '...' : 'Submit'}
                        </Button>
                    )}
                />
            </div>
        </form>
    )
}