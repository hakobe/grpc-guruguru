# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: guruguru.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='guruguru.proto',
  package='',
  syntax='proto3',
  serialized_pb=_b('\n\x0eguruguru.proto\")\n\x06Member\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x11\n\thost_port\x18\x02 \x01(\t\"&\n\x0bJoinRequest\x12\x17\n\x06member\x18\x01 \x01(\x0b\x32\x07.Member\"\x1a\n\x0cJoinResponse\x12\n\n\x02ok\x18\x01 \x01(\x08\"5\n\x0bPokeRequest\x12\x15\n\x04\x66rom\x18\x01 \x01(\x0b\x32\x07.Member\x12\x0f\n\x07message\x18\x02 \x01(\t\"\x1a\n\x0cPokeResponse\x12\n\n\x02ok\x18\x01 \x01(\x08\")\n\x0eSetNextRequest\x12\x17\n\x06member\x18\x01 \x01(\x0b\x32\x07.Member\"\x1d\n\x0fSetNextResponse\x12\n\n\x02ok\x18\x01 \x01(\x08\x32\x32\n\x0b\x42ossService\x12#\n\x04Join\x12\x0c.JoinRequest\x1a\r.JoinResponse2b\n\rMemberService\x12#\n\x04Poke\x12\x0c.PokeRequest\x1a\r.PokeResponse\x12,\n\x07SetNext\x12\x0f.SetNextRequest\x1a\x10.SetNextResponseb\x06proto3')
)




_MEMBER = _descriptor.Descriptor(
  name='Member',
  full_name='Member',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='Member.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='host_port', full_name='Member.host_port', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=18,
  serialized_end=59,
)


_JOINREQUEST = _descriptor.Descriptor(
  name='JoinRequest',
  full_name='JoinRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='member', full_name='JoinRequest.member', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=61,
  serialized_end=99,
)


_JOINRESPONSE = _descriptor.Descriptor(
  name='JoinResponse',
  full_name='JoinResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ok', full_name='JoinResponse.ok', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=101,
  serialized_end=127,
)


_POKEREQUEST = _descriptor.Descriptor(
  name='PokeRequest',
  full_name='PokeRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='from', full_name='PokeRequest.from', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='message', full_name='PokeRequest.message', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=129,
  serialized_end=182,
)


_POKERESPONSE = _descriptor.Descriptor(
  name='PokeResponse',
  full_name='PokeResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ok', full_name='PokeResponse.ok', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=184,
  serialized_end=210,
)


_SETNEXTREQUEST = _descriptor.Descriptor(
  name='SetNextRequest',
  full_name='SetNextRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='member', full_name='SetNextRequest.member', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=212,
  serialized_end=253,
)


_SETNEXTRESPONSE = _descriptor.Descriptor(
  name='SetNextResponse',
  full_name='SetNextResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ok', full_name='SetNextResponse.ok', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=255,
  serialized_end=284,
)

_JOINREQUEST.fields_by_name['member'].message_type = _MEMBER
_POKEREQUEST.fields_by_name['from'].message_type = _MEMBER
_SETNEXTREQUEST.fields_by_name['member'].message_type = _MEMBER
DESCRIPTOR.message_types_by_name['Member'] = _MEMBER
DESCRIPTOR.message_types_by_name['JoinRequest'] = _JOINREQUEST
DESCRIPTOR.message_types_by_name['JoinResponse'] = _JOINRESPONSE
DESCRIPTOR.message_types_by_name['PokeRequest'] = _POKEREQUEST
DESCRIPTOR.message_types_by_name['PokeResponse'] = _POKERESPONSE
DESCRIPTOR.message_types_by_name['SetNextRequest'] = _SETNEXTREQUEST
DESCRIPTOR.message_types_by_name['SetNextResponse'] = _SETNEXTRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Member = _reflection.GeneratedProtocolMessageType('Member', (_message.Message,), dict(
  DESCRIPTOR = _MEMBER,
  __module__ = 'guruguru_pb2'
  # @@protoc_insertion_point(class_scope:Member)
  ))
_sym_db.RegisterMessage(Member)

JoinRequest = _reflection.GeneratedProtocolMessageType('JoinRequest', (_message.Message,), dict(
  DESCRIPTOR = _JOINREQUEST,
  __module__ = 'guruguru_pb2'
  # @@protoc_insertion_point(class_scope:JoinRequest)
  ))
_sym_db.RegisterMessage(JoinRequest)

JoinResponse = _reflection.GeneratedProtocolMessageType('JoinResponse', (_message.Message,), dict(
  DESCRIPTOR = _JOINRESPONSE,
  __module__ = 'guruguru_pb2'
  # @@protoc_insertion_point(class_scope:JoinResponse)
  ))
_sym_db.RegisterMessage(JoinResponse)

PokeRequest = _reflection.GeneratedProtocolMessageType('PokeRequest', (_message.Message,), dict(
  DESCRIPTOR = _POKEREQUEST,
  __module__ = 'guruguru_pb2'
  # @@protoc_insertion_point(class_scope:PokeRequest)
  ))
_sym_db.RegisterMessage(PokeRequest)

PokeResponse = _reflection.GeneratedProtocolMessageType('PokeResponse', (_message.Message,), dict(
  DESCRIPTOR = _POKERESPONSE,
  __module__ = 'guruguru_pb2'
  # @@protoc_insertion_point(class_scope:PokeResponse)
  ))
_sym_db.RegisterMessage(PokeResponse)

SetNextRequest = _reflection.GeneratedProtocolMessageType('SetNextRequest', (_message.Message,), dict(
  DESCRIPTOR = _SETNEXTREQUEST,
  __module__ = 'guruguru_pb2'
  # @@protoc_insertion_point(class_scope:SetNextRequest)
  ))
_sym_db.RegisterMessage(SetNextRequest)

SetNextResponse = _reflection.GeneratedProtocolMessageType('SetNextResponse', (_message.Message,), dict(
  DESCRIPTOR = _SETNEXTRESPONSE,
  __module__ = 'guruguru_pb2'
  # @@protoc_insertion_point(class_scope:SetNextResponse)
  ))
_sym_db.RegisterMessage(SetNextResponse)



_BOSSSERVICE = _descriptor.ServiceDescriptor(
  name='BossService',
  full_name='BossService',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=286,
  serialized_end=336,
  methods=[
  _descriptor.MethodDescriptor(
    name='Join',
    full_name='BossService.Join',
    index=0,
    containing_service=None,
    input_type=_JOINREQUEST,
    output_type=_JOINRESPONSE,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_BOSSSERVICE)

DESCRIPTOR.services_by_name['BossService'] = _BOSSSERVICE


_MEMBERSERVICE = _descriptor.ServiceDescriptor(
  name='MemberService',
  full_name='MemberService',
  file=DESCRIPTOR,
  index=1,
  options=None,
  serialized_start=338,
  serialized_end=436,
  methods=[
  _descriptor.MethodDescriptor(
    name='Poke',
    full_name='MemberService.Poke',
    index=0,
    containing_service=None,
    input_type=_POKEREQUEST,
    output_type=_POKERESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SetNext',
    full_name='MemberService.SetNext',
    index=1,
    containing_service=None,
    input_type=_SETNEXTREQUEST,
    output_type=_SETNEXTRESPONSE,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_MEMBERSERVICE)

DESCRIPTOR.services_by_name['MemberService'] = _MEMBERSERVICE

# @@protoc_insertion_point(module_scope)
