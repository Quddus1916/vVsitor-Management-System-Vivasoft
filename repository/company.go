package repository

import (
	"visitor-management-system/database"
	"visitor-management-system/model"
)

var db = database.GetDB()

func RegisterCompany(company *model.Company) error {
	err := db.Create(&company).Error
	return err
}

func GetAllSubscriber() (all_company []model.Company, err error) {
	err = db.Preload("User").Find(&all_company).Error
	return
}

func UpdateUser(user *model.User) error {
	err := db.Save(&user).Error
	return err
}

func ChangeSubscription(sub *model.Subscription) error {
	// err := db.Where("company_id = ?", sub.CompanyId).Delete(&sub).Error
	// err = db.Save(&sub).Error
	err := db.Model(&sub).Where("company_id = ?", sub.CompanyId).Update("subscription_type", sub.Subscription_type).Error
	return err
}

func CancelSubscription(sub *model.Subscription) error {
	err := db.Model(&sub).Where("company_id = ?", sub.CompanyId).Update("subscription_type", "free").Error
	return err
}

func GetPreviousSubscription(id int) (model.Subscription, error) {
	var subscription model.Subscription
	err := db.Where("company_id = ?", id).Find(&subscription).Error

	return subscription, err
}
