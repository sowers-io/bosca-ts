export enum AnalyticEventType {
  interaction = 'interaction',
  impression = 'impression',
  completion = 'completion',
}

export interface IAnalyticElement {
  id: string
  extras: { [key: string]: string }
}

export interface IAnalyticEvent {
  type: AnalyticEventType
  name: string
  element: IAnalyticElement
}

export abstract class AnalyticElement {
  abstract get id(): string
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
