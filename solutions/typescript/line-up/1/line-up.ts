function ordinalSuffix(n: number): string {
  const lastTwo = n % 100;
  if (lastTwo >= 11 && lastTwo <= 13) return 'th';
  switch (n % 10) {
    case 1: return 'st';
    case 2: return 'nd';
    case 3: return 'rd';
    default: return 'th';
  }
}

export function format(name: string, number: number): string {
  return `${name}, you are the ${number}${ordinalSuffix(number)} customer we serve today. Thank you!`;
}
