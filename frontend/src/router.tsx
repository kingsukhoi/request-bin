import {createRouter, createRootRoute, createRoute, redirect} from '@tanstack/react-router'
import App from './App'
import {ViewRequests} from './pages/ViewRequests.tsx'
import {Login} from './pages/Login'
import {Home} from './pages/Home'
import {checkAuth} from './api'

// Define the root route
const rootRoute = createRootRoute({
    component: App
})

// Define the home route - redirects to /viewRequests
const homeRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: '/',
    component: Home,
})

// Define the view requests route with search params and auth check
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
    beforeLoad: async () => {
        const isAuthenticated = await checkAuth()
        if (!isAuthenticated) {
            const currentPath = window.location.pathname + window.location.search
            throw redirect({
                to: '/login',
                search: {
                    redirect: currentPath,
                },
            })
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
