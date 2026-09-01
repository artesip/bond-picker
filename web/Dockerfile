FROM node:26.8.1-alpine3.24 AS base

FROM base AS builder
WORKDIR /app

COPY package.json pnpm-lock.yaml* pnpm-workspace.yaml ./
RUN npm install -g pnpm && pnpm install --frozen-lockfile
COPY . .
RUN pnpm build

FROM node:26.8.1-alpine3.24 AS runner
WORKDIR /app

ENV NODE_ENV=production

RUN addgroup --system --gid 1001 nodejs && \
    adduser --system --uid 1001 nodejs

RUN mkdir -p /app/data && \
    chown -R nodejs:nodejs /app/data

COPY --from=builder /app/.output ./.output

EXPOSE 3000

USER nodejs

CMD ["node", ".output/server/index.mjs"]