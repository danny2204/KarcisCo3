package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"github.com/danny2204/KarcisCoAPI/middleware"
	"time"
)

type Blog struct {
	Id int `gorm:primary_key`
	BlogTitle string
	BlogDetail string `gorm:"type:text"`
	UploadDate time.Time
	CreatedAt time.Time
	UpdateAt time.Time
	DeletedAt *time.Time `sql:"index"`
	View int
	AuthorName string
	Category string
	ImagePath string
}

func GetAllBlog()([]Blog, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var blogList []Blog
	db.Order("view desc").Find(&blogList)
	return blogList, err
}

func GetAllBlogExcept(id int)([]Blog, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var blogList []Blog
	db.Where("id != ?", id).Order("view desc").Limit(5).Find(&blogList)
	return blogList, err
}

func CreateBlog(blogTitle string, blogDetail string, view int, imagePath string, authorName string, category string)(*Blog, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	blog := &Blog {
		BlogTitle: blogTitle,
		BlogDetail: blogDetail,
		UploadDate: time.Now(),
		View: view,
		ImagePath: imagePath,
		AuthorName: authorName,
		Category: category,
	}
	db.Save(blog)
	return blog, err
}

func UpdateBlog(id int, blogTitle string, blogDetail string, view int, imagePath string, authorName string, category string)(*Blog, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var blog Blog
	db.Model(&blog).Where("id = ?", id).Update(map[string]interface{}{"blog_title": blogTitle, "blog_detail": blogDetail, "upload_date": time.Now(), "view": view, "image_path": imagePath, "author_name": authorName, "category": category})
	db.Where("id = ?", id).Find(&blog)
	return &blog, nil
}

func RemoveBlog(id int)(*Blog, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	blog := Blog{Id: id}
	db.Where("id = ?", id).Find(&blog)
	return &blog, db.Delete(blog).Error
}

func GetBlogById(id int)(*Blog, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var blog Blog
	db.Where("id = ?", id).Find(&blog)
	return &blog, err
}
