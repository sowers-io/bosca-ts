import { FastifyReply, FastifyRequest } from 'fastify'

export interface RequestContext {
  request: FastifyRequest
  reply: FastifyReply
}