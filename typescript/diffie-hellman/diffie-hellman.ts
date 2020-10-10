export default class DiffieHellman {
  constructor(private readonly p: number, private readonly g: number) {
    if (!this.isPrime(p) || !this.isPrime(g)) {
      throw new Error("Arguments must be prime numbers!");
    }
  }

  getPublicKeyFromPrivateKey(privateKey: number) {
    if (privateKey < 2 || privateKey >= this.p) {
      throw new Error("Invalid private key");
    }
    return this.g ** privateKey % this.p;
  }

  getSharedSecret(privateKey: number, publicKey: number) {
    return publicKey ** privateKey % this.p;
  }

  private isPrime(num: number) {
    for (let i = 2; i < num; i++) if (num % i === 0) return false;
    return num > 1;
  }
}
