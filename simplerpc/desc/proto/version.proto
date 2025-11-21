syntax = "proto3";

package version;
option go_package = "./types/version";

message VersionRequest {}

message VersionResponse {
    string version = 1;
    string goVersion = 2;
    string commit = 3;
    string date = 4;
}

service Version {
    rpc Version(VersionRequest) returns(VersionResponse) {};
}