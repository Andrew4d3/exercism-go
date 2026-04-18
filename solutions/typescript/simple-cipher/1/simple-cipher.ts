export class SimpleCipher {
  key: string

  constructor(key?: string) {
    // If no key is provided, generate a random 100-character lowercase key.
    // Each character is derived by picking a random number 0–25 and adding
    // it to 97 (ASCII 'a') to get a letter in the range 'a'–'z'.
    this.key = key ?? Array.from({ length: 100 }, () =>
      String.fromCharCode(97 + Math.floor(Math.random() * 26))
    ).join('')
  }

  encode(plaintext: string): string {
    return plaintext
      .split('')
      .map((char, i) => {
        // Cycle through the key using modulo so it repeats for long messages.
        // Subtract 97 to convert the key character to a 0-based shift (a=0, b=1, …).
        const shift = this.key[i % this.key.length].charCodeAt(0) - 97
        // Shift the plaintext character forward, wrapping around at 'z'.
        return String.fromCharCode(((char.charCodeAt(0) - 97 + shift) % 26) + 97)
      })
      .join('')
  }

  decode(ciphertext: string): string {
    return ciphertext
      .split('')
      .map((char, i) => {
        const shift = this.key[i % this.key.length].charCodeAt(0) - 97
        // Shift backward to reverse the encoding. Add 26 before taking modulo
        // to avoid negative results (e.g. 0 - 1 = -1 would wrap incorrectly).
        return String.fromCharCode(((char.charCodeAt(0) - 97 - shift + 26) % 26) + 97)
      })
      .join('')
  }
}
