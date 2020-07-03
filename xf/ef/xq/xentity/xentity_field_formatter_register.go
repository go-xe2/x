package xentity

func registerFieldFormatters(entry FieldFormatterEntry) {
	// 日期格式化
	entry.Register(FormatDate, NewFieldDateFormatter())
	// 字符串截取格式化
	entry.Register(FormatStrCat, NewFieldStrCatFormatter())
}
