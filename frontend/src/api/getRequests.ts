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


export async function GetRequests() {


    const res = await apiClient.get("rbv1/requests").json<RequestResponse[]>()

    return res.map(req => ({
        ...req,
        timestamp: new Date(req.timestamp),
        content: req.content ? atob(req.content) : null,
    })) as Request[]

}