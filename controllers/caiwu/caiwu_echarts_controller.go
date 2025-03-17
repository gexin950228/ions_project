package caiwu

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/xuri/excelize/v2"
	"ions_project/models/caiwu"
	"ions_project/utils"
	"math"
	"strconv"
	"time"
)

type CaiwuEchartDataController struct {
	beego.Controller
}

func (c *CaiwuEchartDataController) Get() {
	o := orm.NewOrm()
	qs := o.QueryTable("sys_caiwu_data")
	pagePerNum := 8
	currentPage, err := c.GetInt("page")
	if err != nil { // 说明没有获取到当前页
		currentPage = 1
	}
	offsetNum := pagePerNum * (currentPage - 1)
	month := c.GetString("month")
	var count int64 = 0
	var caiwuDatas []models.CaiwuData
	if month != "" { // 有查询条件的
		count, _ = qs.Filter("caiwu_date", month).Count()
		qs.Filter("caiwu_date", month).Limit(pagePerNum).Offset(offsetNum).All(&caiwuDatas)
	} else {
		month = time.Now().Format("2006-01")
		count, _ = qs.Filter("caiwu_date", month).Count()
		qs.Filter("caiwu_date", month).Limit(pagePerNum).Offset(offsetNum).All(&caiwuDatas)
	}
	// 总页数
	countPage := int(math.Ceil(float64(count) / float64(pagePerNum)))
	prePage := 1
	if currentPage == 1 {
		prePage = currentPage
	} else if currentPage > 1 {
		prePage = currentPage - 1
	}
	nextPage := 1
	if currentPage < countPage {
		nextPage = currentPage + 1
	} else if currentPage >= countPage {
		nextPage = currentPage
	}
	page_map := utils.Paginator(currentPage, pagePerNum, count)
	c.Data["caiwu_datas"] = caiwuDatas
	c.Data["prePage"] = prePage
	c.Data["nextPage"] = nextPage
	c.Data["currentPage"] = currentPage
	c.Data["countPage"] = countPage
	c.Data["count"] = count
	c.Data["page_map"] = page_map
	c.Data["month"] = month

	c.TplName = "caiwu/echart_data_list.html"
}

func (c *CaiwuEchartDataController) ToImportExcel() {
	c.TplName = "caiwu/echart_data_import.html"
}

func (c *CaiwuEchartDataController) DoImportExcel() {
	f, h, err := c.GetFile("upload_file")
	var message_map map[string]interface{}
	var errDataArr []string
	defer func() {
		err := f.Close()
		if err != nil {
			logs.Error(fmt.Sprintf("获取上传的财务excel文件失败，错误信息: %s", err.Error()))
		}
	}()
	file_name := h.Filename
	time_unix_int := time.Now().Unix()
	time_unix_str := strconv.FormatInt(time_unix_int, 10)

	file_path := "upload/echart_data_upload/" + time_unix_str + "-" + file_name

	err = c.SaveToFile("upload_file", file_path)
	if err != nil {
		logs.Error(fmt.Sprintf("保存上传的财务文件出错，错误信息:%s", err.Error()))
	}
	file, err1 := excelize.OpenFile(file_path)
	if err1 != nil {
		logs.Error(fmt.Sprintf("解析上传的excel问价报错: %s", err.Error()))
	}

	o := orm.NewOrm()
	i := 0

	rows, _ := file.GetRows("sheet1")
	for _, row := range rows {
		caiwuDate := row[0]
		salesVolume, _ := strconv.ParseFloat(row[1], 64)
		studentIncress := utils.StrToInt(row[2])
		django := utils.StrToInt(row[3])
		vueDjango := utils.StrToInt(row[4])
		celery := utils.StrToInt(row[5])
		echartData := models.CaiwuData{
			CaiwuDate:      caiwuDate,
			SalesVolume:    salesVolume,
			StudentIncress: studentIncress,
			Django:         django,
			VueDjango:      vueDjango,
			Celery:         celery,
		}
		if i == 0 {
			i++
			continue
		}
		qs := o.QueryTable("sys_caiwu_data")
		is_exist := qs.Filter("caiwu_date", caiwuDate).Exist()
		if is_exist {
			qs.Filter("caiwu_date", caiwuDate).Delete()
		}

		// 精确到导入失败的数据信息提示
		_, err := o.Insert(&echartData)

		if err != nil { // 报错的数据
			errDataArr = append(errDataArr, caiwuDate)
		}
		i++
	}
	if len(errDataArr) <= 0 {
		message_map["code"] = 200
		message_map["msg"] = "导入成功"
	} else {
		message_map["code"] = 10002
		message_map["msg"] = "导入失败"
		message_map["err_data"] = errDataArr
	}
	c.Data["json"] = message_map
	c.ServeJSON()
}
