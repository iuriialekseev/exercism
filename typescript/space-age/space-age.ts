enum PlanetRatios {
  Earth = 31557600,
  Mercury = 0.2408467 * Earth,
  Venus = 0.61519726 * Earth,
  Mars = 1.8808158 * Earth,
  Jupiter = 11.862615 * Earth,
  Saturn = 29.447498 * Earth,
  Uranus = 84.016846 * Earth,
  Neptune = 164.79132 * Earth,
}

export default class SpaceAge {
  public seconds: number;

  constructor(seconds: number) {
    this.seconds = seconds;
  }

  onEarth = (): number => this.toYear(PlanetRatios.Earth);
  onMercury = (): number => this.toYear(PlanetRatios.Mercury);
  onVenus = (): number => this.toYear(PlanetRatios.Venus);
  onMars = (): number => this.toYear(PlanetRatios.Mars);
  onJupiter = (): number => this.toYear(PlanetRatios.Jupiter);
  onSaturn = (): number => this.toYear(PlanetRatios.Saturn);
  onUranus = (): number => this.toYear(PlanetRatios.Uranus);
  onNeptune = (): number => this.toYear(PlanetRatios.Neptune);

  private toYear(ratio: number): number {
    return Number((this.seconds / ratio).toFixed(2));
  }
}
