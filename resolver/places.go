package resolver

import (
	"context"

	"github.com/jotacamou/ngingql/db"
)

type placeResolver struct {
	p *db.Place
}

func (r *Resolver) AllPlaces(ctx context.Context) *[]*placeResolver {
	var p []*placeResolver

	wrapper := func(place *db.Place) {
		p = append(p, &placeResolver{place})
	}

	db.AllPlaces(wrapper)

	return &p
}

func (r *placeResolver) ID() int32 {
	return r.p.ID
}

func (r *placeResolver) Name() string {
	return r.p.Name
}

func (r *placeResolver) Zipcode() string {
	return r.p.Zipcode
}

func (r *placeResolver) Country() string {
	return r.p.Country
}
