// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var guruguru_pb = require('./guruguru_pb.js');

function serialize_JoinRequest(arg) {
  if (!(arg instanceof guruguru_pb.JoinRequest)) {
    throw new Error('Expected argument of type JoinRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_JoinRequest(buffer_arg) {
  return guruguru_pb.JoinRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_JoinResponse(arg) {
  if (!(arg instanceof guruguru_pb.JoinResponse)) {
    throw new Error('Expected argument of type JoinResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_JoinResponse(buffer_arg) {
  return guruguru_pb.JoinResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_PokeRequest(arg) {
  if (!(arg instanceof guruguru_pb.PokeRequest)) {
    throw new Error('Expected argument of type PokeRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_PokeRequest(buffer_arg) {
  return guruguru_pb.PokeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_PokeResponse(arg) {
  if (!(arg instanceof guruguru_pb.PokeResponse)) {
    throw new Error('Expected argument of type PokeResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_PokeResponse(buffer_arg) {
  return guruguru_pb.PokeResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_SetNextRequest(arg) {
  if (!(arg instanceof guruguru_pb.SetNextRequest)) {
    throw new Error('Expected argument of type SetNextRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_SetNextRequest(buffer_arg) {
  return guruguru_pb.SetNextRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_SetNextResponse(arg) {
  if (!(arg instanceof guruguru_pb.SetNextResponse)) {
    throw new Error('Expected argument of type SetNextResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_SetNextResponse(buffer_arg) {
  return guruguru_pb.SetNextResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var BossServiceService = exports.BossServiceService = {
  join: {
    path: '/BossService/Join',
    requestStream: false,
    responseStream: false,
    requestType: guruguru_pb.JoinRequest,
    responseType: guruguru_pb.JoinResponse,
    requestSerialize: serialize_JoinRequest,
    requestDeserialize: deserialize_JoinRequest,
    responseSerialize: serialize_JoinResponse,
    responseDeserialize: deserialize_JoinResponse,
  },
};

exports.BossServiceClient = grpc.makeGenericClientConstructor(BossServiceService);
var MemberServiceService = exports.MemberServiceService = {
  poke: {
    path: '/MemberService/Poke',
    requestStream: false,
    responseStream: false,
    requestType: guruguru_pb.PokeRequest,
    responseType: guruguru_pb.PokeResponse,
    requestSerialize: serialize_PokeRequest,
    requestDeserialize: deserialize_PokeRequest,
    responseSerialize: serialize_PokeResponse,
    responseDeserialize: deserialize_PokeResponse,
  },
  setNext: {
    path: '/MemberService/SetNext',
    requestStream: false,
    responseStream: false,
    requestType: guruguru_pb.SetNextRequest,
    responseType: guruguru_pb.SetNextResponse,
    requestSerialize: serialize_SetNextRequest,
    requestDeserialize: deserialize_SetNextRequest,
    responseSerialize: serialize_SetNextResponse,
    responseDeserialize: deserialize_SetNextResponse,
  },
};

exports.MemberServiceClient = grpc.makeGenericClientConstructor(MemberServiceService);
