package db

import (
	"context"
	"strings"
)

func ReverseGecode(ctx context.Context, lat, lon float64) ([]map[string]any, error) {

	// SELECT name, tags, boundary
	//    FROM planet_osm_roads
	//    where name is not null
	//    ORDER BY way <-> ST_SetSRID(ST_MakePoint(${lon}, ${lat}), 4326)
	//    limit 1
	var nearestRoadQuery = db.NewSelect().
		Column("name", "tags", "boundary").
		Table("planet_osm_roads").
		Where("name is not null").
		OrderExpr("way <-> ST_SetSRID(ST_MakePoint(?, ?), 4326)", lon, lat).
		Limit(1)

	// SELECT name, tags, boundary
	//    FROM planet_osm_point
	//    where name is not null
	//    ORDER BY way <-> ST_SetSRID(ST_MakePoint(${lon}, ${lat}), 4326)
	//    limit 1
	var nearestPointQuery = db.NewSelect().
		Column("name", "tags", "boundary").
		Table("planet_osm_point").
		Where("name is not null").
		OrderExpr("way <-> ST_SetSRID(ST_MakePoint(?, ?), 4326)", lon, lat).
		Limit(1)

	// SELECT
	//    p.osm_id,
	//    COALESCE(p.tags -> 'postal_code', np.tags -> 'addr:postcode') AS postcode,
	//    p.name as polygon_name,
	//    COALESCE(nr.name, np.tags -> 'addr:street') as road_name,
	//    np.tags -> 'addr:housenumber' as house_number,
	//    np.name as point_name,
	//    np.tags -> 'addr:city' AS city_point,
	//    np.tags -> 'addr:country' AS country_point,
	//    p.boundary,
	//    p.place
	//FROM planet_osm_polygon p
	//JOIN nearest_road nr on true
	//JOIN nearest_point np on true
	//-- group by boundary, name, tags, way
	//WHERE 1=1
	//--   and place is not null
	//  and p.boundary = 'administrative'
	//  and admin_level is not null
	//  and st_contains(way, st_setsrid(st_makepoint(${lon}, ${lat}), 4326))
	//ORDER BY way <-> ST_SetSRID(ST_MakePoint(${lon}, ${lat}), 4326), p.admin_level::INTEGER desc

	var columns = []string{
		"p.osm_id",
		"COALESCE(p.tags -> 'postal_code', np.tags -> 'addr:postcode') AS postcode",
		"p.name as polygon_name",
		"COALESCE(nr.name, np.tags -> 'addr:street') as road_name",
		"np.tags -> 'addr:housenumber' as house_number",
		"np.name as point_name",
		"np.tags -> 'addr:city' AS city_point",
		"np.tags -> 'addr:country' AS country_point",
		"p.boundary",
		"p.place",
	}

	var nearestPolygonQuery = db.NewSelect().
		With("nearest_road", nearestRoadQuery).
		With("nearest_point", nearestPointQuery).
		ColumnExpr(strings.Join(columns, ", ")).
		TableExpr("planet_osm_polygon p").
		Join("JOIN nearest_road nr on true").
		Join("JOIN nearest_point np on true").
		Where("p.boundary = 'administrative'").
		Where("admin_level is not null").
		Where("st_contains(way, st_setsrid(st_makepoint(?, ?), 4326))", lon, lat).
		OrderExpr("way <-> ST_SetSRID(ST_MakePoint(?, ?), 4326), p.admin_level::INTEGER desc", lon, lat)

	var result = make([]map[string]any, 0)
	var err = nearestPolygonQuery.Scan(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
