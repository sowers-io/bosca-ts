FROM node:20 AS build

WORKDIR /build

COPY . .

RUN npm install
RUN rm .env.dev
RUN npm run build

FROM node:20 AS release

WORKDIR /app

COPY --from=build /build/.output/ /app

EXPOSE 3000

RUN echo "echo 'Starting Server...' && NODE_ENV=production node " > /app/entrypoint

USER 1000

ENTRYPOINT ["sh", "/app/entrypoint"]