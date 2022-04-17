// Copyright (c) 2018 The Jaeger Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mocks

import (
	sarama "github.com/Shopify/sarama"
	mock "github.com/stretchr/testify/mock"
)

// PartitionConsumer is an autogenerated mock type for the PartitionConsumer type
type PartitionConsumer struct {
	mock.Mock
}

// AsyncClose provides a mock function with given fields:
func (_m *PartitionConsumer) AsyncClose() {
	_m.Called()
}

// Close provides a mock function with given fields:
func (_m *PartitionConsumer) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Errors provides a mock function with given fields:
func (_m *PartitionConsumer) Errors() <-chan *sarama.ConsumerError {
	ret := _m.Called()

	var r0 <-chan *sarama.ConsumerError
	if rf, ok := ret.Get(0).(func() <-chan *sarama.ConsumerError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *sarama.ConsumerError)
		}
	}

	return r0
}

// HighWaterMarkOffset provides a mock function with given fields:
func (_m *PartitionConsumer) HighWaterMarkOffset() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Messages provides a mock function with given fields:
func (_m *PartitionConsumer) Messages() <-chan *sarama.ConsumerMessage {
	ret := _m.Called()

	var r0 <-chan *sarama.ConsumerMessage
	if rf, ok := ret.Get(0).(func() <-chan *sarama.ConsumerMessage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *sarama.ConsumerMessage)
		}
	}

	return r0
}

// Partition provides a mock function with given fields:
func (_m *PartitionConsumer) Partition() int32 {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// Topic provides a mock function with given fields:
func (_m *PartitionConsumer) Topic() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NB (yurishkuro): Added the following functions manually as Makefile doesn't show how to auto-gen it.

// Pause mocked.
func (_m *PartitionConsumer) Pause() {
	_m.Called()
}

// Resume mocked.
func (_m *PartitionConsumer) Resume() {
	_m.Called()
}

// IsPaused mocked.
func (_m *PartitionConsumer) IsPaused() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
