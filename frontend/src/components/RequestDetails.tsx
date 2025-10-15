import type {Request} from '../api'
import {GetHeaders} from "../api/getHeaders.ts";
import {GetQueryParams} from "../api/getQueryParams.ts";
import {useEffect, useState} from "react";
import type {Header} from "../api/getHeaders.ts";
import type {QueryParam} from "../api/getQueryParams.ts";
import ReactJson from 'react-json-view';

interface RequestDetailsProps {
    request: Request
    onClose: () => void
}

export function RequestDetails({request, onClose}: RequestDetailsProps) {

    const [headers, setHeaders] = useState<Header[]>([])
    const [queryParams, setQueryParams] = useState<QueryParam[]>([])

    useEffect(() => {
        GetHeaders(request.id).then(setHeaders)
        GetQueryParams(request.id).then(setQueryParams)
    }, [request.id])


    return (
        <div className="request-details">
            <div className="details-header">
                <h2>Request Details</h2>
                <button className="close-button" onClick={onClose}>
                    ×
                </button>
            </div>
            <div className="details-content">
                <div className="detail-section">
                    <h3>Basic Info</h3>
                    <p><strong>Method:</strong> {request.method}</p>
                    <p><strong>Path:</strong> {request.path}</p>
                    <p><strong>Timestamp:</strong> {request.timestamp.toLocaleString()}</p>
                    <p><strong>IP Address:</strong> {request.sourceIp}</p>
                </div>

                <div className="detail-section">
                    <h3>Headers</h3>
                    <table className="headers-table">
                        <thead>
                            <tr>
                                <th>Name</th>
                                <th>Value</th>
                            </tr>
                        </thead>
                        <tbody>
                            {headers.map((header, index) => (
                                <tr key={index}>
                                    <td>{header.name}</td>
                                    <td>{header.value}</td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>

                <div className="detail-section">
                    <h3>Query Parameters</h3>
                    <table className="headers-table">
                        <thead>
                            <tr>
                                <th>Name</th>
                                <th>Value</th>
                            </tr>
                        </thead>
                        <tbody>
                            {queryParams.length > 0 ? (
                                queryParams.map((param, index) => (
                                    <tr key={index}>
                                        <td>{param.name}</td>
                                        <td>{param.value}</td>
                                    </tr>
                                ))
                            ) : (
                                <tr>
                                    <td colSpan={2} style={{textAlign: 'center', fontStyle: 'italic', color: '#8b949e'}}>
                                        No query parameters
                                    </td>
                                </tr>
                            )}
                        </tbody>
                    </table>
                </div>

                {request.content && (
                    <div className="detail-section">
                        <h3>Body</h3>
                        <ReactJson
                            src={JSON.parse(request.content)}
                            theme="monokai"
                            collapsed={false}
                            displayDataTypes={false}
                            displayObjectSize={false}
                            enableClipboard={true}
                            style={{
                                background: '#0d1117',
                                padding: '1rem',
                                borderRadius: '6px',
                                border: '1px solid #30363d'
                            }}
                        />
                    </div>
                )}
            </div>
        </div>
    )
}
