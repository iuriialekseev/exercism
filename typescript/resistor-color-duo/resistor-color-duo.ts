export class ResistorColor {
  private colors: string[];
  private static COLORS_MAP = ["black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"];

  constructor(colors: string[]) {
    if (colors.length < 2) throw "At least two colors need to be present";

    this.colors = colors;
  }

  value = (): number => {
    const colors = this.colors.slice(0, 2);
    const result = colors.map((color: string) => ResistorColor.COLORS_MAP.indexOf(color)).join("");

    return Number(result);
  };
}
