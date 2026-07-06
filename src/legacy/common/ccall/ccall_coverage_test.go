package ccall

import (
	"testing"
)

func TestCallCoverage(t *testing.T) {
	// Void functions
	CallVoidVoid(testPtrVoidVoid())
	CallVoidUPtr(testPtrVoidUPtr(), 1)
	CallVoidUPtr2(testPtrVoidUPtr2(), 1, 2)
	CallVoidUPtr3(testPtrVoidUPtr3(), 1, 2, 3)
	CallVoidUPtr4(testPtrVoidUPtr4(), 1, 2, 3, 4)
	CallVoidUPtr5(testPtrVoidUPtr5(), 1, 2, 3, 4, 5)
	CallVoidUPtr6(testPtrVoidUPtr6(), 1, 2, 3, 4, 5, 6)

	CallVoidPtr(testPtrVoidPtr(), nil)
	CallVoidPtr2(testPtrVoidPtr2(), nil, nil)
	CallVoidPtr3(testPtrVoidPtr3(), nil, nil, nil)
	CallVoidPtr4(testPtrVoidPtr4(), nil, nil, nil, nil)
	CallVoidPtr5(testPtrVoidPtr5(), nil, nil, nil, nil, nil)
	CallVoidPtr6(testPtrVoidPtr6(), nil, nil, nil, nil, nil, nil)

	CallVoidInt(testPtrVoidInt(), 1)
	CallVoidInt2(testPtrVoidInt2(), 1, 2)
	CallVoidInt3(testPtrVoidInt3(), 1, 2, 3)
	CallVoidInt4(testPtrVoidInt4(), 1, 2, 3, 4)
	CallVoidInt5(testPtrVoidInt5(), 1, 2, 3, 4, 5)
	CallVoidInt6(testPtrVoidInt6(), 1, 2, 3, 4, 5, 6)

	// UPtr return functions
	if v := CallUPtrVoid(testPtrUPtrVoid()); v != 1 {
		t.Errorf("CallUPtrVoid = %d", v)
	}
	CallUPtrUPtr(testPtrUPtrUPtr(), 1)
	CallUPtrUPtr2(testPtrUPtrUPtr2(), 1, 2)
	CallUPtrUPtr3(testPtrUPtrUPtr3(), 1, 2, 3)
	CallUPtrUPtr4(testPtrUPtrUPtr4(), 1, 2, 3, 4)
	CallUPtrUPtr5(testPtrUPtrUPtr5(), 1, 2, 3, 4, 5)
	CallUPtrUPtr6(testPtrUPtrUPtr6(), 1, 2, 3, 4, 5, 6)

	CallUPtrPtr(testPtrUPtrPtr(), nil)
	CallUPtrPtr2(testPtrUPtrPtr2(), nil, nil)
	CallUPtrPtr3(testPtrUPtrPtr3(), nil, nil, nil)
	CallUPtrPtr4(testPtrUPtrPtr4(), nil, nil, nil, nil)
	CallUPtrPtr5(testPtrUPtrPtr5(), nil, nil, nil, nil, nil)
	CallUPtrPtr6(testPtrUPtrPtr6(), nil, nil, nil, nil, nil, nil)

	CallUPtrInt(testPtrUPtrInt(), 1)
	CallUPtrInt2(testPtrUPtrInt2(), 1, 2)
	CallUPtrInt3(testPtrUPtrInt3(), 1, 2, 3)
	CallUPtrInt4(testPtrUPtrInt4(), 1, 2, 3, 4)
	CallUPtrInt5(testPtrUPtrInt5(), 1, 2, 3, 4, 5)
	CallUPtrInt6(testPtrUPtrInt6(), 1, 2, 3, 4, 5, 6)

	// Ptr return functions
	CallPtrVoid(testPtrPtrVoid())
	CallPtrUPtr(testPtrPtrUPtr(), 1)
	CallPtrUPtr2(testPtrPtrUPtr2(), 1, 2)
	CallPtrUPtr3(testPtrPtrUPtr3(), 1, 2, 3)
	CallPtrUPtr4(testPtrPtrUPtr4(), 1, 2, 3, 4)
	CallPtrUPtr5(testPtrPtrUPtr5(), 1, 2, 3, 4, 5)
	CallPtrUPtr6(testPtrPtrUPtr6(), 1, 2, 3, 4, 5, 6)

	CallPtrPtr(testPtrPtrPtr(), nil)
	CallPtrPtr2(testPtrPtrPtr2(), nil, nil)
	CallPtrPtr3(testPtrPtrPtr3(), nil, nil, nil)
	CallPtrPtr4(testPtrPtrPtr4(), nil, nil, nil, nil)
	CallPtrPtr5(testPtrPtrPtr5(), nil, nil, nil, nil, nil)
	CallPtrPtr6(testPtrPtrPtr6(), nil, nil, nil, nil, nil, nil)

	CallPtrInt(testPtrPtrInt(), 1)
	CallPtrInt2(testPtrPtrInt2(), 1, 2)
	CallPtrInt3(testPtrPtrInt3(), 1, 2, 3)
	CallPtrInt4(testPtrPtrInt4(), 1, 2, 3, 4)
	CallPtrInt5(testPtrPtrInt5(), 1, 2, 3, 4, 5)
	CallPtrInt6(testPtrPtrInt6(), 1, 2, 3, 4, 5, 6)

	// Int return functions
	if v := CallIntVoid(testPtrIntVoid()); v != 42 {
		t.Errorf("CallIntVoid = %d", v)
	}
	CallIntUPtr(testPtrIntUPtr(), 1)
	CallIntUPtr2(testPtrIntUPtr2(), 1, 2)
	CallIntUPtr3(testPtrIntUPtr3(), 1, 2, 3)
	CallIntUPtr4(testPtrIntUPtr4(), 1, 2, 3, 4)
	CallIntUPtr5(testPtrIntUPtr5(), 1, 2, 3, 4, 5)
	CallIntUPtr6(testPtrIntUPtr6(), 1, 2, 3, 4, 5, 6)

	CallIntPtr(testPtrIntPtr(), nil)
	CallIntPtr2(testPtrIntPtr2(), nil, nil)
	CallIntPtr3(testPtrIntPtr3(), nil, nil, nil)
	CallIntPtr4(testPtrIntPtr4(), nil, nil, nil, nil)
	CallIntPtr5(testPtrIntPtr5(), nil, nil, nil, nil, nil)
	CallIntPtr6(testPtrIntPtr6(), nil, nil, nil, nil, nil, nil)

	CallIntInt(testPtrIntInt(), 1)
	CallIntInt2(testPtrIntInt2(), 1, 2)
	CallIntInt3(testPtrIntInt3(), 1, 2, 3)
	CallIntInt4(testPtrIntInt4(), 1, 2, 3, 4)
	CallIntInt5(testPtrIntInt5(), 1, 2, 3, 4, 5)
	CallIntInt6(testPtrIntInt6(), 1, 2, 3, 4, 5, 6)
}
