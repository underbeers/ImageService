FROM golang:1.19-alpine AS build

WORKDIR /build

COPY . .

RUN go mod download
RUN go build -o /build/image_service /build/cmd/main.go

FROM alpine:latest

WORKDIR /app
COPY ./configs/service.json ./configs/service.json
COPY ./service.json ./service.json
COPY --from=build /build/image_service .

ENV CLOUDINARY_CLOUD_NAME=$CLOUDINARY_CLOUD_NAME
ENV CLOUDINARY_API_KEY=$CLOUDINARY_API_KEY
ENV CLOUDINARY_API_SECRET=$CLOUDINARY_API_SECRET
ENV CLOUDINARY_UPLOAD_FOLDER=$CLOUDINARY_UPLOAD_FOLDER
ENV IMAGESERVICE_PORT=$IMAGESERVICE_PORT
ENV IMAGESERVICE_IP=$IMAGESERVICE_IP
ENV GATEWAY_IP=$GATEWAY_IP
ENV GATEWAY_PORT=$GATEWAY_PORT

CMD [ "./image_service" ]
