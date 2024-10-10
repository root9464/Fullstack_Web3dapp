FROM oven/bun
WORKDIR /app
COPY package*.json ./
COPY bun.lockb ./
RUN bun i
COPY . .
EXPOSE 3000
CMD ["bun", "start"]