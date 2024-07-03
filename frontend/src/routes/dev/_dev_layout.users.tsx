import { DevDataPanel } from '@/components/dev/devDataPanel'
import { DevNewUserForm, DevUserDataTable } from '@/components/dev/user'
import { Card } from '@/components/ui/card'
import { deleteUserById, getUsers } from '@/lib/api/user'
import { Separator } from '@radix-ui/react-separator'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/dev/_dev_layout/users')({
  component: () => <UserCard />
})

export function UserCard() {
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
      <DevUserDataTable users={users.data} deleteUser={deleteUser.mutate} />
      <Separator />
      <p className="text-xl">New User</p>
      <Card className="w-2/5 mx-auto p-4">
        <DevNewUserForm />
      </Card>
    </DevDataPanel>
  )
}