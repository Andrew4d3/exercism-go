type Planet = 'mercury' | 'venus' | 'earth' | 'mars' | 'jupiter' | 'saturn' | 'uranus' | 'neptune';

const EARTH_YEAR_SECONDS = 31_557_600;

const ORBITAL_PERIODS: Record<Planet, number> = {
  mercury: 0.2408467,
  venus:   0.61519726,
  earth:   1.0,
  mars:    1.8808158,
  jupiter: 11.862615,
  saturn:  29.447498,
  uranus:  84.016846,
  neptune: 164.79132,
};

export function age(planet: Planet, seconds: number): number {
  // Multiply by 100, round to integer, divide by 100 — JS's way of rounding to 2 decimal places
  return Math.round(seconds / (EARTH_YEAR_SECONDS * ORBITAL_PERIODS[planet]) * 100) / 100;
}
