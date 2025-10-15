import {apiClient} from "./client";

export interface QueryParam {
    requestId: string;
    name: string;
    value: string;
}

export async function GetQueryParams(requestId: string) {
    return await apiClient.get(`rbv1/requests/queryParams`, {
        searchParams: {request_id: requestId}
    }).json<QueryParam[]>()
}
