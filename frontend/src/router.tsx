import {createRouter, createRootRoute, createRoute} from '@tanstack/react-router'
import App from './App'
import {Index} from './pages/Index'

// Define the root route
const rootRoute = createRootRoute({
    component: App
})

// Define the index route (home page) with search params
const indexRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: '/',
    component: Index,
    validateSearch: (search: Record<string, unknown>) => {
        return {
            request_id: (search.request_id as string) || undefined,
        }
    },
})

// Create the route tree
const routeTree = rootRoute.addChildren([indexRoute])

// Create and export the router
export const router = createRouter({
    routeTree,
    defaultPreload: 'intent',
})

// Register the router for type safety
declare module '@tanstack/react-router' {
    interface Register {
        router: typeof router
    }
}
