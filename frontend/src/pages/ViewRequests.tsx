import {useCallback, useEffect, useState} from 'react'
import {useNavigate, useSearch} from '@tanstack/react-router'
import {Banner} from '../components/Banner'
import {RequestsTable} from '../components/RequestsTable'
import {RequestDetails} from '../components/RequestDetails'
import {PaginationBar} from '../components/PaginationBar'
import {GetRequests, type Request} from "../api"

export function ViewRequests() {
    const navigate = useNavigate()
    const { request_id, nextToken } = useSearch({ from: '/viewRequests' })

    const [requests, setRequests] = useState<Request[]>([])

    const [selectedRequest, setSelectedRequest] = useState<Request | null>(null)
    const [lastRefreshed, setLastRefreshed] = useState<Date | null>(null)
    const [isRefreshing, setIsRefreshing] = useState(false)

    const [pageSize, setPageSize] = useState(20)
    const [currentToken, setCurrentToken] = useState<string | undefined>(nextToken)
    const [previousTokens, setPreviousTokens] = useState<(string | undefined)[]>([])

    const refreshRequests = useCallback(async (token?: string) => {
        if (isRefreshing) return

        setIsRefreshing(true)
        try {
            const requestsData = await GetRequests({
                limit: pageSize,
                nextToken: token
            })
            setRequests(requestsData)
            setCurrentToken(token)
            setLastRefreshed(new Date())

            // Update URL with current token
            navigate({
                to: '/viewRequests',
                search: { request_id, nextToken: token },
                replace: true
            })
        } finally {
          setIsRefreshing(false)
        }
    }, [isRefreshing, pageSize, navigate, request_id])

    useEffect(() => {
        refreshRequests(nextToken)
    }, [])//leave that empty or it spams the server

    useEffect(() => {
        // Reset to first page when page size changes
        setPreviousTokens([])
        refreshRequests()
    }, [pageSize])

    // Sync URL query param with selected request
    useEffect(() => {
        if (request_id) {
            const request = requests.find(r => r.id === request_id)
            // Update selectedRequest if we found it, or if it's a new request_id
            if (request) {
                setSelectedRequest(request)
            }
            // If request not found in current page but request_id matches current selection,
            // keep showing the current selectedRequest (allows details to stay open when paginating)
        } else {
            // Only clear if request_id is explicitly undefined (closing the panel)
            setSelectedRequest(null)
        }
    }, [request_id, requests])

    const handleNextPage = () => {
        // Use the ID of the last request as the next token
        if (requests.length > 0) {
            const lastRequestId = requests[requests.length - 1].id
            // Store current token so we can go back
            setPreviousTokens([...previousTokens, currentToken])
            refreshRequests(lastRequestId)
        }
    }

    const handleFirstPage = () => {
        setPreviousTokens([])
        refreshRequests(undefined)
    }

    const handlePreviousPage = () => {
        if (previousTokens.length > 0) {
            const newPreviousTokens = [...previousTokens]
            const previousToken = newPreviousTokens.pop()
            setPreviousTokens(newPreviousTokens)
            refreshRequests(previousToken)
        } else {
            // Go back to first page
            refreshRequests(undefined)
        }
    }

    const handlePageSizeChange = (size: number) => {
        setPageSize(size)
    }

    const handleSelectRequest = (request: Request) => {
        navigate({
            to: '/viewRequests',
            search: { request_id: request.id, nextToken: currentToken }
        })
    }

    const handleCloseRequest = () => {
        navigate({
            to: '/viewRequests',
            search: { request_id: undefined, nextToken: currentToken }
        })
    }

    return (
        <div className="min-h-screen bg-gh-bg-primary">
            <Banner title="Request Bin" subtitle="Monitor and inspect HTTP requests" />

            <main className="max-w-[1800px] mx-auto md:p-8 p-4 flex md:flex-row flex-col md:gap-8 gap-4 items-start">
                <div className="flex-1 min-w-0 flex flex-col gap-4 w-full">
                    <RequestsTable
                        requests={requests}
                        onSelectRequest={handleSelectRequest}
                        onRefresh={() => refreshRequests()}
                        lastRefreshed={lastRefreshed}
                        isRefreshing={isRefreshing}
                        selectedRequest={selectedRequest}
                    />

                    <PaginationBar
                        pageSize={pageSize}
                        onPageSizeChange={handlePageSizeChange}
                        onFirstPage={handleFirstPage}
                        onPreviousPage={handlePreviousPage}
                        onNextPage={handleNextPage}
                        hasPreviousPage={previousTokens.length > 0}
                        hasNextPage={requests.length >= pageSize}
                    />
                </div>

                {selectedRequest && (
                    <RequestDetails
                        request={selectedRequest}
                        onClose={handleCloseRequest}
                    />
                )}
            </main>
        </div>
    )
}
