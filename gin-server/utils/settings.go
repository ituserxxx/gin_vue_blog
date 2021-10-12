package utils

//
//type App struct {
//	TimeFormat       string `json:"time_format"`
//	JwtSecret        string `json:"jwt_secret"`
//	TokenTimeout     int64  `json:"token_timeout"`
//	StaticBasePath   string `json:"static_base_path"`
//	UploadBasePath   string `json:"upload_base_path"`
//	ImageRelPath     string `json:"image_rel_path"`
//	AvatarRelPath    string `json:"avatar_rel_path"`
//	BlogBaseUrl      string `json:"blog_base_url"`
//	BlogAdminBaseUrl string `json:"blog_admin_base_url"`
//}
//
//type Server struct {
//	RunMode      string        `json:"run_mode"`
//	ServerAddr   string        `json:"server_addr"`
//	ReadTimeout  time.Duration `json:"read_timeout"`
//	WriteTimeout time.Duration `json:"write_timeout"`
//}
//
//type DataBase struct {
//	Mode            string `json:"mode"`
//	Host            string `json:"host"`
//	Port            string `json:"port"`
//	User            string `json:"user"`
//	Password        string `json:"password"`
//	DBName          string `json:"db_name"`
//	Prefix          string `json:"prefix"`
//	MaxOpenConns    int    `json:"max_open_conns"`
//	MaxIdleConns    int    `json:"max_idle_conns"`
//	ConnMaxLifeTime int    `json:"conn_max_life_time"`
//}
//
//type Redis struct {
//	Host      string `json:"host"`
//	Port      string `json:"port"`
//	Password  string `json:"password"`
//	DB        int    `json:"db"`
//	CacheTime int    `json:"cache_time"`
//}
//
//var ServerInfo = &Server{}
//var DBInfo = &DataBase{}
//var AppInfo = &App{}
//var RedisInfo = &Redis{}
//
//func init() {
//	viper.AddConfigPath("config")
//	viper.SetConfigName("config")
//	viper.SetConfigType("yml")
//	if err := viper.ReadInConfig(); err != nil {
//		panic(err)
//	}
//
//	AppInfo.TimeFormat = viper.GetString("app.timeFormat")
//	AppInfo.JwtSecret = viper.GetString("app.jwtSecret")
//	AppInfo.TokenTimeout = viper.GetInt64("app.tokenTimeout")
//	AppInfo.StaticBasePath = viper.GetString("app.staticBasePath")
//	AppInfo.UploadBasePath = viper.GetString("app.uploadBasePath")
//	AppInfo.ImageRelPath = viper.GetString("app.imageRelPath")
//	AppInfo.AvatarRelPath = viper.GetString("app.avatarRelPath")
//	AppInfo.BlogBaseUrl = viper.GetString("app.BlogBaseUrl")
//	AppInfo.BlogAdminBaseUrl = viper.GetString("app.BlogAdminBaseUrl")
//
//	ServerInfo.RunMode = viper.GetString("server.runMode")
//	ServerInfo.ServerAddr = viper.GetString("server.serverAddr")
//	ServerInfo.ReadTimeout = time.Duration(viper.GetInt("server.readTimeout")) * time.Second
//	ServerInfo.WriteTimeout = time.Duration(viper.GetInt("server.writeTimeout")) * time.Second
//
//	DBInfo.Mode = viper.GetString("database.mode")
//	DBInfo.Host = viper.GetString("database.host")
//	DBInfo.Port = viper.GetString("database.port")
//	DBInfo.User = viper.GetString("database.user")
//	DBInfo.Password = viper.GetString("database.password")
//	DBInfo.DBName = viper.GetString("database.dbName")
//	DBInfo.Prefix = viper.GetString("database.prefix")
//	DBInfo.MaxOpenConns = viper.GetInt("database.max_open_conns")
//	DBInfo.MaxIdleConns = viper.GetInt("database.max_idle_conns")
//	DBInfo.ConnMaxLifeTime = viper.GetInt("database.conn_max_lifetime")
//
//	RedisInfo.Host = viper.GetString("redis.host")
//	RedisInfo.Port = viper.GetString("redis.port")
//	RedisInfo.Password = viper.GetString("redis.password")
//	RedisInfo.DB = viper.GetInt("redis.db")
//	RedisInfo.CacheTime = viper.GetInt("redis.cacheTime")
//}
