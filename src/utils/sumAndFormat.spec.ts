import {describe, expect, it} from 'vitest';

import {sumAndFormat} from './formatter';

describe('sumAndFormat', () => {
  it('should sum and format values correctly', () => {
    const result = sumAndFormat(['$10.50', '$20.75', '$5.25']);
    expect(result).toEqual('$36.50');
  });

  it('should handle empty array', () => {
    const result = sumAndFormat([]);
    expect(result).toEqual('$0.00');
  });

  it('should handle values with cents', () => {
    const result = sumAndFormat(['$5.35', '$3.10', '$7.45']);
    expect(result).toEqual('$15.90');
  });
});
