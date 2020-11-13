package main

import (
	"log"
	"os"

	"github.com/casbin/casbin/util"
	"github.com/casbin/casbin/v2"
	defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	"github.com/cdennig/cloudshipper-authz/cmd/server"
	"github.com/cdennig/cloudshipper-authz/internal/config"
)

// Version set current code version
var Version = "1.0.0"

func main() {
	if err := run(); err != nil {
		log.Printf("Fatal Error: %s", err)
		os.Exit(1)
	}
}

func initialize() (*config.Config, error) {

	cfg, err := config.Load()
	if err != nil {
		log.Printf("failed to load application configuration: %s", err)
		return nil, err
	}

	return cfg, nil
}

func run() error {
	// initialize config and logging
	cfg, err := initialize()
	if err != nil {
		log.Printf("failed to initialize: %s", err)
		return err
	}
	// connect to the database
	// db, err := sqlx.Connect("mysql", cfg.DB.Dsn)
	// if err != nil {
	// 	logger.Errorf("failed to connect to database: %s", err)
	// 	return err
	// }

	enforcer, err := casbin.NewCachedEnforcer("casbin/model.conf", "casbin/policy.csv")

	if err != nil {
		log.Panicf("Failed to load Permissions definition/model: %s", err)
	}
	enforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddDomainMatchingFunc("g", util.KeyMatch)
	//enforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("g", util.KeyMatch)

	app := server.Routing(enforcer)
	return server.Start(cfg, app)
}
