import { dtos } from "@wails/go/models";

export enum ResourceTypes {
  VIDEO = "VIDEO",
  IMAGE = "IMAGE",
}

export enum ContentType {
  SCREEN = "SCREEN",
  ADVERTISEMENT = "ADVERTISEMENT",
}

export class Resource {
  id: number;
  name: string;
  type: ResourceTypes;
  contentType: ContentType;

  static fromAdvertisement(ad: dtos.Ad) {
    return new Resource(
      ad.id,
      ad.name,
      ad.type as ResourceTypes,
      ContentType.ADVERTISEMENT
    );
  }

  static fromScreen(screen: dtos.ScreenType) {
    return new Resource(
      screen.id,
      screen.name,
      ResourceTypes.IMAGE,
      ContentType.SCREEN
    );
  }

  static fromDropData(data: unknown) {
    const res = data as Resource;
    return new Resource(res.id, res.name, res.type, res.contentType);
  }

  constructor(
    id: number,
    name: string,
    type: ResourceTypes,
    contentCategory: ContentType
  ) {
    this.id = id;
    this.name = name;
    this.type = type;
    this.contentType = contentCategory;
  }

  get icon(): string {
    return this.type === ResourceTypes.VIDEO ? "film" : "image";
  }
}
