import ky from 'ky';

const baseUrl = import.meta.env.VITE_API_BASE_URL || './';

export const apiClient = ky.create({
  prefixUrl: baseUrl,
  timeout: 5000,
});
