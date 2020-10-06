class Isogram {
  static isIsogram(str: string): boolean {
    const letters = str.replace(/[ -]/g, "").toLowerCase();
    const uniqueLetters = new Set(letters);

    return letters.length == uniqueLetters.size;
  }
}

export default Isogram;
