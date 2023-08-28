FROM node:16-alpine AS builder1

WORKDIR /app

COPY . .

RUN npm ci
RUN npm run build

FROM node:16-alpine

WORKDIR /app

COPY --from=builder1 /app/package*.json ./
RUN npm ci --production --ignore-scripts
COPY --from=builder1 /app/build ./

EXPOSE 3000
CMD ["node", "./index.js"]