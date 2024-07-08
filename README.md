# Bank API

This project provides a simple HTTP API for managing bank accounts, allowing users to perform operations such as retrieving account statements, making deposits, withdrawals, and transferring money between accounts.

## Features

- Retrieve account statements
- Make deposits
- Make withdrawals
- Transfer money between accounts

## Endpoints

### GET /statement

Retrieve the statement for a specific account.

#### Query Parameters:

- `number:` Account number (required)

#### Example Request:

``` bash
GET /statement?number=1001
```

#### Example Response:

```json
{
  "customer": {
    "name": "Jacob",
    "address": "1234 Main St",
    "phone": "123-456-7890"
  },
  "number": 1001,
  "balance": 100.0
}
```

### POST /deposit

Deposit an amount into a specific account.

#### Query Parameters:

- `number:` Account number (required)
- `amount:` Amount to deposit (required)

#### Example Request:

```bash
POST /deposit?number=1001&amount=50
```

#### Example Response:

``` yaml
Balance after deposit: 150.0
```

### POST /withdraw

Withdraw an amount from a specific account.

#### Query Parameters:

- `number:` Account number (required)
- `amount:` Amount to withdraw (required)

#### Example Request:

```bash
POST /withdraw?number=1001&amount=50
```

#### Example Response:

```yaml
Balance after withdrawal: 50.0
```

### POST /transfer

Transfer an amount from one account to another.

#### Query Parameters:

- `from:` Source account number (required)
- `to:` Destination account number (required)
- `amount:` Amount to transfer (required)

#### Example Request:

```bash
POST /transfer?from=1001&to=1002&amount=50
```

#### Example Response:

```yaml
From account balance: 50.0
To account balance: 150.0
```

## How to Run

1. Ensure you have Go installed.
2. Install the bankcore package from GitHub:

```sh
go get github.com/Jacob-Krueckl/bankcore/bank
```

3. Create a file main.go with the provided code.
4. Run the application:

```sh
go run main.go
```

The server will start on localhost:8000.

## Dependencies

- encoding/json
- fmt
- log
- net/http
- strconv
- github.com/Jacob-Krueckl/bankcore/bank

## Example Accounts

The project initializes with two example accounts:

1. Account 1001
    - Customer: Jacob
    - Address: 1234 Main St
    - Phone: 123-456-7890
2. Account 1002
    - Customer: Emily
    - Address: 1234 Main St
    - Phone: 123-456-7890

## Error Handling

- If required query parameters are missing or invalid, the server responds with 400 Bad Request and an appropriate error message.
- If the account number or amount is invalid, an error message is returned.
- If an account is not found, an appropriate error message is returned.
