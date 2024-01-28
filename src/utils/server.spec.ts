/* eslint-disable @typescript-eslint/unbound-method */
import axios, {AxiosInstance} from 'axios';
import {
  afterAll,
  beforeAll,
  describe,
  expect,
  MockedFunction,
  test,
  vi,
} from 'vitest';

vi.mock('axios', () => {
  return {
    default: {
      post: vi.fn(),
      get: vi.fn(),
      delete: vi.fn(),
      put: vi.fn(),
      create: vi.fn().mockReturnThis(),
      interceptors: {
        request: {
          use: vi.fn(),
          eject: vi.fn(),
        },
        response: {
          use: vi.fn(),
          eject: vi.fn(),
        },
      },
    },
  };
});

describe('Server', () => {
  beforeAll(() => {
    import.meta.env.VITE_SERVER_URL = 'https://example.com';
    import.meta.env.VITE_USE_BROWSER_TZ = 'true';
  });

  afterAll(() => {
    delete import.meta.env.VITE_SERVER_URL;
    delete import.meta.env.VITE_USE_BROWSER_TZ;
  });

  test('Server instance has correct configuration', async () => {
    vi.mock('./formatter', () => ({
      browserUtcOffset: vi.fn(() => 'UTC+00:00'),
    }));

    (axios.create as MockedFunction<typeof axios.create>).mockImplementation(
      (config) => {
        expect(config?.baseURL).toBe('https://example.com');
        expect(config?.timeout).toBe(5000);
        return {} as AxiosInstance;
      },
    );

    const axiosSpy = vi.spyOn(axios, 'create');
    const {ServerClient} = await import('./axios');
    ServerClient();

    expect(axiosSpy).toHaveBeenCalled();
  });
});
