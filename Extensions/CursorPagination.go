package Extensions

type (
	CursorPagination struct {
		pagingInfo CursorPagingInfo
	}
	CursorPagingInfo struct {
		Page        int  `json:"page"`
		TotalCount  int  `json:"totalCount,omitempty"`
		Limit       int  `json:"limit"`
		HasNext     bool `json:"hasNext"`
		StartCursor int  `json:"startCursor"`
		EndCursor   int  `json:"endCursor"`
	}
)

//CursorOffset
func (c *CursorPagination) CursorOffset() {

}

//GetPagingInfo
func (c *CursorPagination) GetPagingInfo() {

}

//startCursor
func (c *CursorPagination) startCursor(page int, limit int) int {
	return (page - 1) * limit
}

//endCursor
func (c *CursorPagination) endCursor(page int, limit int) int {
	return page * limit
}

//hasNext
func (c *CursorPagination) hasNext(numOfRows int, limit int) bool {
	return numOfRows >= limit
}
