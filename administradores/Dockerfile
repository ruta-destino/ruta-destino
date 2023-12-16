FROM node:21.4-alpine as builder
WORKDIR /app
COPY ./package.json ./package-lock.json ./
RUN npm ci
COPY . .
RUN npm run build

FROM node:21.4-alpine
WORKDIR /app
COPY --from=builder /app/package.json /app/package-lock.json ./
COPY --from=builder /app/build /app/build
CMD [ "node", "build" ]
