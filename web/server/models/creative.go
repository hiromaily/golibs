package models

import (
	"time"
)

//-----------------------------------------------------------------------------
// Creative
//-----------------------------------------------------------------------------

const (
	BannerCreativeType CreativeType = "banner"
	HTMLCreativeType   CreativeType = "html"
)

type CreativeType string

type Creative struct {
	ID             int
	Type           CreativeType
	ClickURL       string
	StandardWidth  *int
	StandardHeight *int
	CustomWidth    *int
	CustomHeight   *int
	StartAt        *time.Time
	EndAt          *time.Time
}

func (c Creative) Width() int {
	if c.StandardWidth != nil {
		return *c.StandardWidth
	} else if c.CustomWidth != nil {
		return *c.CustomWidth
	}
	return 0
}

func (c Creative) Height() int {
	if c.StandardHeight != nil {
		return *c.StandardHeight
	} else if c.CustomHeight != nil {
		return *c.CustomHeight
	}
	return 0
}

//-----------------------------------------------------------------------------
// CreativeDetail
//-----------------------------------------------------------------------------

type CreativeDetail struct {
	Creative
	AdvertiserID   int
	FilePath       string
	ActualFilePath string
	HTML           string
}

//-----------------------------------------------------------------------------
// CreativeDetailRepository
//-----------------------------------------------------------------------------

type CreativeDetailRepository interface {
	Fetch(creativeID string) (*CreativeDetail, error)
}
