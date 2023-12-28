# cloud-functions-hello-go
This is a simple example of a Google Cloud Function written in Go.

## Deploying
To deploy this function, you can use the `gcloud` command line tool:

```bash
gcloud functions deploy go-http-function \
    --gen2 \
    --runtime go121 \
    --region=asia-northeast1 \
    --source=. \
    --entry-point HelloHTTP \
    --trigger-http \
    --allow-unauthenticated
```

## Testing
To test this function, access the URL that is printed after deploying the function.
