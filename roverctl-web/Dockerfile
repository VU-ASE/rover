# Source: https://gist.github.com/aradalvand/04b2cad14b00e5ffe8ec96a3afbb34fb
FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json .
RUN npm ci
COPY . .
# These are runtime arguments, but they do need to exist at built time to be able to import from public in svelte
ARG PUBLIC_ROVERD_HOST=""
ARG PUBLIC_ROVERD_PORT=""
ARG PUBLIC_ROVERD_USERNAME=""
ARG PUBLIC_ROVERD_PASSWORD=""
ARG PUBLIC_PASSTHROUGH_HOST=""
ARG PUBLIC_PASSTHROUGH_PORT=""

RUN npm run build
RUN npm prune --production

FROM node:20-alpine
WORKDIR /app
COPY --from=builder /app/build build/
COPY --from=builder /app/node_modules node_modules/
COPY package.json .
EXPOSE 3000
ENV NODE_ENV=production
CMD [ "node", "build" ]