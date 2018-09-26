# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/admin/common.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from flyteidl.core import execution_pb2 as flyteidl_dot_core_dot_execution__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl/admin/common.proto',
  package='flyteidl.admin',
  syntax='proto3',
  serialized_pb=_b('\n\x1b\x66lyteidl/admin/common.proto\x12\x0e\x66lyteidl.admin\x1a\x1d\x66lyteidl/core/execution.proto\";\n\nIdentifier\x12\x0f\n\x07project\x18\x01 \x01(\t\x12\x0e\n\x06\x64omain\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\"W\n\x15IdentifierListRequest\x12\x0f\n\x07project\x18\x01 \x01(\t\x12\x0e\n\x06\x64omain\x18\x02 \x01(\t\x12\r\n\x05limit\x18\x03 \x01(\r\x12\x0e\n\x06offset\x18\x04 \x01(\r\">\n\x0eIdentifierList\x12,\n\x08\x65ntities\x18\x01 \x03(\x0b\x32\x1a.flyteidl.admin.Identifier\"\x1f\n\x10ObjectGetRequest\x12\x0b\n\x03urn\x18\x01 \x01(\t\"m\n\x13ResourceListRequest\x12&\n\x02id\x18\x01 \x01(\x0b\x32\x1a.flyteidl.admin.Identifier\x12\r\n\x05limit\x18\x02 \x01(\r\x12\x0e\n\x06offset\x18\x03 \x01(\r\x12\x0f\n\x07\x66ilters\x18\x04 \x01(\t\"\xb3\x01\n\x0cNotification\x12/\n\x04type\x18\x01 \x01(\x0e\x32!.flyteidl.admin.Notification.Type\x12\x35\n\x06phases\x18\x02 \x03(\x0e\x32%.flyteidl.core.WorkflowExecutionPhase\";\n\x04Type\x12\r\n\tUNDEFINED\x10\x00\x12\t\n\x05\x45MAIL\x10\x01\x12\x0e\n\nPAGER_DUTY\x10\x02\x12\t\n\x05SLACK\x10\x03\x42\x33Z1github.com/lyft/flyteidl/gen/pb-go/flyteidl/adminb\x06proto3')
  ,
  dependencies=[flyteidl_dot_core_dot_execution__pb2.DESCRIPTOR,])



_NOTIFICATION_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='flyteidl.admin.Notification.Type',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNDEFINED', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='EMAIL', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PAGER_DUTY', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SLACK', index=3, number=3,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=557,
  serialized_end=616,
)
_sym_db.RegisterEnumDescriptor(_NOTIFICATION_TYPE)


_IDENTIFIER = _descriptor.Descriptor(
  name='Identifier',
  full_name='flyteidl.admin.Identifier',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='project', full_name='flyteidl.admin.Identifier.project', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='domain', full_name='flyteidl.admin.Identifier.domain', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='flyteidl.admin.Identifier.name', index=2,
      number=3, type=9, cpp_type=9, label=1,
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
  serialized_start=78,
  serialized_end=137,
)


_IDENTIFIERLISTREQUEST = _descriptor.Descriptor(
  name='IdentifierListRequest',
  full_name='flyteidl.admin.IdentifierListRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='project', full_name='flyteidl.admin.IdentifierListRequest.project', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='domain', full_name='flyteidl.admin.IdentifierListRequest.domain', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='limit', full_name='flyteidl.admin.IdentifierListRequest.limit', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='offset', full_name='flyteidl.admin.IdentifierListRequest.offset', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=139,
  serialized_end=226,
)


_IDENTIFIERLIST = _descriptor.Descriptor(
  name='IdentifierList',
  full_name='flyteidl.admin.IdentifierList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='entities', full_name='flyteidl.admin.IdentifierList.entities', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=228,
  serialized_end=290,
)


_OBJECTGETREQUEST = _descriptor.Descriptor(
  name='ObjectGetRequest',
  full_name='flyteidl.admin.ObjectGetRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='urn', full_name='flyteidl.admin.ObjectGetRequest.urn', index=0,
      number=1, type=9, cpp_type=9, label=1,
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
  serialized_start=292,
  serialized_end=323,
)


