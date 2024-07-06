import { BookIdentificationCode, BookIdentificationCodes } from '../identification'
import { Attributes, CodeFactoryFilter, UsxContext, UsxItemContainer, UsxItemFactory } from './item'
import { Text, TextFactory } from './text'

export class BookIdentification extends UsxItemContainer<Text> {
  id!: string
  code: BookIdentificationCode

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.id = attributes.STYLE.toString()
    this.code = attributes.CODE.toString() as BookIdentificationCode
  }
}

export class BookIdentificationFactory extends UsxItemFactory<BookIdentification> {

  static readonly instance = new BookIdentificationFactory()

  private constructor() {
    super('book', new CodeFactoryFilter<BookIdentificationCode>(BookIdentificationCodes))
  }

  protected onInitialize() {
    this.register(TextFactory.instance)
  }

  create(context: UsxContext, attributes: Attributes): BookIdentification {
    return new BookIdentification(context, attributes)
  }
}