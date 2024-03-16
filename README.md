# Snippetbox

Simple web application that used to paste and share snippets of text.

## Requirements

- Go (v1.20)
- MySQL database

## Installation

### Database

Create `snippetbox` database:

```sql
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

Create `snippets` table and add and index on `created` column:

```sql
CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);
```

Create a new user:

```sql
CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT ON snippetbox.* TO 'web'@'localhost';

-- Important: Make sure to swap 'pass' with a password of your own choosing.
ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';
```

Create `users` table:

```sql
CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created DATETIME NOT NULL
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);
```

### TLS Certificate

For production server, it is recommend using [Let's Encrypt](https://letsencrypt.org/)
to create TLS certificate. But for development purpose, we can generate a
self-signed TLS certificate with Go's standard library:

```sh
mkdir tls
cd tls
go run {your-go-path}/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```

### Web Server

Run the web server:

```sh
go run cmd/web/!(*_test).go
```

Testing:

```sh
go test -v ./cmd/web
```

## API Routes

| Method | Pattern           | Handler             | Action                            |
|--------|-------------------|---------------------|-----------------------------------|
| ANY    | `/`               | `home`              | Display the home page             |
| GET    | `/snippet/:id`    | `showSnippet`       | Display a specific snippet        |
| GET    | `/snippet/create` | `createSnippetForm` | Display the snippet creation form |
| POST   | `/snippet/create` | `createSnippet`     | Create a new snippet              |
| GET    | `/user/signup`    | `signupUserForm`    | Display the sign up form          |
| POST   | `/user/signup`    | `signupUser`        | Create a new user                 |
| GET    | `/user/login`     | `loginUserForm`     | Display the login form            |
| POST   | `/user/login`     | `loginUser`         | User login authentication         |
| POST   | `/user/logout`    | `logoutUser`        | User logout                       |
| GET    | `/static/`        | `http.FileServer`   | Serve a specific static file      |
