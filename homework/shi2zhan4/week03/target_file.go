// Package knast go ast example code
// @author KelipuTe
/* @multiple line1
line2
*/
package week03

import "context"

type (
	// KNFunc is a test func
	// @type func()
	/* @multiple line1
	line2
	*/
	KNFunc func()
)

type (
	// EmptyStruct is a test struct
	// @type empty struct
	EmptyStruct struct {
	}

	// NormalStruct is a test struct
	// @type normal struct
	NormalStruct struct {
		// PublicInt is a test public field
		// @type int
		PublicInt int

		// PublicString is a test public field
		// @type string
		PublicString string
	}
)

type (
	// NormalInterface is a test interface
	// @type normal interface
	/* @multiple line1
	line2
	*/
	NormalInterface interface {
		// funcWithoutInAndOut is a test func
		// @type normal func
		funcWithoutInAndOut()

		// funcWithOnly1In is a test func
		// @type normal func
		// @parameter i int
		funcWithOnly1In(i int)

		// funcWithOnly1Out is a test func
		// @type normal func
		// @return int
		funcWithOnly1Out() int

		// funcWith2InAnd2Out is a test func
		// @type normal func
		// @parameter i int
		// @parameter i64 int64
		// @return int
		funcWith2InAnd2Out(i int, i64 int64) (int, int64)

		// funcWithNamed2Out is a test func
		// @type normal func
		// @return int
		// @return int64
		funcWithNamed2Out() (i int, i64 int64)

		// funcWithCtxInAndErrOut is a test func
		// @type normal func
		// @return context.Context
		// @return error
		funcWithCtxInAndErrOut(ctx context.Context) error
	}
)