# IBAN API

This api is responsible in validating IBAN numbers.

## <b>Algorithm </b>
1. Check that the total IBAN length is correct as per the country. If not, the IBAN is invalid.
2. Move the four initial characters to the end of the string.
3. Replace each letter in the string with two digits, thereby expanding the string, where A = 10, B = 11, ..., Z = 35.
4. Interpret the string as a decimal integer and compute the remainder of that number on division by 97

If the remainder is 1, the check digit test is passed and the IBAN might be valid. 

Reference:  https://en.wikipedia.org/wiki/International_Bank_Account_Number#Structure

## <b>Build</b>

### <b>Makefile</b>
Build binary for specific OS.
```
# Build for Windows
make build_windows

# Build for Linux
make build_linux

# Build for Darwin (macOS)
make build_darwin
```
Build binary based on local OS.

```
make build

# Could build and run binary
make build_and_run
```

Remove binaaries built
```
make clean
```
### <b>Docker</b>
Creating a container and running the API.
```
make docker_run
```
Stop and remove container.
```
make docker_rm
```

## <b>Tests</b>

Run the unit tests using make.

```
make test
```

## <b>Usage</b>

```
POST 127.0.0.1:3000/valid/iban
```
Header
```
Content-Type: application/json
```
Request Body

```
{
    "iban":"VG96VPVG00000L2345678901"
}
```

Success Response - 200
```
{
    "valid": false
}
```

Error Response - 400
```
{
    "type": "",
    "code": "",
    "message": ""
}
```