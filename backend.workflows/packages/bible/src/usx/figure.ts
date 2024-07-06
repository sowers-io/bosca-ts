import { Attributes, UsxContext, UsxItemContainer, UsxItemFactory } from './item'
import { Text } from './text'

export class Figure extends UsxItemContainer<Text> {

  style: string
  alt?: string
  file: string
  size?: string
  loc?: string
  copy?: string
  ref?: string

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString()
    this.alt = attributes.ALT?.toString()
    this.file = attributes.FILE.toString()
    this.size = attributes.SIZE?.toString()
    this.loc = attributes.LOC?.toString()
    this.copy = attributes.COPY?.toString()
    this.ref = attributes.REF?.toString()
  }
}

export class FigureFactory extends UsxItemFactory<Figure> {

  static readonly instance = new FigureFactory()

  private constructor() {
    super('figure')
  }

  protected onInitialize() {
  }

  create(context: UsxContext, attributes: Attributes): Figure {
    return new Figure(context, attributes)
  }
}