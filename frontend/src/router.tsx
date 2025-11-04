import {createRouter, createRootRoute, createRoute} from '@tanstack/react-router'
import App from './App'
import {ViewRequests} from './pages/ViewRequests.tsx'
import {Login} from './pages/Login'
import {Home} from './pages/Home'

// Define the root route
const rootRoute = createRootRoute({
    component: App
})

// Define the home route
const homeRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: '/',
    component: Home,
})

// Define the view requests route with search params
const viewRequestsRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: '/viewRequests',
    component: ViewRequests,
    validateSearch: (search: Record<string, unknown>) => {
        return {
            request_id: (search.request_id as string) || undefined,
            nextToken: (search.nextToken as string) || undefined,
        }
    },
})

// Define the login route
const loginRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: '/login',
    component: Login,
    validateSearch: (search: Record<string, unknown>) => {
        return {
            redirect: (search.redirect as string) || undefined,
        }
    },
})

// Create the route tree
const routeTree = rootRoute.addChildren([homeRoute, viewRequestsRoute, loginRoute])

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
