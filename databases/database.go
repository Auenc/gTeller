package databases

import (
	"database/sql"

	"github.com/astaxie/beego/session"
	"github.com/auenc/gTeller-core/config"
	"github.com/auenc/gTeller-core/discounts"
	"github.com/auenc/gTeller-core/email"
	"github.com/auenc/gTeller-core/items"
	"github.com/auenc/gTeller-core/logging"
	"github.com/auenc/gTeller-core/orders"
	"github.com/auenc/gTeller-core/requirements"
	"github.com/auenc/gTeller-core/shipping"
	"github.com/auenc/gTeller-core/statuses"
	"github.com/auenc/gTeller-core/users"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	ItemRepository        items.ItemRepository
	SessionRepository     *session.Manager
	UserRepository        users.UserRepository
	ShippingRepository    shipping.ShippingRepository
	StatusRepository      status.StatusRepository
	OrderRepository       orders.OrderRepository
	EmailRepository       email.EmailRepository
	DiscountRepository    discounts.DiscountRepository
	RequirementRepository requirements.Repository
	db                    *sql.DB
	logger                logging.Logger
}

func (db *Database) ReqRepo() requirements.Repository {
	return db.RequirementRepository
}

func (db *Database) ItemRepo() items.ItemRepository {
	return db.ItemRepository
}

func (db *Database) SessionRepo() *session.Manager {
	return db.SessionRepository
}

func (db *Database) UserRepo() users.UserRepository {
	return db.UserRepository
}

func (db *Database) ShippingRepo() shipping.ShippingRepository {
	return db.ShippingRepository
}

func (db *Database) StatusRepo() status.StatusRepository {
	return db.StatusRepository
}

func (db *Database) OrderRepo() orders.OrderRepository {
	return db.OrderRepository
}

func (db *Database) EmailRepo() email.EmailRepository {
	return db.EmailRepository
}

func (db *Database) DiscountRepo() discounts.DiscountRepository {
	return db.DiscountRepository
}

func (db *Database) DB() *sql.DB {
	return db.db
}

func (db *Database) Database(sqlURI string, sqlUser string, sqlPass string, dbName string) (*sql.DB, error) {
	if db.db == nil {
		db.Logger().LogLine("Creating connection")

		tmp, err := sql.Open("mysql", sqlUser+":"+sqlPass+"@unix("+sqlURI+")/"+dbName)
		if err != nil {
			return nil, err
		}
		db.db = tmp

		stmt, err := db.db.Prepare("SET SESSION wait_timeout = ?")
		if err != nil {
			return db.db, err
		}

		_, err = stmt.Exec(30)
		if err != nil {
			return db.db, err
		}
	}
	return db.db, nil
}

func (db *Database) Logger() logging.Logger {
	return db.logger
}

func New(config config.Config) (Database, error) {
	database := Database{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}

	database.logger = &logging.MemLogger{&config}
	err := database.logger.LoadLogFile()
	if err != nil {
		panic(err)
	}
	//Creating connection to sql database
	var dbName string
	if !config.Testing {
		dbName = "eposdb"
	} else {
		dbName = "eposdb_test"
	}
	db, err := database.Database(config.SqlSock, config.SqlUser, config.SqlPass, dbName)
	if err != nil {
		database.logger.LogLine(err)
		return database, err
	}
	database.RequirementRepository = &requirements.SQLRepository{db, database.logger}
	database.DiscountRepository = &discounts.SQLRepository{db}
	database.ItemRepository = &SQLItemRepository{database.RequirementRepository,
		db, database.logger}
	database.UserRepository = &users.SQLUserRepository{db, database.logger}
	database.EmailRepository = &email.SQLEmailRepository{db, database.logger}
	database.StatusRepository = &status.SQLStatusRepository{db, database.logger, database.EmailRepository}
	database.ShippingRepository = &shipping.SQLShippingRepository{db, database.logger}
	database.OrderRepository = &orders.SQLOrderRepository{db, database.logger, database.StatusRepository, database.ShippingRepository, database.DiscountRepository}
	seshRepo, _ := session.NewManager("memory", `{"cookieName":"gosessionid","gclifetime":3600}`)

	database.SessionRepository = seshRepo
	go seshRepo.GC()
	return database, nil
}
