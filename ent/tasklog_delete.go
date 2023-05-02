// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-job/ent/predicate"
	"github.com/suyuan32/simple-admin-job/ent/tasklog"
)

// TaskLogDelete is the builder for deleting a TaskLog entity.
type TaskLogDelete struct {
	config
	hooks    []Hook
	mutation *TaskLogMutation
}

// Where appends a list predicates to the TaskLogDelete builder.
func (tld *TaskLogDelete) Where(ps ...predicate.TaskLog) *TaskLogDelete {
	tld.mutation.Where(ps...)
	return tld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tld *TaskLogDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tld.sqlExec, tld.mutation, tld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tld *TaskLogDelete) ExecX(ctx context.Context) int {
	n, err := tld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tld *TaskLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(tasklog.Table, sqlgraph.NewFieldSpec(tasklog.FieldID, field.TypeUint64))
	if ps := tld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tld.mutation.done = true
	return affected, err
}

// TaskLogDeleteOne is the builder for deleting a single TaskLog entity.
type TaskLogDeleteOne struct {
	tld *TaskLogDelete
}

// Where appends a list predicates to the TaskLogDelete builder.
func (tldo *TaskLogDeleteOne) Where(ps ...predicate.TaskLog) *TaskLogDeleteOne {
	tldo.tld.mutation.Where(ps...)
	return tldo
}

// Exec executes the deletion query.
func (tldo *TaskLogDeleteOne) Exec(ctx context.Context) error {
	n, err := tldo.tld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tasklog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tldo *TaskLogDeleteOne) ExecX(ctx context.Context) {
	if err := tldo.Exec(ctx); err != nil {
		panic(err)
	}
}
