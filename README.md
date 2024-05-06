# OTP Service

OTP Service is a gRPC server designed to send Time-based One-Time Passwords (TOTPs) via SMS to a phone number.
It utilizes the Twilio API for sending SMS messages and features a RabbitMQ listener for asynchronous SMS 
requests.

### Features:
1. **TOTP Generation**: Generates TOTPs for secure authentication.
2. **SMS Sending**: Sends TOTPs to a phone number via SMS using the Twilio API.
3. **Asynchronous SMS Requests**: Listens to RabbitMQ for asynchronous SMS requests, allowing for efficient processing of SMS messages.


### API Definitions

### 1. GenerateOTP
The API generates a TOTP and dispatches the code to a mobile number via SMS using the Twilio API.

input
```yaml
  string requestId = 1;
  int32 countryCode = 2;
  string phoneNumber = 3;
```
output:
```yaml
  bool isSuccess = 1;
  OtpError error = 2;
```

### Requirements
1. Go version of at least 1.22
2. Twilio account for SMS sending.
3. RabbitMQ server for asynchronous SMS requests.

```shell
 go version
```

### Install Dependencies
download dependencies into vendor folder
```shell
make mod
```

install dependencies
```shell
make install
```

### Running Tests

running tests with code coverage

```shell
make test
```

### Running the app
```shell
make mod
make run
```

