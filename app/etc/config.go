package etc

type Config struct {
	ServerConf
	Mysql     Mysql
	Statics   Statics
	Swagger   Swagger
	BasicAuth BasicAuth
	Zap       Zap
	Redis     Redis
}

type (
	ServerConf struct {
		Name     string
		ListenOn string
		Mode     string
	}

	Mysql struct {
		Url string
	}

	Statics struct {
		Url string
	}

	Swagger struct {
		Title    string
		Desc     string
		Host     string
		Url      string
		BasePath string
		Schemes  []string
	}

	BasicAuth struct {
		UserName string
		Password string
	}

	Zap struct {
		Format     string
		Dir        string
		MaxSize    int
		MaxAge     int
		MaxBackups int
		Compress   bool
		Localtime  bool
		ShowLine   bool
	}

	Redis struct {
		/** redis 数据服务器ip和端口 */
		Addr string

		/** 指定连接的数据库 默认连数据库0 */
		DB int

		/** 连接密码 */
		Password string

		/** 最大重试次数 */
		MaxRetries int

		/** 最大空闲连接数 */
		MinIdleConns int

		/** 连接池大小 */
		PoolSize int
	}
)
