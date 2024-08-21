export enum AnalyticEventType {
  interaction,
  impression,
  completion,
}

export abstract class AnalyticSubject {
  abstract get id(): string
}

export abstract class AnalyticEvent {
  abstract get type(): AnalyticEventType
  abstract get name(): string
  abstract get created(): Date
  abstract get subject(): AnalyticSubject

  abstract toParameters(): any
  abstract clone(): AnalyticEvent
}

