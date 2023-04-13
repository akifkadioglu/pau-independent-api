// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"pamukkale_university/ent/departments"
	"pamukkale_university/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DepartmentsDelete is the builder for deleting a Departments entity.
type DepartmentsDelete struct {
	config
	hooks    []Hook
	mutation *DepartmentsMutation
}

// Where appends a list predicates to the DepartmentsDelete builder.
func (dd *DepartmentsDelete) Where(ps ...predicate.Departments) *DepartmentsDelete {
	dd.mutation.Where(ps...)
	return dd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dd *DepartmentsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, DepartmentsMutation](ctx, dd.sqlExec, dd.mutation, dd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dd *DepartmentsDelete) ExecX(ctx context.Context) int {
	n, err := dd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dd *DepartmentsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(departments.Table, sqlgraph.NewFieldSpec(departments.FieldID, field.TypeUUID))
	if ps := dd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dd.mutation.done = true
	return affected, err
}

// DepartmentsDeleteOne is the builder for deleting a single Departments entity.
type DepartmentsDeleteOne struct {
	dd *DepartmentsDelete
}

// Where appends a list predicates to the DepartmentsDelete builder.
func (ddo *DepartmentsDeleteOne) Where(ps ...predicate.Departments) *DepartmentsDeleteOne {
	ddo.dd.mutation.Where(ps...)
	return ddo
}

// Exec executes the deletion query.
func (ddo *DepartmentsDeleteOne) Exec(ctx context.Context) error {
	n, err := ddo.dd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{departments.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ddo *DepartmentsDeleteOne) ExecX(ctx context.Context) {
	if err := ddo.Exec(ctx); err != nil {
		panic(err)
	}
}