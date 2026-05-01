const COLORS = ['black', 'brown', 'red', 'orange', 'yellow', 'green', 'blue', 'violet', 'grey', 'white'] as const;

type Color = typeof COLORS[number];

export function decodedResistorValue(colors: Color[]): string {
  const value = (COLORS.indexOf(colors[0]) * 10 + COLORS.indexOf(colors[1])) * Math.pow(10, COLORS.indexOf(colors[2]));

  if (value >= 1_000_000_000) return `${value / 1_000_000_000} gigaohms`;
  if (value >= 1_000_000) return `${value / 1_000_000} megaohms`;
  if (value >= 1_000) return `${value / 1_000} kiloohms`;
  return `${value} ohms`;
}
