package model

import (
	"time"

	"github.com/s21platform/materials-service/pkg/materials"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MaterialList []Material

type Material struct {
	UUID            string     `json:"uuid"`
	OwnerUUID       string     `json:"owner_uuid"`
	Title           string     `json:"title"`
	CoverImageURL   string     `json:"cover_image_url"`
	Description     string     `json:"description"`
	Content         *string    `json:"content"`
	ReadTimeMinutes int32      `json:"read_time_minutes"`
	Status          string     `json:"status"`
	CreatedAt       *time.Time `json:"created_at"`
	EditedAt        *time.Time `json:"edited_at"`
	PublishedAt     *time.Time `json:"published_at"`
	ArchivedAt      *time.Time `json:"archived_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
	LikesCount      int32      `json:"likes_count"`
}

func FromProto(protoList []*materials.Material) MaterialList {
	result := make(MaterialList, 0, len(protoList))
	for _, proto := range protoList {
		m := Material{
			UUID:            proto.Uuid,
			OwnerUUID:       proto.OwnerUuid,
			Title:           proto.Title,
			CoverImageURL:   proto.CoverImageUrl,
			Description:     proto.Description,
			ReadTimeMinutes: proto.ReadTimeMinutes,
			Status:          proto.Status,
			LikesCount:      proto.LikesCount,
		}
		if proto.Content != "" {
			content := proto.Content
			m.Content = &content
		}
		if proto.CreatedAt != nil {
			t := proto.CreatedAt.AsTime()
			m.CreatedAt = &t
		}
		if proto.EditedAt != nil {
			t := proto.EditedAt.AsTime()
			m.EditedAt = &t
		}
		if proto.PublishedAt != nil {
			t := proto.PublishedAt.AsTime()
			m.PublishedAt = &t
		}
		if proto.ArchivedAt != nil {
			t := proto.ArchivedAt.AsTime()
			m.ArchivedAt = &t
		}
		if proto.DeletedAt != nil {
			t := proto.DeletedAt.AsTime()
			m.DeletedAt = &t
		}
		result = append(result, m)
	}
	return result
}

func (a *MaterialList) ListFromDTO() []*materials.Material {
	result := make([]*materials.Material, 0, len(*a))

	for _, material := range *a {
		m := &materials.Material{
			Uuid:            material.UUID,
			OwnerUuid:       material.OwnerUUID,
			Title:           material.Title,
			CoverImageUrl:   material.CoverImageURL,
			Description:     material.Description,
			ReadTimeMinutes: material.ReadTimeMinutes,
			Status:          material.Status,
			LikesCount:      material.LikesCount,
		}

		if material.Content != nil {
			m.Content = *material.Content
		}
		if material.EditedAt != nil {
			m.EditedAt = timestamppb.New(*material.EditedAt)
		}
		if material.PublishedAt != nil {
			m.PublishedAt = timestamppb.New(*material.PublishedAt)
		}
		if material.ArchivedAt != nil {
			m.ArchivedAt = timestamppb.New(*material.ArchivedAt)
		}
		if material.DeletedAt != nil {
			m.DeletedAt = timestamppb.New(*material.DeletedAt)
		}

		result = append(result, m)
	}

	return result
}
