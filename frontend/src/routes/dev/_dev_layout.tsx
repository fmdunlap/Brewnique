import { Link, Outlet, createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/dev/_dev_layout')({
  component: LayoutComponent,
})

function NotDevMode() {
  return <div>Not in development mode</div>
}

function DevNavBar() {
  return (
    <>
      <div className="p-2 flex gap-2 bg-white">
        <Link to="/dev/recipes" className="[&.active]:font-bold">
          Recipes
        </Link>{' '}
        <Link to="/dev/users" className="[&.active]:font-bold">
          Users
        </Link>
      </div>
      <hr />
    </>
  )
}

function LayoutComponent() {
  const notInDev = process.env.NODE_ENV !== 'development'

  if (notInDev) {
    return <NotDevMode />
  }

  return (
    <>
      <DevNavBar />
      <div className="px-4 py-2">
        <Outlet />
      </div>
    </>
  )
}
