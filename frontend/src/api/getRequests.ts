import {apiClient} from "./client";

interface RequestResponse {
    id: string;
    method: string;
    content: string;
    sourceIp: string;
    responseCode: number;
    timestamp: string;
    path: string;
}

export interface Request {
    id: string;
    method: string;
    content: string;
    sourceIp: string;
    responseCode: number;
    timestamp: Date;
    path: string;
}

export interface GetRequestsParams {
    limit?: number;
    nextToken?: string;
}

export async function GetRequests(params?: GetRequestsParams): Promise<Request[]> {
    const searchParams = new URLSearchParams();

    if (params?.limit) {
        searchParams.append('limit', params.limit.toString());
    }

    if (params?.nextToken) {
        searchParams.append('next_token', params.nextToken);
    }

    const url = `rbv1/requests${searchParams.toString() ? `?${searchParams.toString()}` : ''}`;

    const res = await apiClient.get(url).json<RequestResponse[]>();

    return res.map(req => ({
        ...req,
        timestamp: new Date(req.timestamp),
        content: req.content ? atob(req.content) : null,
    })) as Request[];
}