import axios from 'axios';

const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
});

export const checkFlight = (flightNumber, date) => {
  return apiClient.post('/api/check', { flightNumber, date });
};

export const generateVoucher = (payload) => {
  return apiClient.post('/api/generate', payload);
};