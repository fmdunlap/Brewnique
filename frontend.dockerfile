FROM node:20-slim AS base
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
COPY . /app
WORKDIR /app

RUN pnpm i && pnpm i -g serve
RUN pnpm run build --mode staging

EXPOSE 8000
CMD [ "serve", "-p", "8000", "-C", "frontend/dist" ]