FROM golang:1.22.3-alpine AS build

# Set the working directory
WORKDIR /forum-project

# Install required packages
RUN apk add --no-cache gcc musl-dev

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o forum .

# Stage 2: Runtime
FROM alpine:latest

# Set the working directory
WORKDIR /myProject

# Copy the built application from the build stage
COPY --from=build /forum-project/forum /myProject/
COPY --from=build /forum-project/frontend /myProject/frontend
COPY --from=build /forum-project/base.db /myProject/

# Metadata for the image
LABEL version="1.0"
LABEL projectname="EduTalks"

# Set the command to run the application
CMD ["./forum", "base.db"]
