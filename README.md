# PLAN

```
         P urposeful
         L ogistics
         A rchitecture
P  L  A  N etwork
```

[PLAN](http://plan-systems.org) is a free and open platform for groups to securely communicate, collaborate, and coordinate projects and activities.

## About

This repo contains PLAN's backend infrastrcuture that ultimately hosts [plan-client-unity](https://github.com/plan-systems/plan-client-unity) clients.  Since PLAN uses open standards, protocols, and data structures, other implementations may someday also exist, so `plan-core` is regarded as the official reference implementation.

The [PLAN Design & Engineering Docs](https://github.com/plan-systems/design-docs) lay out PLAN's architecture in detail.

## Milestone Log


|    Date     | Status                                                                  |
|:-----------:|-------------------------------------------------------------------------|
| Jul 2019 | pre-alpha prototype demonstrates `ch/talk` and `ch/space` at [DWeb 2019](https://www.plan-systems.org/2019/05/13/decentralized-web-camp-and-beyond/)               |
| May 2019 | pnode community genesis and basic channel ops implmented               |
| Feb 2019 | pdi-datastore now implements [StorageProvider](https://github.com/plan-systems/design-docs/blob/master/PLAN-API-Documentation.md#Persistent-Data-Interface) using a centralized db |
| Dec 2018 | Architecture docs complete; resuming active development                 |
| Sep 2018 | Pausing active development to work on PLAN PoC                          |
| Aug 2018 | SKI makeover following permissions model proof of concept               |
| Jul 2018 | Active development begins                                               |



## Links

|                           |                                                          |
|--------------------------:|----------------------------------------------------------|
|                 Team Lead | [Drew O'Meara](mailto:drew2019@plan-systems.org)         |
| Design & Engineering Docs | https://github.com/plan-systems/design-docs              |
|              PLAN Website | http://plan-systems.org                                  |
|                   License | [GPL-v3](https://www.gnu.org/licenses/gpl-3.0.en.htmlm)  |


## Developer Setup

This library uses protobuf definitions found in the [plan-protobufs](https://github.com/plan-systems/plan-protobufs) repo. The `*.pb.go` source code files generated from those protobuf definitions has been vendored into this repository at the appropriate subpackage location. This means that if you are consuming this library for use in another application, you can simply:

```
go get github.com/plan-systems/plan-core
```


If you are contributing new features to `plan-core`, you may end up needing to update the generated `*.pb.go` source. To do that, you'll need to do the following:

1. Set up **Gprc.Tools**:
    1. Download the [latest Grpc.Tools nuget package](https://www.nuget.org/packages/Grpc.Tools/)

       ```
       curl -Lso "grpc.tools.1.22.0.nupkg" \
           "https://www.nuget.org/api/v2/package/Grpc.Tools/1.22.0"
       ```

    2. Unzip the nuget pkg zip file using unzip.

       _Don't use macOS's default unarchiver since it does not restore `.nupkg` filenames properly. Instead:_

        ```
        unzip grpc.tools.1.22.0.nupkg
        ```

    3. Copy the binary appropriate for your platform to somewhere on your `$PATH` and give it executable permissions:
        ```
        cp ./Grpc.Tools/tools/macosx_x64/protoc /usr/local/bin
        chmod +x /usr/local/bin/protoc
        ```

2. Ensure your `$PATH` contains your `$GOPATH`'s `bin` directory:

    ```
    # if you have GOPATH set:
    PATH="${PATH}:${GOPATH}/bin"

    # if you don't have GOPATH set:
    PATH="${PATH}:~/go/bin"
    ```

3. Install [gogo protobufs](https://github.com/gogo/protobuf/):

    `go get -u github.com/gogo/protobuf/protoc-gen-gofast`

4. Install [gRPC](https://grpc.io/):

    `go get -u google.golang.org/grpc`

5. Invoke the build scripts, giving it the path to the protobufs files and this repo. Then commit the changes:

    ```
    ./build-protobufs --protos ../plan-protobufs/pkg --dest .
    git add .
    git commit -m "updated protobufs from v1.2.3"
    ```

6. Pick up your new lambo.
