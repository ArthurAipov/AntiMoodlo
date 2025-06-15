package Models

type User struct {
	UserID       uint   `gorm:"primaryKey;column:userid" json:"userid"`
	UserLogin    string `gorm:"column:userlogin" json:"userlogin"`
	UserPassword string `gorm:"column:userpassword" json:"userpassword"`
	UserRole     uint   `gorm:"column:userrole" json:"userrole"`
}

func (User) TableName() string {
	return "users"
}

type UserRole struct {
	RoleID   uint   `gorm:"primaryKey;column:roleid" json:"roleid"`
	RoleName string `gorm:"column:rolename" json:"rolename"`
}

func (UserRole) TableName() string {
	return "userroles"
}
