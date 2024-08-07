import { useMutation, useQueryClient } from "@tanstack/react-query"
import { Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table.tsx"
import { Button } from "@/components/ui/button.tsx"
import { Label } from "@/components/ui/label"
import { Input } from "@/components/ui/input.tsx"
import { useForm } from "@tanstack/react-form"
import { NewUser, User } from "@/lib/types"
import { createUser } from "@/lib/api/user"

export function DevUserDataTable({ users, deleteUser }: { users?: User[], deleteUser: (userId: number) => void }) {
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

export function DevNewUserForm() {
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