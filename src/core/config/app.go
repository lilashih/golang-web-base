package config

type AppConfig struct {
	Name    string `envDefault:"App"`
	Version string `envDefault:"1.0.0"`
	Mode    string `env:"APP_MODE" envDefault:"release"` // release, debug, test

	Port int `env:"APP_PORT" envDefault:"8000"`

	BasePath   string `env:"APP_BASE_PATH" envDefault:"./"`
	PublicPath string `env:"APP_PUBLIC_PATH" envDefault:"./public"` // 前端檔案都放在release/public目錄下,所以開發時env要特別設置此參數
}
