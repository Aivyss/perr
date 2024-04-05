package test

import (
	"fmt"
	pErrPkg "github.com/aivyss/perr"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPreparedError(t *testing.T) {
	pErr := pErrPkg.NewPreparedError("test-error")
	differentPErr := pErrPkg.NewPreparedError("different-test-error")
	modifiedPErr := pErr.AddMessage("test-message")

	t.Run("same", func(t *testing.T) {
		assert.True(t, pErr.Same(modifiedPErr))
		assert.True(t, modifiedPErr.Same(pErr))
		assert.False(t, differentPErr.Same(pErr))
		assert.False(t, differentPErr.Same(modifiedPErr))

		wrappedModifiedPErrOneTime := fmt.Errorf("wrap 1: %w", modifiedPErr)
		wrappedModifiedPErrTwoTime := fmt.Errorf("wrap 2: %w", wrappedModifiedPErrOneTime)

		assert.True(t, pErr.Same(wrappedModifiedPErrOneTime))
		assert.True(t, pErr.Same(wrappedModifiedPErrTwoTime))
		assert.True(t, modifiedPErr.Same(wrappedModifiedPErrOneTime))
		assert.True(t, modifiedPErr.Same(wrappedModifiedPErrTwoTime))
		assert.False(t, differentPErr.Same(wrappedModifiedPErrOneTime))
		assert.False(t, differentPErr.Same(wrappedModifiedPErrTwoTime))

		errs := make([]error, 0, 100)
		for i := 0; i < 100; i++ {
			errs = append(errs, pErr.AddMessage("times: %d", i+1))
		}

		for _, err := range errs {
			assert.True(t, pErr.Same(err))
			assert.True(t, modifiedPErr.Same(err))
			assert.False(t, differentPErr.Same(err))
		}
	})

	t.Run("in", func(t *testing.T) {
		assert.True(t, pErr.In(differentPErr, modifiedPErr))
		assert.True(t, modifiedPErr.In(differentPErr, pErr))
		assert.False(t, differentPErr.In(pErr, modifiedPErr))
	})

	t.Run("contains", func(t *testing.T) {
		a := pErrPkg.NewPreparedError("a")
		b := pErrPkg.NewPreparedError("b")
		c := pErrPkg.NewPreparedError("c")

		abcGroup := pErrPkg.ErrorGroup(a, b, c)
		abGroup := pErrPkg.ErrorGroup(a, b)
		aGroup := pErrPkg.ErrorGroup(a)
		emptyGroup := pErrPkg.ErrorGroup()

		assert.True(t, abcGroup.Contains(a))
		assert.True(t, abcGroup.Contains(b))
		assert.True(t, abcGroup.Contains(c))

		assert.True(t, abGroup.Contains(a))
		assert.True(t, abGroup.Contains(b))
		assert.False(t, abGroup.Contains(c))

		assert.True(t, aGroup.Contains(a))
		assert.False(t, aGroup.Contains(b))
		assert.False(t, aGroup.Contains(c))

		assert.False(t, emptyGroup.Contains(a))
		assert.False(t, emptyGroup.Contains(b))
		assert.False(t, emptyGroup.Contains(c))
	})

	t.Run("add message", func(t *testing.T) {
		original := pErrPkg.NewPreparedError("original")

		firstMessage := original.AddMessage("first-message")
		secondMessage := firstMessage.AddMessage("second-message")

		fmt.Println(original.Error())
		assert.True(t, strings.Contains(original.Error(), "original"))
		fmt.Println(firstMessage.Error())
		assert.True(t, strings.Contains(firstMessage.Error(), "original\n[first-message]\n[stacks]"))
		fmt.Println(secondMessage.Error())
		assert.True(t, strings.Contains(secondMessage.Error(), "original\n[first-message second-message]\n[stacks]"))
	})
}
