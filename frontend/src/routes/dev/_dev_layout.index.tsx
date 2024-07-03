import { Button } from '@/components/ui/button'
import { Card, CardContent } from '@/components/ui/card'
import { Link } from '@tanstack/react-router'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/dev/_dev_layout/')({
    component: Index,
})

function Index() {
    return (<div className="flex flex-row gap-x-4 w-full">
        <Link to='/dev/recipes'>
            <Card>
                <CardContent className="m-0 p-0">
                    <Button variant="link" className="m-auto p-8">Recipes</Button>
                </CardContent>
            </Card>
        </Link>
        <Link to='/dev/users'>
            <Card>
                <CardContent className="m-0 p-0">
                    <Button variant="link" className="m-auto p-8">Users</Button>
                </CardContent>
            </Card>
        </Link>
    </div>)
}