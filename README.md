# [WORK IN PROGRESS] A Golang client for the Anchore Engine API

Uses a slightly modified swagger spec and based on generated code. No executables provided, only intended for use as a library in other projects.

To pull the swagger definition and re-generate all client go code:
```bash
make 
```

Note: the version of the engine where the swagger spec is pulled from is pinned in the Makefile. 
