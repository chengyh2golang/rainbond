
// RAINBOND, Application Management Platform
// Copyright (C) 2014-2017 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package dao

import (
	"github.com/goodrain/rainbond/pkg/db/model"

	"github.com/jinzhu/gorm"
	"github.com/Sirupsen/logrus"
	"encoding/json"
)


//AddModel AddModel
func (c *AppPublishDaoImpl) AddModel(mo model.Interface) error {
	result := mo.(*model.AppPublish)
	result.Status="success"
	r,_:=json.Marshal(result)
	logrus.Infof("creating new app publish recode ,details %s",string(r))
	var oldResult model.AppPublish
	if ok := c.DB.Where("service_key=? and app_version=?", result.ServiceKey,result.AppVersion).Find(&oldResult).RecordNotFound(); ok {
		if err := c.DB.Create(result).Error; err != nil {
			return err
		}
	}
	return nil
}

//UpdateModel UpdateModel
func (c *AppPublishDaoImpl) UpdateModel(mo model.Interface) error {
	return nil
}
//EventLogMessageDaoImpl EventLogMessageDaoImpl
type AppPublishDaoImpl struct {
	DB *gorm.DB
}

//GetEventLogMessages get event log message
func (c *AppPublishDaoImpl) GetAppPublish(serviceKey,appVersion string) (*model.AppPublish, error) {
	var result model.AppPublish
	if err := c.DB.Where("service_key = ? and app_version =?", serviceKey,appVersion).Find(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			//return messageRaw, nil
			return &model.AppPublish{
				Status:"failure",
			},nil
		}
		return nil, err
	}
	return &result, nil
}
