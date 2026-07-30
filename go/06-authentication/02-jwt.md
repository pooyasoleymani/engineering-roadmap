---
Date: 2026-07-27
tags:
  - software_engineering
---
---
# JWT Authentication in Go

JSON Web Token (JWT) is a compact and secure method for transmitting information between parties as a JSON object. It is commonly used for authentication and authorization in web applications, allowing users to access protected resources without repeatedly providing credentials.


## Characteristics

- ***Stateless Authentication:*** Stores user information in a token, reducing the need for server-side session storage.
- ***Secure Data Exchange:*** Uses digital signatures to verify that the token has not been altered.
- ***Compact Format:*** Consists of three parts—Header, Payload, and Signature—encoded into a single string.
- ***Widely Supported:*** Works across different programming languages and platforms.


## What Problem Does JWT Solve?

Without authentication:

```text
POST /books
DELETE /books/1
```

Anyone can call your API.

We need to know:

```text
Who is the user?
Are they logged in?
```

---

## Traditional Session Authentication

Old web applications:

```text
Login
  ↓
Server creates session
  ↓
Session stored in DB/Redis
  ↓
Cookie sent to browser
```

Works well, but requires server-side session storage.

---

## JWT Authentication

JWT = JSON Web Token

Flow:

```text
Login
  ↓
Server verifies username/password
  ↓
Server generates JWT
  ↓
Client stores JWT
  ↓
Client sends JWT on every request
```

Example request:

```http
Authorization: Bearer eyJhbGciOi...
```

---

## JWT Structure

A JWT has three parts:

```text
HEADER.PAYLOAD.SIGNATURE
```

Example:

```text
eyJhbGciOiJIUzI1NiJ9
.
eyJ1c2VyX2lkIjoxfQ
.
abc123signature
```

Normally you never parse this manually.

---

## Install JWT Library

A popular library:

```bash
go get github.com/golang-jwt/jwt/v5
```

---

## Creating a Token

```go
func GenerateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": userID,
			"exp":     time.Now().Add(24 * time.Hour).Unix(),
		},
	)

	return token.SignedString(
		[]byte("super-secret-key"),
	)
}
```

---

## Login Handler

Request:

```json
{
  "username": "pooya",
  "password": "1234"
}
```

After verifying credentials:

```go
token, err := GenerateToken(user.ID)
```

Response:

```json
{
  "token": "eyJhbGc..."
}
```

---

## Validating Tokens

```go
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := strings.TrimPrefix(
			r.Header.Get("Authorization"),
			"Bearer ",
		)

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (any, error) {
				return []byte("super-secret-key"), nil
			},
		)

		if err != nil || !token.Valid {
			http.Error(
				w,
				"unauthorized",
				http.StatusUnauthorized,
			)
			return
		}

		next(w, r)
	}
}
```

---

## Extracting User ID

```go
claims := token.Claims.(jwt.MapClaims)

userID := int(
	claims["user_id"].(float64),
)
```

---

## Store User ID in Context

Instead of parsing the token again later:

```go
ctx := context.WithValue(
	r.Context(),
	UserIDKey,
	userID,
)

next(w, r.WithContext(ctx))
```

Handler:

```go
userID := r.Context().Value(UserIDKey)
```

Use a typed key:

```go
type contextKey string

const UserIDKey contextKey = "user_id"
```

---

## Token Expiration

Tokens should expire:

```go
"exp": time.Now().
	Add(24*time.Hour).
	Unix()
```

Expired tokens should return:

```text
401 Unauthorized
```

---

## Secret Management

Never hardcode secrets:

```go
[]byte("super-secret-key")
```

Instead:

```env
JWT_SECRET=my-secret
```

Load it through your configuration package.

---

## Password Storage

Never store plain-text passwords.

Incorrect:

```text
password = "1234"
```

Use bcrypt:

```go
hash, err := bcrypt.GenerateFromPassword(
	[]byte(password),
	bcrypt.DefaultCost,
)
```

Verify passwords:

```go
bcrypt.CompareHashAndPassword(
	[]byte(hash),
	[]byte(password),
)
```

---

## Typical Architecture

Login flow:

```text
POST /login
        ↓
UserService
        ↓
Verify Password
        ↓
Generate JWT
        ↓
Return Token
```

Protected request:

```text
JWT Middleware
        ↓
Validate Token
        ↓
Store UserID in Context
        ↓
Handler
```

---

## Common Beginner Mistakes

### Store plain passwords

Never do:

```sql
password = "1234"
```

---

### Create long-lived tokens

Avoid:

```text
exp = 365 days
```

Prefer hours or a few days depending on your security requirements.

---

### Store sensitive data inside JWT

Bad example:

```json
{
  "password": "1234"
}
```

JWT payloads are **not encrypted** by default. Anyone holding the token can read its payload.

---

### Skip signature verification

Always verify:

```go
token.Valid
```

before trusting any claims.

---

# Exercise

Implement the following endpoints:

```text
POST /register
POST /login
GET  /profile
```

### Register

- Hash the user's password using bcrypt.
- Save the user in the database.

### Login

- Verify the password.
- Generate and return a JWT.

### Profile

- Protect the route using JWT middleware.
- Read the authenticated user's ID from the request context.
- Return the user's profile information.