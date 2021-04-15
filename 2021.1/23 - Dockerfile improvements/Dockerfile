# Compile stage
FROM golang:1.16.2-alpine3.13 as build

RUN go version

# Step 1. Observe the line separatator from below.
# Step 2. Observe that you can fold/unfold the contents of this file.
# Step 3. If the container building takes too long, you can now cancel it.
# Step 4. Invoke completion in the container name for any of the stage.
# Step 5. Edit the Run Configuration and change the Bind mounts option.

# Run stage
FROM alpine:3.13.2

CMD ["echo", "Hello!"]