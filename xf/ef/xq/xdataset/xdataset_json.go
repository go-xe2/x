package xdataset

import (
	"bytes"
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/encoding/xjson"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/xf/ef/xqi"
)

func (ds *tDataset) MarshalJSON() (data []byte, err error) {
	// 序列化
	writer := bytes.NewBufferString("")
	writer.WriteString(`{ "fields": `)
	if _, err := ds.fields.WriteJson(writer); err != nil {
		return nil, err
	}
	writer.WriteString(`, "rows": [`)
	var innerErr error
	ds.Iterator(func(row xqi.DatasetRow) bool {
		if row.RowIndex() > 0 {
			if _, innerErr = writer.Write([]byte{','}); innerErr != nil {
				return false
			}
		}
		if _, innerErr = row.WriteJson(writer); innerErr != nil {
			return false
		}
		return true
	})
	if innerErr != nil {
		return nil, innerErr
	}
	writer.WriteString(`]}`)
	return writer.Bytes(), nil
}

func (ds *tDataset) UnmarshalJSON(data []byte) error {
	const szErrJsonData = "不是dataset的有效json字符串数据"
	json, err := xjson.DecodeToJson(data)
	if err != nil {
		return err
	}
	fields := json.GetMap("fields")
	if fields == nil {
		return exception.NewText(szErrJsonData)
	}
	colCount := t.Int(fields["count"])
	if colCount <= 0 {
		return exception.NewText(szErrJsonData)
	}
	ds.fields.Clear()
	for i := 0; i < colCount; i++ {
		colInfo := t.Map(fields[t.String(i)])
		if colInfo == nil {
			return exception.NewText(szErrJsonData)
		}
		fdName := t.String(colInfo["N"])
		szType := t.String(colInfo["T"])
		if fdName == "" || szType == "" {
			return exception.NewText(szErrJsonData)
		}
		fdType, isOk := xqi.JsFieldDataTypeToType(szType)
		if !isOk {
			return exception.NewText(szErrJsonData)
		}
		ds.fields.Add(fdName, fdType)
	}
	// 解析数据
	rows := json.GetArray("rows")
	if rows == nil {
		return exception.NewText(szErrJsonData)
	}
	for _, v := range rows {
		row, ok := v.(map[string]interface{})
		if !ok {
			return exception.NewText(szErrJsonData)
		}
		rowData := make([]interface{}, colCount)
		for j := 0; j < colCount; j++ {
			if v, ok := row[t.String(j)]; ok {
				rowData[j] = v
			}
		}
		ds.Append(rowData...)
	}
	// 解析
	return nil
}
