type DnaNucleotide = 'G' | 'C' | 'T' | 'A';
type RnaNucleotide = 'C' | 'G' | 'A' | 'U';

// Maps each DNA nucleotide to its RNA complement
const DNA_TO_RNA: Record<DnaNucleotide, RnaNucleotide> = { G: 'C', C: 'G', T: 'A', A: 'U' };

function isDna(n: string): n is DnaNucleotide {
  return n in DNA_TO_RNA;
}

export function toRna(dna: string): string {
  // Translate each nucleotide, throwing on anything not in the map
  return dna.split('').map((nucleotide) => {
    if (!isDna(nucleotide)) throw new Error('Invalid input DNA.');
    return DNA_TO_RNA[nucleotide];
  }).join('');
}
