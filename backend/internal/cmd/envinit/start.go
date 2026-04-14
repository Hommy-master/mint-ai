package envinit

import (
	"context"
	"fmt"

	"cozeos/internal/config"
	"cozeos/internal/model"
	"cozeos/internal/pkg/db"

	"github.com/gogf/gf/v2/os/glog"
	"gorm.io/gorm"
)

func start(ctx context.Context, cfgPath string) error {
	if err := config.Init(cfgPath); err != nil {
		glog.Errorf(ctx, "config.Init failed: %v", err)
		return err
	}

	conn := db.NewDB()
	models := []interface{}{
		&model.User{},
		&model.Upload{},
		&model.Order{},
		&model.PluginBalanceLog{},
	}

	if err := conn.AutoMigrate(models...); err != nil {
		glog.Errorf(ctx, "auto migrate failed: %v", err)
		return err
	}

	// 同步删除已废弃字段，确保users表不再保留api_key列。
	if conn.Migrator().HasColumn(&model.User{}, "api_key") {
		if err := conn.Migrator().DropColumn(&model.User{}, "api_key"); err != nil {
			glog.Errorf(ctx, "drop users.api_key failed: %v", err)
			return err
		}
	}

	for _, m := range models {
		if err := setAutoIncrementStart(conn, m, model.AUTO_INCREMENT); err != nil {
			glog.Warningf(ctx, "set auto increment start failed, model: %T, err: %v", m, err)
		}
	}

	glog.Infof(ctx, "env init done, migrated %d models", len(models))
	return nil
}

func setAutoIncrementStart(conn *gorm.DB, m interface{}, start int) error {
	tableNamer, ok := m.(interface{ TableName() string })
	if !ok {
		return fmt.Errorf("model does not implement TableName: %T", m)
	}

	tableName := tableNamer.TableName()
	switch conn.Dialector.Name() {
	case "mysql":
		return conn.Exec(fmt.Sprintf("ALTER TABLE %s AUTO_INCREMENT = %d", tableName, start)).Error
	case "postgres":
		return conn.Exec(
			"SELECT setval(pg_get_serial_sequence(?, 'id'), ?, false)",
			tableName,
			start,
		).Error
	default:
		return nil
	}
}
