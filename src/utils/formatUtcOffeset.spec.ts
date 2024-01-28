import {describe, expect, it} from 'vitest';

import {formatUtcOffset} from './formatter';

describe('formatUtcOffset', () => {
  it('should format positive offset correctly', () => {
    const result = formatUtcOffset(120);
    expect(result).toEqual('UTC-02:00');
  });

  it('should format negative offset correctly', () => {
    const result = formatUtcOffset(-180);
    expect(result).toEqual('UTC+03:00');
  });

  it('should format zero offset correctly', () => {
    const result = formatUtcOffset(0);
    expect(result).toEqual('UTC+00:00');
  });
});
