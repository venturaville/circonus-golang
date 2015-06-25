circonus-golang
========================

About
-----------

Circonus API for golang

Setup
------------

Need to define (at a minimum):
    CIRCONUS_APITOKEN="aaaaaaaaaaaaaaaaaaaaa"


either:
    c := NewCirconus(&Circonus{ApiToken: "myapitoken"})
or:
    os.SetEnv("CIRCONUS_APITOKEN","myapitoken")
    c := NewCirconus(&Circonus{})


