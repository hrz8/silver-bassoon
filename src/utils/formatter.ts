export const formatUtcOffset = (offsetMinutes: number): string => {
  const sign = offsetMinutes > 0 ? '-' : '+';
  const absoluteOffset = Math.abs(offsetMinutes);
  const hours = Math.floor(absoluteOffset / 60);
  const minutes = absoluteOffset % 60;

  return `UTC${sign}${hours.toString().padStart(2, '0')}:${minutes.toString().padStart(2, '0')}`;
};

export const sumAndFormat = (values: string[]): string => {
  const numericValues = values.map((value) =>
    parseFloat(value.replace('$', '')),
  );

  const sum = numericValues.reduce((acc, val) => acc + val, 0);

  return `$${sum.toFixed(2)}`;
};
