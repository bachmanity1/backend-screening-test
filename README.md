# Terraform Labs

## Description
This code is solution for [Terraform Labs Backend Screening Test](https://github.com/terra-project/screening-test-backend).\
Postman documentation is available at [this link](https://documenter.getpostman.com/view/15324968/TzJoF1Z1).\
Postman collection is available at [this link](https://www.getpostman.com/collections/8230d42d3b62261dbb7c). 

## Run in Docker
```bash
    $ docker-compose up
```

## Run in Local Env
1. Make sure to reflect parameters for local instance of MariaDB(MySQL) in the `./conf/.env.dev.json` file.
2. ```bash
    $ make build
    $ bin/terra
   ```
