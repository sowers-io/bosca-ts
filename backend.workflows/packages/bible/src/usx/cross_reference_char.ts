import { Reference, ReferenceFactory } from './reference'
import { Text, TextFactory } from './text'
import { Attributes, StyleFactoryFilter, UsxContext, UsxItemContainer, UsxItemFactory } from './item'
import { CrossReferenceCharStyle, CrossReferenceCharStyles } from './styles'
import { Char, CharFactory } from './char'

type CrossReferenceCharType = Char | CrossReferenceChar | Reference | Text

export class CrossReferenceChar extends UsxItemContainer<CrossReferenceCharType> {

  style: CrossReferenceCharStyle
  //char.link?
  //char.closed?

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString() as CrossReferenceCharStyle
  }
}

export class CrossReferenceCharFactory extends UsxItemFactory<CrossReferenceChar> {

  static readonly instance = new CrossReferenceCharFactory()

  private constructor() {
    super('char', new StyleFactoryFilter(CrossReferenceCharStyles))
  }

  protected onInitialize() {
    this.register(CharFactory.instance)
    this.register(CrossReferenceCharFactory.instance)
    this.register(ReferenceFactory.instance)
    this.register(TextFactory.instance)
  }

  create(context: UsxContext, attributes: Attributes): CrossReferenceChar {
    return new CrossReferenceChar(context, attributes)
  }
}