package skiplist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetFirstInRange(t *testing.T) {

	skl := Create()

	ret := skl.GetFirstInRange(1)
	assert.Nil(t, ret)

	skl.Insert(1, "1")
	ret = skl.GetFirstInRange(1)
	assert.Equal(t, "1", ret.val)

	skl.Insert(2, "2")
	ret = skl.GetFirstInRange(2)
	assert.Equal(t, "2", ret.val)

	skl.Insert(3, "3")
	ret = skl.GetFirstInRange(3)
	assert.Equal(t, "3", ret.val)

	skl.Insert(4, "4")
	ret = skl.GetFirstInRange(4)
	assert.Equal(t, "4", ret.val)

	skl.Insert(5, "5")
	ret = skl.GetFirstInRange(5)
	assert.Equal(t, "5", ret.val)

	ret = skl.GetFirstInRange(6)
	assert.Nil(t, ret)
}
