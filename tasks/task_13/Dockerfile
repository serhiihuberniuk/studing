FROM golang
RUN mkdir newFolder
COPY . / newFolder/
ADD . / newFolder/
WORKDIR /newFolder
RUN go build main.go