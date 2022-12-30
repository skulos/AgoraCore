FROM node:14.20.1-buster

RUN npm i -g pnpm

RUN pnpm -v