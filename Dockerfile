FROM golang:1.19-alpine AS build

WORKDIR /build

COPY . .

RUN go mod download
RUN go build -o /build/image_service /build/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=build /build/image_service .

ENV CLOUDINARY_CLOUD_NAME=$CLOUDINARY_CLOUD_NAME
ENV CLOUDINARY_API_KEY=$CLOUDINARY_API_KEY
ENV CLOUDINARY_API_SECRET=$CLOUDINARY_API_SECRET
ENV CLOUDINARY_UPLOAD_FOLDER=$CLOUDINARY_UPLOAD_FOLDER
ENV IMAGESERVICE_PORT=$IMAGESERVICE_PORT

CMD [ "./image_service" ]