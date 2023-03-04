FROM node:18 as node
FROM golang:1.20.1

ENV APP_ROOT /app

COPY --from=node /usr/local/bin/node /usr/local/bin/
COPY --from=node /opt/yarn* /opt/yarn
RUN ln -s /opt/yarn/bin/yarn /usr/local/bin/yarn

WORKDIR $APP_ROOT

COPY package.json yarn.lock ./
RUN yarn install

COPY go.mod go.sum ./
RUN go mod download && go mod verify

EXPOSE 1323
EXPOSE 3000
