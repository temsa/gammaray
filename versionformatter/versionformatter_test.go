package versionformatter

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSemverVersion(t *testing.T) {
	res, _ := Format("1.0.0-rc1")
	fmt.Println("TestSemverVersion result:", res)
	if diff := cmp.Diff(res, "1.0.0-rc1"); diff != "" {
		t.Errorf("TestSemverVersion: after Format : (-got +want)\n%s", diff)
	}
}

func TestNotSemverVersion(t *testing.T) {
	res, _ := Format("1.0.rc.1")
	fmt.Println("TestNotSemverVersion result:", res)
	if diff := cmp.Diff(res, "1.0.0-rc.1"); diff != "" {
		t.Errorf("TestNotSemverVersion: after Format : (-got +want)\n%s", diff)
	}
}

func TestNotSemverVersionAgain(t *testing.T) {
	res, _ := Format("0.8.beta-1")
	fmt.Println("TestNotSemverVersionAgain result:", res)
	if diff := cmp.Diff(res, "0.8.0-beta-1"); diff != "" {
		t.Errorf("TestNotSemverVersionAgain: after Format : (-got +want)\n%s", diff)
	}
}

func TestNotSemverSimpleExpression(t *testing.T) {
	res, _ := Format("<0.8")
	fmt.Println("TestNotSemverSimpleExpression result:", res)
	if diff := cmp.Diff(res, "<0.8.0"); diff != "" {
		t.Errorf("TestNotSemverSimpleExpression: after Format : (-got +want)\n%s", diff)
	}
}

func TestPrimitiveRange(t *testing.T) {
	res, _ := Format(">=1.1.0 <=1.1.1")
	fmt.Println("TestPrimitiveRange result:", res)
	if diff := cmp.Diff(res, ">=1.1.0, <=1.1.1"); diff != "" {
		t.Errorf("TestPrimitiveRange: after Format : (-got +want)\n%s", diff)
	}
}

func TestPrimitiveRangeWithComma(t *testing.T) {
	res, _ := Format(">=1.1.0, <=1.1.1")
	fmt.Println("TestPrimitiveRangeWithComma result:", res)
	if diff := cmp.Diff(res, ">=1.1.0, <=1.1.1"); diff != "" {
		t.Errorf("TestPrimitiveRangeWithComma: after Format : (-got +want)\n%s", diff)
	}
}

func TestPrimitiveRangeWithPipe(t *testing.T) {
	res, _ := Format(">=1.1.0|<=1.1.1")
	fmt.Println("TestPrimitiveRangeWithPipe result:", res)
	if diff := cmp.Diff(res, ">=1.1.0, <=1.1.1"); diff != "" {
		t.Errorf("TestPrimitiveRangeWithPipe: after Format : (-got +want)\n%s", diff)
	}
}

func TestReversedPrimitiveRange(t *testing.T) {
	res, _ := Format("<=1.1.1, >=1.1.0")
	fmt.Println("TestReversedPrimitiveRange result:", res)
	if diff := cmp.Diff(res, ">=1.1.0, <=1.1.1"); diff != "" {
		t.Errorf("TestReversedPrimitiveRange: after Format : (-got +want)\n%s", diff)
	}
}

func TestHyphenRange(t *testing.T) {
	res, _ := Format("1.1.0 - 1.1.1")
	fmt.Println("TestHyphenRange result:", res)
	if diff := cmp.Diff(res, "1.1.0 - 1.1.1"); diff != "" {
		t.Errorf("TestHyphenRange: after Format : (-got +want)\n%s", diff)
	}
}

func TestCaretRange(t *testing.T) {
	res, _ := Format("^1.1")
	fmt.Println("TestCaretRange result:", res)
	if diff := cmp.Diff(res, "^1.1.0"); diff != "" {
		t.Errorf("TestCaretRange: after Format : (-got +want)\n%s", diff)
	}
}

