import {DevDataPanel} from "./devDataPanel"
import {Comment} from "@/lib/types"
import {useQuery} from "@tanstack/react-query"
import {getComments} from "@/lib/api/comment"

export function CommentCard({ commentRecipeId }: { commentRecipeId: number | null }) {
    const comments = useQuery({
        queryKey: ['comments', commentRecipeId], queryFn: async () => {
            if (commentRecipeId === null) {
                return []
            }
            return await getComments(commentRecipeId)
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