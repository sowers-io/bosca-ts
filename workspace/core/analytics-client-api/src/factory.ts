import { AnalyticElement, AnalyticEvent, AnalyticEventType, IAnalyticElement, IAnalyticEvent } from "./event"


export interface AnalyticEventFactory {

  createEvent(event: IAnalyticEvent): Promise<AnalyticEvent>
}

let factory: AnalyticEventFactory = {
  async createEvent(event) {
    return new DefaultAnalyticEvent(event, new DefaultAnalyticElement(event.element))
  },
}

export function getAnalyticEventFactory(): AnalyticEventFactory {
  return factory
}

export function setAnalyticEventFactory(newFactory: AnalyticEventFactory) {
  factory = newFactory
}

export class DefaultAnalyticElement extends AnalyticElement {
  private element: IAnalyticElement

  constructor(element: IAnalyticElement) {
    super()
    this.element = element
  }

  get id(): string {
    return this.element.id
  }

  clone(): AnalyticElement {
    return new DefaultAnalyticElement(this.element)
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
    return {
      type: this.type.toString(),
      elementId: this.element.id,
    }
  }

  clone(): AnalyticEvent {
    return new DefaultAnalyticEvent(this.event, this.element.clone())
  }
}

