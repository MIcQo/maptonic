# maptonic

**Join the MapTonic Team: Revolutionize Mapping & Location Services!**

We’re building MapTonic, a cutting-edge platform that blends reverse geocoding with dynamic map tile services (vector & raster), providing powerful location insights and seamless map rendering. If you’re passionate about mapping, geospatial data, and want to be part of an innovative project, we want you on our team!


### Infra Requirements

- Go 1.23+
- Prefers stdlib
- golangci-lint
- Docker & Kubernetes compatible

### APP Requirements

- Reverse geocoding
- Address Search
- Map tile server from mbtiles for vector and raster
- All done by using OSM data

#### Used components
- osm2pgsql (https://github.com/osm2pgsql-dev/osm2pgsql)
- Postgres (https://www.postgresql.org/)
- PostGIS (https://postgis.net/)
- Bun (https://github.com/uptrace/bun)
- Fiber HTTP Framework (https://docs.gofiber.io/) (because of fasthttp)
- Huma (https://github.com/danielgtaylor/huma) (for OpenAPI documentation)
- mbtiles (https://github.com/mapbox/mbtiles-tools) TODO
