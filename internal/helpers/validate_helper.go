package helpers

import (
	"fmt"
	"net/mail"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// ValidateStruct
// reflect.StructField => lấy các field có trong struct
// reflect.Value => lấy các thông tin của field đó
// reflect.StructField.Tag.Get("something") => lấy giá trị ở tag something
// reflect.Value.Kind() => kiểm tra kiểu dữ liệu của value ví dụ string, int , bool , struct, interface và so sánh với reflect.
// reflect.Value.String() => trả về dữ liệu dạng chữ
// reflect.Value.Int() => trả về dữ liệu dạng số
// reflect.Value.IsZero() => kiểm tra khi dữ liệu có phải là dữ liệu mặc định khi khởi tạo hay không và chưa có các thay đổi gì
func ValidateStruct(obj interface{}) map[string][]string {
	errs := make(map[string][]string)
	typ := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)
	//Duyệt qua các field trong struct
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i) // reflect lấy tên field
		value := val.Field(i) // reflect lấy giá trị
		//Lấy field name => return string Username
		fieldName := field.Name
		//Lấy dữ liệu ở tag validate => return string required,min=6,max=24
		fieldTagValidate := field.Tag.Get("validate")
		//Tách các giá trị của tag bằng ,
		separateTagValue := strings.Split(fieldTagValidate, ",")
		mapTagValue := make(map[string]string)
		trimValue := strings.TrimSpace(value.String())
		for _, v := range separateTagValue {
			if strings.Contains(v, "=") {
				separateValue := strings.Split(v, "=")
				mapTagValue[separateValue[0]] = separateValue[1]
			} else {
				mapTagValue[v] = ""
			}
		}
		for key, valueMap := range mapTagValue {

			if mapTagValue[key] == "" {
				switch key {
				case "required":
					err := RequiredTag(fieldName, trimValue)
					if err != nil {
						errs[fieldName] = append(errs[fieldName], err.Error())
					}
				case "email":
					err := EmailTag(fieldName, trimValue)
					if err != nil {
						errs[fieldName] = append(errs[fieldName], err.Error())
					}
				case "vietnamPhoneNumber":
					err := VietnamPhoneNumberTag(fieldName, trimValue)
					if err != nil {
						errs[fieldName] = append(errs[fieldName], err.Error())
					}
				case "password":
					err := PasswordTag(fieldName, trimValue)
					if err != nil {
						errs[fieldName] = append(errs[fieldName], err.Error())
					}
				}
			} else {
				switch key {
				case "min":
					err := MinValue(fieldName, value.Int(), valueMap)
					if err != nil {
						errs[fieldName] = append(errs[fieldName], err.Error())
					}
				case "max":
					err := MaxValue(fieldName, value.Int(), valueMap)
					if err != nil {
						errs[fieldName] = append(errs[fieldName], err.Error())
					}
				case "minLength":
					err := MinLength(fieldName, trimValue, valueMap)
					if err != nil {
						errs[fieldName] = append(errs[fieldName], err.Error())
					}
				case "maxLength":
					err := MaxLength(fieldName, trimValue, valueMap)
					if err != nil {
						errs[fieldName] = append(errs[fieldName], err.Error())
					}
				}

			}
		}
		//Lấy loại dữ liệu của dữ liệu => return type dữ liệu ví dụ : string, int
		//value.Kind()
		//Lấy dữ liệu đươc gán vào nếu ở dạng string
		//value.String()
		//Lấy dữ liệu nhập vào nếu ở dạng số
		//value.Int()
	}
	return errs
}

func MinValue(fieldName string, fieldValue int64, tagValue string) error {
	numConv, err := strconv.Atoi(tagValue)
	if err != nil {
		return err
	}
	if fieldValue < int64(numConv) {
		return fmt.Errorf("%s must be greater than %d", fieldName, numConv)
	}
	return nil
}

func MaxValue(fieldName string, fieldValue int64, tagValue string) error {
	numConv, err := strconv.Atoi(tagValue)
	if err != nil {
		return err
	}
	if fieldValue > int64(numConv) {
		return fmt.Errorf("%s must be less than %d", fieldName, numConv)
	}
	return nil
}

func RequiredTag(tagName string, tagValue string) error {
	if tagValue == "" {
		return fmt.Errorf("%s is required", tagName)
	}
	if reflect.ValueOf(tagValue).IsZero() {
		return fmt.Errorf("%s is required", tagName)
	}

	return nil
}

func EmailTag(tagName string, tagValue string) error {
	_, err := mail.ParseAddress(tagValue)
	if err != nil {
		return fmt.Errorf("%s is not a valid email address", tagName)
	}
	return nil
}

func PasswordTag(tagName string, tagValue string) error {
	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	for _, char := range tagValue {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	if hasUpper && hasLower && hasNumber && hasSpecial {
		return nil
	}
	return fmt.Errorf("%s is not a valid password", tagName)
}

func VietnamPhoneNumberTag(tagName string, tagValue string) error {
	const PhoneNumber string = `^[0|+84][3|5|7|9]0-9{8}$`
	re := regexp.MustCompile(PhoneNumber)
	if !re.MatchString(tagValue) {
		return fmt.Errorf("%s is not a valid vietnam phone number", tagName)
	}
	return nil
}

func MaxLength(fieldName string, fieldValue string, tagValue string) error {
	numConv, err := strconv.Atoi(tagValue)
	if err != nil {
		return err
	}
	if len(fieldValue) > numConv {
		return fmt.Errorf("%s must be less than %d characters", fieldName, numConv)
	}
	return nil
}

func MinLength(fieldName string, fieldValue string, tagValue string) error {

	numConv, err := strconv.Atoi(tagValue)
	if err != nil {
		return err
	}
	if len(fieldValue) < numConv {
		return fmt.Errorf("%s must be at least %d characters", fieldName, numConv)
	}

	return nil
}
