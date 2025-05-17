import { apiClient } from "./apiClient";

export const authApi = {
  login: (credentials) => apiClient.post("/login", credentials),
  logout: () => apiClient.post("/logout"),
  test: () => apiClient.get("/test"),
};

export const logApi = {
  getLogs: ({ page = 1, pageSize = 50 } = {}) =>
    apiClient.get(`/getlogs?page=${page}&pageSize=${pageSize}`),
};

// Add more API endpoints here as needed, organized by feature
export const userApi = {
  // Example for future endpoints
  // getProfile: () => apiClient.get('/user/profile'),
};
