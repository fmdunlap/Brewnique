import { NewRecipe, Recipe, RecipeCategory } from "@/lib/types"
import { Button } from "@/components/ui/button.tsx"
import { Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table.tsx"
import { HoverCard, HoverCardContent, HoverCardTrigger } from "@/components/ui/hover-card"
import { DatabaseIcon } from "lucide-react"
import { ColumnDef } from "@tanstack/react-table"
import { DataTable } from "../ui/data-table"
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "../ui/accordion"
import { useMutation, useQueryClient } from "@tanstack/react-query"
import { createRecipe } from "@/lib/api/recipe"
import { useForm } from "@tanstack/react-form"
import { User } from "@/lib/types";
import { Label } from "@radix-ui/react-label"
import { Input } from "../ui/input"
import { Card, CardContent } from "../ui/card"

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
        mutationFn: (newRecipe: NewRecipe) => {
            return createRecipe(newRecipe)
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
            author_id: 0,
            ingredients: [],
            instructions: [],
            category_id: 0,
            subcategory_id: 0,
            attributes: [],
            tags: []
        },
        onSubmit: async (values) => {
            submitRecipe.mutate(values.value)
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
                        </>
                    )
                }} />
            </div>
        </form>
    )
}