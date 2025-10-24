package config

type DBConfig struct {
	Main   Main
	Second Second
}

type Main struct {
	Path string `env:"DB_MAIN_PATH" envDefault:"./database/db.sqlite"`
}

type Second struct {
	Path string `env:"DB_SECOND_PATH" envDefault:"./database/db_second.sqlite"`
}
