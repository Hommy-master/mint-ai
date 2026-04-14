package consts

// 用户角色
type UserRole string

const (
	UserRoleNormal  UserRole = "normal"  // 普通用户
	UserRoleCreator UserRole = "creator" // 智能体创作者，能添加智能体，赚取收入
	UserRoleAdmin   UserRole = "admin"   // 管理员
)
