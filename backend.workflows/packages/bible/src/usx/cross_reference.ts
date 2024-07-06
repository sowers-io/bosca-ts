import { CrossReferenceChar, CrossReferenceCharFactory } from './cross_reference_char'
import { Text, TextFactory } from './text'
import { Attributes, StyleFactoryFilter, UsxContext, UsxItemContainer, UsxItemFactory } from './item'
import { CrossReferenceStyle, CrossReferenceStyles } from './styles'

type CrossReferenceType = CrossReferenceChar | Text

export class CrossReference extends UsxItemContainer<CrossReferenceType> {

  style: CrossReferenceStyle
  caller: string

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString() as CrossReferenceStyle
    this.caller = attributes.CALLER.toString()
  }
}

export class CrossReferenceFactory extends UsxItemFactory<CrossReference> {

  static readonly instance = new CrossReferenceFactory()

  private constructor() {
    super('note', new StyleFactoryFilter(CrossReferenceStyles))
  }

  protected onInitialize() {
    this.register(CrossReferenceCharFactory.instance)
    this.register(TextFactory.instance)
  }

  create(context: UsxContext, attributes: Attributes): CrossReference {
    return new CrossReference(context, attributes)
  }
}