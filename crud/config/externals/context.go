package externals

import (
	"database/sql"
	"log"
	"os"

	"github.com/gloryof/go-crud-practice/crud/config"
	"github.com/gloryof/go-crud-practice/crud/externals/gorp/tables"
	"github.com/go-gorp/gorp"

	// sql/driver内で依存しているためインポート
	_ "github.com/lib/pq"
)

// Context externalのコンテキスト
type Context struct {
	// DB DBに関するコンテキスト
	DB DBContext
}

// CreateContext コンテキストを生成する
func CreateContext(config config.DBConfig) (Context, error) {

	db, der := createDBContext(config)

	if der != nil {

		return Context{}, der
	}

	return Context{
		DB: db,
	}, nil
}

// Close クローズが必要な値に対してクローズ処理を行う
func (c Context) Close() {

	c.DB.close()
}

// DBContext DBに関するコンテキスト
type DBContext struct {

	// DBMap GoopのDBMap
	DBMap *gorp.DbMap
}

func createDBContext(c config.DBConfig) (DBContext, error) {

	db, er := initDb(c)

	if er != nil {

		return DBContext{}, er
	}

	return DBContext{
		DBMap: db,
	}, nil
}

func (c DBContext) close() {

	c.DBMap.Db.Close()
}

func initDb(c config.DBConfig) (*gorp.DbMap, error) {

	db, err := sql.Open("postgres", c.ToDatasourceParameter())

	if err != nil {

		log.Fatal(err)
		return &gorp.DbMap{}, err
	}

	pe := db.Ping()

	if pe != nil {

		log.Fatal(pe)
		return &gorp.DbMap{}, pe
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

	dbmap.AddTableWithName(tables.Users{}, "users")

	return dbmap, nil
}
