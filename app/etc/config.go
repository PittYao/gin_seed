package etc

type Config struct {
	ServerConf
	Mysql     Mysql
	Statics   Statics
	Swagger   Swagger
	BasicAuth BasicAuth
	Zap       Zap
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
)
