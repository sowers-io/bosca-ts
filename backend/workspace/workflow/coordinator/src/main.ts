import { fastify } from "fastify";
import { fastifyConnectPlugin } from "@connectrpc/connect-fastify";
import routes from "./service";
import { logger, newLoggingInterceptor } from "@bosca/common";

async function main() {
  const server = fastify({
    http2: true,
    logger: {
      level: 'debug',
    },
  });
  server.setErrorHandler((error, request, reply) => {
    logger.error({ error, request }, "uncaught error");
    reply.status(500).send({ ok: false });
  });
  await server.register(fastifyConnectPlugin, {
    routes,
    interceptors: [newLoggingInterceptor()],
  });
  await server.listen({ host: "0.0.0.0", port: 7200 });
}

void main();
