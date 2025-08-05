package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mremperor-atwork/rpg1/api1/internal/repository"
)

type SRDHandlers struct {
	srdRepo *repository.SRDRepository
}

func NewSRDHandlers() *SRDHandlers {
	return &SRDHandlers{
		srdRepo: repository.NewSRDRepository(),
	}
}

func (h *SRDHandlers) GetSRDEntries(c *gin.Context) {
	category := c.Query("category")
	search := c.Query("search")
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	limit := 0
	offset := 0

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
			offset = o
		}
	}

	entries, err := h.srdRepo.GetAllEntries(category, search, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve SRD entries",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"entries": entries,
		"count":   len(entries),
	})
}

func (h *SRDHandlers) GetSRDEntryByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Entry ID is required",
		})
		return
	}

	entry, err := h.srdRepo.GetEntryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "SRD entry not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"entry": entry,
	})
}

func (h *SRDHandlers) GetSRDEntriesByCategory(c *gin.Context) {
	category := c.Param("category")
	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Category is required",
		})
		return
	}

	entries, err := h.srdRepo.GetEntriesByCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve SRD entries",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"category": category,
		"entries":  entries,
		"count":    len(entries),
	})
}

func (h *SRDHandlers) GetSRDCategories(c *gin.Context) {
	categories, err := h.srdRepo.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve categories",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
		"count":      len(categories),
	})
}

func (h *SRDHandlers) SearchSRDEntries(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Search query is required",
		})
		return
	}

	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	limit := 0
	offset := 0

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
			offset = o
		}
	}

	entries, err := h.srdRepo.SearchEntries(query, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to search SRD entries",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"query":   query,
		"entries": entries,
		"count":   len(entries),
	})
}

func (h *SRDHandlers) GetSRDContent(c *gin.Context) {
	category := c.Query("category")
	search := c.Query("search")
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	limit := 0
	offset := 0

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
			offset = o
		}
	}

	content, err := h.srdRepo.GetAllContent(category, search, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve SRD content",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"content": content,
		"count":   len(content),
	})
}

func (h *SRDHandlers) GetSRDContentByTitle(c *gin.Context) {
	title := c.Param("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Content title is required",
		})
		return
	}

	content, err := h.srdRepo.GetContentByTitle(title)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "SRD content not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"content": content,
	})
}
