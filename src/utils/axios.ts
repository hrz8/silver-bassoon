import axios from 'axios';

import {formatUtcOffset} from './formatter';

const browserUtcOffset = (): string =>
  formatUtcOffset(new Date().getTimezoneOffset());

const {VITE_SERVER_URL, VITE_USE_BROWSER_TZ} = import.meta.env;

console.info(VITE_SERVER_URL, VITE_USE_BROWSER_TZ);

export const Server = axios.create({
  baseURL: VITE_SERVER_URL,
  timeout: 5000,
  headers: {
    'X-Time-Zone':
      VITE_USE_BROWSER_TZ === 'true' ? browserUtcOffset() : 'UTC+11:00',
  },
});
