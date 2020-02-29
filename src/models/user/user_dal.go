package user

import (
	"github.com/LTNB/go-dal/helper/sql"
	"github.com/LTNB/go-dal/postgres"
)

const (
	tableName = "account"
)

var accountHelper *AccountHelper

type Account struct {
	ID       string `json:"id" primary:"id" form:"id"`
	Email    string `json:"email" form:"email"`
	FullName string `json:"full_name" form:"full_name"`
	Role     string `json:"role" form:"role"`
	Active   bool   `json:"active" form:"active"`
	Password string `json:"password" form:"password"`
}

type AccountHelper struct {
	*sql.Helper
}

func GetAccountHelper() AccountHelper {
	return *accountHelper
}

func (aHelper *AccountHelper) Init() {
	helper := sql.Helper{
		TableName:      tableName,
		Bo:             Account{},
		DefaultTagName: "json",
		Handler:        postgres.Helper{},
	}
	aHelper.Helper = &helper
	accountHelper = aHelper
}

func (accountHelper AccountHelper) FindByEmail(email string) (*Account, error) {
	conditions := make(map[string]interface{})
	conditions["email"] = email
	orderBy := make(map[string]string)
	orderBy["id"] = "ASC"
	limit := 1
	offset := 0
	result, err := accountHelper.GetByConditions(conditions, orderBy, limit, offset, "")
	if err != nil || len(result) == 0 {
		return nil, err
	}
	return result[0].(*Account), nil
}

func (accountHelper AccountHelper) EmailIsExisted(email string) (bool, error) {
	conditions := make(map[string]interface{})
	conditions["email"] = email
	result, err := accountHelper.GetByConditions(conditions, nil, 0, -1, "")
	if err != nil {
		panic(err)
	}
	if len(result) > 0 {
		return true, nil
	}
	return false, nil

}

func (accountHelper AccountHelper) Login(email, password string) bool {
	conditions := make(map[string]interface{})
	conditions["email"] = email
	conditions["password"] = password
	conditions["active"] = true
	result, err := accountHelper.GetByConditions(conditions, nil, 0, -1, "")
	if err != nil {
		panic(err)
	}
	if len(result) > 0 {
		return true
	}
	return false
}
