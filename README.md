# Spanner Go client not building on Bazel

This is a minimal reproduction of the Spanner go client not building with Bazel.

To reproduce:

```sh
bazelisk build //spanner
```

### Unsolved: protoc-gen-validate can't find rules_cc or rules_java

The build currently errors with:

```
error loading package '@@gazelle~~go_deps~com_github_envoyproxy_protoc_gen_validate//validate':
Unable to find package for @@[unknown repo 'rules_java' requested from @@gazelle~~go_deps~com_github_envoyproxy_protoc_gen_validate]//java:defs.bz
```


### Solved: CNCF proto file error

To work around proto files in https://github.com/cncf/xds/tree/main/xds/service/orca/v3 that result in the following
error:

```
ERROR: no such package '@@gazelle~~go_deps~com_github_cncf_xds_go//xds/service/orca/v3':
BUILD file not found in directory 'xds/service/orca/v3' of external repository @@gazelle~~go_deps~com_github_cncf_xds_go.
Add a BUILD file to a directory to mark it as a package.
```

We add:

```bazel
go_deps.gazelle_default_attributes(
    directives = [
        "gazelle:proto disable_global",
    ],
)
```

### Solved: CEL build error

To work around https://github.com/google/cel-spec/issues/378, we used:

```bazel
go_deps.gazelle_override(
    # Force Gazelle to wipe out the existing build files before regenerate them.
    build_file_generation = "clean",
    directives = [
        "gazelle:resolve proto go google/rpc/status.proto @org_golang_google_genproto_googleapis_rpc//status",
        "gazelle:resolve proto proto google/rpc/status.proto @googleapis//google/rpc:status_proto",
    ],
    path = "cel.dev/expr",
)
```

To work around 
