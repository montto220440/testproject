package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Global struct {
	Th map[string]string `json:"th" bson:"th"`
	En map[string]string `json:"en" bson:"en"`
}

func Resp(c *gin.Context, status int, msg string, data interface{}) {
	// var language string = os.Getenv("LANGUAGE")

	var BYPASS_REDIS string = os.Getenv("BYPASS_REDIS")
	var msgGlobal string
	var code int = 1
	var statusCode int = 200
	var statusMsg string = "ok"
	if status == 0 || status == 200 {
		code = 0
	} else if status == 101 {
		code = 101
	} else if status == 1 {
		code = 1
	} else if status == 400 {
		statusCode = 400
		statusMsg = "invalid data"
	} else if status == 401 {
		statusCode = 401
		// c.AbortWithStatusJSON(statusCode, gin.H{
		// 	"code":    code,
		// 	"msg":     msg,
		// 	"service": os.Getenv("SERVICE"),
		// 	"payload": data,
		// })
		c.AbortWithStatusJSON(statusCode, gin.H{
			"status": map[string]interface{}{
				"code":    statusCode,
				"message": "unauthorized",
			},
			"message": map[string]interface{}{
				"code":    1,
				"message": msg,
			},
			"data": data,
		})
		return
		// } else if status == 404 {
		// 	statusCode = 404
		// 	c.AbortWithStatusJSON(statusCode, gin.H{
		// 		"status": map[string]interface{}{
		// 			"code":    statusCode,
		// 			"message": msg,
		// 		},
		// 	})
		// 	return
	} else if status == 500 {
		statusCode = 500
		statusMsg = "server error"
	}

	// file, _ := ioutil.ReadFile("src/static/global.json")
	obj := Global{}
	obj.En = map[string]string{
		"invalidData":            "Invalid Data",
		"bankNumberInvalid":      "Invalid Bank Number!!",
		"phoneNunberDuplicate":   "Duplicate, Phone Number",
		"usernameDuplicate":      "Duplicate, Username",
		"bankNumberDuplicate":    "Duplicate, Bank Number",
		"cannotCreateUsername":   "Cannot create UserDetail, Please contact staff",
		"registerSuccess":        "Register Successfully",
		"cannotConnectServer":    "Cannot connect server, Please contact staff",
		"bankAccountLimit":       "Back Account cannot exceed more than 5, Please contact staff",
		"bankUpdateSuccess":      "Update bank Successfully",
		"amountCannotExceed2000": "Deposit amount cannot exceed more than 2,000, Please try again",
		"withdrawInprogress":     "Withdraw In Progress, Waiting",
		"invalidAmount":          "Invalid amount, Please connact staff",
		"userNotFound":           "User detail not found",
		"registerFail":           "Register got an trouble, Please try again",
		"maintenanceSystem":      "Maintenance System",
		"invalidImg":             "Invalid Image",
		"duplicateBankStatement": "Already Deposit, Please contact staff.",
		"cannotVerifySlip":       "Cannot verify slip, Please contact staff.",
		"successfully":           "Successfully",
	}

	obj.Th = map[string]string{
		"invalidData":            "ส่งข้อมูลไม่ถูกต้อง กรุณาติดต่อพนักงาน",
		"bankNumberInvalid":      "เลขบัญชีธนาคารไม่ถูกต้อง กรุณาทำรายการใหม่อีกครั้ง!!",
		"phoneNunberDuplicate":   "ไม่สามารถใช้เบอนี้ได้ เบอร์นี้ถูกใช้งานแล้ว!!",
		"usernameDuplicate":      "ไม่สามารถใช้ยูสเซอร์เนมนี้ได้ ยูสเซอร์เนมนี้ถูกใช้งานแล้ว!!",
		"bankNumberDuplicate":    "บัญชีนี้ถูกใช้ไปแล้ว กรุณาติดต่อพนักงาน!!",
		"cannotCreateUsername":   "ไม่สามารถสร้างยูสเซอร์ได้ กรุณาติดต่อพนักงาน",
		"registerSuccess":        "สมัครสมาชิกสำเร็จ",
		"cannotConnectServer":    "ไม่สามารถเชื่อมต่อระบบได้ กรุณาติดต่อพนักงาน",
		"bankAccountLimit":       "บัญชีเพิ่่มได้สูงสุด 5 บัญชี กรุณาติดต่อพนักงาน",
		"bankUpdateSuccess":      "เพิ่มบัญชีธนาคารสำเร็จ",
		"amountCannotExceed2000": "ระบบทศนิยมฝากสูงสุด 2000 บาท กรุณากรอกจำนวนอีกครั้ง",
		"withdrawInprogress":     "มีรายการถอนค้างอยู่ กรุณารอซักครู๋",
		"invalidAmount":          "ยอดเงินไม่ถูกต้อง กรุณาติดต่อพนักงาน",
		"userNotFound":           "ไม่พบข้อมูลลูกค้า",
		"registerFail":           "ไม่สามารถสมัครสมาชิกได้ กรุณาตรวจสอบข้อมูลใหม่อีกครั้ง",
		"maintenanceSystem":      "ระบบปิดปรับปรุงชั่วคราว กรุณาติดต่อพนักงาน",
		"invalidImg":             "กรุณาส่งรูปให้ถูกต้อง",
		"duplicateBankStatement": "รายการฝากนี้มีอยู่ในระบบแล้ว หากยอดไม่เข้ากรุณาติดต่อพนักงาน!",
		"cannotVerifySlip":       "ไม่สามารถตรวจสอบสลิปได้ กรุณาติดต่อพนักงาน",
		"successfully":           "ทำรายการสำเร็จ",
	}

	// _ = json.Unmarshal([]byte(file), &obj)
	// fmt.Println("Unmarshal")
	// fmt.Println("obj[]", obj)

	// if language == "EN" {
	if BYPASS_REDIS == "TRUE" {
		if _, ok := obj.En[msg]; ok {
			msgGlobal = obj.En[msg]
		}

	} else {
		if _, ok := obj.Th[msg]; ok {
			msgGlobal = obj.Th[msg]
		}

	}
	//  else if language == "TH" {
	// 	fmt.Println("TH", obj.Th[msg])
	// 	msgGlobal = obj.Th[msg]
	// }
	//

	if msgGlobal == "" {
		msgGlobal = msg
	}

	// c.JSON(statusCode, gin.H{
	// 	"code":    code,
	// 	"msg":     msg,
	// 	"service": os.Getenv("SERVICE"),
	// 	"payload": data,
	// })
	c.JSON(statusCode, gin.H{
		"status": map[string]interface{}{
			"code":    statusCode,
			"message": statusMsg,
		},
		"message": map[string]interface{}{
			"code":    code,
			"message": msgGlobal,
		},
		"data": data,
	})
}

func Find(slice interface{}, f func(value interface{}) bool) interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() == reflect.Slice {
		for index := 0; index < s.Len(); index++ {
			if f(s.Index(index).Interface()) {
				return s.Index(index).Interface()
			}
		}
	}
	return nil
}

func FindIndex(slice interface{}, f func(value interface{}) bool) int {
	s := reflect.ValueOf(slice)
	if s.Kind() == reflect.Slice {
		for index := 0; index < s.Len(); index++ {
			if f(s.Index(index).Interface()) {
				return index
			}
		}
	}
	return -1
}

func OpenFileToStruct(path string, obj interface{}) error {

	jsonFile, err := os.Open(fmt.Sprintf("./static%s", path))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, obj)

	return err
}
