package domain

type ArchiveTagMapping struct {
	ID        int64 `storm:"id,increment"`
	TagID     int64 `storm:"index"`
	ArchiveID int64 `storm:"index"`
	CreatedAt Time
	UpdatedAt Time
}
