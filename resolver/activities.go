package resolver

import (
	"context"

	"github.com/jotacamou/ngingql/db"
)

type activityResolver struct {
	p *db.Activity
}

func (r *Resolver) AllActivities(ctx context.Context) *[]*activityResolver {
	var p []*activityResolver

	wrapper := func(activity *db.Activity) {
		p = append(p, &activityResolver{activity})
	}

	db.AllActivities(wrapper)

	return &p
}

func (r *activityResolver) ID() int32 {
	return r.p.ID
}

func (r *activityResolver) Title() string {
	return r.p.Title
}

func (r *activityResolver) Description() string {
	return r.p.Description
}

func (r *activityResolver) Owner() string {
	return r.p.Owner
}

func (r *activityResolver) Price() float64 {
	return r.p.Price
}

func (r *activityResolver) Fee() float64 {
	return r.p.Fee
}

func (r *activityResolver) MainPicture() []string {
	return r.p.MainPicture
}

func (r *activityResolver) SlideshowPicture() []string {
	return r.p.SlideshowPicture
}

func (r *activityResolver) Duration() string {
	return r.p.Duration
}

func (r *activityResolver) PlaceID() int32 {
	return r.p.PlaceID
}

func (r *activityResolver) Hour() []string {
	return r.p.Hour
}

func (r *activityResolver) Tag() []string {
	return r.p.Tag
}
