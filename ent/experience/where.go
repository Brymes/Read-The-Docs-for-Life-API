// Code generated by entc, DO NOT EDIT.

package experience

import (
	"LifeDocs/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v uuid.UUID) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUID), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// Views applies equality check predicate on the "views" field. It's identical to ViewsEQ.
func Views(v int) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldViews), v))
	})
}

// Likes applies equality check predicate on the "likes" field. It's identical to LikesEQ.
func Likes(v int) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLikes), v))
	})
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v uuid.UUID) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUID), v))
	})
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v uuid.UUID) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUUID), v))
	})
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...uuid.UUID) predicate.Experience {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Experience(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUUID), v...))
	})
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...uuid.UUID) predicate.Experience {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Experience(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUUID), v...))
	})
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v uuid.UUID) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUUID), v))
	})
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v uuid.UUID) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUUID), v))
	})
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v uuid.UUID) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUUID), v))
	})
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v uuid.UUID) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUUID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Experience {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Experience(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Experience {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Experience(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Experience {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Experience(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Experience {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Experience(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Experience {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Experience(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Experience {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Experience(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContent), v))
	})
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContent), v))
	})
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContent), v))
	})
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContent), v))
	})
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContent), v))
	})
}

// ViewsEQ applies the EQ predicate on the "views" field.
func ViewsEQ(v int) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldViews), v))
	})
}

// ViewsNEQ applies the NEQ predicate on the "views" field.
func ViewsNEQ(v int) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldViews), v))
	})
}

// ViewsIn applies the In predicate on the "views" field.
func ViewsIn(vs ...int) predicate.Experience {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Experience(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldViews), v...))
	})
}

// ViewsNotIn applies the NotIn predicate on the "views" field.
func ViewsNotIn(vs ...int) predicate.Experience {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Experience(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldViews), v...))
	})
}

// ViewsGT applies the GT predicate on the "views" field.
func ViewsGT(v int) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldViews), v))
	})
}

// ViewsGTE applies the GTE predicate on the "views" field.
func ViewsGTE(v int) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldViews), v))
	})
}

// ViewsLT applies the LT predicate on the "views" field.
func ViewsLT(v int) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldViews), v))
	})
}

// ViewsLTE applies the LTE predicate on the "views" field.
func ViewsLTE(v int) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldViews), v))
	})
}

// LikesEQ applies the EQ predicate on the "likes" field.
func LikesEQ(v int) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLikes), v))
	})
}

// LikesNEQ applies the NEQ predicate on the "likes" field.
func LikesNEQ(v int) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLikes), v))
	})
}

// LikesIn applies the In predicate on the "likes" field.
func LikesIn(vs ...int) predicate.Experience {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Experience(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLikes), v...))
	})
}

// LikesNotIn applies the NotIn predicate on the "likes" field.
func LikesNotIn(vs ...int) predicate.Experience {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Experience(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLikes), v...))
	})
}

// LikesGT applies the GT predicate on the "likes" field.
func LikesGT(v int) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLikes), v))
	})
}

// LikesGTE applies the GTE predicate on the "likes" field.
func LikesGTE(v int) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLikes), v))
	})
}

// LikesLT applies the LT predicate on the "likes" field.
func LikesLT(v int) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLikes), v))
	})
}

// LikesLTE applies the LTE predicate on the "likes" field.
func LikesLTE(v int) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLikes), v))
	})
}

// HasComments applies the HasEdge predicate on the "comments" edge.
func HasComments() predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CommentsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentsWith applies the HasEdge predicate on the "comments" edge with a given conditions (other predicates).
func HasCommentsWith(preds ...predicate.Comments) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CommentsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Experience) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Experience) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Experience) predicate.Experience {
	return predicate.Experience(func(s *sql.Selector) {
		p(s.Not())
	})
}
