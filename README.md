# mongo
[![Travis branch](https://img.shields.io/travis/com/I1820/mongo/master.svg?style=flat-square)](https://travis-ci.com/I1820/mongo)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/edd16ac391684c138a1903c064832ace)](https://www.codacy.com/app/i1820/mongo?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=I1820/mongo&amp;utm_campaign=Badge_Grade)

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
