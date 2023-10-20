package models

import "gorm.io/gorm"

type Project struct {
	Model
	Name string `json:"name"`
}

// ExistProjectByID checks if an project exists based on ID
func ExistProjectByID(id int) (bool, error) {
	var project Project
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&project).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if project.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetProjectTotal gets the total number of projects based on the constraints
func GetProjectTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Project{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetProjects gets a list of projects based on paging constraints
func GetProjects(pageNum int, pageSize int, maps interface{}) ([]*Project, error) {
	var projects []*Project
	maps = getProjectMaps()

	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&projects).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return projects, nil
}

// GetProject Get a single project based on ID
func GetProject(id int) (*Project, error) {
	var project Project
	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&project).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// err = db.Model(&project).Related(&project.Tag).Error
	err = db.Model(&project).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &project, nil
}

// EditProject modify a single project
func EditProject(id int, data interface{}) error {
	if err := db.Model(&Project{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// AddProject add a single project
func AddProject(project *Project) error {
	// project := Project{
	// 	Name: data["name"].(string),
	// }
	if err := db.Create(&project).Error; err != nil {
		return err
	}

	return nil
}

// DeleteProject delete a single project
func DeleteProject(id int) error {
	if err := db.Where("id = ?", id).Delete(Project{}).Error; err != nil {
		return err
	}

	return nil
}

// CleanAllProject clear all project
func CleanAllProject() error {
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Project{}).Error; err != nil {
		return err
	}

	return nil
}

func getProjectMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0

	return maps
}
