/*
Copyright 2017 Ahmed Zaher

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mathex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRound(t *testing.T) {

	var precision float64

	assert.InDelta(t, 0.00000001, Round(0.000000012543275, 8, 0.5), precision)
	assert.InDelta(t, 0.00000002, Round(0.000000018543275, 8, 0.5), precision)
	assert.InDelta(t, 1.00000001, Round(1.000000012543275, 8, 0.5), precision)
	assert.InDelta(t, 1.00000002, Round(1.000000018543275, 8, 0.5), precision)
	assert.InDelta(t, 1.65326, Round(1.653264418543275, 5, 0.5), precision)
	assert.InDelta(t, 1.65327, Round(1.653265718543275, 5, 0.5), precision)
	assert.InDelta(t, 0.6, Round(0.62, 1, 0.5), precision)
	assert.InDelta(t, 0.7, Round(0.62, 1, 0.2), precision)
	assert.InDelta(t, 1, Round(0.62, 0, 0.5), precision)
	assert.InDelta(t, 90, Round(92, -1, 0.5), precision)
	assert.InDelta(t, 100, Round(92, -1, 0.1), precision)
}
