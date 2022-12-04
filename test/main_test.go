package test

import (
	"fmt"

	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/information_schema"
	"github.com/syuparn/sqlboilerpractice/config"
	"github.com/syuparn/sqlboilerpractice/di"
	simmodels "github.com/syuparn/sqlboilerpractice/simulator_models"
	"go.uber.org/dig"
)

type sqlSimulatorDB struct {
	DB            *memory.Database
	CategoryTable *memory.Table
	ProductTable  *memory.Table
}

func newSimulatorDB() *sqlSimulatorDB {
	db := memory.NewDatabase("practice")

	categoryTable := simmodels.CreateDummyCategoryTable(db)
	db.AddTable("category", categoryTable)
	productTable := simmodels.CreateDummyProductTable(db)
	db.AddTable("product", productTable)

	return &sqlSimulatorDB{
		DB:            db,
		CategoryTable: categoryTable,
		ProductTable:  productTable,
	}
}

type sqlSimulator struct {
	DB     *sqlSimulatorDB
	Server *server.Server
}

func newSimulator(port int) (*sqlSimulator, error) {
	db := newSimulatorDB()
	engine := sqle.NewDefault(
		sql.NewDatabaseProvider(
			db.DB,
			information_schema.NewInformationSchemaDatabase(),
		))
	engine.Analyzer.Catalog.MySQLDb.AddSuperUser("root", "localhost", "")

	config := server.Config{
		Protocol: "tcp",
		Address:  fmt.Sprintf("localhost:%d", port),
	}
	srv, err := server.NewDefaultServer(config, engine)
	if err != nil {
		return nil, err
	}

	return &sqlSimulator{
		DB:     db,
		Server: srv,
	}, nil
}

func (s *sqlSimulator) Run() {
	go func() {
		if err := s.Server.Start(); err != nil {
			panic(err)
		}
	}()
}

func (s *sqlSimulator) Stop() error {
	return s.Server.Close()
}

const port = 12345

func mockContainer() *dig.Container {
	c := di.NewContainer()
	c.Decorate(func() *config.Config {
		return &config.Config{
			DBPort: port,
		}
	})
	return c
}
