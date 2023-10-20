package models

import "github.com/jinzhu/gorm"

type Organization struct {
	Model
	Name      string  `json:"name"`
	ArticleID int     `json:"article_id"`
	Article   Article `json:"article"`
}

// ExistOrganizationByID checks if an organization exists based on ID
func ExistOrganizationByID(id int) (bool, error) {
	var organization Organization
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&organization).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if organization.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetOrganizationTotal gets the total number of organizations based on the constraints
func GetOrganizationTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Organization{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetOrganizations gets a list of organizations based on paging constraints
func GetOrganizations(pageNum int, pageSize int, maps interface{}) ([]*Organization, error) {
	var organizations []*Organization
	err := db.Preload("Article").Where(maps).Offset(pageNum).Limit(pageSize).Find(&organizations).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return organizations, nil
}

// GetOrganization Get a single organization based on ID
func GetOrganization(id int) (*Organization, error) {
	var organization Organization
	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&organization).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	err = db.Model(&organization).Related(&organization.Article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &organization, nil
}

// EditOrganization modify a single organization
func EditOrganization(id int, data interface{}) error {
	if err := db.Model(&Organization{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// AddOrganization add a single organization
func AddOrganization(data map[string]interface{}) error {
	organization := Organization{
		Name: data["name"].(string),
	}
	if err := db.Create(&organization).Error; err != nil {
		return err
	}

	return nil
}

// DeleteOrganization delete a single organization
func DeleteOrganization(id int) error {
	if err := db.Where("id = ?", id).Delete(Organization{}).Error; err != nil {
		return err
	}

	return nil
}

// CleanAllOrganization clear all organization
func CleanAllOrganization() error {
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Organization{}).Error; err != nil {
		return err
	}

	return nil
}
