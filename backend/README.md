# Backend

All backend implementation directory.

## About

The main API goal is to handle database with millions of records efficiently.

## Features

- Users API
- Accounts API
- Security with simple token

## Usage

Generate access token using default admin user

    curl -XPOST \
        -H "Content-type: application/json" \
        -d '{ "username": "admin", "password": "123456" }' \
        'http://<backend-addres>/token'


Create another account:

    curl -XPOST \
        -H "Content-type: application/json" \
        -H "X-Token: <generated-access-token>" \
        -d '{ "username": "admin", "password": "123456" }' \
        'http://<backend-addres>/accounts'


Retrieve some users:

    curl \
        -H "Content-type: application/json" \
        -H "X-Token: <generated-access-token>" \
        'http://<backend-addres>/users'


Applying filter:

      curl \
          -H "Content-type: application/json" \
          -H "X-Token: <generated-access-token>" \
          'http://<backend-addres>/users?keyword=<keyword>'


Pagination:

      curl \
          -H "Content-type: application/json" \
          -H "X-Token: <generated-access-token>" \
          'http://<backend-addres>/users?since=<username>'


Putting all together. Querying users where username starts with "roger" since "rogeriolino" username:

      curl \
          -H "Content-type: application/json" \
          -H "X-Token: <generated-access-token>" \
          'http://<backend-addres>/users?keyword=roger&since=rogeriolino'
