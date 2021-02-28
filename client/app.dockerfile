# build environment
FROM node:13.12.0-alpine as build
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
COPY client/package.json ./
COPY client/package-lock.json ./
RUN npm ci --silent
RUN npm install react-scripts@3.4.1 -g --silent
COPY client ./
RUN npm run build

# start app
# dev server not for prod---shivan
CMD npm run start
