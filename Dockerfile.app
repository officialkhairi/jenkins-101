# Image for the Go app. The binary is built by the pipeline's Build stage.
# (The plain "Dockerfile" in this folder builds Jenkins itself.)
FROM gcr.io/distroless/static-debian12:nonroot
COPY app /app
EXPOSE 9090
ENTRYPOINT ["/app"]
