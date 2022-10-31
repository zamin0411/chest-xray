package model

import "github.com/google/uuid"

type AreaDetail struct {
	ID                    uuid.UUID `gorm:"column:area_detail_id" json:"id"`
	Area                  string    `gorm:"column:area_detail_area" json:"area"`
	Detail                string    `gorm:"column:area_detail_detail" json:"detail"`
	ClinicalExaminationID uuid.UUID `gorm:"column:clinical_examination_id" json:"clinicalExaminationId"`
}
