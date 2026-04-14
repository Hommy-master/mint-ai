package db

import (
	"cozeos/internal/config"
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB // 数据库连接实例
	once     sync.Once
)

// connectPostgreSQL 连接PostgreSQL数据库
func connectPostgreSQL() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(config.PostgreSQL), &gorm.Config{})
}

// connectMySQL 连接MySQL数据库
func connectMySQL() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(config.MySQL), &gorm.Config{})
}

// setupConnectionPool 设置数据库连接池
func setupConnectionPool(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying database connection: %v", err)
	}

	sqlDB.SetMaxIdleConns(1)  // 设置空闲连接池大小
	sqlDB.SetMaxOpenConns(10) // 设置最大连接数
	return nil
}

// NewDB 创建一个新的数据库连接（单例模式）
// 如果PostgreSQL的DSN不为空，优先连接PostgreSQL数据库，否则连接MySQL数据库
// 如果DSN不为空但连接失败，直接报错
func NewDB() *gorm.DB {
	once.Do(func() {
		var err error

		// 检查PostgreSQL DSN是否为空
		if config.PostgreSQL != "" {
			// PostgreSQL DSN不为空，尝试连接PostgreSQL
			instance, err = connectPostgreSQL()
			if err != nil {
				panic(fmt.Errorf("failed to connect to PostgreSQL: %v", err))
			}
			fmt.Println("Successfully connected to PostgreSQL database")
		} else {
			// PostgreSQL DSN为空，连接MySQL
			instance, err = connectMySQL()
			if err != nil {
				panic(fmt.Errorf("failed to connect to MySQL: %v", err))
			}
			fmt.Println("Successfully connected to MySQL database")
		}

		// 设置数据库连接池
		if err = setupConnectionPool(instance); err != nil {
			panic(err)
		}
	})

	return instance
}
