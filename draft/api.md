# API Design

## Objective

* To design building and my notes.

## Design

| Method | Path | Description |
| ---- | ---- | ---- |
| **GET** | /AS_NUMBER | returns ipfs hash |
| **GET** | /AS_NUMBER/prefixes/4 | Get IPv4 Prefixes |
| **GET** | /AS_NUMBER/prefixes/6 | Get IPv6 Prefixes |
| **POST** | /AS_NUMBER | Post IPFS Content |

## Examples

### GET `/AS_NUMBER/`

```json:retval.json
{
  "as_number": 65000;
  "owner": "0xdeadbeef";
  "ipfs_hash": "0x12345678";
  "prefixes": {
    "ipv4": [
      "10.0.0.0/16",
      "10.1.0.0/16",
    ],
    "ipv6": [
      "fc00:200::/32",
      "fc00:201::/32"
    ],
  }
}
```

* `prefixes` field will generate automatically from `ipfs_hash` content.

### POST `/AS_NUMBER`

```
{
  owner: "0xdeadbeef";
  object: "blah blah blah";
}
```

* I exactly do not know how to write data to smart contract via API server.
  * Does owner of AS directly write a value to the contract?
