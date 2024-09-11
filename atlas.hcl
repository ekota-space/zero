data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./db/loader",
  ]
}

env "gorm" {
  schemas = [ "public" ]
  src = data.external_schema.gorm.url
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