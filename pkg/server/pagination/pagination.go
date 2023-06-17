package pagination

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ecshreve/jepp/pkg/models"
	"github.com/gin-gonic/gin"
)

const (
	DEFAULT_PAGE_TEXT    = "page"
	DEFAULT_SIZE_TEXT    = "size"
	DEFAULT_PAGE         = "0"
	DEFAULT_PAGE_SIZE    = "10"
	DEFAULT_MIN_PAGESIZE = 10
	DEFAULT_MAX_PAGESIZE = 100
)

// Default create a new pagination middleware with default values.
func Default() gin.HandlerFunc {
	return New(
		DEFAULT_PAGE_TEXT,
		DEFAULT_SIZE_TEXT,
		DEFAULT_PAGE,
		DEFAULT_PAGE_SIZE,
		DEFAULT_MIN_PAGESIZE,
		DEFAULT_MAX_PAGESIZE,
	)
}

// New creates a new pagniation middleware with custom values.
func New(pageText, sizeText, defaultPage, defaultPageSize string, minPageSize, maxPageSize int) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the page from the query string and convert it to an integer
		pageStr := c.DefaultQuery(pageText, defaultPage)
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "page number must be an integer"})
			return
		}

		// Validate for positive page number
		if page < 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "page number must be positive"})
			return
		}

		// Extract the size from the query string and convert it to an integer
		sizeStr := c.DefaultQuery(sizeText, defaultPageSize)
		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "page size must be an integer"})
			return
		}

		// Validate for min and max page size
		if size < minPageSize || size > maxPageSize {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "page size must be between " + strconv.Itoa(minPageSize) + " and " + strconv.Itoa(maxPageSize)})
			return
		}

		// Set the page and size in the gin context
		c.Set(pageText, page)
		c.Set(sizeText, size)

		c.Next()
	}
}

// Response is the response for pagination query request
type Response struct {
	Data  any   `json:"data"`
	Links links `json:"_links,omitempty"`
}

type links struct {
	Next string `json:"next,omitempty"`
	Prev string `json:"prev,omitempty"`
}

func GetLinks(ctx *gin.Context, total int64, q *models.PaginationParams) links {
	url := fmt.Sprintf("%v", ctx.Request.URL)
	baseURL := strings.Split(url, "?")[0]
	l := links{}
	if int64(total) == int64(q.PageSize) {
		l.Next = fmt.Sprintf(
			"%s?page=%d&size=%d",
			baseURL,
			q.Page+1,
			q.PageSize,
		)
	}

	if q.Page != 1 {
		prevStart := q.Page - 1
		if prevStart < 0 {
			prevStart = 0
		}

		l.Prev = fmt.Sprintf(
			"%s?page=%d&size=%d",
			baseURL,
			q.Page,
			q.PageSize,
		)
	}

	return l
}
