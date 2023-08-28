
# Iptables-GO

A rest API using GOlang to manage linux iptables



## API Reference

#### Get all rules

```http
  GET /v1/input
```

#### Add Rule

```http
  POST /v1/input
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `sourceAddress` | `string` | **Optional**. source address |
| `destinationAddress` | `string` | **Optional**. destination address |
| `protocol` | `string` | **Optional**. protocol |
| `sourcePort` | `int` | **Optional**. source port |
| `destinationPort` | `int` | **Optional**. destination port |
| `target` | `string` | **Optional**. target |
| `ruleNumber` | `int` | **Optional**. rule number |

#### Update Rule

```http
  PATCH /v1/input/{ruleNumber}
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `sourceAddress` | `string` | **Optional**. source address |
| `destinationAddress` | `string` | **Optional**. destination address |
| `protocol` | `string` | **Optional**. protocol |
| `sourcePort` | `int` | **Optional**. source port |
| `destinationPort` | `int` | **Optional**. destination port |
| `target` | `string` | **Optional**. target |


#### Delete Rule

```http
  DELETE /v1/input/{ruleNumber}
```

## Run Locally

Clone the project

```bash
  git clone https://github.com/akashkola/iptables-go.git
```

Go to the project directory

```bash
  cd iptables-go
```

Install dependencies

```bash
  go get ./..
```

Start the server

```bash
  make run
```


## Contributing

Contributions are always welcome!

## Tech Stack

**Server:** GOlang


## Authors

- [akashkola](https://www.github.com/akashkola)


