# hypermedia

An collection of opinionated hypermedia wrappers. Resources and collections. Useful for decorating
arbitrary structs with links.

## Install

    go get github.com/jagregory/hypermedia

## Usage

    import (
      hm "github.com/jagregory/hypermedia"
    )


Resource:

    res := hm.NewResource(product, hm.Selff("/products/%d", product.Id))

JSON:

    {
      "links": {
        "self": { "href": "/products/123" }
      },
      "entity": {
        "id": "123",
        "name": "Soap"
      }
    }

Collection:

    col := hm.NewCollection(
      []hm.Resource{
        hm.NewResource(product, hm.Selff("/products/%d", product.Id)),
      },
      hm.Self("/archive/2"),
      hm.Linkf("next", "/archive/3"),
      hm.Linkf("prev", "/archive/1"),
    )

JSON:

    {
      "link": {
        "self": { "href": "/archive/2" },
        "next": { "href": "/archive/3" },
        "prev": { "href": "/archive/1" }
      },
      "collection": [
        {
          "links": {
            "self": { "href": "/products/123" }
          },
          "entity": {
            "id": "123",
            "name": "Soap"
          }
        }
      ]
    }