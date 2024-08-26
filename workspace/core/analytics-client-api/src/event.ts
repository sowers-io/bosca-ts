export enum AnalyticEventType {
  interaction = 'interaction',
  impression = 'impression',
  completion = 'completion',
}

export interface IContentElement {
  id: string
  type: string
  percent: number
}

export interface IAnalyticElement {
  id: string
  type: string
  content: IContentElement[]
  extras: { [key: string]: string }
}

export interface IAnalyticEvent {
  type: AnalyticEventType
  name: string
  element: IAnalyticElement
}

export abstract class ContentElement {
  abstract get id(): string
  abstract get type(): string
  abstract get percent(): number

  abstract clone(): ContentElement
}

export abstract class AnalyticElement {
  abstract get id(): string
  abstract get type(): string
  abstract get content(): ContentElement[]
  
  abstract clone(): AnalyticElement
}

export abstract class AnalyticEvent {

  abstract get type(): AnalyticEventType
  abstract get name(): string
  abstract get created(): Date
  abstract get element(): AnalyticElement

  abstract toParameters(): any
  abstract clone(): AnalyticEvent
}
