# sqlboiler-practice
sample app made by SQLBoiler

# usage

```bash
# prepare: run DB
$ docker-compose up -d

# control categories
$ ./sqlboilerpractice createCategory --name stationary
$ ./sqlboilerpractice createCategory --name book

$ ./sqlboilerpractice listCategory
id: 01GEXZYGXKGFKX87A9TYKE407A, name: book
id: 01GEXZXT694DD0TNTR7FBGQ0KE, name: stationary

$ ./sqlboilerpractice deleteCategory --id 01GEXZYGXKGFKX87A9TYKE407A

$ ./sqlboilerpractice listCategory
id: 01GEXZXT694DD0TNTR7FBGQ0KE, name: stationary

# control products
$ ./sqlboilerpractice createProduct --name pencil --categoryid 01GEXZXT694DD0TNTR7FBGQ0KE
$ ./sqlboilerpractice createProduct --name eraser --categoryid 01GEXZXT694DD0TNTR7FBGQ0KE
$ ./sqlboilerpractice createProduct --name stapler --categoryid 01GEXZXT694DD0TNTR7FBGQ0KE

$ ./sqlboilerpractice listProduct
id: 01GEY00VMPK9YSR68SP2JBPEE0, name: pencil, category_id: 01GEXZXT694DD0TNTR7FBGQ0KE
id: 01GEY0140PWB68FX4CQFZBVTFF, name: eraser, category_id: 01GEXZXT694DD0TNTR7FBGQ0KE
id: 01GEY01BCRA89AKXM6P4W14W1D, name: stapler, category_id: 01GEXZXT694DD0TNTR7FBGQ0KE

$ ./sqlboilerpractice deleteProduct --id 01GEY01BCRA89AKXM6P4W14W1D

$ ./sqlboilerpractice listProduct
id: 01GEY00VMPK9YSR68SP2JBPEE0, name: pencil, category_id: 01GEXZXT694DD0TNTR7FBGQ0KE
id: 01GEY0140PWB68FX4CQFZBVTFF, name: eraser, category_id: 01GEXZXT694DD0TNTR7FBGQ0KE

$ ./sqlboilerpractice countProduct --categoryid 01GEXZXT694DD0TNTR7FBGQ0KE
category_id: 01GEXZXT694DD0TNTR7FBGQ0KE, category_name: stationary, num_products: 2
```

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
