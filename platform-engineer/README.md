# Platform Assignment

HTTP service which exposes an API `/redact` which will be used to redact a JSON based on a JSON path specified.

The JSON path uses "." as a delimiter for nested JSON objects.
In case the JSON object is a list, instead of using a key, one can use:

1. "*" indicating that all the elements in the list must be redacted
2. "index", eg. 0,1 which indicates a particular index to be redacted.

eg.

Consider the JSON:

```json
{
  "a": {
    "b": {
      "c": "va12345l",
      "d": "val",
      "e": "val"
    },
    "l": [
      {
        "k": "v1"
      },
      {
        "k": "v2"
      },
      {
        "k": "v3"
      }
    ]
  }
}
```

Path a.b.c indicates that "va12345l" should be redacted.

Path a.l.*.k indicates that "v1", "v2" and "v3" will be redacted.

Path a.l.0.k indicates that "v1" will be redacted.


## Getting Started

### Setup

Clone Repo:

`git pull https://gitlab.com/vernacularai/platform-engg-assignment.git`

No dependencies to download. Uses Go stdlib.

### Running

Build Docker

`docker build -t platform-assignment .`

Run Docker

`docker run -d -p 8001:8003 platform-assignment`

Use this CURL to test the redaction:

```
curl --location --request POST 'http://localhost:8001/redact/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "json_to_redact": {
        "a": {
            "b": {
                "c": "va12345l",
                "d": "val",
                "e": "val"
            },
            "l": [
                {"k": "a"},
                {"k": "b"},
                {"k": "c"}
            ]
        }
    },
    "redact_regexes": [
        {
            "path": "a.b.c",
            "regexes": [
                "[0-9]{5}"
            ]
        },
        {
            "path": "a.b.d",
            "regexes": [
                "[0-9]{5}"
            ]
        }
    ],
    "redact_completely": [
        "a.b.e",
        "a.b.f",
        "a.l.*.k"
    ]
}'
```

Response to the CURL:

```
{
    "a": {
        "b": {
            "c": "va*****l",
            "d": "***",
            "e": "***"
        },
        "l": [
            {
                "k": "*"
            },
            {
                "k": "*"
            },
            {
                "k": "*"
            }
        ]
    }
}
```