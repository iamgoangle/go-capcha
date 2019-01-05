package captcha_test

import (
	"testing"

	"github.com/iamgoangle/go-capcha/captcha"
)

func TestRightOperandOverLimit(t *testing.T) {
	got := captcha.Encode(1, 1, 1, 10)
	want := "righ operand should not greater than 9"

	if got != want {
		t.Errorf("it should say %v but got %v", got, want)
	}
}

func TestOperandOverLimit(t *testing.T) {
	got := captcha.Encode(1, 1, 5, 1)
	want := "operator should not greater than 4"

	if got != want {
		t.Errorf("it should say %v but got %v", got, want)
	}
}

func TestCaptchaPatternOne(t *testing.T) {
	got := captcha.Encode(1, 1, 1, 1)
	want := "1 + one"

	if got != want {
		t.Errorf("it should say %v but got %v", got, want)
	}
}
