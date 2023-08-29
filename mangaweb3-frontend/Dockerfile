FROM node:16-alpine

WORKDIR /app

COPY . .

RUN npm ci
RUN npm run build

EXPOSE 5173
CMD ["node", "-r", "detenv/config" "build"]