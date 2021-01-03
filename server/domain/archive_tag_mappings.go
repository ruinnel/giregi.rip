package domain

import "reflect"

type ArchiveTagMapping struct {
	ID        int64 `json:"id" mysql:"id" storm:"id,increment"`
	TagID     int64 `json:"tagId" mysql:"tag_id" storm:"index" storm:"index"`
	ArchiveID int64 `json:"archiveId" mysql:"archive_id" storm:"index" storm:"index"`
	CreatedAt Time  `json:"createdAt" mysql:"created_at"`
	UpdatedAt Time  `json:"updatedAt" mysql:"updated_at"`
}

type archiveTagMappingField struct {
	ID        reflect.StructField
	TagID     reflect.StructField
	ArchiveID reflect.StructField
	CreatedAt reflect.StructField
	UpdatedAt reflect.StructField
}

var ArchiveTagMappingField = makeFields(&ArchiveTagMapping{}, &archiveTagMappingField{}).(*archiveTagMappingField)
