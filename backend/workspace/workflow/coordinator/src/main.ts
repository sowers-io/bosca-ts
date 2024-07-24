import { fastify } from "fastify";
import { fastifyConnectPlugin } from "@connectrpc/connect-fastify";
import routes from "./service";

async function main() {
  const server = fastify({
    http2: true,
  });
  await server.register(fastifyConnectPlugin, { routes });
  await server.listen({ host: "0.0.0.0", port: 5000 });
  console.log("server is listening at", server.addresses()[0].address + ':' + server.addresses()[0].port);
}

void main();
