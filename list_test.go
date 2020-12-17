package skiplist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Set_Get_Del(t *testing.T) {

	skl := Create()

	val := skl.GetByIndex(0)
	assert.Nil(t, val)
	val = skl.GetByScore(1)
	assert.Nil(t, val)

	skl.Set(1, "1")
	val = skl.GetByScore(1)
	assert.Equal(t, "1", val)

	skl.Set(2, "2")
	val = skl.GetByScore(2)
	assert.Equal(t, "2", val)

	skl.Set(3, "3")
	val = skl.GetByScore(3)
	assert.Equal(t, "3", val)

	skl.Set(4, "4")
	val = skl.GetByScore(4)
	assert.Equal(t, "4", val)

	skl.Set(5, "5")
	val = skl.GetByScore(5)
	assert.Equal(t, "5", val)

	val = skl.GetByScore(6)
	assert.Nil(t, val)

	skl.Set(5, "55")
	val = skl.GetByScore(5)
	assert.Equal(t, "55", val)

	skl.Del(6)
	skl.Del(5)
	val = skl.GetByScore(5)
	assert.Nil(t, val)
}

func Test_Len(t *testing.T) {

	skl := Create()
	len := skl.Len()
	assert.Equal(t, 0, len)

	skl.Set(4, "4")
	len = skl.Len()
	assert.Equal(t, 1, len)

	skl.Set(1, "1")
	len = skl.Len()
	assert.Equal(t, 2, len)

	skl.Set(15, "15")
	len = skl.Len()
	assert.Equal(t, 3, len)

	skl.Del(4)
	len = skl.Len()
	assert.Equal(t, 2, len)

	skl.Del(1)
	len = skl.Len()
	assert.Equal(t, 1, len)

	skl.Del(15)
	len = skl.Len()
	assert.Equal(t, 0, len)

	skl.Del(15)
	len = skl.Len()
	assert.Equal(t, 0, len)
}

func Test_GetByIndex(t *testing.T) {

	skl := Create()

	val := skl.GetByIndex(0)
	assert.Nil(t, val)

	skl.Set(1, "1")
	val = skl.GetByIndex(0)
	assert.Equal(t, "1", val)

	skl.Set(5, "5")
	val = skl.GetByIndex(1)
	assert.Equal(t, "5", val)

	skl.Set(2, "2")
	val = skl.GetByIndex(1)
	assert.Equal(t, "2", val)

	skl.Set(3, "3")
	val = skl.GetByIndex(2)
	assert.Equal(t, "3", val)

	skl.Set(4, "4")
	val = skl.GetByIndex(3)
	assert.Equal(t, "4", val)

	skl.Set(6, "6")
	val = skl.GetByIndex(5)
	assert.Equal(t, "6", val)

	val = skl.GetByIndex(4)
	assert.Equal(t, "5", val)
}
