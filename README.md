# sqlboiler-practice
sample app made by SQLBoiler

# for developer

## add a new subcommand

```bash
$ cobra-cli add ${subcommand}
```

See https://github.com/spf13/cobra-cli/blob/main/README.md for details

## update ORM schema

```bash
# run DB
$ docker-compose up -d
# generate SQL Boiler models from DB
$ go generate
```
