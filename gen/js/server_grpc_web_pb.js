/**
 * @fileoverview gRPC-Web generated client stub for server
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v6.30.2
// source: server.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.server = require('./server_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.server.ExampleServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.server.ExampleServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.server.MessageRequest,
 *   !proto.server.MessageResponse>}
 */
const methodDescriptor_ExampleService_SendMessage = new grpc.web.MethodDescriptor(
  '/server.ExampleService/SendMessage',
  grpc.web.MethodType.UNARY,
  proto.server.MessageRequest,
  proto.server.MessageResponse,
  /**
   * @param {!proto.server.MessageRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.server.MessageResponse.deserializeBinary
);


/**
 * @param {!proto.server.MessageRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.server.MessageResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.server.MessageResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.server.ExampleServiceClient.prototype.sendMessage =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/server.ExampleService/SendMessage',
      request,
      metadata || {},
      methodDescriptor_ExampleService_SendMessage,
      callback);
};


/**
 * @param {!proto.server.MessageRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.server.MessageResponse>}
 *     Promise that resolves to the response
 */
proto.server.ExampleServicePromiseClient.prototype.sendMessage =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/server.ExampleService/SendMessage',
      request,
      metadata || {},
      methodDescriptor_ExampleService_SendMessage);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.server.PoliceRequest,
 *   !proto.server.PoliceResponse>}
 */
const methodDescriptor_ExampleService_SendPolice = new grpc.web.MethodDescriptor(
  '/server.ExampleService/SendPolice',
  grpc.web.MethodType.UNARY,
  proto.server.PoliceRequest,
  proto.server.PoliceResponse,
  /**
   * @param {!proto.server.PoliceRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.server.PoliceResponse.deserializeBinary
);


/**
 * @param {!proto.server.PoliceRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.server.PoliceResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.server.PoliceResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.server.ExampleServiceClient.prototype.sendPolice =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/server.ExampleService/SendPolice',
      request,
      metadata || {},
      methodDescriptor_ExampleService_SendPolice,
      callback);
};


/**
 * @param {!proto.server.PoliceRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.server.PoliceResponse>}
 *     Promise that resolves to the response
 */
proto.server.ExampleServicePromiseClient.prototype.sendPolice =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/server.ExampleService/SendPolice',
      request,
      metadata || {},
      methodDescriptor_ExampleService_SendPolice);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.server.PoliceRequest,
 *   !proto.server.PoliceStreamResponse>}
 */
const methodDescriptor_ExampleService_StreamPoliceUpdates = new grpc.web.MethodDescriptor(
  '/server.ExampleService/StreamPoliceUpdates',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.server.PoliceRequest,
  proto.server.PoliceStreamResponse,
  /**
   * @param {!proto.server.PoliceRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.server.PoliceStreamResponse.deserializeBinary
);


/**
 * @param {!proto.server.PoliceRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.server.PoliceStreamResponse>}
 *     The XHR Node Readable Stream
 */
proto.server.ExampleServiceClient.prototype.streamPoliceUpdates =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/server.ExampleService/StreamPoliceUpdates',
      request,
      metadata || {},
      methodDescriptor_ExampleService_StreamPoliceUpdates);
};


/**
 * @param {!proto.server.PoliceRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.server.PoliceStreamResponse>}
 *     The XHR Node Readable Stream
 */
proto.server.ExampleServicePromiseClient.prototype.streamPoliceUpdates =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/server.ExampleService/StreamPoliceUpdates',
      request,
      metadata || {},
      methodDescriptor_ExampleService_StreamPoliceUpdates);
};


module.exports = proto.server;

