import { AnalyticElement, AnalyticEvent, AnalyticEventType, ContentElement, IAnalyticElement, IAnalyticEvent, IContentElement } from './event'


export interface AnalyticEventFactory {

  createEvent(event: IAnalyticEvent): Promise<AnalyticEvent>
}

let factory: AnalyticEventFactory = {
  async createEvent(event) {
    return new DefaultAnalyticEvent(event, new DefaultAnalyticElement(event.element, event.element && event.element.content ? event.element.content.map((c) => new DefaultContentElement(c)) : []))
  },
}

export function getAnalyticEventFactory(): AnalyticEventFactory {
  return factory
}

export function setAnalyticEventFactory(newFactory: AnalyticEventFactory) {
  factory = newFactory
}

export class DefaultContentElement extends ContentElement {

  readonly id: string
  readonly type: string
  readonly percent: number

  constructor(element: IContentElement) {
    super()
    this.id = element.id
    this.type = element.type
    this.percent = element.percent
  }

  clone(): ContentElement {
    return new DefaultContentElement({ id: this.id, type: this.type, percent: this.percent })
  }
}

export class DefaultAnalyticElement extends AnalyticElement {
  private readonly element: IAnalyticElement
  readonly content: ContentElement[]

  constructor(element: IAnalyticElement, content: ContentElement[]) {
    super()
    this.element = element
    this.content = content
  }

  get id(): string {
    return this.element.id
  }

  get type(): string {
    return this.element.type
  }

  clone(): AnalyticElement {
    return new DefaultAnalyticElement(this.element, this.content.map((c) => c.clone()))
  }
}

export class DefaultAnalyticEvent extends AnalyticEvent {

  private readonly event: IAnalyticEvent
  readonly element: AnalyticElement
  readonly created: Date = new Date()

  constructor(event: IAnalyticEvent, element: AnalyticElement) {
    super()
    this.event = event
    this.element = element
  }

  get type(): AnalyticEventType {
    return this.event.type
  }

  get name(): string {
    return this.event.name
  }

  toParameters(): any {
    const parameters: { [key: string ]: string } = {
      type: this.type.toString(),
      element_id: this.element.id,
      created: this.created.toISOString(),
    }
    if (this.element.content) {
      let ix = 0
      for (const content of this.element.content) {
        parameters['content_id_' + (ix++)] = content.id
        parameters['content_id_type_' + (ix++)] = content.type
        if (content.percent) {
          parameters['content_id_percent_' + (ix++)] = content.percent.toFixed(1)
        }
      }
    }
    return parameters
  }

  clone(): AnalyticEvent {
    return new DefaultAnalyticEvent(this.event, this.element.clone())
  }
}

