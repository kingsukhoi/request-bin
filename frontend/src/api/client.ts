import ky, {isHTTPError} from 'ky';
import {router} from '../router';

const baseUrl = import.meta.env.VITE_API_BASE_URL || './';

export const apiClient = ky.create({
  prefix: baseUrl,
  timeout: 10000,
  credentials: 'include',
  hooks: {
    beforeError: [
      async ({error}) => {
        // Redirect to login on authentication errors
        if (isHTTPError(error) && error.response.status === 401) {
          const currentPath = window.location.pathname + window.location.search;
          void router.navigate({
            to: '/login',
            search: {
              redirect: currentPath
            }
          });
        }

        return error;
      }
    ]
  }
});
