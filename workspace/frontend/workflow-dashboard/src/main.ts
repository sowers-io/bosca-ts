import { createBullBoard } from '@bull-board/api'
import { BullMQAdapter } from '@bull-board/api/bullMQAdapter'
import { FastifyAdapter } from '@bull-board/fastify'
import { Queue as QueueMQ } from 'bullmq'
import fastify from 'fastify'
import { logger } from '@bosca/common'

async function main() {
  const queueNames = [
    'metadata',
    'traits',
    'transition',
    'bible',
    'bible-verse',
    'bible-book',
    'bible-ai',
    'search-index',
  ]
  const app = fastify()
  const serverAdapter = new FastifyAdapter()
  createBullBoard({
    queues: queueNames.map(
      (name) =>
        new BullMQAdapter(
          new QueueMQ(name, {
            connection: {
              port: 6379,
              host: 'localhost',
            },
          }),
        ),
    ),
    serverAdapter,
  })
  serverAdapter.setBasePath('/ui')
  app.register(serverAdapter.registerPlugin(), { prefix: '/ui', basePath: '/ui' })

  await app.listen({ port: 3001 })
  logger.info('For the UI, open http://localhost:3001/ui')
}

main().catch((e) => {
  logger.error(e)
  process.exit(1)
})
