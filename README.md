# Git Log to JSON
- Looks for subjects starting with Change
- Builds a json object

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
