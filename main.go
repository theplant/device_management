package main

import (
	"bitbucket.org/sunfmin/hunthub/profile/db"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/qor/qor"
	"github.com/qor/qor/admin"
	"html/template"
	"log"
	"net/http"
)

func main() {
	adm := admin.New(&qor.Config{DB: &db.DB})

	pgr := adm.AddResource(&db.ProfileGithubRepository{}, &admin.Config{
		Name: "Github Repositories", Menu: []string{"Profile"}, PageCount: 100,
	})
	pgr.IndexAttrs("Login", "Repository", "Skill", "RepositoryStargazersCount")

	pgr.Meta(&admin.Meta{
		Name:  "Repository",
		Label: "Repository",
		Valuer: func(resource interface{}, ctx *qor.Context) interface{} {
			repo := resource.(*db.ProfileGithubRepository).Repository
			return template.HTML(fmt.Sprintf(`<a href="http://github.com/%s" target="_blank">%s</a>`, repo, repo))
		},
	})

	pgr.Scope(&admin.Scope{
		Name:    "default",
		Label:   "Default",
		Default: true,
		Handle: func(db *gorm.DB, ctx *qor.Context) *gorm.DB {
			return db.Order("repository_stargazers_count desc")
		},
	})

	pgr.Scope(&admin.Scope{
		Name:    "hangzhou",
		Label:   "Hangzhou",
		Default: false,
		Handle: func(db *gorm.DB, ctx *qor.Context) *gorm.DB {
			return db.Order("repository_stargazers_count desc").Joins("INNER JOIN github_profiles ON github_profiles.login = profile_github_repositories.login").Where("github_profiles.location LIKE ?", "%hangzhou%")
		},
	})
	pgr.Scope(&admin.Scope{
		Name:    "dalian",
		Label:   "Dalian",
		Default: false,
		Handle: func(db *gorm.DB, ctx *qor.Context) *gorm.DB {
			return db.Order("repository_stargazers_count desc").Joins("INNER JOIN github_profiles ON github_profiles.login = profile_github_repositories.login").Where("github_profiles.location LIKE ?", "%dalian%")
		},
	})

	pgr.SearchAttrs("skill")

	adm.MountTo("/admin", http.DefaultServeMux)

	log.Println("Starting Server at 9000.")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
