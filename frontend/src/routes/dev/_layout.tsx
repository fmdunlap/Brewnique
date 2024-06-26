import { Outlet, createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/dev/_layout')({
  component: LayoutComponent,
})

function NotDevMode() {
  return <div>Not in development mode</div>
}

function LayoutComponent() {
  const notInDev = process.env.NODE_ENV !== 'development'

  if (notInDev) {
    return <NotDevMode />
  }

  return (
    <div className="px-4 py-2">
      <Outlet />
    </div>
  )
}
