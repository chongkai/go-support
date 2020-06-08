//
// Copyright 2019 Lightbend Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package eventsourced

type (
	ServiceName string
	EntityId    string
	CommandId   int64
)

func (sn ServiceName) String() string {
	return string(sn)
}

func (id CommandId) Value() int64 {
	return int64(id)
}

// The EntityInstance represents a concrete instance of
// a event sourced entity
type EntityInstance struct {
	// Instance is an instance of the EventSourcedEntity.Entity
	Instance interface{}
	// EventSourcedEntity describes the instance
	EventSourcedEntity *EventSourcedEntity
	eventSequence      int64
}

func (e *EntityInstance) shouldSnapshot() bool {
	return e.eventSequence >= e.EventSourcedEntity.SnapshotEvery
}

func (e *EntityInstance) resetSnapshotEvery() {
	e.eventSequence = 0
}

// A EntityInstanceContext represents a event sourced entity together with its
// associated service.
// Commands are dispatched through this context.
type EntityInstanceContext struct { // TODO: EntityInstanceContext might be actually a EntityInstance
	context *Context
	// EntityInstance is the entity instance of this context
	EntityInstance *EntityInstance
	// active indicates if this context is active
	active bool // TODO: inactivate a context in case of errors
}

// ServiceName returns the contexts service name.
func (c EntityInstanceContext) ServiceName() string {
	return c.EntityInstance.EventSourcedEntity.ServiceName.String()
}