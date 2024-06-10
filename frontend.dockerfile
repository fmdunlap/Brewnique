FROM node:20-slim AS base
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
COPY . /app
WORKDIR /app

RUN pnpm i && pnpm i -g serve
RUN pnpm run build

EXPOSE 8000
CMD [ "pnpm", "run", "serve" ]