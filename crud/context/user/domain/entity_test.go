package domain

import (
	"reflect"
	"testing"

	"github.com/gloryof/go-crud-practice/crud/context/base"
)

func TestNewName(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name  string
		args  args
		want  Name
		want1 *base.ValidationResults
	}{
		{
			name: "正常系",
			args: args{
				value: "テスト",
			},
			want:  Name{value: "テスト"},
			want1: nil,
		},
		{
			name: "正常系（最大文字数）",
			args: args{
				value: "１２３４５６７８９０1234567890ABCDEFGHIJあいうえおかきくけこ",
			},
			want:  Name{value: "１２３４５６７８９０1234567890ABCDEFGHIJあいうえおかきくけこ"},
			want1: nil,
		},
		{
			name: "未入力",
			args: args{
				value: "",
			},
			want:  Name{},
			want1: createNameRequiredError(),
		},
		{
			name: "文字数オーバー",
			args: args{
				value: "１２３４５６７８９０1234567890ABCDEFGHIJあいうえおかきくけこさ",
			},
			want:  Name{},
			want1: createNameLengthError(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := NewName(tt.args.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewName() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewName() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewBirthDay(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name  string
		args  args
		want  BirthDay
		want1 *base.ValidationResults
	}{
		{
			name: "正常パターン",
			args: args{
				value: "2018-05-01",
			},
			want:  createBirthDay("2018-05-01"),
			want1: nil,
		},
		{
			name: "正常パターン（閏年）",
			args: args{
				value: "2016-02-29",
			},
			want:  createBirthDay("2016-02-29"),
			want1: nil,
		},
		{
			name: "正常パターン（400年ごとの閏年）",
			args: args{
				value: "2000-02-29",
			},
			want:  createBirthDay("2000-02-29"),
			want1: nil,
		},
		{
			name: "エラーパターン（未入力）",
			args: args{
				value: "",
			},
			want:  BirthDay{},
			want1: createBirthDayRequiredError(),
		},
		{
			name: "エラーパターン（日付形式にアルファベットを含む）",
			args: args{
				value: "201a-05-20",
			},
			want:  BirthDay{},
			want1: createBirthDayDateFormatError(),
		},
		{
			name: "エラーパターン（年数が4桁を超える）",
			args: args{
				value: "20018-05-20",
			},
			want:  BirthDay{},
			want1: createBirthDayDateFormatError(),
		},
		{
			name: "エラーパターン（月が0）",
			args: args{
				value: "2018-00-20",
			},
			want:  BirthDay{},
			want1: createBirthDayDateFormatError(),
		},
		{
			name: "エラーパターン（月が13）",
			args: args{
				value: "2018-13-20",
			},
			want:  BirthDay{},
			want1: createBirthDayDateFormatError(),
		},
		{
			name: "エラーパターン（月の桁数が1）",
			args: args{
				value: "2018-1-20",
			},
			want:  BirthDay{},
			want1: createBirthDayDateFormatError(),
		},
		{
			name: "エラーパターン（日が0）",
			args: args{
				value: "2018-01-00",
			},
			want:  BirthDay{},
			want1: createBirthDayDateFormatError(),
		},
		{
			name: "エラーパターン（日が32）",
			args: args{
				value: "2018-01-32",
			},
			want:  BirthDay{},
			want1: createBirthDayDateFormatError(),
		},
		{
			name: "日が31（30日までの月）",
			args: args{
				value: "2018-04-31",
			},
			want:  BirthDay{},
			want1: createBirthDayDateFormatError(),
		},
		{
			name: "日が29（閏年ではない2月）",
			args: args{
				value: "2018-02-29",
			},
			want:  BirthDay{},
			want1: createBirthDayDateFormatError(),
		},
		{
			name: "日が1桁",
			args: args{
				value: "2018-05-1",
			},
			want:  BirthDay{},
			want1: createBirthDayDateFormatError(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := NewBirthDay(tt.args.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBirthDay() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewBirthDay() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func createNameRequiredError() *base.ValidationResults {
	r := base.NewValidationResults()

	r.AddRequiredError("名前")

	return &r
}

func createNameLengthError() *base.ValidationResults {
	r := base.NewValidationResults()

	r.AddMaxLengthError("名前", 40)

	return &r
}

func createBirthDayRequiredError() *base.ValidationResults {
	r := base.NewValidationResults()

	r.AddRequiredError("誕生日")

	return &r
}

func createBirthDayDateFormatError() *base.ValidationResults {
	r := base.NewValidationResults()

	r.AddDateFormatError("誕生日")

	return &r
}

func createBirthDay(value string) BirthDay {

	r, _ := base.ConvertDateTime(value)

	return BirthDay{value: r}
}
