# Day 1 - Health Check API

## Overview

Simple health check API for Day 1 assignment.

Endpoint:

```http
GET /health-check
```

## Response:
``` json
{
  "message": "ok",
  "service_name": "golang-backend-ebvn-k04-day-1",
  "instance_id": "local-instance-01"
}
```

## Environment Variables

| Variable           | Default                         | Description  |
| ------------------ | ------------------------------- | ------------ |
| `API_APP_PORT`     | `8081`                          | Server port  |
| `API_SERVICE_NAME` | `golang-backend-ebvn-k04-day-1` | Service name |
| `API_INSTANCE_ID`  | auto-generated UUID             | Instance ID  |
