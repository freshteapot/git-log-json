# Git Log to JSON
- Looks for subjects starting with Change
- Builds a json object

# Usage
- argument = path to git repo
```sh
go run main.go ~/git/learnalist-api
```

## Example where I pipe it thru jq
```sh
go run main.go ~/git/learnalist-api | jq > changelog.json
```

# Example of json output

```json
[
  {
    "hash": "6511c7",
    "what": "The body of the git message",
    "when": "2021-03-29T14:29:39+02:00",
    "pr": "235"
  }
]
```
