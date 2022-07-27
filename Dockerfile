FROM golang:1.17.8 as build

WORKDIR /api
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .
# copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /api/main /main
ENTRYPOINT [ "/main" ]   