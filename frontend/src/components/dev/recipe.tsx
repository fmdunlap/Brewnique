import {NewRecipe, Recipe, RecipeCategory} from "@/lib/types"
import {Button} from "@/components/ui/button.tsx"
import {Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow} from "@/components/ui/table.tsx"
import {HoverCard, HoverCardContent, HoverCardTrigger} from "@/components/ui/hover-card"
import {DatabaseIcon, PlusIcon, TrashIcon} from "lucide-react"
import {ColumnDef} from "@tanstack/react-table"
import {DataTable} from "../ui/data-table"
import {Accordion, AccordionContent, AccordionItem, AccordionTrigger} from "../ui/accordion"
import {useMutation, useQuery, useQueryClient} from "@tanstack/react-query"
import {useForm} from "@tanstack/react-form"
import {Label} from "@radix-ui/react-label"
import {Input} from "../ui/input"
import {Card, CardContent} from "../ui/card"
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "../ui/select"
import {getUsers} from "@/lib/api/user"
import {getCategories, getSubcategories} from "@/lib/api/category"
import {getTags} from "@/lib/api/tags"
import MultipleSelector from "../ui/multi-select"
import {getAttributes} from "@/lib/api/attributes"
import {createRecipe} from "@/lib/api/recipe"

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
                <p>Author: {recipe.author.username} - <a href={`mailto:${recipe.author.email}`}>{recipe.author.email}</a> - {recipe.author.id ?? "NO_ID"}</p>
                <p>Ingredients: {recipe.ingredients.join(', ')}</p>
                <p>Instructions: {recipe.instructions.join(', ')}</p>
                <p>Category: {recipe.category}</p>
                <p>Subcategory: {recipe.subcategory}</p>
                <p>Attributes: </p>
                <ul className="pl-4 list-disc list-inside">
                    {recipe.attributes.map(attribute => (
                        <li key={attribute.name}>{attribute.name}: {attribute.value}</li>
                    ))}
                </ul>
                <p>Tags: </p>
                <ul className="pl-4 list-disc list-inside">
                    {recipe.tags.map((tag, i) => (
                        <li key={"tag-" + i}>{tag}</li>
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
                {recipes && recipes.map(recipe => {
                    return (<TableRow key={recipe.id}>
                        <TableCell><RecipeDetailHoverCard recipe={recipe} /></TableCell>
                        <TableCell>{recipe.id}</TableCell>
                        <TableCell>{recipe.name}</TableCell>
                        <TableCell>{recipe.author.username}</TableCell>
                        <TableCell>{recipe.created_at != null ? recipe.created_at.toString() : ""}</TableCell>
                        <TableCell>{recipe.updated_at != null ? recipe.updated_at.toString() : ""}</TableCell>
                        <TableCell><Button variant={'outline'} size={'sm'} onClick={() => onShowRecipeComments(recipe.id)}>Comments</Button></TableCell>
                        <TableCell><Button variant={'destructive'} size={'sm'} onClick={() => onShowRecipeComments(recipe.id)}>Delete</Button></TableCell>
                    </TableRow>)
                })}
            </TableBody>
        </Table>
    )
}

function NewRecipeForm() {
    const queryClient = useQueryClient()

    const submitRecipe = useMutation({
        mutationFn: async (newRecipe: NewRecipe) => {
            // Fake resolve chain to mock server response
            await createRecipe(newRecipe)
        },
        onSuccess: async  () => {
            await queryClient.invalidateQueries({
                queryKey: ['recipes']
            })
        }
    })

    type RecipeFormValues = {
        name: string,
        author_id: number,
        ingredients: string[],
        instructions: string[],
        category_id: number,
        subcategory_id: string,
        attribute_ids: number[],
        tag_ids: number[]
    }

    const recipeForm = useForm({
        defaultValues: {
            name: '',
            author_id: 1,
            ingredients: [],
            instructions: [],
            category_id: 1,
            subcategory_id: '7', // Points at "American Ale"
            attribute_ids: [],
            tag_ids: []
        } as RecipeFormValues,
        onSubmit: async (values) => {
            await submitRecipe.mutateAsync({
                name: values.value.name,
                author_id: values.value.author_id,
                ingredients: values.value.ingredients,
                instructions: values.value.instructions,
                category_id: values.value.category_id,
                subcategory_id: parseInt(values.value.subcategory_id),
                attribute_ids: values.value.attribute_ids,
                tag_ids: values.value.tag_ids
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
            return await getSubcategories(recipeForm.state.values.category_id)
        }
    })
    const tagOptions = useQuery({
        queryKey: ['tags'], queryFn: getTags
    })
    const attributes = useQuery({
        queryKey: ['attributes'], queryFn: getAttributes
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
                <Label className="my-auto" htmlFor="ingredients">Ingredients</Label>
                <recipeForm.Field name="ingredients" mode="array">
                    {(field) => {
                        return (
                            <div className="py-2 flex flex-col gap-2">
                                {field.state.value.map((_, i) => {
                                    return (
                                        <recipeForm.Field key={i} name={`ingredients[${i}]`}>
                                            {(subField) => {
                                                return (
                                                    <div className="flex flex-row gap-2">
                                                        <Label className="my-auto" htmlFor={subField.name}>{i}.</Label>
                                                        <Input
                                                            value={subField.state.value}
                                                            onChange={(e) =>
                                                                subField.handleChange(e.target.value)
                                                            }
                                                        />
                                                        <Button onClick={() => field.removeValue(i)} variant={"destructive"} type="button">
                                                            <TrashIcon className="h-6 w-6" />
                                                        </Button>
                                                    </div>
                                                )
                                            }}
                                        </recipeForm.Field>
                                    )
                                })}
                                <Button
                                    className="mr-auto"
                                    onClick={() => field.pushValue("")}
                                    type="button"
                                >
                                    <div className="flex flex-row gap-2">
                                        <Label className="my-auto" htmlFor="ingredients">Add Ingredient</Label>
                                        <PlusIcon className="h-6 w-6" />
                                    </div>
                                </Button>
                            </div>
                        )
                    }}
                </recipeForm.Field>
                <Label className="my-auto" htmlFor="instructions">Instructions</Label>
                <recipeForm.Field name="instructions" mode="array">
                    {(field) => {
                        return (
                            <div className="py-2 flex flex-col gap-2">
                                {field.state.value.map((_, i) => {
                                    return (
                                        <recipeForm.Field key={i} name={`instructions[${i}]`}>
                                            {(subField) => {
                                                return (
                                                    <div className="flex flex-row gap-2">
                                                        <Label className="my-auto" htmlFor={subField.name}>{i}.</Label>
                                                        <Input
                                                            value={subField.state.value}
                                                            onChange={(e) =>
                                                                subField.handleChange(e.target.value)
                                                            }
                                                        />
                                                        <Button onClick={() => field.removeValue(i)} variant={"destructive"} type="button">
                                                            <TrashIcon className="h-6 w-6" />
                                                        </Button>
                                                    </div>
                                                )
                                            }}
                                        </recipeForm.Field>
                                    )
                                })}
                                <Button
                                    className="mr-auto"
                                    onClick={() => field.pushValue("")}
                                    type="button"
                                >
                                    <div className="flex flex-row gap-2">
                                        <Label className="my-auto" htmlFor="instructions">Add Instruction</Label>
                                        <PlusIcon className="h-6 w-6" />
                                    </div>
                                </Button>
                            </div>
                        )
                    }}
                </recipeForm.Field>
                <recipeForm.Field name={"tag_ids"} mode="array" children={(field) => (
                    <>
                        <Label className="my-auto" htmlFor="tag_ids">Tags</Label>
                        {tagOptions.data && <MultipleSelector
                            defaultOptions={tagOptions.data?.map(tag => ({
                                value: tag.id.toString(),
                                label: tag.name
                            }))}
                            className="w-full"
                            placeholder="Select Tags"
                            emptyIndicator={
                                <p className="text-center text-lg leading-10 text-gray-600 dark:text-gray-400">
                                    no results found.
                                </p>
                            }
                            onChange={(options) => {
                                field.setValue(options.map(option => parseInt(option.value)))
                            }}
                        />}
                    </>
                    )}
                />
                <Label className="my-auto" htmlFor="attribute_ids">Attributes</Label>
                <recipeForm.Field name="attribute_ids" mode="array" children={(field) => {
                    return (
                        <div className="mx-auto grid grid-cols-3 gap-4 w-full">
                            {attributes.data && attributes.data.map(attribute => (
                                <div key={attribute.id}>
                                    <Label className="my-auto" htmlFor={attribute.name}>{attribute.name}</Label>
                                    <div className="flex flex-row gap-2">
                                        <Select
                                            onValueChange={(v) => {
                                                const thisAttrValues = attribute.values.map(value => value.id)
                                                for (const existingValue of field.state.value) {
                                                    if (thisAttrValues.includes(existingValue)) {
                                                        field.removeValue(field.state.value.indexOf(existingValue))
                                                    }
                                                }
                                                if (v !== '0') {
                                                    field.pushValue(parseInt(v))
                                                }
                                            }}
                                            defaultValue={field.state.value.toString()}
                                        >
                                            <SelectTrigger>
                                                <SelectValue placeholder="" />
                                            </SelectTrigger>
                                            <SelectContent>
                                                {attribute.values.map(value => (
                                                    <SelectItem key={value.id} value={value.id.toString()}>{value.value}</SelectItem>
                                                ))}
                                                <SelectItem className="h-7" value="0">{" "}</SelectItem>
                                            </SelectContent>
                                        </Select>
                                    </div>
                                </div>
                            ))}
                        </div>
                    )
                }} />
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