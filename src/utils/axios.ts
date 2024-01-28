import axios, {AxiosInstance} from 'axios';

import {browserUtcOffset} from './formatter';

const {VITE_SERVER_URL, VITE_USE_BROWSER_TZ} = import.meta.env;

export const ServerClient = (): AxiosInstance => {
  return axios.create({
    baseURL: VITE_SERVER_URL,
    timeout: 5000,
    headers: {
      'X-Time-Zone':
        VITE_USE_BROWSER_TZ === 'true' ? browserUtcOffset() : 'UTC+11:00',
    },
  });
};
