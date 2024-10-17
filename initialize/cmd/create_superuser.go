package cmd

import (
	"fmt"
	"log"

	"github.com/PandaPy/pginer/template/initialize/db"
	"github.com/PandaPy/pginer/template/models"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var username, password string

var createSuperUserCmd = &cobra.Command{
	Use:   "create-superuser",
	Short: "创建超级管理员账户 (--username=<用户名> --password=<密码>)",
	Run: func(cmd *cobra.Command, args []string) {

		var count int64
		err := db.DB().Model(&models.UserModel{}).Where("is_superuser = ?", true).Count(&count).Error
		if err != nil {
			log.Fatalf("查询超级管理员失败: %v", err)
		}
		if count > 0 {
			fmt.Println("超级管理员已存在，无需重复创建。")
			return
		}

		superuser := models.UserModel{
			Nickname:    username,
			Username:    username,
			IsSuperuser: boolPtr(true),
			Status:      1,
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("加密密码失败: %v", err)
		}
		superuser.Password = strPtr(string(hashedPassword))

		err = db.DB().Create(&superuser).Error
		if err != nil {
			log.Fatalf("创建超级管理员失败: %v", err)
		}
		fmt.Println("超级管理员创建成功！")
	},
}

func init() {
	createSuperUserCmd.Flags().StringVarP(&username, "username", "u", "", "超级管理员用户名")
	createSuperUserCmd.Flags().StringVarP(&password, "password", "p", "", "超级管理员密码")
	createSuperUserCmd.MarkFlagRequired("username")
	createSuperUserCmd.MarkFlagRequired("password")
	rootCmd.AddCommand(createSuperUserCmd)
}

func strPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}

func intPtr(i int) *int {
	return &i
}
