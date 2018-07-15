package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

// DBConfig DBに関わる設定
type DBConfig struct {
	// host ホスト
	Host string `json:"host"`
	// port ポート
	Port uint64 `json:"port"`
	// database データベース名
	Database string `json:"database"`
	// useraname ユーザ名
	Useraname string `json:"username"`
	// password パスワード
	Password string `json:"password"`
	// connectionTimeout コネクションタイムアウト
	ConnectionTimeout int64 `json:"connectionTimeout"`
}

// defaultConfig デフォルトの設定
func defaultConfig() DBConfig {
	return DBConfig{
		Host:              "localhost",
		Port:              5432,
		ConnectionTimeout: 300,
	}
}

// ToDatasourceParameter データソースのパラメータに変換する
func (c DBConfig) ToDatasourceParameter() string {

	p := toParameter("host", c.Host)
	p += toParameter("port", fmt.Sprint(c.Port))
	p += toParameter("dbname", c.Database)
	p += toParameter("user", c.Useraname)
	p += toParameter("password", c.Password)
	p += toParameter("connect_timeout", fmt.Sprint(c.ConnectionTimeout))
	p += toParameter("sslmode", "disable")

	return p
}

func toParameter(key string, value string) string {

	return key + "=" + value + " "
}

// LoadDBConfig DBの設定をロードする
// ファイル名は「db.json」
func LoadDBConfig(param CrudParameter) (DBConfig, error) {

	c := defaultConfig()

	bs, be := ioutil.ReadFile(param.BasePath + "db.json")

	if be != nil {

		log.Fatal(be)
		return DBConfig{}, be
	}

	if err := json.Unmarshal(bs, &c); err != nil {

		log.Fatal(err)
		return DBConfig{}, err
	}

	if ve := validate(c); ve != nil {

		log.Fatal(ve)
		return DBConfig{}, ve
	}

	return c, nil
}

// validate 設定値の検証を行う
func validate(c DBConfig) error {

	if len(c.Database) < 1 {

		return errors.New("databaseの設定がされていません。")
	}

	if len(c.Useraname) < 1 {

		return errors.New("usernameの設定がされていません。")
	}

	if len(c.Password) < 1 {

		return errors.New("passwordの設定がされていません。")
	}

	return nil
}
