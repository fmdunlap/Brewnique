import { createFileRoute } from '@tanstack/react-router'
import { useState } from 'react'
import { UserCard } from '@/components/dev/userCard'
import { RecipeCard } from '@/components/dev/recipeCard'
import { CommentCard } from '@/components/dev/commentCard'

export const Route = createFileRoute('/dev/_dev_layout/')({
    component: Index,
})

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