# JavaScript with Docker
This is a basic generalized implementation of using Docker with an application built with JavaScript.

## Steps
* Code your JavaScript application
* In your terminal run npm install
* This will generate your package.json and package-lock.json
* Be sure all dependencies are correctly within your package-lock.json
* Build and run your docker image (steps below)

## Build Docker image
* $ docker build -t javascript-image .
* $ docker image ls
* You should be able to see your new javascript-image image

## Run Docker Image
* $ docker run javascript-image

## My Implementation
* [modern-warfare](https://github.com/abspen1/modern-warfare)


![Alt Text](https://miro.medium.com/max/1000/1*_f186TrKOUrpvoYSQMSP6A.jpeg "JavaScript with Docker")