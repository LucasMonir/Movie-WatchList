# Dependencies:

- Docker
    * Environment variables needed: 
        PSQL_USER
        PSQL_DB
        PSQL_PASS 
        PSQL_ADDRESS_PROD
        PSQL_ADDRESS_HML
        RABBIT_MQ_PROD

## How to run:
- On the root folder, run "docker-compose up --build -d", might take a while!

- It's important to be sure that the environmet variables are being recovered properly, check the terminal after running the compose command, if there is any messages about "Environment variable being null", then check them and if any changes are made to the environment variables, restart your computer

- It's a project that I made purely for fun, the code may not be the sharpest but the idea was to practice, I intend to improve this code during the next weeks... I'm open for hiring :)