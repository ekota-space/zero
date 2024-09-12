data "hcl_schema" "zero" {
  paths = glob("db/schema/*.pg.hcl")
}

env "local" {
  schemas = [ "public" ]
  src = data.hcl_schema.zero.url
  // This is just to create diff
  dev = "docker://postgres/16-alpine/zero?search_path=public"
  migration {
    dir = "file://db/migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}