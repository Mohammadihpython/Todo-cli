package memorystore

import "todocli/entity"

type Category struct {
	categories []entity.Category
}

func (c Category) DoseThisUserHaveThisCategoryID(userID int, categoryId int) bool {
	isFound := false
	for _, c := range c.categories {
		if c.ID == categoryId && c.UserID == userID {
			isFound = true
			break
		}

	}
	return isFound

}
