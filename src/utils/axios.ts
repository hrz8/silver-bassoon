import axios from 'axios';

// import {formatUtcOffset} from './formatter';

// const browserUtcOffset = (): string =>
//   formatUtcOffset(new Date().getTimezoneOffset());

export const Server = axios.create({
  baseURL: import.meta.env.VITE_SERVER_URL,
  timeout: 5000,
  // headers: {'X-Time-Zone': browserUtcOffset()},
  headers: {'X-Time-Zone': 'UTC+11:00'},
});
