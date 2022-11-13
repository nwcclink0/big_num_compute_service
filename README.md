## Big number compute service 
Big number compute service is a jsonrpc service for create/update/delete number object and compute big number. It support account to differentiate user.
## Content
- [Content](#content)
- [Design](#design)
- [big_num db schema](#big_num-db-schema)
- [Run on native machine](#run-on-native-machine)
- [Run for load balancer with docker](#run-for-load-balancer-with-docker)
- [Run big_num_compute_service](#run-big_num_compute_service)
- [Number object jsonrpc API](#number-object-jsonrpc-api)
- [Account jsonrpc API](#account-jsonrpc-api)

## Design

---
### YAML config example
Big number compute service use config.yml to configurate service.
example:
```
core:
  worker_num: 4 # default worker number is runtime.NumCPU()
  queue_num: 8192 # default queue number is 8192
  address: ""
  port: "8090"
  mode: "docker" # docker/localhost

log:
  format: "string" # string or json
  access_log: "stdout" # stdout: output to console,or define log path like "log/access_log"
  access_level: "debug"
  error_log: "stderr" # stderr: output to console,or define log path like "log/error_log"
  error_level: "error"
```
### Account authentication mechanism
Account authentication contains hashed password/JWT/TOTP mechanism
#### Hashed password
Account password will be hashed to store in database. 
#### Account token
Login account by email and password will return JWT token 
#### Validate account
Big number compute service will send TOTP passcode to account email's mail box to validate account. 
### .env file
```.env``` file need to provide before run big number compute service.
```.env``` file contains
```
JWT_SECRET=your_jwt_secret
MAIL_ACCOUNT=your_sender_gmail_account
MAIL_AUTH=your_sender_gmail_auth
DB_USER=your_db_user
DB_PASSWORD=your_db_password
MAIL_SMTP_HOST=your_smtp_host
MAIL_SMTP_PORT=your_smtp_port

```

## big_num db schema

---
### *Numbers*

| Name        | DataType           |
|-------------|--------------------|
| Name        | text (primary key) |
| Number      | float64            |
| create_at   | Date               |
| updated_at  | Date               |
| deleted_at  | Date               |
### *Accounts*

| Name          | DataType           |
|---------------|--------------------|
| Id            | UUID (primary key) |
| email         | text               |
| activated     | bool               |
| hash_password | text               |
| create_at     | Date               |
| updated_at    | Date               |
| deleted_at    | Date               |

## Run on native machine

---

```
./build.sh native
```
## Run for load balancer with docker

---
Go to *load_balancer* subdirectory and run docker compose.[](https://) *N* is http api service number for load balancing
docker compose will up and run multiple services including: 
- Two big_num_compute_service 
- Postgresql container
- Traefik service for load balance.

```
./build.sh lb
```

## Run big_num_compute_service

---
```
$ big_num_compute_service -c config.yml
```

## Number object jsonrpc API

---

Before operating number object related api. account token must be taken by login account.
### Create number object 

Request

```
{"jsonrpc":"1.0","method":"create","params":["grav_const", "0.000000000066731039356729","your@email.com", "account_token"],"id":1}
```

Response

```
{"jsonrpc":"1.0","result":"success","error":null,"id":1}
```

### Update number object
Request

```
{"jsonrpc":"1.0","method":"update","params":["grav_const", "0.000000000066731039356729","your@email.com", "account_token"],"id":1}
```

Response

```
{"jsonrpc":"1.0","result":"success","error":null,"id":1}
```
### Delete number object

Request

```
{"jsonrpc":"1.0","method":"delete","params":["grav_const", "your@email.com", "account_token"],"id":1}
```

Response

```
{"jsonrpc":"1.0","result":"success","error":null,"id":1}
```

### Number object addition

Request

```
{"jsonrpc":"1.0","method":"add","params":["grav_const", "1000.123", "your@email.com", "account_token"],"id":1}
```
or
```
{"jsonrpc":"1.0","method":"add","params":["grav_const", "mars_const", "your@email.com", "account_token"],"id":1}
```

Response

```
{"jsonrpc":"1.0","result":"your_add_result","error":null,"id":1}
```

### Number object subtraction

Request

```
{"jsonrpc":"1.0","method":"subtract","params":["grav_const", "1000.123", "your@email.com", "account_token"],"id":1}
```
or
```
{"jsonrpc":"1.0","method":"subtract","params":["grav_const", "mars_const", "your@email.com", "account_token"],"id":1}
```

Response

```
{"jsonrpc":"1.0","result":"your_subtract_result","error":null,"id":1}
```

### Number object multiplication

Request

```
{"jsonrpc":"1.0","method":"multiply","params":["grav_const", "1000.123", "your@email.com", "account_token"],"id":1}
```
or
```
{"jsonrpc":"1.0","method":"multiply","params":["grav_const", "mars_const", "your@email.com", "account_token"],"id":1}
```

Response

```
{"jsonrpc":"1.0","result":"your_multiply_result","error":null,"id":1}
```

### Number object division

Request

```
{"jsonrpc":"1.0","method":"divide","params":["grav_const", "1000.123", "your@email.com", "account_token"],"id":1}
```
or
```
{"jsonrpc":"1.0","method":"divide","params":["grav_const", "mars_const", "your@email.com", "account_token"],"id":1}
```

Response

```
{"jsonrpc":"1.0","result":"your_divide_result","error":null,"id":1}
```



## Account jsonrpc API

---

### Create account  
*passcode mail will send to your mail box to validate your email and activate your account.*

Request

```
{"jsonrpc":"1.0","method":"createaccount","params":["your@email.com", "your_password"],"id":1}
```

Response

```
{"jsonrpc":"1.0","result":"your_passcode","error":null,"id":1}
```

### Validate account email
Request

```
{"jsonrpc":"1.0","method":"","params":["your@email.com", "your_passcode"],"id":1}
```

Response

```
{"jsonrpc":"1.0","result":"success","error":null,"id":1}
```

### Login account and get token
Request

```
{"jsonrpc":"1.0","method":"","params":["your@email.com", "password"],"id":1}
```

Response

```
{"jsonrpc":"1.0","result":"your_account_token","error":null,"id":1}
```

### Delete account
Request

```
{"jsonrpc":"1.0","method":","params":["your@email.com", "account_token"],"id":1}
```

Response

```
{"jsonrpc":"1.0","result":"success","error":null,"id":1}
```

