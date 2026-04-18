export enum Color {
  black = 0,
  brown = 1,
  red = 2,
  orange = 3,
  yellow = 4,
  green = 5,
  blue = 6,
  violet = 7,
  grey = 8,
  white = 9,
}

export const COLORS = Object.keys(Color).filter(k =>
  isNaN(Number(k))
) as Array<keyof typeof Color>;

export const colorCode = (color: keyof typeof Color): number => Color[color];
