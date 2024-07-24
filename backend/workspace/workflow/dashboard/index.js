const { createBullBoard } = require('@bull-board/api');
const { BullMQAdapter } = require('@bull-board/api/bullMQAdapter');
const { FastifyAdapter } = require('@bull-board/fastify');
const { Queue: QueueMQ, Worker } = require('bullmq');
const fastify = require('fastify');

const redisOptions = {
  port: 6379,
  host: 'localhost'
};

const createQueueMQ = (name) => new QueueMQ(name, { connection: redisOptions });

const run = async () => {
  const queueNames = ['metadata', 'traits', 'transition', 'bible', 'bible-verse', 'bible-book', 'bible-ai', 'search-index'];
  const app = fastify();
  const serverAdapter = new FastifyAdapter();
  createBullBoard({
    queues: queueNames.map((n) => new BullMQAdapter(createQueueMQ(n))),
    serverAdapter,
  });
  serverAdapter.setBasePath('/ui');
  app.register(serverAdapter.registerPlugin(), { prefix: '/ui' });

  await app.listen({ port: 3000 });
  // eslint-disable-next-line no-console
  console.log('Running on 3000...');
  console.log('For the UI, open http://localhost:3000/ui');
};

run().catch((e) => {
  // eslint-disable-next-line no-console
  console.error(e);
  process.exit(1);
});