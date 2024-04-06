package test

import (
	"github.com/aivyss/perr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPreparedErrorSwitch(t *testing.T) {
	//  given
	isInvoked := false
	pErr := perr.NewPreparedError("origin")
	differentPErr1 := perr.NewPreparedError("different origin - 1")
	differentPErr2 := perr.NewPreparedError("different origin - 2")
	modifiedPErr := pErr.AddMessage("my-message")

	// when
	t.Run("different origin", func(t *testing.T) {
		perr.Switch(modifiedPErr).Case(differentPErr1, func() {
			assert.Fail(t, "should not run")
		})
		perr.Switch(modifiedPErr).Case(differentPErr2, func() {
			assert.Fail(t, "should not run")
		})
	})

	t.Run("same origin", func(t *testing.T) {
		perr.Switch(modifiedPErr).Case(pErr, func() {
			isInvoked = true
		}).Case(differentPErr1, func() {
			assert.Fail(t, "should not run")
		})
		assert.True(t, isInvoked)
		isInvoked = false

		perr.Switch(modifiedPErr).Case(differentPErr1, func() {
			assert.Fail(t, "should not run")
		}).Case(pErr, func() {
			isInvoked = true
		}).Case(pErr, func() {
			assert.Fail(t, "should not run")
		})
		assert.True(t, isInvoked)
		isInvoked = false
	})

	t.Run("group", func(t *testing.T) {
		perr.Switch(modifiedPErr).CaseGroup(perr.ErrorGroup(differentPErr1, differentPErr2), func() {
			assert.Fail(t, "should not run")
		}).CaseGroup(perr.ErrorGroup(pErr), func() {
			isInvoked = true
		}).CaseGroup(perr.ErrorGroup(modifiedPErr), func() {
			assert.Fail(t, "should not run")
		})

		assert.True(t, isInvoked)
		isInvoked = false
	})

	t.Run("mix", func(t *testing.T) {
		perr.Switch(modifiedPErr).CaseGroup(perr.ErrorGroup(differentPErr1), func() {
			assert.Fail(t, "should not run")
		}).Case(differentPErr2, func() {
			assert.Fail(t, "should not run")
		}).Case(pErr, func() {
			isInvoked = true
		}).Case(modifiedPErr, func() {
			assert.Fail(t, "should not run")
		}).CaseGroup(perr.ErrorGroup(differentPErr2), func() {
			assert.Fail(t, "should not run")
		})

		assert.True(t, isInvoked)
		isInvoked = false
	})
}
