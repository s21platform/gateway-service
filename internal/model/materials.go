package model

type EditMaterial struct {
	MaterialUUID  string `json:"material_uuid"`
	Title         string `json:"title"`
	CoverImageURL string `json:"cover_image_url"`
	Description   string `json:"description"`
	Content       string `json:"content"`
}