_RESOURCELISTREQUEST = _descriptor.Descriptor(
  name='ResourceListRequest',
  full_name='flyteidl.admin.ResourceListRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='flyteidl.admin.ResourceListRequest.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='limit', full_name='flyteidl.admin.ResourceListRequest.limit', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='offset', full_name='flyteidl.admin.ResourceListRequest.offset', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='filters', full_name='flyteidl.admin.ResourceListRequest.filters', index=3,
      number=4, type=9, cpp_type=9, label=1,
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
  serialized_start=325,
  serialized_end=434,
)


_NOTIFICATION = _descriptor.Descriptor(
  name='Notification',
  full_name='flyteidl.admin.Notification',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='flyteidl.admin.Notification.type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='phases', full_name='flyteidl.admin.Notification.phases', index=1,
      number=2, type=14, cpp_type=8, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _NOTIFICATION_TYPE,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=437,
  serialized_end=616,
)

_IDENTIFIERLIST.fields_by_name['entities'].message_type = _IDENTIFIER
_RESOURCELISTREQUEST.fields_by_name['id'].message_type = _IDENTIFIER
_NOTIFICATION.fields_by_name['type'].enum_type = _NOTIFICATION_TYPE
_NOTIFICATION.fields_by_name['phases'].enum_type = flyteidl_dot_core_dot_execution__pb2._WORKFLOWEXECUTIONPHASE
_NOTIFICATION_TYPE.containing_type = _NOTIFICATION
DESCRIPTOR.message_types_by_name['Identifier'] = _IDENTIFIER
DESCRIPTOR.message_types_by_name['IdentifierListRequest'] = _IDENTIFIERLISTREQUEST
DESCRIPTOR.message_types_by_name['IdentifierList'] = _IDENTIFIERLIST
DESCRIPTOR.message_types_by_name['ObjectGetRequest'] = _OBJECTGETREQUEST
DESCRIPTOR.message_types_by_name['ResourceListRequest'] = _RESOURCELISTREQUEST
DESCRIPTOR.message_types_by_name['Notification'] = _NOTIFICATION
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Identifier = _reflection.GeneratedProtocolMessageType('Identifier', (_message.Message,), dict(
  DESCRIPTOR = _IDENTIFIER,
  __module__ = 'flyteidl.admin.common_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.Identifier)
  ))
_sym_db.RegisterMessage(Identifier)

IdentifierListRequest = _reflection.GeneratedProtocolMessageType('IdentifierListRequest', (_message.Message,), dict(
  DESCRIPTOR = _IDENTIFIERLISTREQUEST,
  __module__ = 'flyteidl.admin.common_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.IdentifierListRequest)
  ))
_sym_db.RegisterMessage(IdentifierListRequest)

IdentifierList = _reflection.GeneratedProtocolMessageType('IdentifierList', (_message.Message,), dict(
  DESCRIPTOR = _IDENTIFIERLIST,
  __module__ = 'flyteidl.admin.common_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.IdentifierList)
  ))
_sym_db.RegisterMessage(IdentifierList)

ObjectGetRequest = _reflection.GeneratedProtocolMessageType('ObjectGetRequest', (_message.Message,), dict(
  DESCRIPTOR = _OBJECTGETREQUEST,
  __module__ = 'flyteidl.admin.common_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.ObjectGetRequest)
  ))
_sym_db.RegisterMessage(ObjectGetRequest)

ResourceListRequest = _reflection.GeneratedProtocolMessageType('ResourceListRequest', (_message.Message,), dict(
  DESCRIPTOR = _RESOURCELISTREQUEST,
  __module__ = 'flyteidl.admin.common_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.ResourceListRequest)
  ))
_sym_db.RegisterMessage(ResourceListRequest)

Notification = _reflection.GeneratedProtocolMessageType('Notification', (_message.Message,), dict(
  DESCRIPTOR = _NOTIFICATION,
  __module__ = 'flyteidl.admin.common_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.Notification)
  ))
_sym_db.RegisterMessage(Notification)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z1github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin'))
# @@protoc_insertion_point(module_scope)