import { Reference } from './reference'
import { Milestone } from './milestone'
import { Footnote } from './footnote'
import { Break } from './break'
import { Text } from './text'
import { Attributes, StyleFactoryFilter, UsxContext, UsxItemContainer, UsxItemFactory } from './item'
import { IntroCharStyle, IntroCharStyles } from './styles'
import { Char } from './char'

type IntroCharType = Reference | Char | IntroChar | Milestone | Footnote | Break | Text

export class IntroChar extends UsxItemContainer<IntroCharType> {

  style: IntroCharStyle

  // char.closed?

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString() as IntroCharStyle
  }
}

export class IntroCharFactory extends UsxItemFactory<IntroChar> {
  static readonly instance = new IntroCharFactory()

  private constructor() {
    super('char', new StyleFactoryFilter(IntroCharStyles))
  }

  protected onInitialize() {
  }

  create(context: UsxContext, attributes: Attributes): IntroChar {
    return new IntroChar(context, attributes)
  }
}