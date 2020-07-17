# Start from the latest golang base image
FROM golang:latest

# Copy the local package files to the container's workspace.
ADD . /CoCo/kyma-test/graphql-server

# Set the Current Working Directory inside the container
WORKDIR /CoCo/kyma-test/graphql-server

# Expose port
EXPOSE 3030

# Run the app
CMD ["go", "run", "server.go"]



