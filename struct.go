package main

// Fundamental 指数基本面返回数据
type Fundamental struct {
	Code    int
	Message string
	Data    []FundamentalItem
}

// FundamentalItem 基本面数据项
type FundamentalItem struct {
	Date string `json:"date"`
	PeTtm MixMetricsType `json:"pe_ttm,omitempty"` // pe_ttm
	Pb MixMetricsType `json:"pb,omitempty"` // pb
	PsTtm MixMetricsType `json:"ps_ttm,omitempty"` // ps_ttm
	Dyr  FloatMetricsType `json:"dyr,omitempty"`// 股息率
	Cp float64 `json:"cp,omitempty"` // 收盘点位
	Rcp float64 `json:"r_cp,omitempty"` // 全收益收盘点位
	Cpc float64 `json:"cpc,omitempty"` // 涨跌幅
	Rcpc float64 `json:"r_cpc,omitempty"` // 全收益收盘点位涨跌幅
	Mc float64 `json:"mc,omitempty"` // 市值
	Cmc float64 `json:"cmc,omitempty"` // 流通市值
	Fb float64 `json:"fb,omitempty"` // 融资余额
	Sb float64 `json:"sb,omitempty"` // 融券余额
	HaShm float64 `json:"ha_shm,omitempty"` // 陆股通持仓金额
}

// FloatMetricsType 浮点型指标数据
type FloatMetricsType struct {
	Ew     float64 `json:"ew,omitempty"` // 等权
	Mcw    float64 `json:"mcw,omitempty"` // 市值加权
	Ewpvo  float64 `json:"ewpvo,omitempty"` // 正数等权
	Avg    float64 `json:"avg,omitempty"` // 平均值
	Median float64 `json:"median,omitempty"` // 中位数
}

// StructMetricsType 结构化指标数据
type StructMetricsType struct {
	Ew     Estimated `json:"ew,omitempty"` // 等权
	Mcw    Estimated `json:"mcw,omitempty"` // 市值加权
	Ewpvo  Estimated `json:"ewpvo,omitempty"` // 正数等权
	Avg    Estimated `json:"avg,omitempty"` // 平均值
	Median Estimated `json:"median,omitempty"` // 中位数
}

// MixMetricsType 混合化指标数据
type MixMetricsType struct {
	Granularity // 年份类型
	FloatMetricsType // 浮点型指标数据
}

// Granularity 年份类型
type Granularity struct {
	Fs StructMetricsType `json:"fs,omitempty"` // 所有
	Y20 StructMetricsType `json:"y20,omitempty"` // 20 年
	Y10 StructMetricsType `json:"y10,omitempty"` // 10 年
	Y5 StructMetricsType `json:"y5,omitempty"` // 5 年
	Y3 StructMetricsType `json:"y3,omitempty"` // 3 年
}

// Estimated 估值数据
type Estimated struct {
	Cv    float64 `json:"cv,omitempty"` // 当前值
	Cvpos float64 `json:"cvpos,omitempty"` // 当前分位点
	Minv  float64 `json:"minv,omitempty"` // 最小值
	Maxv  float64 `json:"maxv,omitempty"` // 最大值
	Maxpv float64 `json:"maxpv,omitempty"` // 最大正值
	Q5v   float64 `json:"q5v,omitempty"` // 50%分位点
	Q8v   float64 `json:"q8v,omitempty"` // 80%分位点
	Q2v   float64 `json:"q2v,omitempty"` // 20%分位点
	Avgv  float64 `json:"avgv,omitempty"` // 平均值
}

