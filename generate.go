package main

// NOTE: templates cannot be written in sqlboiler.toml because it does not recognize $GOPATH
//go:generate sqlboiler mysql --templates ${GOPATH}/pkg/mod/github.com/volatiletech/sqlboiler/v4@v4.13.0/templates/main,$GOPATH/pkg/mod/github.com/volatiletech/sqlboiler/v4@v4.13.0/templates/test,templates