func TestTildeRange(t *testing.T) {
	res, _ := Format("~1.1")
	fmt.Println("TestTildeRange result:", res)
	if diff := cmp.Diff(res, "~1.1.0"); diff != "" {
		t.Errorf("TestTildeRange: after Format : (-got +want)\n%s", diff)
	}
}

func TestXRangeWithX(t *testing.T) {
	res, _ := Format("1.1.x")
	fmt.Println("TestXRangeWithX result:", res)
	if diff := cmp.Diff(res, "1.1.x"); diff != "" {
		t.Errorf("TestXRangeWithX: after Format : (-got +want)\n%s", diff)
	}
}

func TestXRangeWithUpperX(t *testing.T) {
	res, _ := Format("1.1.X")
	fmt.Println("TestXRangeWithUpperX result:", res)
	if diff := cmp.Diff(res, "1.1.X"); diff != "" {
		t.Errorf("TestXRangeWithUpperX: after Format : (-got +want)\n%s", diff)
	}
}

func TestXRangeWithStar(t *testing.T) {
	res, _ := Format("1.1.*")
	fmt.Println("TestXRangeWithStar result:", res)
	if diff := cmp.Diff(res, "1.1.*"); diff != "" {
		t.Errorf("TestXRangeWithStar: after Format : (-got +want)\n%s", diff)
	}
}

func TestUnaryAndVersion(t *testing.T) {
	res, _ := Format("<=1.1.1")
	fmt.Println("TestUnaryAndVersion result:", res)
	if diff := cmp.Diff(res, "<=1.1.1"); diff != "" {
		t.Errorf("TestUnaryAndVersion: after Format : (-got +want)\n%s", diff)
	}
}

func TestUnaryAndVersionWithSpace(t *testing.T) {
	res, _ := Format(">= 2.0.1")
	fmt.Println("TestUnaryAndVersion result:", res)
	if diff := cmp.Diff(res, ">=2.0.1"); diff != "" {
		t.Errorf("TestUnaryAndVersion: after Format : (-got +want)\n%s", diff)
	}
}

func TestOrExpression(t *testing.T) {
	res, _ := Format(">=1.1.0 || <=1.1.1")
	fmt.Println("TestOrExpression result:", res)
	if diff := cmp.Diff(res, ">=1.1.0 || <=1.1.1"); diff != "" {
		t.Errorf("TestOrExpression: after Format : (-got +want)\n%s", diff)
	}
}

func TestComplexExpression(t *testing.T) {
	res, _ := Format(" >=0.08beta-1 || !1.rc.1 <=1.rc+build.543 || >2 ")
	fmt.Println("TestComplexExpression result:", res)
	if diff := cmp.Diff(res, ">=0.8.0-beta-1 || !1.0.0-rc.1 || <=1.0.0-rc+build.543 || >2.0.0"); diff != "" {
		t.Errorf("TestComplexExpression: after Format : (-got +want)\n%s", diff)
	}
}

func TestMultiRange(t *testing.T) {
	res, _ := Format(">=1.3.0 <1.3.2 || >=1.4.0 <1.4.11 || >=1.5.0 <1.5.2")
	fmt.Println("TestMultiRange result:", res)
	if diff := cmp.Diff(res, ">=1.3.0, <1.3.2 || >=1.4.0, <1.4.11 || >=1.5.0, <1.5.2"); diff != "" {
		t.Errorf("TestMultiRange: after Format : (-got +want)\n%s", diff)
	}
}

func TestSimplePipeExpression(t *testing.T) {
	res, _ := Format("2.0.0 2.0.0-x | 2.1.0-x 2.1.1 2.1.2")
	fmt.Println("TestSimplePipeExpression result:", res)
	if diff := cmp.Diff(res, "2.0.0 || 2.0.0-x || 2.1.0-x || 2.1.1 || 2.1.2"); diff != "" {
		t.Errorf("TestComplexExpression: after Format : (-got +want)\n%s", diff)
	}
}
