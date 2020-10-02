# Get the GCC preinstalled image from Docker Hub
FROM gcc:4.9

# Copy encode folder into encode folder
COPY encode /encode

# Set the working directory
WORKDIR /encode/

# Optional -- Check our txt file
# RUN cat haiku1.txt

# Compile program using makefile specification
RUN make encode

# Run the application inside the image
CMD ["/bin/sh", "-c", "./encode insertion <haiku1.txt"]