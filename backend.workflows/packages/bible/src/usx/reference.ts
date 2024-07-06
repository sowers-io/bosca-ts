import { Attributes, UsxContext, UsxItemContainer, UsxItemFactory } from './item'
import { Text } from './text'

export class Reference extends UsxItemContainer<Text> {

  loc: string

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.loc = attributes.LOC.toString()
  }
}

export class ReferenceFactory extends UsxItemFactory<Reference> {

  static readonly instance = new ReferenceFactory()

  private constructor() {
    super('ref')
  }

  protected onInitialize() {
  }

  create(context: UsxContext, attributes: Attributes): Reference {
    return new Reference(context, attributes)
  }
}
