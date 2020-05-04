package blog

// Repository is an adapter for persistence implementation (port interface)
type Repository interface {
	// GetPageBySlug returns page by it's slug or error if page not found
	FindPages() []*Page

	// GetPageBySlug returns page by it's slug or error if page not found
	GetPageBySlug(string) (*Page, error)
}
