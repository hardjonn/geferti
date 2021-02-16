ARG GOLANG_VERSION="1.15.8"
FROM golang:${GOLANG_VERSION} as golang

ENV USER_ID 1000
ENV GROUP_ID 1000

ENV APP_USER geferti
ENV APP_GROUP geferti
ENV APP_HOME /app


RUN groupadd --gid $GROUP_ID $APP_GROUP && useradd -m -l --uid $USER_ID --gid $GROUP_ID $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_GROUP $APP_HOME

USER $APP_USER
WORKDIR $APP_HOME


COPY . .

# RUN go get github.com/githubnemo/CompileDaemon

# ENTRYPOINT CompileDaemon

# RUN go get -d -v ./...
# RUN go install -v ./...

# it will keep the container alive
CMD tail -f /dev/null

EXPOSE 443
