package base

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValidationResults(t *testing.T) {

	const item1 = "Test1"
	const item2 = "Test2"
	const item3 = "Test3"

	sut := NewValidationResults()

	if len(sut.GetResults()) != 0 {

		t.Fatal("エラーが0件ではない")
	}
	if sut.HasError() != false {

		t.Fatal("HasErroがfalseではない")
	}

	sut.AddRequiredError(item1)
	if sut.HasError() != true {

		t.Fatal("HasErroがtrueではない")
	}

	sut.AddDateFormatError(item2)
	sut.AddRequiredError(item3)
	if len(sut.GetResults()) != 3 {

		t.Fatal("エラーが3件ではない")
	}

	actual := sut.GetResults()

	const failMessage = "%d件目のメッセージが不正。[expected:%s, actual:%s]"

	const expectMessage1 = item1 + "を入力してください"
	if actual[0] != expectMessage1 {

		t.Fatal(fmt.Sprintf(failMessage, 1, expectMessage1, actual[0]))
	}

	const expectMessage2 = item2 + "は日付形式で入力してください"
	if actual[1] != expectMessage2 {

		t.Fatal(fmt.Sprintf(failMessage, 2, expectMessage2, actual[1]))
	}

	const expectMessage3 = item3 + "を入力してください"
	if actual[2] != expectMessage3 {

		t.Fatal(fmt.Sprintf(failMessage, 3, expectMessage3, actual[2]))
	}
}

func TestMerge(t *testing.T) {

	data1 := NewValidationResults()
	data2 := NewValidationResults()

	data3 := NewValidationResults()
	data3.AddRequiredError("Test1")

	data4 := NewValidationResults()
	data4.AddRequiredError("Test2")
	data4.AddRequiredError("Test3")

	data5 := NewValidationResults()
	data5.AddRequiredError("Test3")
	data5.AddRequiredError("Test4")

	type args struct {
		results []ValidationResults
	}
	tests := []struct {
		name string
		args args
		want ValidationResults
	}{
		{
			name: "空配列が指定された場合",
			args: args{results: []ValidationResults{}},
			want: expectedMerged1(),
		},
		{
			name: "エラーがない要素が一つ指定された場合",
			args: args{results: []ValidationResults{data1}},
			want: expectedMerged1(),
		},
		{
			name: "エラーがない要素が複数指定された場合",
			args: args{results: []ValidationResults{data1, data2}},
			want: expectedMerged1(),
		},
		{
			name: "エラーがある要素が一つ指定された場合",
			args: args{results: []ValidationResults{data3}},
			want: expectedMerged2(),
		},
		{
			name: "エラーがあり要素が複数指定され、重複するエラーがない場合",
			args: args{results: []ValidationResults{data3, data4}},
			want: expectedMerged3(),
		},
		{
			name: "エラーがあり要素が複数指定され、重複するエラーがある場合",
			args: args{results: []ValidationResults{data3, data4, data5}},
			want: expectedMerged4(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.results...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func expectedMerged1() ValidationResults {

	expected := NewValidationResults()

	return expected
}

func expectedMerged2() ValidationResults {

	expected := NewValidationResults()

	expected.AddRequiredError("Test1")

	return expected
}

func expectedMerged3() ValidationResults {

	expected := NewValidationResults()

	expected.AddRequiredError("Test1")
	expected.AddRequiredError("Test2")
	expected.AddRequiredError("Test3")

	return expected
}

func expectedMerged4() ValidationResults {

	expected := NewValidationResults()

	expected.AddRequiredError("Test1")
	expected.AddRequiredError("Test2")
	expected.AddRequiredError("Test3")
	expected.AddRequiredError("Test4")

	return expected
}
