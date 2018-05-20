package base

import (
	"fmt"
)

const (
	// required 必須メッセージのパターン
	required MessagePattern = iota
	// dateFormat 日付形式メッセージのパターン
	dateFormat
)

// MessagePattern メッセージパターン
type MessagePattern int

// String Stringerの実装関数
func (p MessagePattern) String() string {
	switch p {
	case required:
		return "%sを入力してください"
	case dateFormat:
		return "%sは日付形式で入力してください"
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

// GetResults 結果のリストを取得する
func (r *ValidationResults) GetResults() []string { return r.results }

// HasError エラーがあるかどうかを判定する
func (r *ValidationResults) HasError() bool { return len(r.results) > 0 }

// AddRequiredError 必須エラーを追加する
func (r *ValidationResults) AddRequiredError(items string) {

	r.results = append(r.results, fmt.Sprintf(fmt.Sprint(required), items))
}

// AddDateFormatError 日付形式エラーを追加する
func (r *ValidationResults) AddDateFormatError(items string) {

	r.results = append(r.results, fmt.Sprintf(fmt.Sprint(dateFormat), items))
}

// Merge 全ての入力エラーを繋げる
// 同じエラーがあった場合は一つにまとめられる
func Merge(results ...ValidationResults) ValidationResults {

	messages := []string{}

	for _, r := range results {

		for _, v := range r.GetResults() {

			if contains(messages, v) == false {

				messages = append(messages, v)
			}
		}
	}

	return ValidationResults{results: messages}
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
