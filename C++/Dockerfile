# Download ubuntu from docker hub
FROM amytabb/docker_ubuntu16_essentials

# Copy this folder to the folder inside the image
COPY HelloWorld /HelloWorld

# Set the working directory
WORKDIR /HelloWorld/

# Compile program using g++
RUN g++ -o HelloWorld Hello.cpp

# Run the application inside the image
CMD ["./HelloWorld"]