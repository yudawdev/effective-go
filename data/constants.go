package data

import "fmt"

const (
	_ = iota
	PENDING
	NEW
)

const name = "NAME"

const (
	PF = "PF"
	F  = "F"
)

// 常量
// PENDING: 1
// NEW: 2
// name: NAME
// PF: PF
// F: F
func constant() {
	fmt.Printf("PENDING: %v\n", PENDING)
	fmt.Printf("NEW: %v\n", NEW)

	fmt.Printf("name: %v\n", name)

	fmt.Printf("PF: %v\n", PF)
	fmt.Printf("F: %v\n", F)
}

// Example: 实现类似 Java 这样的枚举 `AU("AU", AUD, "Australia/Canberra")`

// Country 定义国家或者地区结构体
type Country struct {
	name string
}

func (c Country) String() string {
	return c.name
}

// 定义 Country 枚举实例
var (
	AUCountry = Country{"AU"}
	HKCountry = Country{"HK"}
)

// Currency 定义币种类型
type Currency string

// 定义币种枚举常量
const (
	HKD = Currency("HKD")
	AUD = Currency("AUD")
)

// BizCountry 定义业务地区的结构体
type BizCountry struct {
	country  Country
	currency Currency
}

// Country Get country
func (b BizCountry) Country() Country {
	return b.country
}

// Currency Get currency
func (b BizCountry) Currency() Currency {
	return b.currency
}

// 实例 - 业务地区枚举实例
var (
	AU = BizCountry{
		country:  AUCountry,
		currency: AUD,
	}

	HK = BizCountry{
		country:  HKCountry,
		currency: HKD,
	}
)

// 最终用户使用
func useCase() {
	country := AU.Country()
	currency := AU.Currency()

	fmt.Printf("country: %v, currency: %v", country, currency)
}
