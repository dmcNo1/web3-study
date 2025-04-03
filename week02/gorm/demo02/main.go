package main

import (
	"demo02/models"
	"demo02/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 初始化表
	router.GET("/init_tab", func(ctx *gin.Context) {
		db := models.DB

		db.AutoMigrate(&models.Permission{}, &models.Profile{}, &models.PermissionProfile{}, &models.Role{})

		permissions := []*models.Permission{}
		permissions = append(permissions, &models.Permission{Name: "rm -rf *"},
			&models.Permission{Name: "mv"}, &models.Permission{Name: "vi"})
		db.CreateInBatches(&permissions, 10)

		profile := models.Profile{PermissionId: int(permissions[0].ID)}
		db.Create(&profile)

		permissionProfiles := make([]*models.PermissionProfile, 0)
		permissionProfiles = append(permissionProfiles, &models.PermissionProfile{ProfileId: profile.ID, Name: "profile1"})
		permissionProfiles = append(permissionProfiles, &models.PermissionProfile{ProfileId: profile.ID, Name: "profile2"})
		db.CreateInBatches(&permissionProfiles, 10)

		roles := []*models.Role{}
		roles = append(roles, &models.Role{Name: "admin"}, &models.Role{Name: "bad guy"})
		db.CreateInBatches(&roles, 10)
	})

	routers.InitPermissionRouter(router)
	routers.InitProfileRouter(router)

	router.Run(":8081")
}
