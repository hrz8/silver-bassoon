export type MetaResponse = null | {
  count: number;
  total: number;
};

export type Response<T = any> = {
  message: string;
  result: T;
  meta: MetaResponse;
};
