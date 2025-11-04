import { apiClient } from "./client";

export interface LoginRequest {
    username: string;
    password: string;
}

export async function login(credentials: LoginRequest): Promise<void> {
    // Returns 200 on success and sets an HTTP-only cookie
    await apiClient.post('rbv1/login', {
        json: credentials,
    });
}

export async function checkAuth(): Promise<boolean> {
    try {
        // Returns 200 if authenticated, otherwise throws
        await apiClient.get('rbv1/checkAuth');
        return true;
    } catch {
        return false;
    }
}
