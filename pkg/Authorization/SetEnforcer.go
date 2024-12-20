package Authorization

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

// SetupRoutes : all the routes are defined here
func NewEnforcerByDB(db *gorm.DB) *casbin.Enforcer {

	// Initialize  Authorization adapter
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize Authorization adapter: %v", err))
	}

	// Load model configuration file and policy store adapter
	enforcer, err := casbin.NewEnforcer("pkg/Authorization/rbac_model.conf", adapter)
	if err != nil {
		panic(fmt.Sprintf("failed to create Authorization enforcer: %v", err))
	}
	return enforcer

}
