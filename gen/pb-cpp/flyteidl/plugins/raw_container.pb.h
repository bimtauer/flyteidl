// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: flyteidl/plugins/raw_container.proto

#ifndef PROTOBUF_INCLUDED_flyteidl_2fplugins_2fraw_5fcontainer_2eproto
#define PROTOBUF_INCLUDED_flyteidl_2fplugins_2fraw_5fcontainer_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3007000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3007000 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/inlined_string_field.h>
#include <google/protobuf/metadata.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/generated_enum_reflection.h>
#include <google/protobuf/unknown_field_set.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_flyteidl_2fplugins_2fraw_5fcontainer_2eproto

// Internal implementation detail -- do not use these members.
struct TableStruct_flyteidl_2fplugins_2fraw_5fcontainer_2eproto {
  static const ::google::protobuf::internal::ParseTableField entries[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::AuxillaryParseTableField aux[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::ParseTable schema[1]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::FieldMetadata field_metadata[];
  static const ::google::protobuf::internal::SerializationTable serialization_table[];
  static const ::google::protobuf::uint32 offsets[];
};
void AddDescriptors_flyteidl_2fplugins_2fraw_5fcontainer_2eproto();
namespace flyteidl {
namespace plugins {
class CoPilot;
class CoPilotDefaultTypeInternal;
extern CoPilotDefaultTypeInternal _CoPilot_default_instance_;
}  // namespace plugins
}  // namespace flyteidl
namespace google {
namespace protobuf {
template<> ::flyteidl::plugins::CoPilot* Arena::CreateMaybeMessage<::flyteidl::plugins::CoPilot>(Arena*);
}  // namespace protobuf
}  // namespace google
namespace flyteidl {
namespace plugins {

enum CoPilot_MetadataFormat {
  CoPilot_MetadataFormat_JSON = 0,
  CoPilot_MetadataFormat_YAML = 1,
  CoPilot_MetadataFormat_PROTO = 2,
  CoPilot_MetadataFormat_CoPilot_MetadataFormat_INT_MIN_SENTINEL_DO_NOT_USE_ = std::numeric_limits<::google::protobuf::int32>::min(),
  CoPilot_MetadataFormat_CoPilot_MetadataFormat_INT_MAX_SENTINEL_DO_NOT_USE_ = std::numeric_limits<::google::protobuf::int32>::max()
};
bool CoPilot_MetadataFormat_IsValid(int value);
const CoPilot_MetadataFormat CoPilot_MetadataFormat_MetadataFormat_MIN = CoPilot_MetadataFormat_JSON;
const CoPilot_MetadataFormat CoPilot_MetadataFormat_MetadataFormat_MAX = CoPilot_MetadataFormat_PROTO;
const int CoPilot_MetadataFormat_MetadataFormat_ARRAYSIZE = CoPilot_MetadataFormat_MetadataFormat_MAX + 1;

const ::google::protobuf::EnumDescriptor* CoPilot_MetadataFormat_descriptor();
inline const ::std::string& CoPilot_MetadataFormat_Name(CoPilot_MetadataFormat value) {
  return ::google::protobuf::internal::NameOfEnum(
    CoPilot_MetadataFormat_descriptor(), value);
}
inline bool CoPilot_MetadataFormat_Parse(
    const ::std::string& name, CoPilot_MetadataFormat* value) {
  return ::google::protobuf::internal::ParseNamedEnum<CoPilot_MetadataFormat>(
    CoPilot_MetadataFormat_descriptor(), name, value);
}
// ===================================================================

class CoPilot final :
    public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:flyteidl.plugins.CoPilot) */ {
 public:
  CoPilot();
  virtual ~CoPilot();

  CoPilot(const CoPilot& from);

  inline CoPilot& operator=(const CoPilot& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  CoPilot(CoPilot&& from) noexcept
    : CoPilot() {
    *this = ::std::move(from);
  }

  inline CoPilot& operator=(CoPilot&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor() {
    return default_instance().GetDescriptor();
  }
  static const CoPilot& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const CoPilot* internal_default_instance() {
    return reinterpret_cast<const CoPilot*>(
               &_CoPilot_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  void Swap(CoPilot* other);
  friend void swap(CoPilot& a, CoPilot& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline CoPilot* New() const final {
    return CreateMaybeMessage<CoPilot>(nullptr);
  }

  CoPilot* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<CoPilot>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const CoPilot& from);
  void MergeFrom(const CoPilot& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  #if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  static const char* _InternalParse(const char* begin, const char* end, void* object, ::google::protobuf::internal::ParseContext* ctx);
  ::google::protobuf::internal::ParseFunc _ParseFunc() const final { return _InternalParse; }
  #else
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) final;
  #endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const final;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      ::google::protobuf::uint8* target) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(CoPilot* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return nullptr;
  }
  inline void* MaybeArenaPtr() const {
    return nullptr;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  typedef CoPilot_MetadataFormat MetadataFormat;
  static const MetadataFormat JSON =
    CoPilot_MetadataFormat_JSON;
  static const MetadataFormat YAML =
    CoPilot_MetadataFormat_YAML;
  static const MetadataFormat PROTO =
    CoPilot_MetadataFormat_PROTO;
  static inline bool MetadataFormat_IsValid(int value) {
    return CoPilot_MetadataFormat_IsValid(value);
  }
  static const MetadataFormat MetadataFormat_MIN =
    CoPilot_MetadataFormat_MetadataFormat_MIN;
  static const MetadataFormat MetadataFormat_MAX =
    CoPilot_MetadataFormat_MetadataFormat_MAX;
  static const int MetadataFormat_ARRAYSIZE =
    CoPilot_MetadataFormat_MetadataFormat_ARRAYSIZE;
  static inline const ::google::protobuf::EnumDescriptor*
  MetadataFormat_descriptor() {
    return CoPilot_MetadataFormat_descriptor();
  }
  static inline const ::std::string& MetadataFormat_Name(MetadataFormat value) {
    return CoPilot_MetadataFormat_Name(value);
  }
  static inline bool MetadataFormat_Parse(const ::std::string& name,
      MetadataFormat* value) {
    return CoPilot_MetadataFormat_Parse(name, value);
  }

  // accessors -------------------------------------------------------

  // string input_path = 1;
  void clear_input_path();
  static const int kInputPathFieldNumber = 1;
  const ::std::string& input_path() const;
  void set_input_path(const ::std::string& value);
  #if LANG_CXX11
  void set_input_path(::std::string&& value);
  #endif
  void set_input_path(const char* value);
  void set_input_path(const char* value, size_t size);
  ::std::string* mutable_input_path();
  ::std::string* release_input_path();
  void set_allocated_input_path(::std::string* input_path);

  // uint32 output_path = 2;
  void clear_output_path();
  static const int kOutputPathFieldNumber = 2;
  ::google::protobuf::uint32 output_path() const;
  void set_output_path(::google::protobuf::uint32 value);

  // .flyteidl.plugins.CoPilot.MetadataFormat format = 3;
  void clear_format();
  static const int kFormatFieldNumber = 3;
  ::flyteidl::plugins::CoPilot_MetadataFormat format() const;
  void set_format(::flyteidl::plugins::CoPilot_MetadataFormat value);

  // @@protoc_insertion_point(class_scope:flyteidl.plugins.CoPilot)
 private:
  class HasBitSetters;

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::ArenaStringPtr input_path_;
  ::google::protobuf::uint32 output_path_;
  int format_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  friend struct ::TableStruct_flyteidl_2fplugins_2fraw_5fcontainer_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// CoPilot

// string input_path = 1;
inline void CoPilot::clear_input_path() {
  input_path_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& CoPilot::input_path() const {
  // @@protoc_insertion_point(field_get:flyteidl.plugins.CoPilot.input_path)
  return input_path_.GetNoArena();
}
inline void CoPilot::set_input_path(const ::std::string& value) {
  
  input_path_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:flyteidl.plugins.CoPilot.input_path)
}
#if LANG_CXX11
inline void CoPilot::set_input_path(::std::string&& value) {
  
  input_path_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:flyteidl.plugins.CoPilot.input_path)
}
#endif
inline void CoPilot::set_input_path(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  
  input_path_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:flyteidl.plugins.CoPilot.input_path)
}
inline void CoPilot::set_input_path(const char* value, size_t size) {
  
  input_path_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:flyteidl.plugins.CoPilot.input_path)
}
inline ::std::string* CoPilot::mutable_input_path() {
  
  // @@protoc_insertion_point(field_mutable:flyteidl.plugins.CoPilot.input_path)
  return input_path_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* CoPilot::release_input_path() {
  // @@protoc_insertion_point(field_release:flyteidl.plugins.CoPilot.input_path)
  
  return input_path_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void CoPilot::set_allocated_input_path(::std::string* input_path) {
  if (input_path != nullptr) {
    
  } else {
    
  }
  input_path_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), input_path);
  // @@protoc_insertion_point(field_set_allocated:flyteidl.plugins.CoPilot.input_path)
}

// uint32 output_path = 2;
inline void CoPilot::clear_output_path() {
  output_path_ = 0u;
}
inline ::google::protobuf::uint32 CoPilot::output_path() const {
  // @@protoc_insertion_point(field_get:flyteidl.plugins.CoPilot.output_path)
  return output_path_;
}
inline void CoPilot::set_output_path(::google::protobuf::uint32 value) {
  
  output_path_ = value;
  // @@protoc_insertion_point(field_set:flyteidl.plugins.CoPilot.output_path)
}

// .flyteidl.plugins.CoPilot.MetadataFormat format = 3;
inline void CoPilot::clear_format() {
  format_ = 0;
}
inline ::flyteidl::plugins::CoPilot_MetadataFormat CoPilot::format() const {
  // @@protoc_insertion_point(field_get:flyteidl.plugins.CoPilot.format)
  return static_cast< ::flyteidl::plugins::CoPilot_MetadataFormat >(format_);
}
inline void CoPilot::set_format(::flyteidl::plugins::CoPilot_MetadataFormat value) {
  
  format_ = value;
  // @@protoc_insertion_point(field_set:flyteidl.plugins.CoPilot.format)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)

}  // namespace plugins
}  // namespace flyteidl

namespace google {
namespace protobuf {

template <> struct is_proto_enum< ::flyteidl::plugins::CoPilot_MetadataFormat> : ::std::true_type {};
template <>
inline const EnumDescriptor* GetEnumDescriptor< ::flyteidl::plugins::CoPilot_MetadataFormat>() {
  return ::flyteidl::plugins::CoPilot_MetadataFormat_descriptor();
}

}  // namespace protobuf
}  // namespace google

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // PROTOBUF_INCLUDED_flyteidl_2fplugins_2fraw_5fcontainer_2eproto
