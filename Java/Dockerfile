# the first stage of our build will use a parent image
FROM alpine:latest

# Add our class file
ADD App.class App.class

# Run update and use the jdk 11 since that's what local compile uses
RUN apk --update add openjdk11-jre

# Add our cmds to run
ENTRYPOINT [ "java", "-Djava.security.egd=file:/dev/./urandom", "App" ]