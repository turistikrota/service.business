package config

type MongoBusiness struct {
	Host       string `env:"MONGO_BUSINESS_HOST" envDefault:"localhost"`
	Port       string `env:"MONGO_BUSINESS_PORT" envDefault:"27017"`
	Username   string `env:"MONGO_BUSINESS_USERNAME" envDefault:""`
	Password   string `env:"MONGO_BUSINESS_PASSWORD" envDefault:""`
	Database   string `env:"MONGO_BUSINESS_DATABASE" envDefault:"empty"`
	Collection string `env:"MONGO_BUSINESS_COLLECTION" envDefault:"empties"`
	Query      string `env:"MONGO_BUSINESS_QUERY" envDefault:""`
}
type MongoInvite struct {
	Collection string `env:"MONGO_INVITE_COLLECTION" envDefault:"invite"`
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

type Rpc struct {
	AccountHost    string `env:"RPC_ACCOUNT_HOST" envDefault:"localhost:3001"`
	AccountUsesSsl bool   `env:"RPC_ACCOUNT_USES_SSL" envDefault:"localhost:3001"`
	Port           int    `env:"GRPC_PORT" envDefault:"3001"`
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
	Business BusinessTopics
	Account  AccountEvents
	Notify   NotifyTopics
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

type NotifyTopics struct {
	SendSpecialEmail string `env:"STREAMING_TOPIC_NOTIFY_SEND_SPECIAL_EMAIL"`
}

type BusinessTopics struct {
	Created               string `env:"STREAMING_TOPIC_BUSINESS_CREATED"`
	UserRemoved           string `env:"STREAMING_TOPIC_BUSINESS_USER_REMOVED"`
	UserAdded             string `env:"STREAMING_TOPIC_BUSINESS_USER_ADDED"`
	UserPermissionRemoved string `env:"STREAMING_TOPIC_BUSINESS_USER_PERMISSION_REMOVED"`
	UserPermissionAdded   string `env:"STREAMING_TOPIC_BUSINESS_USER_PERMISSION_ADDED"`
	VerifiedByAdmin       string `env:"STREAMING_TOPIC_BUSINESS_VERIFIED_BY_ADMIN"`
	DeletedByAdmin        string `env:"STREAMING_TOPIC_BUSINESS_DELETED_BY_ADMIN"`
	RecoverByAdmin        string `env:"STREAMING_TOPIC_BUSINESS_RECOVER_BY_ADMIN"`
	RejectedByAdmin       string `env:"STREAMING_TOPIC_BUSINESS_REJECTED_BY_ADMIN"`
	Disabled              string `env:"STREAMING_TOPIC_BUSINESS_DISABLED"`
	Enabled               string `env:"STREAMING_TOPIC_BUSINESS_ENABLED"`

	InviteCreate string `env:"STREAMING_TOPIC_INVITE_CREATE"`
	InviteDelete string `env:"STREAMING_TOPIC_INVITE_DELETE"`
	InviteUse    string `env:"STREAMING_TOPIC_INVITE_USE"`
}

type Nats struct {
	Url     string   `env:"NATS_URL" envDefault:"nats://localhost:4222"`
	Streams []string `env:"NATS_STREAMS" envDefault:""`
}

type Cipher struct {
	Key string `env:"CIPHER_KEY"`
	IV  string `env:"CIPHER_IV"`
}

type RSA struct {
	PrivateKeyFile string `env:"RSA_PRIVATE_KEY"`
	PublicKeyFile  string `env:"RSA_PUBLIC_KEY"`
}

type Vkn struct {
	Username string `env:"VKN_USERNAME"`
	Password string `env:"VKN_PASSWORD"`
}

type Urls struct {
	InviteAccept string `env:"URL_INVITE_ACCEPT"`
}

type App struct {
	Protocol string `env:"PROTOCOL" envDefault:"http"`
	DB       struct {
		MongoBusiness MongoBusiness
		MongoInvite   MongoInvite
	}
	Cipher      Cipher
	Rpc         Rpc
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
	Urls        Urls
}
