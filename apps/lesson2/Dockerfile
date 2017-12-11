FROM golang:1.9


RUN mkdir -p /go/src/agato/go-exercise/apps/lesson2
WORKDIR /go/src/agato/go-exercise/apps/lesson2
COPY . /go/src/agato/go-exercise/apps/lesson2

# third party
# gin auto build tool
RUN go get github.com/codegangsta/gin
RUN go get github.com/gorilla/websocket
RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

# Set the PORT environment variable inside the container
ENV PORT 8080
# Expose port 3000 to the host so we can access the gin proxy
EXPOSE 3000


# ["app"]
CMD ["gin", "run"]

