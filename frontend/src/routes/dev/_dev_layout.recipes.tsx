import { CommentCard } from '@/components/dev/commentCard'
import { DevDataPanel } from '@/components/dev/devDataPanel'
import { RecipeDataAccordion } from '@/components/dev/recipe'
import { getCategories, getSubcategories } from '@/lib/api/category'
import { getRecipes } from '@/lib/api/recipe'
import { RecipeCategory } from '@/lib/types'
import { useQuery } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'
import { useState } from 'react'

export const Route = createFileRoute('/dev/_dev_layout/recipes')({
  component: () => {
    const [commentRecipeId, setCommentRecipeId] = useState<number | null>(null)

    return (
      <div className='flex flex-col gap-4'>
        <RecipeCard setCommentRecipeId={setCommentRecipeId} />
        <CommentCard commentRecipeId={commentRecipeId} />
      </div>
    )
  }
})

function RecipeCard({ setCommentRecipeId }: { setCommentRecipeId: (recipeId: number) => void }) {
  const recipes = useQuery({
    queryKey: ['recipes'], queryFn: getRecipes
  })

  const {
    data: categories
  } = useQuery({
    queryKey: ['recipeCategories'],
    queryFn: async () => {
      const categories: RecipeCategory[] = []

      const parentCategories = await getCategories()
      for (const category of parentCategories) {
        const subcategories = await getSubcategories(category.id)
        categories.push(category)
        for (const sc of subcategories) {
          categories.push(sc)
        }
      }

      return categories
    }
  })

  return (
    <DevDataPanel title="Recipes">
      <RecipeDataAccordion
        categories={categories} recipes={recipes.data} onShowRecipeComments={setCommentRecipeId} />
    </DevDataPanel>
  )
}