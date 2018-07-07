package externals

import (
	"database/sql"

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
func CreateContext() (Context, error) {

	db, der := createDBContext()

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

func createDBContext() (DBContext, error) {

	db, er := initDb()

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

func initDb() (*gorp.DbMap, error) {

	db, err := sql.Open("postgres", "user=crud-user dbname=go-crud password=crud-user sslmode=disable")

	if err != nil {

		return &gorp.DbMap{}, err
	}

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	return dbmap, nil
}
