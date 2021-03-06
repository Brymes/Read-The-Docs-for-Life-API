// Code generated by entc, DO NOT EDIT.

package ent

import (
	"LifeDocs/ent/comments"
	"LifeDocs/ent/experience"
	"LifeDocs/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CommentsCreate is the builder for creating a Comments entity.
type CommentsCreate struct {
	config
	mutation *CommentsMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (cc *CommentsCreate) SetUUID(u uuid.UUID) *CommentsCreate {
	cc.mutation.SetUUID(u)
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *CommentsCreate) SetCreatedAt(t time.Time) *CommentsCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CommentsCreate) SetNillableCreatedAt(t *time.Time) *CommentsCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetComment sets the "comment" field.
func (cc *CommentsCreate) SetComment(s string) *CommentsCreate {
	cc.mutation.SetComment(s)
	return cc
}

// SetLikes sets the "likes" field.
func (cc *CommentsCreate) SetLikes(i int) *CommentsCreate {
	cc.mutation.SetLikes(i)
	return cc
}

// SetPostsID sets the "posts" edge to the Experience entity by ID.
func (cc *CommentsCreate) SetPostsID(id int) *CommentsCreate {
	cc.mutation.SetPostsID(id)
	return cc
}

// SetNillablePostsID sets the "posts" edge to the Experience entity by ID if the given value is not nil.
func (cc *CommentsCreate) SetNillablePostsID(id *int) *CommentsCreate {
	if id != nil {
		cc = cc.SetPostsID(*id)
	}
	return cc
}

// SetPosts sets the "posts" edge to the Experience entity.
func (cc *CommentsCreate) SetPosts(e *Experience) *CommentsCreate {
	return cc.SetPostsID(e.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cc *CommentsCreate) SetUserID(id int) *CommentsCreate {
	cc.mutation.SetUserID(id)
	return cc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (cc *CommentsCreate) SetNillableUserID(id *int) *CommentsCreate {
	if id != nil {
		cc = cc.SetUserID(*id)
	}
	return cc
}

// SetUser sets the "user" edge to the User entity.
func (cc *CommentsCreate) SetUser(u *User) *CommentsCreate {
	return cc.SetUserID(u.ID)
}

// Mutation returns the CommentsMutation object of the builder.
func (cc *CommentsCreate) Mutation() *CommentsMutation {
	return cc.mutation
}

// Save creates the Comments in the database.
func (cc *CommentsCreate) Save(ctx context.Context) (*Comments, error) {
	var (
		err  error
		node *Comments
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CommentsCreate) SaveX(ctx context.Context) *Comments {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CommentsCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CommentsCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CommentsCreate) defaults() {
	if _, ok := cc.mutation.UUID(); !ok {
		v := comments.DefaultUUID()
		cc.mutation.SetUUID(v)
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := comments.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CommentsCreate) check() error {
	if _, ok := cc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "uuid"`)}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := cc.mutation.Comment(); !ok {
		return &ValidationError{Name: "comment", err: errors.New(`ent: missing required field "comment"`)}
	}
	if v, ok := cc.mutation.Comment(); ok {
		if err := comments.CommentValidator(v); err != nil {
			return &ValidationError{Name: "comment", err: fmt.Errorf(`ent: validator failed for field "comment": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Likes(); !ok {
		return &ValidationError{Name: "likes", err: errors.New(`ent: missing required field "likes"`)}
	}
	if v, ok := cc.mutation.Likes(); ok {
		if err := comments.LikesValidator(v); err != nil {
			return &ValidationError{Name: "likes", err: fmt.Errorf(`ent: validator failed for field "likes": %w`, err)}
		}
	}
	return nil
}

func (cc *CommentsCreate) sqlSave(ctx context.Context) (*Comments, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CommentsCreate) createSpec() (*Comments, *sqlgraph.CreateSpec) {
	var (
		_node = &Comments{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: comments.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: comments.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: comments.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: comments.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.Comment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comments.FieldComment,
		})
		_node.Comment = value
	}
	if value, ok := cc.mutation.Likes(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: comments.FieldLikes,
		})
		_node.Likes = value
	}
	if nodes := cc.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comments.PostsTable,
			Columns: []string{comments.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: experience.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.experience_comments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comments.UserTable,
			Columns: []string{comments.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_comments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CommentsCreateBulk is the builder for creating many Comments entities in bulk.
type CommentsCreateBulk struct {
	config
	builders []*CommentsCreate
}

// Save creates the Comments entities in the database.
func (ccb *CommentsCreateBulk) Save(ctx context.Context) ([]*Comments, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Comments, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommentsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CommentsCreateBulk) SaveX(ctx context.Context) []*Comments {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CommentsCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CommentsCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
