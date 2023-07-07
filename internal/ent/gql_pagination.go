// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/ecshreve/jepp/internal/ent/category"
	"github.com/ecshreve/jepp/internal/ent/clue"
	"github.com/ecshreve/jepp/internal/ent/game"
	"github.com/ecshreve/jepp/internal/ent/season"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Common entgql types.
type (
	Cursor         = entgql.Cursor[int]
	PageInfo       = entgql.PageInfo[int]
	OrderDirection = entgql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == entgql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// CategoryEdge is the edge representation of Category.
type CategoryEdge struct {
	Node   *Category `json:"node"`
	Cursor Cursor    `json:"cursor"`
}

// CategoryConnection is the connection containing edges to Category.
type CategoryConnection struct {
	Edges      []*CategoryEdge `json:"edges"`
	PageInfo   PageInfo        `json:"pageInfo"`
	TotalCount int             `json:"totalCount"`
}

func (c *CategoryConnection) build(nodes []*Category, pager *categoryPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Category
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Category {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Category {
			return nodes[i]
		}
	}
	c.Edges = make([]*CategoryEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &CategoryEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// CategoryPaginateOption enables pagination customization.
type CategoryPaginateOption func(*categoryPager) error

// WithCategoryOrder configures pagination ordering.
func WithCategoryOrder(order *CategoryOrder) CategoryPaginateOption {
	if order == nil {
		order = DefaultCategoryOrder
	}
	o := *order
	return func(pager *categoryPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultCategoryOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithCategoryFilter configures pagination filter.
func WithCategoryFilter(filter func(*CategoryQuery) (*CategoryQuery, error)) CategoryPaginateOption {
	return func(pager *categoryPager) error {
		if filter == nil {
			return errors.New("CategoryQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type categoryPager struct {
	reverse bool
	order   *CategoryOrder
	filter  func(*CategoryQuery) (*CategoryQuery, error)
}

func newCategoryPager(opts []CategoryPaginateOption, reverse bool) (*categoryPager, error) {
	pager := &categoryPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultCategoryOrder
	}
	return pager, nil
}

func (p *categoryPager) applyFilter(query *CategoryQuery) (*CategoryQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *categoryPager) toCursor(c *Category) Cursor {
	return p.order.Field.toCursor(c)
}

func (p *categoryPager) applyCursors(query *CategoryQuery, after, before *Cursor) (*CategoryQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultCategoryOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *categoryPager) applyOrder(query *CategoryQuery) *CategoryQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultCategoryOrder.Field {
		query = query.Order(DefaultCategoryOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *categoryPager) orderExpr(query *CategoryQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultCategoryOrder.Field {
			b.Comma().Ident(DefaultCategoryOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Category.
func (c *CategoryQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...CategoryPaginateOption,
) (*CategoryConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newCategoryPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if c, err = pager.applyFilter(c); err != nil {
		return nil, err
	}
	conn := &CategoryConnection{Edges: []*CategoryEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := c.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if c, err = pager.applyCursors(c, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		c.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := c.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	c = pager.applyOrder(c)
	nodes, err := c.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// CategoryOrderFieldID orders Category by id.
	CategoryOrderFieldID = &CategoryOrderField{
		Value: func(c *Category) (ent.Value, error) {
			return c.ID, nil
		},
		column: category.FieldID,
		toTerm: category.ByID,
		toCursor: func(c *Category) Cursor {
			return Cursor{
				ID:    c.ID,
				Value: c.ID,
			}
		},
	}
	// CategoryOrderFieldName orders Category by name.
	CategoryOrderFieldName = &CategoryOrderField{
		Value: func(c *Category) (ent.Value, error) {
			return c.Name, nil
		},
		column: category.FieldName,
		toTerm: category.ByName,
		toCursor: func(c *Category) Cursor {
			return Cursor{
				ID:    c.ID,
				Value: c.Name,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f CategoryOrderField) String() string {
	var str string
	switch f.column {
	case CategoryOrderFieldID.column:
		str = "ID"
	case CategoryOrderFieldName.column:
		str = "NAME"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f CategoryOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *CategoryOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("CategoryOrderField %T must be a string", v)
	}
	switch str {
	case "ID":
		*f = *CategoryOrderFieldID
	case "NAME":
		*f = *CategoryOrderFieldName
	default:
		return fmt.Errorf("%s is not a valid CategoryOrderField", str)
	}
	return nil
}

// CategoryOrderField defines the ordering field of Category.
type CategoryOrderField struct {
	// Value extracts the ordering value from the given Category.
	Value    func(*Category) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) category.OrderOption
	toCursor func(*Category) Cursor
}

// CategoryOrder defines the ordering of Category.
type CategoryOrder struct {
	Direction OrderDirection      `json:"direction"`
	Field     *CategoryOrderField `json:"field"`
}

// DefaultCategoryOrder is the default ordering of Category.
var DefaultCategoryOrder = &CategoryOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &CategoryOrderField{
		Value: func(c *Category) (ent.Value, error) {
			return c.ID, nil
		},
		column: category.FieldID,
		toTerm: category.ByID,
		toCursor: func(c *Category) Cursor {
			return Cursor{ID: c.ID}
		},
	},
}

// ToEdge converts Category into CategoryEdge.
func (c *Category) ToEdge(order *CategoryOrder) *CategoryEdge {
	if order == nil {
		order = DefaultCategoryOrder
	}
	return &CategoryEdge{
		Node:   c,
		Cursor: order.Field.toCursor(c),
	}
}

// ClueEdge is the edge representation of Clue.
type ClueEdge struct {
	Node   *Clue  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// ClueConnection is the connection containing edges to Clue.
type ClueConnection struct {
	Edges      []*ClueEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *ClueConnection) build(nodes []*Clue, pager *cluePager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Clue
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Clue {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Clue {
			return nodes[i]
		}
	}
	c.Edges = make([]*ClueEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &ClueEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// CluePaginateOption enables pagination customization.
type CluePaginateOption func(*cluePager) error

// WithClueOrder configures pagination ordering.
func WithClueOrder(order *ClueOrder) CluePaginateOption {
	if order == nil {
		order = DefaultClueOrder
	}
	o := *order
	return func(pager *cluePager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultClueOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithClueFilter configures pagination filter.
func WithClueFilter(filter func(*ClueQuery) (*ClueQuery, error)) CluePaginateOption {
	return func(pager *cluePager) error {
		if filter == nil {
			return errors.New("ClueQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type cluePager struct {
	reverse bool
	order   *ClueOrder
	filter  func(*ClueQuery) (*ClueQuery, error)
}

func newCluePager(opts []CluePaginateOption, reverse bool) (*cluePager, error) {
	pager := &cluePager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultClueOrder
	}
	return pager, nil
}

func (p *cluePager) applyFilter(query *ClueQuery) (*ClueQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *cluePager) toCursor(c *Clue) Cursor {
	return p.order.Field.toCursor(c)
}

func (p *cluePager) applyCursors(query *ClueQuery, after, before *Cursor) (*ClueQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultClueOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *cluePager) applyOrder(query *ClueQuery) *ClueQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultClueOrder.Field {
		query = query.Order(DefaultClueOrder.Field.toTerm(direction.OrderTermOption()))
	}
	switch p.order.Field.column {
	case CategoryOrderFieldCategoryName.column:
	default:
		if len(query.ctx.Fields) > 0 {
			query.ctx.AppendFieldOnce(p.order.Field.column)
		}
	}
	return query
}

func (p *cluePager) orderExpr(query *ClueQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	switch p.order.Field.column {
	case CategoryOrderFieldCategoryName.column:
		query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	default:
		if len(query.ctx.Fields) > 0 {
			query.ctx.AppendFieldOnce(p.order.Field.column)
		}
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultClueOrder.Field {
			b.Comma().Ident(DefaultClueOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Clue.
func (c *ClueQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...CluePaginateOption,
) (*ClueConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newCluePager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if c, err = pager.applyFilter(c); err != nil {
		return nil, err
	}
	conn := &ClueConnection{Edges: []*ClueEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := c.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if c, err = pager.applyCursors(c, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		c.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := c.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	c = pager.applyOrder(c)
	nodes, err := c.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// ClueOrderFieldID orders Clue by id.
	ClueOrderFieldID = &ClueOrderField{
		Value: func(c *Clue) (ent.Value, error) {
			return c.ID, nil
		},
		column: clue.FieldID,
		toTerm: clue.ByID,
		toCursor: func(c *Clue) Cursor {
			return Cursor{
				ID:    c.ID,
				Value: c.ID,
			}
		},
	}
	// ClueOrderFieldQuestion orders Clue by question.
	ClueOrderFieldQuestion = &ClueOrderField{
		Value: func(c *Clue) (ent.Value, error) {
			return c.Question, nil
		},
		column: clue.FieldQuestion,
		toTerm: clue.ByQuestion,
		toCursor: func(c *Clue) Cursor {
			return Cursor{
				ID:    c.ID,
				Value: c.Question,
			}
		},
	}
	// ClueOrderFieldAnswer orders Clue by answer.
	ClueOrderFieldAnswer = &ClueOrderField{
		Value: func(c *Clue) (ent.Value, error) {
			return c.Answer, nil
		},
		column: clue.FieldAnswer,
		toTerm: clue.ByAnswer,
		toCursor: func(c *Clue) Cursor {
			return Cursor{
				ID:    c.ID,
				Value: c.Answer,
			}
		},
	}
	// CategoryOrderFieldCategoryName orders by CATEGORY_NAME.
	CategoryOrderFieldCategoryName = &ClueOrderField{
		Value: func(c *Clue) (ent.Value, error) {
			return c.Value("category_name")
		},
		column: "category_name",
		toTerm: func(opts ...sql.OrderTermOption) clue.OrderOption {
			return clue.ByCategoryField(
				category.FieldName,
				append(opts, sql.OrderSelectAs("category_name"))...,
			)
		},
		toCursor: func(c *Clue) Cursor {
			cv, _ := c.Value("category_name")
			return Cursor{
				ID:    c.ID,
				Value: cv,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f ClueOrderField) String() string {
	var str string
	switch f.column {
	case ClueOrderFieldID.column:
		str = "ID"
	case ClueOrderFieldQuestion.column:
		str = "QUESTION"
	case ClueOrderFieldAnswer.column:
		str = "ANSWER"
	case CategoryOrderFieldCategoryName.column:
		str = "CATEGORY_NAME"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f ClueOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *ClueOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("ClueOrderField %T must be a string", v)
	}
	switch str {
	case "ID":
		*f = *ClueOrderFieldID
	case "QUESTION":
		*f = *ClueOrderFieldQuestion
	case "ANSWER":
		*f = *ClueOrderFieldAnswer
	case "CATEGORY_NAME":
		*f = *CategoryOrderFieldCategoryName
	default:
		return fmt.Errorf("%s is not a valid ClueOrderField", str)
	}
	return nil
}

// ClueOrderField defines the ordering field of Clue.
type ClueOrderField struct {
	// Value extracts the ordering value from the given Clue.
	Value    func(*Clue) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) clue.OrderOption
	toCursor func(*Clue) Cursor
}

// ClueOrder defines the ordering of Clue.
type ClueOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *ClueOrderField `json:"field"`
}

// DefaultClueOrder is the default ordering of Clue.
var DefaultClueOrder = &ClueOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &ClueOrderField{
		Value: func(c *Clue) (ent.Value, error) {
			return c.ID, nil
		},
		column: clue.FieldID,
		toTerm: clue.ByID,
		toCursor: func(c *Clue) Cursor {
			return Cursor{ID: c.ID}
		},
	},
}

// ToEdge converts Clue into ClueEdge.
func (c *Clue) ToEdge(order *ClueOrder) *ClueEdge {
	if order == nil {
		order = DefaultClueOrder
	}
	return &ClueEdge{
		Node:   c,
		Cursor: order.Field.toCursor(c),
	}
}

// GameEdge is the edge representation of Game.
type GameEdge struct {
	Node   *Game  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// GameConnection is the connection containing edges to Game.
type GameConnection struct {
	Edges      []*GameEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *GameConnection) build(nodes []*Game, pager *gamePager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Game
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Game {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Game {
			return nodes[i]
		}
	}
	c.Edges = make([]*GameEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &GameEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// GamePaginateOption enables pagination customization.
type GamePaginateOption func(*gamePager) error

// WithGameOrder configures pagination ordering.
func WithGameOrder(order *GameOrder) GamePaginateOption {
	if order == nil {
		order = DefaultGameOrder
	}
	o := *order
	return func(pager *gamePager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultGameOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithGameFilter configures pagination filter.
func WithGameFilter(filter func(*GameQuery) (*GameQuery, error)) GamePaginateOption {
	return func(pager *gamePager) error {
		if filter == nil {
			return errors.New("GameQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type gamePager struct {
	reverse bool
	order   *GameOrder
	filter  func(*GameQuery) (*GameQuery, error)
}

func newGamePager(opts []GamePaginateOption, reverse bool) (*gamePager, error) {
	pager := &gamePager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultGameOrder
	}
	return pager, nil
}

func (p *gamePager) applyFilter(query *GameQuery) (*GameQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *gamePager) toCursor(ga *Game) Cursor {
	return p.order.Field.toCursor(ga)
}

func (p *gamePager) applyCursors(query *GameQuery, after, before *Cursor) (*GameQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultGameOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *gamePager) applyOrder(query *GameQuery) *GameQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultGameOrder.Field {
		query = query.Order(DefaultGameOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *gamePager) orderExpr(query *GameQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultGameOrder.Field {
			b.Comma().Ident(DefaultGameOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Game.
func (ga *GameQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...GamePaginateOption,
) (*GameConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newGamePager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if ga, err = pager.applyFilter(ga); err != nil {
		return nil, err
	}
	conn := &GameConnection{Edges: []*GameEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := ga.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if ga, err = pager.applyCursors(ga, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		ga.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := ga.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	ga = pager.applyOrder(ga)
	nodes, err := ga.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// GameOrderFieldID orders Game by id.
	GameOrderFieldID = &GameOrderField{
		Value: func(ga *Game) (ent.Value, error) {
			return ga.ID, nil
		},
		column: game.FieldID,
		toTerm: game.ByID,
		toCursor: func(ga *Game) Cursor {
			return Cursor{
				ID:    ga.ID,
				Value: ga.ID,
			}
		},
	}
	// GameOrderFieldShow orders Game by show.
	GameOrderFieldShow = &GameOrderField{
		Value: func(ga *Game) (ent.Value, error) {
			return ga.Show, nil
		},
		column: game.FieldShow,
		toTerm: game.ByShow,
		toCursor: func(ga *Game) Cursor {
			return Cursor{
				ID:    ga.ID,
				Value: ga.Show,
			}
		},
	}
	// GameOrderFieldAirDate orders Game by airDate.
	GameOrderFieldAirDate = &GameOrderField{
		Value: func(ga *Game) (ent.Value, error) {
			return ga.AirDate, nil
		},
		column: game.FieldAirDate,
		toTerm: game.ByAirDate,
		toCursor: func(ga *Game) Cursor {
			return Cursor{
				ID:    ga.ID,
				Value: ga.AirDate,
			}
		},
	}
	// GameOrderFieldTapeDate orders Game by tapeDate.
	GameOrderFieldTapeDate = &GameOrderField{
		Value: func(ga *Game) (ent.Value, error) {
			return ga.TapeDate, nil
		},
		column: game.FieldTapeDate,
		toTerm: game.ByTapeDate,
		toCursor: func(ga *Game) Cursor {
			return Cursor{
				ID:    ga.ID,
				Value: ga.TapeDate,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f GameOrderField) String() string {
	var str string
	switch f.column {
	case GameOrderFieldID.column:
		str = "ID"
	case GameOrderFieldShow.column:
		str = "SHOW"
	case GameOrderFieldAirDate.column:
		str = "AIR_DATE"
	case GameOrderFieldTapeDate.column:
		str = "TAPE_DATE"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f GameOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *GameOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("GameOrderField %T must be a string", v)
	}
	switch str {
	case "ID":
		*f = *GameOrderFieldID
	case "SHOW":
		*f = *GameOrderFieldShow
	case "AIR_DATE":
		*f = *GameOrderFieldAirDate
	case "TAPE_DATE":
		*f = *GameOrderFieldTapeDate
	default:
		return fmt.Errorf("%s is not a valid GameOrderField", str)
	}
	return nil
}

// GameOrderField defines the ordering field of Game.
type GameOrderField struct {
	// Value extracts the ordering value from the given Game.
	Value    func(*Game) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) game.OrderOption
	toCursor func(*Game) Cursor
}

// GameOrder defines the ordering of Game.
type GameOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *GameOrderField `json:"field"`
}

// DefaultGameOrder is the default ordering of Game.
var DefaultGameOrder = &GameOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &GameOrderField{
		Value: func(ga *Game) (ent.Value, error) {
			return ga.ID, nil
		},
		column: game.FieldID,
		toTerm: game.ByID,
		toCursor: func(ga *Game) Cursor {
			return Cursor{ID: ga.ID}
		},
	},
}

// ToEdge converts Game into GameEdge.
func (ga *Game) ToEdge(order *GameOrder) *GameEdge {
	if order == nil {
		order = DefaultGameOrder
	}
	return &GameEdge{
		Node:   ga,
		Cursor: order.Field.toCursor(ga),
	}
}

// SeasonEdge is the edge representation of Season.
type SeasonEdge struct {
	Node   *Season `json:"node"`
	Cursor Cursor  `json:"cursor"`
}

// SeasonConnection is the connection containing edges to Season.
type SeasonConnection struct {
	Edges      []*SeasonEdge `json:"edges"`
	PageInfo   PageInfo      `json:"pageInfo"`
	TotalCount int           `json:"totalCount"`
}

func (c *SeasonConnection) build(nodes []*Season, pager *seasonPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Season
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Season {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Season {
			return nodes[i]
		}
	}
	c.Edges = make([]*SeasonEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &SeasonEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// SeasonPaginateOption enables pagination customization.
type SeasonPaginateOption func(*seasonPager) error

// WithSeasonOrder configures pagination ordering.
func WithSeasonOrder(order *SeasonOrder) SeasonPaginateOption {
	if order == nil {
		order = DefaultSeasonOrder
	}
	o := *order
	return func(pager *seasonPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultSeasonOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithSeasonFilter configures pagination filter.
func WithSeasonFilter(filter func(*SeasonQuery) (*SeasonQuery, error)) SeasonPaginateOption {
	return func(pager *seasonPager) error {
		if filter == nil {
			return errors.New("SeasonQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type seasonPager struct {
	reverse bool
	order   *SeasonOrder
	filter  func(*SeasonQuery) (*SeasonQuery, error)
}

func newSeasonPager(opts []SeasonPaginateOption, reverse bool) (*seasonPager, error) {
	pager := &seasonPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultSeasonOrder
	}
	return pager, nil
}

func (p *seasonPager) applyFilter(query *SeasonQuery) (*SeasonQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *seasonPager) toCursor(s *Season) Cursor {
	return p.order.Field.toCursor(s)
}

func (p *seasonPager) applyCursors(query *SeasonQuery, after, before *Cursor) (*SeasonQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultSeasonOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *seasonPager) applyOrder(query *SeasonQuery) *SeasonQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultSeasonOrder.Field {
		query = query.Order(DefaultSeasonOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *seasonPager) orderExpr(query *SeasonQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultSeasonOrder.Field {
			b.Comma().Ident(DefaultSeasonOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Season.
func (s *SeasonQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...SeasonPaginateOption,
) (*SeasonConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newSeasonPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if s, err = pager.applyFilter(s); err != nil {
		return nil, err
	}
	conn := &SeasonConnection{Edges: []*SeasonEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := s.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if s, err = pager.applyCursors(s, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		s.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := s.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	s = pager.applyOrder(s)
	nodes, err := s.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// SeasonOrderFieldID orders Season by id.
	SeasonOrderFieldID = &SeasonOrderField{
		Value: func(s *Season) (ent.Value, error) {
			return s.ID, nil
		},
		column: season.FieldID,
		toTerm: season.ByID,
		toCursor: func(s *Season) Cursor {
			return Cursor{
				ID:    s.ID,
				Value: s.ID,
			}
		},
	}
	// SeasonOrderFieldNumber orders Season by number.
	SeasonOrderFieldNumber = &SeasonOrderField{
		Value: func(s *Season) (ent.Value, error) {
			return s.Number, nil
		},
		column: season.FieldNumber,
		toTerm: season.ByNumber,
		toCursor: func(s *Season) Cursor {
			return Cursor{
				ID:    s.ID,
				Value: s.Number,
			}
		},
	}
	// SeasonOrderFieldStartDate orders Season by startDate.
	SeasonOrderFieldStartDate = &SeasonOrderField{
		Value: func(s *Season) (ent.Value, error) {
			return s.StartDate, nil
		},
		column: season.FieldStartDate,
		toTerm: season.ByStartDate,
		toCursor: func(s *Season) Cursor {
			return Cursor{
				ID:    s.ID,
				Value: s.StartDate,
			}
		},
	}
	// SeasonOrderFieldEndDate orders Season by endDate.
	SeasonOrderFieldEndDate = &SeasonOrderField{
		Value: func(s *Season) (ent.Value, error) {
			return s.EndDate, nil
		},
		column: season.FieldEndDate,
		toTerm: season.ByEndDate,
		toCursor: func(s *Season) Cursor {
			return Cursor{
				ID:    s.ID,
				Value: s.EndDate,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f SeasonOrderField) String() string {
	var str string
	switch f.column {
	case SeasonOrderFieldID.column:
		str = "ID"
	case SeasonOrderFieldNumber.column:
		str = "NUMBER"
	case SeasonOrderFieldStartDate.column:
		str = "START_DATE"
	case SeasonOrderFieldEndDate.column:
		str = "END_DATE"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f SeasonOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *SeasonOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("SeasonOrderField %T must be a string", v)
	}
	switch str {
	case "ID":
		*f = *SeasonOrderFieldID
	case "NUMBER":
		*f = *SeasonOrderFieldNumber
	case "START_DATE":
		*f = *SeasonOrderFieldStartDate
	case "END_DATE":
		*f = *SeasonOrderFieldEndDate
	default:
		return fmt.Errorf("%s is not a valid SeasonOrderField", str)
	}
	return nil
}

// SeasonOrderField defines the ordering field of Season.
type SeasonOrderField struct {
	// Value extracts the ordering value from the given Season.
	Value    func(*Season) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) season.OrderOption
	toCursor func(*Season) Cursor
}

// SeasonOrder defines the ordering of Season.
type SeasonOrder struct {
	Direction OrderDirection    `json:"direction"`
	Field     *SeasonOrderField `json:"field"`
}

// DefaultSeasonOrder is the default ordering of Season.
var DefaultSeasonOrder = &SeasonOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &SeasonOrderField{
		Value: func(s *Season) (ent.Value, error) {
			return s.ID, nil
		},
		column: season.FieldID,
		toTerm: season.ByID,
		toCursor: func(s *Season) Cursor {
			return Cursor{ID: s.ID}
		},
	},
}

// ToEdge converts Season into SeasonEdge.
func (s *Season) ToEdge(order *SeasonOrder) *SeasonEdge {
	if order == nil {
		order = DefaultSeasonOrder
	}
	return &SeasonEdge{
		Node:   s,
		Cursor: order.Field.toCursor(s),
	}
}
