// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/iot-synergy/synergy-message-center/ent/predicate"
	"github.com/iot-synergy/synergy-message-center/ent/smsprovider"
)

// SmsProviderDelete is the builder for deleting a SmsProvider entity.
type SmsProviderDelete struct {
	config
	hooks    []Hook
	mutation *SmsProviderMutation
}

// Where appends a list predicates to the SmsProviderDelete builder.
func (spd *SmsProviderDelete) Where(ps ...predicate.SmsProvider) *SmsProviderDelete {
	spd.mutation.Where(ps...)
	return spd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (spd *SmsProviderDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, spd.sqlExec, spd.mutation, spd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (spd *SmsProviderDelete) ExecX(ctx context.Context) int {
	n, err := spd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (spd *SmsProviderDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(smsprovider.Table, sqlgraph.NewFieldSpec(smsprovider.FieldID, field.TypeUint64))
	if ps := spd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, spd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	spd.mutation.done = true
	return affected, err
}

// SmsProviderDeleteOne is the builder for deleting a single SmsProvider entity.
type SmsProviderDeleteOne struct {
	spd *SmsProviderDelete
}

// Where appends a list predicates to the SmsProviderDelete builder.
func (spdo *SmsProviderDeleteOne) Where(ps ...predicate.SmsProvider) *SmsProviderDeleteOne {
	spdo.spd.mutation.Where(ps...)
	return spdo
}

// Exec executes the deletion query.
func (spdo *SmsProviderDeleteOne) Exec(ctx context.Context) error {
	n, err := spdo.spd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{smsprovider.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (spdo *SmsProviderDeleteOne) ExecX(ctx context.Context) {
	if err := spdo.Exec(ctx); err != nil {
		panic(err)
	}
}
