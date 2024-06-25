import { createLazyFileRoute } from '@tanstack/react-router'

export const Route = createLazyFileRoute('/another')({
  component: () => <div>Hello /another!</div>
})