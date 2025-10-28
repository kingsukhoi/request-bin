import {useCallback, useEffect, useState} from 'react'
import { useSearch, useNavigate } from '@tanstack/react-router'
import '../App.css'
import {Banner} from '../components/Banner'
import {RequestsTable} from '../components/RequestsTable'
import {RequestDetails} from '../components/RequestDetails'
import {PaginationBar} from '../components/PaginationBar'
import {GetRequests, type Request} from "../api"

export function Index() {
    const navigate = useNavigate()
    const { request_id } = useSearch({ from: '/' })

    const [requests, setRequests] = useState<Request[]>([])

    const [selectedRequest, setSelectedRequest] = useState<Request | null>(null)
    const [lastRefreshed, setLastRefreshed] = useState<Date | null>(null)
    const [isRefreshing, setIsRefreshing] = useState(false)

    const [pageSize, setPageSize] = useState(20)
    const [currentToken, setCurrentToken] = useState<string | undefined>(undefined)
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
        } finally {
          setIsRefreshing(false)
        }
    }, [isRefreshing, pageSize])

    useEffect(() => {
        refreshRequests()
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
            to: '/',
            search: { request_id: request.id }
        })
    }

    const handleCloseRequest = () => {
        navigate({
            to: '/',
            search: { request_id: undefined }
        })
    }

    return (
        <div className="app-container">
            <Banner title="Request Bin" subtitle="Monitor and inspect HTTP requests" />

            <main className="main-content">
                <div className="requests-section">
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
