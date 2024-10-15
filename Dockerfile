FROM golang:1.23-alpine as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN ls
COPY . ./
RUN ls -a
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/bluesky-backend cmd/main.go

FROM scratch as release
WORKDIR /production
COPY --from=build /app/.env /app/bin/bluesky-backend ./
EXPOSE 2000
ENTRYPOINT [ "./bluesky-backend" ]
