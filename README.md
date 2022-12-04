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
id: 01GKDKKTHS0TE0SQXZ3E93W6QN, name: book
id: 01GKDKKP6HH4794GVE1EPHKJ8R, name: stationary

$ ./sqlboilerpractice deleteCategory --id 01GKDKKTHS0TE0SQXZ3E93W6QN

$ ./sqlboilerpractice listCategory
id: 01GKDKKP6HH4794GVE1EPHKJ8R, name: stationary

# control products
$ ./sqlboilerpractice createProduct --name pencil --price 100 --categoryid 01GKDKKP6HH4794GVE1EPHKJ8R
$ ./sqlboilerpractice createProduct --name eraser --price 70 --categoryid 01GKDKKP6HH4794GVE1EPHKJ8R
$ ./sqlboilerpractice createProduct --name stapler --price 300 --categoryid 01GKDKKP6HH4794GVE1EPHKJ8R

$ ./sqlboilerpractice listProduct
id: 01GKDM2D8WE9RDKB17NAJWR3H9, name: pencil, price: 100, category_id: 01GKDKKP6HH4794GVE1EPHKJ8R
id: 01GKDM2J5NKKNK1XSTPPCAXKBG, name: eraser, price: 70, category_id: 01GKDKKP6HH4794GVE1EPHKJ8R
id: 01GKDM2QJCEPQ50B40Z312NTVG, name: stapler, price: 300, category_id: 01GKDKKP6HH4794GVE1EPHKJ8R

$ ./sqlboilerpractice deleteProduct --id 01GKDM2QJCEPQ50B40Z312NTVG

$ ./sqlboilerpractice listProduct
id: 01GKDM2D8WE9RDKB17NAJWR3H9, name: pencil, price: 100, category_id: 01GKDKKP6HH4794GVE1EPHKJ8R
id: 01GKDM2J5NKKNK1XSTPPCAXKBG, name: eraser, price: 70, category_id: 01GKDKKP6HH4794GVE1EPHKJ8R

$ ./sqlboilerpractice countProduct --categoryid 01GKDKKP6HH4794GVE1EPHKJ8R
category_id: 01GKDKKP6HH4794GVE1EPHKJ8R, category_name: stationary, num_products: 2
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
