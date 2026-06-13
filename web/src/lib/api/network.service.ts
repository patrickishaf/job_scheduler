import axios from "axios";

const API_BASE_URL = "http://localhost:8000";

const client = axios.create({
  baseURL: API_BASE_URL,
});

export const networkService = {
  get: async (url: string, query?: object) => {
    const res = await client.get(url, query && query);
    return res.data;
  },

  patch: async (url: string, body: object = {}) => {
    const res = await client.patch(url, body);
    return res.data;
  },

  post: async (url: string, body: object) => {
    const res = await client.post(url, body);
    return res.data;
  },
};
