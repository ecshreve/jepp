// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// ListCategoryParams is parameters of listCategory operation.
type ListCategoryParams struct {
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

func unpackListCategoryParams(packed middleware.Parameters) (params ListCategoryParams) {
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "itemsPerPage",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ItemsPerPage = v.(OptInt)
		}
	}
	return params
}

func decodeListCategoryParams(args [0]string, argsEscaped bool, r *http.Request) (params ListCategoryParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Page.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           1,
							MaxSet:        false,
							Max:           0,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: itemsPerPage.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "itemsPerPage",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotItemsPerPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotItemsPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ItemsPerPage.SetTo(paramsDotItemsPerPageVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.ItemsPerPage.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           1,
							MaxSet:        true,
							Max:           255,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "itemsPerPage",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ListCategoryCluesParams is parameters of listCategoryClues operation.
type ListCategoryCluesParams struct {
	// ID of the Category.
	ID int
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

func unpackListCategoryCluesParams(packed middleware.Parameters) (params ListCategoryCluesParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int)
	}
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "itemsPerPage",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ItemsPerPage = v.(OptInt)
		}
	}
	return params
}

func decodeListCategoryCluesParams(args [1]string, argsEscaped bool, r *http.Request) (params ListCategoryCluesParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: itemsPerPage.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "itemsPerPage",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotItemsPerPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotItemsPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ItemsPerPage.SetTo(paramsDotItemsPerPageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "itemsPerPage",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ListClueParams is parameters of listClue operation.
type ListClueParams struct {
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

func unpackListClueParams(packed middleware.Parameters) (params ListClueParams) {
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "itemsPerPage",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ItemsPerPage = v.(OptInt)
		}
	}
	return params
}

func decodeListClueParams(args [0]string, argsEscaped bool, r *http.Request) (params ListClueParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Page.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           1,
							MaxSet:        false,
							Max:           0,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: itemsPerPage.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "itemsPerPage",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotItemsPerPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotItemsPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ItemsPerPage.SetTo(paramsDotItemsPerPageVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.ItemsPerPage.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           1,
							MaxSet:        true,
							Max:           255,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "itemsPerPage",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ListGameParams is parameters of listGame operation.
type ListGameParams struct {
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

func unpackListGameParams(packed middleware.Parameters) (params ListGameParams) {
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "itemsPerPage",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ItemsPerPage = v.(OptInt)
		}
	}
	return params
}

func decodeListGameParams(args [0]string, argsEscaped bool, r *http.Request) (params ListGameParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Page.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           1,
							MaxSet:        false,
							Max:           0,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: itemsPerPage.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "itemsPerPage",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotItemsPerPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotItemsPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ItemsPerPage.SetTo(paramsDotItemsPerPageVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.ItemsPerPage.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           1,
							MaxSet:        true,
							Max:           255,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "itemsPerPage",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ListGameCluesParams is parameters of listGameClues operation.
type ListGameCluesParams struct {
	// ID of the Game.
	ID int
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

func unpackListGameCluesParams(packed middleware.Parameters) (params ListGameCluesParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int)
	}
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "itemsPerPage",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ItemsPerPage = v.(OptInt)
		}
	}
	return params
}

func decodeListGameCluesParams(args [1]string, argsEscaped bool, r *http.Request) (params ListGameCluesParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: itemsPerPage.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "itemsPerPage",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotItemsPerPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotItemsPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ItemsPerPage.SetTo(paramsDotItemsPerPageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "itemsPerPage",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ListSeasonParams is parameters of listSeason operation.
type ListSeasonParams struct {
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

func unpackListSeasonParams(packed middleware.Parameters) (params ListSeasonParams) {
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "itemsPerPage",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ItemsPerPage = v.(OptInt)
		}
	}
	return params
}

func decodeListSeasonParams(args [0]string, argsEscaped bool, r *http.Request) (params ListSeasonParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Page.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           1,
							MaxSet:        false,
							Max:           0,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: itemsPerPage.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "itemsPerPage",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotItemsPerPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotItemsPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ItemsPerPage.SetTo(paramsDotItemsPerPageVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.ItemsPerPage.Get(); ok {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           1,
							MaxSet:        true,
							Max:           255,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "itemsPerPage",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ListSeasonGamesParams is parameters of listSeasonGames operation.
type ListSeasonGamesParams struct {
	// ID of the Season.
	ID int
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

func unpackListSeasonGamesParams(packed middleware.Parameters) (params ListSeasonGamesParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int)
	}
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "itemsPerPage",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ItemsPerPage = v.(OptInt)
		}
	}
	return params
}

func decodeListSeasonGamesParams(args [1]string, argsEscaped bool, r *http.Request) (params ListSeasonGamesParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: itemsPerPage.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "itemsPerPage",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotItemsPerPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotItemsPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ItemsPerPage.SetTo(paramsDotItemsPerPageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "itemsPerPage",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ReadCategoryParams is parameters of readCategory operation.
type ReadCategoryParams struct {
	// ID of the Category.
	ID int
}

func unpackReadCategoryParams(packed middleware.Parameters) (params ReadCategoryParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int)
	}
	return params
}

func decodeReadCategoryParams(args [1]string, argsEscaped bool, r *http.Request) (params ReadCategoryParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ReadClueParams is parameters of readClue operation.
type ReadClueParams struct {
	// ID of the Clue.
	ID int
}

func unpackReadClueParams(packed middleware.Parameters) (params ReadClueParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int)
	}
	return params
}

func decodeReadClueParams(args [1]string, argsEscaped bool, r *http.Request) (params ReadClueParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ReadGameParams is parameters of readGame operation.
type ReadGameParams struct {
	// ID of the Game.
	ID int
}

func unpackReadGameParams(packed middleware.Parameters) (params ReadGameParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int)
	}
	return params
}

func decodeReadGameParams(args [1]string, argsEscaped bool, r *http.Request) (params ReadGameParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ReadSeasonParams is parameters of readSeason operation.
type ReadSeasonParams struct {
	// ID of the Season.
	ID int
}

func unpackReadSeasonParams(packed middleware.Parameters) (params ReadSeasonParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int)
	}
	return params
}

func decodeReadSeasonParams(args [1]string, argsEscaped bool, r *http.Request) (params ReadSeasonParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}