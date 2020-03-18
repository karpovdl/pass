FROM karpovdl/golang:alpine

LABEL author="Denis Karpov" \
      site="github.com/karpovdl" \
      email="karpovdl@hotmail.com" \
      version="1.0.0" \
      release-date="2020-03-18"

ENV TZ="Europe/Moscow" \
### APP SRC
    APP_NAME="github.com/karpovdl/pass" \
    APP_NAME_BRANCH="master" \
    APP_VERSION="1.0.0"

COPY /Dockerfile.run $GOPATH/bin/Dockerfile.run

CMD export WORK_DIR=$GOPATH/src \
 && go get -d $APP_NAME \
###
 && cd $GOPATH/src/$APP_NAME \
 && git checkout $APP_NAME_BRANCH > null \
 && cd $GOPATH/src \
### BUILD APP
 && go install -ldflags '-s -w' -v $APP_NAME \
 && cd $GOPATH/bin && tar -C ./ -cf - .
