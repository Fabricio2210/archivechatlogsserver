package pagination

func PaginationInfo(page int) map[string]int {
    objPagination := map[string]int{
        "nextPage":     page + 1,
        "previousPage": page - 1,
    }
    return objPagination
}