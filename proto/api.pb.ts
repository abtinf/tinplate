/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "./fetch.pb"
import * as GoogleApiHttpbody from "./google/api/httpbody.pb"
import * as GoogleProtobufTimestamp from "./google/protobuf/Timestamp.pb"
import * as GoogleProtobufEmpty from "./google/protobuf/empty.pb"
export type ExampleRequest = {
  name?: string
}

export type ExampleReply = {
  message?: string
}

export type Migration = {
  id?: number
  name?: string
  query?: string
  createdAt?: GoogleProtobufTimestamp.Timestamp
}

export type MigrationList = {
  migrations?: Migration[]
}

export class API {
  static StartupProbe(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<GoogleProtobufEmpty.Empty> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, GoogleProtobufEmpty.Empty>(`/probe/startup?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static LivenessProbe(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<GoogleProtobufEmpty.Empty> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, GoogleProtobufEmpty.Empty>(`/probe/live?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static ReadinessProbe(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<GoogleProtobufEmpty.Empty> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, GoogleProtobufEmpty.Empty>(`/probe/ready?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static ExampleGet(req: ExampleRequest, initReq?: fm.InitReq): Promise<ExampleReply> {
    return fm.fetchReq<ExampleRequest, ExampleReply>(`/ExampleGet?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static ExamplePost(req: ExampleRequest, initReq?: fm.InitReq): Promise<ExampleReply> {
    return fm.fetchReq<ExampleRequest, ExampleReply>(`/ExamplePost`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static Download(req: ExampleRequest, entityNotifier?: fm.NotifyStreamEntityArrival<GoogleApiHttpbody.HttpBody>, initReq?: fm.InitReq): Promise<void> {
    return fm.fetchStreamingRequest<ExampleRequest, GoogleApiHttpbody.HttpBody>(`/download?${fm.renderURLSearchParams(req, [])}`, entityNotifier, {...initReq, method: "GET"})
  }
  static GetMigrations(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<MigrationList> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, MigrationList>(`/migrations?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
}