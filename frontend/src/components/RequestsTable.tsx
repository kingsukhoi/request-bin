import type {Request} from "../api"

interface RequestsTableProps {
    requests: Request[]
    onSelectRequest: (request: Request) => void
}

export function RequestsTable({requests, onSelectRequest}: RequestsTableProps) {
    const getMethodColor = (method: string) => {
        const colors: Record<string, string> = {
            GET: '#61affe',
            POST: '#49cc90',
            PUT: '#fca130',
            DELETE: '#f93e3e',
            PATCH: '#50e3c2'
        }
        return colors[method] || '#999'
    }

    requests.sort((a, b) => new Date(b.timestamp).getTime() - new Date(a.timestamp).getTime())


    return (
        <div className="requests-table-container">
            <h2>Incoming Requests</h2>
            <table className="requests-table">
                <thead>
                <tr>
                    <th>Method</th>
                    <th>Path</th>
                    <th>Timestamp</th>
                    <th>IP Address</th>
                    <th>Actions</th>
                </tr>
                </thead>
                <tbody>
                {requests.length === 0 ? (
                    <tr>
                        <td colSpan={5} className="empty-state">
                            No requests yet. Start sending requests to your bin!
                        </td>
                    </tr>
                ) : (
                    requests.map((request) => (
                        <tr key={request.id}>
                            <td>
                  <span
                      className="method-badge"
                      style={{backgroundColor: getMethodColor(request.method)}}
                  >
                    {request.method}
                  </span>
                            </td>
                            <td className="path-cell">{request.path}</td>
                            <td>{request.timestamp.toLocaleString()}</td>
                            <td>{request.sourceIp}</td>
                            <td>
                                <button
                                    className="view-button"
                                    onClick={() => onSelectRequest(request)}
                                >
                                    View Details
                                </button>
                            </td>
                        </tr>
                    ))
                )}
                </tbody>
            </table>
        </div>
    )
}
