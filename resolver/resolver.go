package resolver

import (
	"context"

	"github.com/jotacamou/ngingql/db"
)

type Resolver struct{}

type placeResolver struct {
	p *db.Place
}

func (r *Resolver) AllPlaces(ctx context.Context) *[]*placeResolver {
	//if fetchUserFromAuthenticationToken(ctx) == nil {
	//	return nil
	//}

	var p []*placeResolver

	wrapper := func(place *db.Place) {
		p = append(p, &placeResolver{place})
	}

	db.AllPlaces(wrapper)

	return &p
}

func (r *placeResolver) ID() string {
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

//func fetchUserFromAuthenticationToken(ctx context.Context) *db.User {
//	token, ok := ctx.Value("AuthorizationToken").(string)
//	if !ok {
//		log.Println("Error occurred while parsing authorization token")
//		return nil
//	}
//
//	return db.FindUserByID(token)
//}
