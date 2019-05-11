# mongo
[![CircleCI](https://circleci.com/gh/aoacademy/LetsGo.svg?style=svg)](https://circleci.com/gh/aoacademy/LetsGo)

## Introduction
Do you need multiple copies of your data? Let's do replication with mongo style.

## How to use?
To run this replication set, you only need to enter the following command.

```sh
docker-compose up -d
```

This command runs 3 mongo instance and then configure first of them as master with the configuration that is provided in `scripts/setup.sh`.
To add more nodes into replication set, you must add them in docker-copmpose.yml and the script.
Global configurations of replication set (like name, etc.) are accessible in env file.
