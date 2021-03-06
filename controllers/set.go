package controllers

import (
	"fmt"
	"strconv"
	"github.com/shwinpiocess/cc/models"
)

type SetController struct {
	BaseController
}

func (this *SetController) GetAllSetInfo() {
	out := make(map[string]interface{})
	
	if applicationID, err := this.GetInt("ApplicationID"); err == nil {
		var fields []string
		var sortby []string
		var order []string
		var query map[string]string = make(map[string]string)
		var limit int64 = 0
		var offset int64 = 0
	
		query["owner"] = strconv.Itoa(this.userId)
		query["application_id"] = strconv.Itoa(applicationID)
	
		if s, err := models.GetAllSet(query, fields, sortby, order, offset, limit); err == nil {
			out["success"] = true
			out["set"] = s
			this.jsonResult(out)
		} else {
			out["success"] = false
			out["errInfo"] = "获取当前业务的集群列表失败！"
			this.jsonResult(out)
		}
	} else {
		out["success"] = false
		out["errInfo"] = "请求参数错误！"
		this.jsonResult(out)
	}
}

func (this *SetController) NewSet() {
	out := make(map[string]interface{})
	
	s := new(models.Set)
	s.ApplicationID, _ = this.GetInt("ApplicationID")
	s.SetName = this.GetString("SetName")
	s.EnviType, _ = this.GetInt("EnviType")
	s.ServiceStatus, _ = this.GetInt("ServiceStatus")
	s.ChnName = this.GetString("ChnName")
	s.Capacity, _ = this.GetInt("Capacity")
	s.Description = this.GetString("Des")
	s.OpenStatus = this.GetString("Openstatus")
	s.Owner = this.userId
	
	if _, err := models.AddSet(s); err == nil {
		out["success"] = true
		this.jsonResult(out)
	} else {
		out["success"] = false
		out["errInfo"] = "重复的集群名！"
		this.jsonResult(out)
	}
}

func (this *SetController) GetSetInfoById() {
	out := make(map[string]interface{})
	
	if setID, err := this.GetInt("SetID"); err == nil {
		if s, err := models.GetSetById(setID); err == nil {
			out["success"] = true
			out["set"] = s
			this.jsonResult(out)
		} else {
			out["success"] = false
			out["errInfo"] = "集群不存在！"
			this.jsonResult(out)
		}
	} else {
		out["success"] = false
		out["errInfo"] = "请求参数不合法！"
		this.jsonResult(out)
	}
}

func (this *SetController) EditSet() {
	out := make(map[string]interface{})
	
	s := new(models.Set)
	s.ApplicationID, _ = this.GetInt("ApplicationID")
	s.SetName = this.GetString("SetName")
	s.EnviType, _ = this.GetInt("EnviType")
	s.ServiceStatus, _ = this.GetInt("ServiceStatus")
	s.ChnName = this.GetString("ChnName")
	s.Capacity, _ = this.GetInt("Capacity")
	s.Description = this.GetString("Des")
	s.OpenStatus = this.GetString("Openstatus")
	s.SetID, _ = this.GetInt("SetID")
	
	if err := models.UpdateSetById(s); err == nil {
		out["success"] = true
		this.jsonResult(out)
	} else {
		fmt.Println("err=", err)
		out["success"] = false
		out["errInfo"] = err
		this.jsonResult(out)
	}
}

func (this *SetController) DelSet() {
	out := make(map[string]interface{})

	if setID, err := this.GetInt("SetID"); err == nil {
		if err := models.DeleteSet(setID); err == nil {
			out["success"] = true
			this.jsonResult(out)
		} else {
			out["success"] = false
			out["errInfo"] = err
			this.jsonResult(out)
		}
	} else {
		out["success"] = false
		out["errInfo"] = "请求参数不合法！"
		this.jsonResult(out)
	}
}