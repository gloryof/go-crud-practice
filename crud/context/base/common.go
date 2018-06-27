package base

import (
	"fmt"
	"strings"
	"time"
)

const (
	// required 必須メッセージのパターン
	required MessagePattern = iota
	// maxLength 最大文字数メッセージのパターン
	maxLength
	// dateFormat 日付形式メッセージのパターン
	dateFormat
	// dataNotExists データが存在しないメッセージのパターン
	dataNotExists
)

// MessagePattern メッセージパターン
type MessagePattern int

// String Stringerの実装関数
func (p MessagePattern) String() string {
	switch p {
	case required:
		return "%sを入力してください"
	case maxLength:
		return "%sは%d文字以内で入力してください"
	case dateFormat:
		return "%sは日付形式で入力してください"
	case dataNotExists:
		return "%sが存在しません。"
	default:
		return "入力値が不正です"
	}
}

// ValidationResults 入力検証結果の集合
type ValidationResults struct {
	results []string
}

// NewValidationResults 新しい入力検証結果を作成する
func NewValidationResults() ValidationResults {
	return ValidationResults{results: []string{}}
}

// Error errorの実装関数
func (r *ValidationResults) Error() string {
	return strings.Join(r.results, ";")
}

// GetResults 結果のリストを取得する
func (r *ValidationResults) GetResults() []string { return r.results }

// HasError エラーがあるかどうかを判定する
func (r *ValidationResults) HasError() bool { return len(r.results) > 0 }

// AddRequiredError 必須エラーを追加する
func (r *ValidationResults) AddRequiredError(items string) {

	r.results = append(r.results, fmt.Sprintf(fmt.Sprint(required), items))
}

// AddMaxLengthError 最大文字数エラーを追加する
func (r *ValidationResults) AddMaxLengthError(items string, length uint64) {

	r.results = append(r.results, fmt.Sprintf(fmt.Sprint(maxLength), items, length))
}

// AddDateFormatError 日付形式エラーを追加する
func (r *ValidationResults) AddDateFormatError(items string) {

	r.results = append(r.results, fmt.Sprintf(fmt.Sprint(dateFormat), items))
}

// AddDataNotExistsError データが存在しないエラーを追加する
func (r *ValidationResults) AddDataNotExistsError(items string) {

	r.results = append(r.results, fmt.Sprintf(fmt.Sprint(dataNotExists), items))
}

// Add パラメータで渡された入力エラーを追加する
// 同じエラーがあった場合は一つにまとめられる
func (r *ValidationResults) Add(results ValidationResults) {

	for _, rs := range results.results {

		if contains(r.results, rs) == false {
			r.results = append(r.results, rs)
		}
	}
}

// contains 対象の値が存在しているかを判定する
// 存在する場合：true、存在しない場合：false
func contains(messages []string, value string) bool {

	for _, v := range messages {

		if v == value {

			return true
		}
	}

	return false
}

// ConvertDateTime 日付型に変換する
// 日付の形式は「yyyy-MM-dd」形式の東京ロケール。
// 変換に失敗した場合はerrorを返す
func ConvertDateTime(value string) (time.Time, error) {

	loc, _ := time.LoadLocation("Asia/Tokyo")
	return time.ParseInLocation("2006-01-02", value, loc)
}
