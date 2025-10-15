import {apiClient} from "./client";

export interface Header {
    requestId: string;
    name: string;
    value: string;
}

export async function GetHeaders(requestId: string) {
    return await apiClient.get(`rbv1/requests/headers`, {
        searchParams: {request_id: requestId}
    }).json<Header[]>()
}
