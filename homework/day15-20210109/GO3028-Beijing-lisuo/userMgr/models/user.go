package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// UserQuit represent quit exit status code
const (
	AdminName string = "admin"
	AdminPass string = "admin123"
	PassCost         = 8
	AdminID   int64  = 1
	UserQuit  int    = 1
)

var (
	dsn = beego.AppConfig.String("db::dsn")
)

// Role stuff
// User Password : 1 -- 1
// User Role     : 1 -- n
// Role URL      : m -- n

// Role is role
type Role struct {
	ID   int64
	Name string

	Users []*User `orm:"reverse(many)"`
	Urls  []*Url  `orm:"rel(m2m)"`
}

// User is user
type User struct {
	ID      int64  `orm:"column(id);pk;auto"`
	Name    string `orm:"size(100)"`
	Role    *Role  `orm:"rel(fk)"`
	Sex     int
	Address string     `orm:"size(512)"`
	Cell    string     `orm:"size(512)"`
	Born    *time.Time `orm:"size(512)"`
	//Password  string     `orm:"size(512)"`
	Password  *Password  `orm:"reverse(one)"`
	CreatedAt *time.Time `orm:"auto_now_add"`
	UpdatedAt *time.Time `orm:"auto_now"`
	DeletedAt *time.Time `orm:"null;"`
}

// UserList contains users
type UserList []User

type Password struct {
	ID   int64
	Hash string

	User *User `orm:"rel(one)"`
}

type Url struct {
	ID   int64
	Name string
	Path string

	Roles []*Role `orm:"reverse(many)"`
}

// RegisterDB get db ready
func RegisterDB() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	if err := orm.RegisterDataBase("default", "mysql", dsn); err != nil {
		panic(err)
	}
	orm.RegisterModel(new(Role), new(User), new(Password), new(Url))
	orm.RunSyncdb("default", false, true)
}
