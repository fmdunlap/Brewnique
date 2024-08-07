import { createRootRoute, Link, Outlet } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/router-devtools'

export const Route = createRootRoute({
    component: () => (
        <>
            <div className="p-2 flex gap-2 bg-white">
                <Link to="/" className="[&.active]:font-bold">
                    Home
                </Link>{' '}
                <Link to="/about" className="[&.active]:font-bold">
                    About
                </Link>
                <Link to="/another" className="[&.active]:font-bold">
                    Another
                </Link>
                {process.env.NODE_ENV === 'development' && <Link to="/dev" className="[&.active]:font-bold">
                    Dev</Link>}
            </div>
            <hr />
            <div>
                <Outlet />
            </div>
            <TanStackRouterDevtools />
        </>
    ),
})