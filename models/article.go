package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	Model
	Name      string  `json:"name"`
	Protocol  string  `json:"protocol"`
	ProjectID int     `json:"project_id"`
	Project   Project `json:"project"`
}

// ExistArticleByID checks if an article exists based on ID
func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetArticleTotal gets the total number of articles based on the constraints
func GetArticleTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Article{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetArticles gets a list of articles based on paging constraints
func GetArticles(pageNum int, pageSize int, projectId string) ([]*Article, error) {
	var articles []*Article
	maps := getArticleMaps(projectId)

	err := db.Preload("Project").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articles, nil
}

// GetArticle Get a single article based on ID
func GetArticle(projectId int, id int) (*Article, error) {
	var article Article
	err := db.Where("project_id = ? AND id = ? AND deleted_on = ?", projectId, id, 0).First(&article).Error
	if err != nil {
		return nil, err
	}

	err = db.Model(&article).Related(&article.Project).Error
	if err != nil {
		return nil, err
	}

	return &article, nil
}

// EditArticle modify a single article
func EditArticle(id int, data interface{}) error {
	if err := db.Model(&Article{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// AddArticle add a single article
func AddArticle(article *Article) error {
	// article := Article{
	// 	Name:      data["name"].(string),
	// 	Protocol:  data["protocol"].(string),
	// 	ProjectID: data["project_id"].(int),
	// }
	if err := db.Create(&article).Error; err != nil {
		return err
	}

	return nil
}

// DeleteArticle delete a single article
func DeleteArticle(id int) error {
	if err := db.Where("id = ?", id).Delete(Article{}).Error; err != nil {
		return err
	}

	return nil
}

// CleanAllArticle clear all article
func CleanAllArticle() error {
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{}).Error; err != nil {
		return err
	}

	return nil
}

func getArticleMaps(projectId string) map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	maps["project_id"] = projectId

	return maps
}
