package client

type DBInfo struct {
	From     string `json:"from"` // local remote
	Address  string `json:"address"`
	Port     uint   `json:"port"`
	Username string `json:"userName"`
	Password string `json:"password"`

	Timeout uint `json:"timeout"` // second
}

type CreateInfo struct {
	Name       string `json:"name"`
	Format     string `json:"format"`
	Version    string `json:"version"`
	Username   string `json:"userName"`
	Password   string `json:"password"`
	Permission string `json:"permission"`

	Timeout uint `json:"timeout"` // second
}

type DeleteInfo struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	Username   string `json:"userName"`
	Permission string `json:"permission"`

	ForceDelete bool `json:"forceDelete"`
	Timeout     uint `json:"timeout"` // second
}

type PasswordChangeInfo struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	Username   string `json:"userName"`
	Password   string `json:"password"`
	Permission string `json:"permission"`

	Timeout uint `json:"timeout"` // second
}

type AccessChangeInfo struct {
	Name          string `json:"name"`
	Version       string `json:"version"`
	Username      string `json:"userName"`
	OldPermission string `json:"oldPermission"`
	Permission    string `json:"permission"`

	Timeout uint `json:"timeout"` // second
}

var formatMap = map[string]string{
	"utf8":    "utf8_general_ci",
	"utf8mb4": "utf8mb4_general_ci",
	"gbk":     "gbk_chinese_ci",
	"big5":    "big5_chinese_ci",
}
