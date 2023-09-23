package config

type MongoOwner struct {
	Host       string `env:"MONGO_OWNER_HOST" envDefault:"localhost"`
	Port       string `env:"MONGO_OWNER_PORT" envDefault:"27017"`
	Username   string `env:"MONGO_OWNER_USERNAME" envDefault:""`
	Password   string `env:"MONGO_OWNER_PASSWORD" envDefault:""`
	Database   string `env:"MONGO_OWNER_DATABASE" envDefault:"empty"`
	Collection string `env:"MONGO_OWNER_COLLECTION" envDefault:"empties"`
	Query      string `env:"MONGO_OWNER_QUERY" envDefault:""`
}
type MongoAccount struct {
	Host       string `env:"MONGO_ACCOUNT_HOST" envDefault:"localhost"`
	Port       string `env:"MONGO_ACCOUNT_PORT" envDefault:"27017"`
	Username   string `env:"MONGO_ACCOUNT_USERNAME" envDefault:""`
	Password   string `env:"MONGO_ACCOUNT_PASSWORD" envDefault:""`
	Database   string `env:"MONGO_ACCOUNT_DATABASE" envDefault:"empty"`
	Collection string `env:"MONGO_ACCOUNT_COLLECTION" envDefault:"empties"`
	Query      string `env:"MONGO_ACCOUNT_QUERY" envDefault:""`
}

type I18n struct {
	Fallback string   `env:"I18N_FALLBACK_LANGUAGE" envDefault:"en"`
	Dir      string   `env:"I18N_DIR" envDefault:"./src/locales"`
	Locales  []string `env:"I18N_LOCALES" envDefault:"en,tr"`
}

type Server struct {
	Host  string `env:"SERVER_HOST" envDefault:"localhost"`
	Port  int    `env:"SERVER_PORT" envDefault:"3000"`
	Group string `env:"SERVER_GROUP" envDefault:"verify"`
}

type Redis struct {
	Host string `env:"REDIS_HOST"`
	Port string `env:"REDIS_PORT"`
	Pw   string `env:"REDIS_PASSWORD"`
	Db   int    `env:"REDIS_DB"`
}

type Cors struct {
	AllowedOrigins   string `env:"CORS_ALLOWED_ORIGINS" envDefault:"*"`
	AllowedMethods   string `env:"CORS_ALLOWED_METHODS" envDefault:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowedHeaders   string `env:"CORS_ALLOWED_HEADERS" envDefault:"*"`
	AllowCredentials bool   `env:"CORS_ALLOW_CREDENTIALS" envDefault:"true"`
}

type Topics struct {
	Owner   OwnerTopics
	Account AccountEvents
}

type AccountEvents struct {
	Deleted  string `env:"STREAMING_TOPIC_ACCOUNT_DELETED"`
	Created  string `env:"STREAMING_TOPIC_ACCOUNT_CREATED"`
	Updated  string `env:"STREAMING_TOPIC_ACCOUNT_UPDATED"`
	Disabled string `env:"STREAMING_TOPIC_ACCOUNT_DISABLED"`
	Enabled  string `env:"STREAMING_TOPIC_ACCOUNT_ENABLED"`
}

type TokenSrv struct {
	Expiration int    `env:"TOKEN_EXPIRATION" envDefault:"3600"`
	Project    string `env:"TOKEN_PROJECT" envDefault:"empty"`
}

type Session struct {
	Topic string `env:"SESSION_TOPIC"`
}

type HttpHeaders struct {
	AllowedOrigins   string `env:"CORS_ALLOWED_ORIGINS" envDefault:"*"`
	AllowedMethods   string `env:"CORS_ALLOWED_METHODS" envDefault:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowedHeaders   string `env:"CORS_ALLOWED_HEADERS" envDefault:"*"`
	AllowCredentials bool   `env:"CORS_ALLOW_CREDENTIALS" envDefault:"true"`
	Domain           string `env:"HTTP_HEADER_DOMAIN" envDefault:"*"`
}

type OwnerTopics struct {
	Created               string `env:"STREAMING_TOPIC_OWNER_CREATED"`
	UserRemoved           string `env:"STREAMING_TOPIC_OWNER_USER_REMOVED"`
	UserAdded             string `env:"STREAMING_TOPIC_OWNER_USER_ADDED"`
	UserPermissionRemoved string `env:"STREAMING_TOPIC_OWNER_USER_PERMISSION_REMOVED"`
	UserPermissionAdded   string `env:"STREAMING_TOPIC_OWNER_USER_PERMISSION_ADDED"`
	VerifiedByAdmin       string `env:"STREAMING_TOPIC_OWNER_VERIFIED_BY_ADMIN"`
	DeletedByAdmin        string `env:"STREAMING_TOPIC_OWNER_DELETED_BY_ADMIN"`
	RecoverByAdmin        string `env:"STREAMING_TOPIC_OWNER_RECOVER_BY_ADMIN"`
	Disabled              string `env:"STREAMING_TOPIC_OWNER_DISABLED"`
	Enabled               string `env:"STREAMING_TOPIC_OWNER_ENABLED"`
}

type Nats struct {
	Url     string   `env:"NATS_URL" envDefault:"nats://localhost:4222"`
	Streams []string `env:"NATS_STREAMS" envDefault:""`
}

type RSA struct {
	PrivateKeyFile string `env:"RSA_PRIVATE_KEY"`
	PublicKeyFile  string `env:"RSA_PUBLIC_KEY"`
}

type Vkn struct {
	Username string `env:"VKN_USERNAME"`
	Password string `env:"VKN_PASSWORD"`
}

type App struct {
	Protocol string `env:"PROTOCOL" envDefault:"http"`
	DB       struct {
		MongoOwner   MongoOwner
		MongoAccount MongoAccount
	}
	Vkn         Vkn
	HttpHeaders HttpHeaders
	Server      Server
	Session     Session
	I18n        I18n
	Topics      Topics
	Nats        Nats
	Redis       Redis
	TokenSrv    TokenSrv
	Rsa         RSA
}
