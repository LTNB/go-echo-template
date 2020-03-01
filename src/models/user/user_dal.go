package user

import (
	"github.com/LTNB/go-dal/helper/sql"
	"github.com/LTNB/go-dal/postgres"
)

const (
	tableName = "account"
)

/*
 * sample demo using go-dal
 *  refer: github.com/LTNB/go-dal
 */

var accountHelper *AccountHelper

/*
 * model
 */
type Account struct {
	ID       string `json:"id" primary:"id" form:"id"`
	Email    string `json:"email" form:"email"`
	FullName string `json:"full_name" form:"full_name"`
	Role     string `json:"role" form:"role"`
	Active   bool   `json:"active" form:"active"`
	Password string `json:"password" form:"password"`
}

/*
 * init helper
 */
type AccountHelper struct {
	*sql.Helper
}

/*
 * get Account helper
 */
func GetAccountHelper() AccountHelper {
	return *accountHelper
}

/*
 * init account helper
 */
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

/*
 * find by email
 */
func (aHelper AccountHelper) FindByEmail(email string) (*Account, error) {
	conditions := make(map[string]interface{})
	conditions["email"] = email
	orderBy := make(map[string]string)
	orderBy["id"] = "ASC"
	limit := 1
	offset := 0
	result, err := aHelper.GetByConditions(conditions, orderBy, limit, offset, "")
	if err != nil || len(result) == 0 {
		return nil, err
	}
	return result[0].(*Account), nil
}

/*
 * check email is existed
 */
func (aHelper AccountHelper) EmailIsExisted(email string) (bool, error) {
	conditions := make(map[string]interface{})
	conditions["email"] = email
	result, err := aHelper.GetByConditions(conditions, nil, 0, -1, "")
	if err != nil {
		panic(err)
	}
	if len(result) > 0 {
		return true, nil
	}
	return false, nil

}

