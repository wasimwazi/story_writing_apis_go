FROM golang:latest as backend

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o story-api

FROM alpine
WORKDIR /app
COPY --from=backend /app/story-api .
CMD ./story-api

