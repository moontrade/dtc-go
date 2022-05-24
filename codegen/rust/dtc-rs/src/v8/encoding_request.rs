// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-24 10:12:33.526761 +0800 WITA m=+0.004576126
use super::*;
use crate::message::*;

const ENCODING_REQUEST_FIXED_SIZE: usize = 16;

/// size              u16           = EncodingRequestFixedSize  (16)
/// type              u16           = ENCODING_REQUEST  (6)
/// protocol_version  i32           = CURRENT_VERSION  (8)
/// encoding          EncodingEnum  = BINARY_ENCODING  (0)
/// protocol_type     string4       = "DTC"
const ENCODING_REQUEST_FIXED_DEFAULT: [u8; 16] = [16, 0, 6, 0, 8, 0, 0, 0, 0, 0, 0, 0, 68, 84, 67, 0];

/// Requirements: Not required for Servers. Required for Clients if the Client
/// needs to discover the encoding the Server uses.
///
/// The EncodingRequestFixed message is a message requesting to change the
/// DTC encoding for messages.
///
/// For the procedure to work with this message, refer to Encoding Request
/// Sequence.
pub trait EncodingRequest {
    /// The protocol version supported by the Client. Automatically set by constructor.
    /// The protocol version supported by the Client. Automatically set by constructor.
    fn protocol_version(&self) -> i32;

    /// The DTC message encoding the Client is requesting the Server to use.
    fn encoding(&self) -> EncodingEnum;

    /// The ProtocolType field needs to be set to the text string "DTC".
    ///
    /// This field is automatically set with the binary encoding data structures.
    /// This field is automatically set with the binary encoding data structures.
    ///
    /// This field is used for the Server to know that it is communicating with
    /// a DTC compliant Client.
    fn protocol_type(&self) -> &str;

    /// The protocol version supported by the Client. Automatically set by constructor.
    /// The protocol version supported by the Client. Automatically set by constructor.
    fn set_protocol_version(&mut self, value: i32) -> &mut Self;

    /// The DTC message encoding the Client is requesting the Server to use.
    fn set_encoding(&mut self, value: EncodingEnum) -> &mut Self;

    /// The ProtocolType field needs to be set to the text string "DTC".
    ///
    /// This field is automatically set with the binary encoding data structures.
    /// This field is automatically set with the binary encoding data structures.
    ///
    /// This field is used for the Server to know that it is communicating with
    /// a DTC compliant Client.
    fn set_protocol_type(&mut self, value: &str) -> &mut Self;

    fn copy_to(&self, to: &mut impl EncodingRequest) {
        to.set_protocol_version(self.protocol_version());
        to.set_encoding(self.encoding());
        to.set_protocol_type(self.protocol_type());
    }
}

/// Requirements: Not required for Servers. Required for Clients if the Client
/// needs to discover the encoding the Server uses.
///
/// The EncodingRequestFixed message is a message requesting to change the
/// DTC encoding for messages.
///
/// For the procedure to work with this message, refer to Encoding Request
/// Sequence.
pub struct EncodingRequestFixed {
    data: *const EncodingRequestFixedData
}

pub struct EncodingRequestFixedUnsafe {
    data: *const EncodingRequestFixedData
}

#[repr(packed, C)]
pub struct EncodingRequestFixedData {
    size: u16,
    r#type: u16,
    protocol_version: i32,
    encoding: EncodingEnum,
    protocol_type: [u8; 4],
}

impl EncodingRequestFixedData {
    pub fn new() -> Self {
        Self {
            size: 16u16.to_le(),
            r#type: ENCODING_REQUEST.to_le(),
            protocol_version: CURRENT_VERSION.to_le(),
            encoding: EncodingEnum::BinaryEncoding.to_le(),
            protocol_type: [68,84,67,0],
        }
    }
}

unsafe impl Send for EncodingRequestFixed {}
unsafe impl Send for EncodingRequestFixedUnsafe {}
unsafe impl Send for EncodingRequestFixedData {}

impl Drop for EncodingRequestFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for EncodingRequestFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for EncodingRequestFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for EncodingRequestFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for EncodingRequestFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for EncodingRequestFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for EncodingRequestFixed {
    type Target = EncodingRequestFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for EncodingRequestFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for EncodingRequestFixedUnsafe {
    type Target = EncodingRequestFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for EncodingRequestFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for EncodingRequestFixed {
    type Safe = EncodingRequestFixed;
    type Unsafe = EncodingRequestFixedUnsafe;
    type Data = EncodingRequestFixedData;
    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, EncodingRequestFixedData::new())
        }
    }

    #[inline]
    fn to_safe(self) -> Self::Safe {
        self
    }

    #[inline]
    fn size(&self) -> u16 {
        u16::from_le(self.size)
    }

    #[inline]
    fn r#type(&self) -> u16 {
        u16::from_le(self.r#type)
    }

    #[inline]
    fn base_size(&self) -> u16 {
        u16::from_le(self.size)
    }

    #[inline]
    fn capacity(&self) -> u16 {
        u16::from_le(self.size)
    }

    #[inline]
    unsafe fn as_ptr(&self) -> *const Self::Data {
        self.data
    }

    #[inline]
    unsafe fn from_raw_parts(data: *const u8, _: usize) -> Self {
        Self {
            data: data as *const EncodingRequestFixedData
        }
    }

}

impl crate::Message for EncodingRequestFixedUnsafe {
    type Safe = EncodingRequestFixed;
    type Unsafe = EncodingRequestFixedUnsafe;
    type Data = EncodingRequestFixedData;
    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, EncodingRequestFixedData::new())
        }
    }

    #[inline]
    fn to_safe(self) -> Self::Safe {
        let mut s = Self::Safe::new();
        self.copy_to(&mut s);
        s
    }

    #[inline]
    fn size(&self) -> u16 {
        u16::from_le(self.size)
    }

    #[inline]
    fn r#type(&self) -> u16 {
        u16::from_le(self.r#type)
    }

    #[inline]
    fn base_size(&self) -> u16 {
        u16::from_le(self.size)
    }

    #[inline]
    fn capacity(&self) -> u16 {
        u16::from_le(self.size)
    }

    #[inline]
    unsafe fn as_ptr(&self) -> *const Self::Data {
        self.data
    }

    #[inline]
    unsafe fn from_raw_parts(data: *const u8, _: usize) -> Self {
        Self {
            data: data as *const EncodingRequestFixedData
        }
    }

}

/// Requirements: Not required for Servers. Required for Clients if the Client
/// needs to discover the encoding the Server uses.
///
/// The EncodingRequestFixed message is a message requesting to change the
/// DTC encoding for messages.
///
/// For the procedure to work with this message, refer to Encoding Request
/// Sequence.
impl EncodingRequest for EncodingRequestFixed {
    /// The protocol version supported by the Client. Automatically set by constructor.
    /// The protocol version supported by the Client. Automatically set by constructor.
    fn protocol_version(&self) -> i32 {
        i32::from_le(self.protocol_version)
    }

    /// The DTC message encoding the Client is requesting the Server to use.
    fn encoding(&self) -> EncodingEnum {
        EncodingEnum::from_le(self.encoding)
    }

    /// The ProtocolType field needs to be set to the text string "DTC".
    ///
    /// This field is automatically set with the binary encoding data structures.
    /// This field is automatically set with the binary encoding data structures.
    ///
    /// This field is used for the Server to know that it is communicating with
    /// a DTC compliant Client.
    fn protocol_type(&self) -> &str {
        crate::get_fixed(&self.protocol_type[..])
    }

    /// The protocol version supported by the Client. Automatically set by constructor.
    /// The protocol version supported by the Client. Automatically set by constructor.
    fn set_protocol_version(&mut self, value: i32) -> &mut Self {
        self.protocol_version = value.to_le();
        self
    }


    /// The DTC message encoding the Client is requesting the Server to use.
    fn set_encoding(&mut self, value: EncodingEnum) -> &mut Self {
        self.encoding = unsafe { core::mem::transmute((value as i32).to_le()) };
        self
    }


    /// The ProtocolType field needs to be set to the text string "DTC".
    ///
    /// This field is automatically set with the binary encoding data structures.
    /// This field is automatically set with the binary encoding data structures.
    ///
    /// This field is used for the Server to know that it is communicating with
    /// a DTC compliant Client.
    fn set_protocol_type(&mut self, value: &str) -> &mut Self {
        crate::set_fixed(&mut self.protocol_type[..], value);
        self
    }

}

/// Requirements: Not required for Servers. Required for Clients if the Client
/// needs to discover the encoding the Server uses.
///
/// The EncodingRequestFixed message is a message requesting to change the
/// DTC encoding for messages.
///
/// For the procedure to work with this message, refer to Encoding Request
/// Sequence.
impl EncodingRequest for EncodingRequestFixedUnsafe {
    /// The protocol version supported by the Client. Automatically set by constructor.
    /// The protocol version supported by the Client. Automatically set by constructor.
    fn protocol_version(&self) -> i32 {
        if self.is_out_of_bounds(8) {
            CURRENT_VERSION.to_le()
        } else {
            i32::from_le(self.protocol_version)
        }
    }

    /// The DTC message encoding the Client is requesting the Server to use.
    fn encoding(&self) -> EncodingEnum {
        if self.is_out_of_bounds(12) {
            EncodingEnum::BinaryEncoding.to_le()
        } else {
            EncodingEnum::from_le(self.encoding)
        }
    }

    /// The ProtocolType field needs to be set to the text string "DTC".
    ///
    /// This field is automatically set with the binary encoding data structures.
    /// This field is automatically set with the binary encoding data structures.
    ///
    /// This field is used for the Server to know that it is communicating with
    /// a DTC compliant Client.
    fn protocol_type(&self) -> &str {
        if self.is_out_of_bounds(16) {
            ""
        } else {
            crate::get_fixed(&self.protocol_type[..])
        }
    }

    /// The protocol version supported by the Client. Automatically set by constructor.
    /// The protocol version supported by the Client. Automatically set by constructor.
    fn set_protocol_version(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.protocol_version = value.to_le();
        }
        self
    }


    /// The DTC message encoding the Client is requesting the Server to use.
    fn set_encoding(&mut self, value: EncodingEnum) -> &mut Self {
        if !self.is_out_of_bounds(12) {
            self.encoding = unsafe { core::mem::transmute((value as i32).to_le()) };
        }
        self
    }


    /// The ProtocolType field needs to be set to the text string "DTC".
    ///
    /// This field is automatically set with the binary encoding data structures.
    /// This field is automatically set with the binary encoding data structures.
    ///
    /// This field is used for the Server to know that it is communicating with
    /// a DTC compliant Client.
    fn set_protocol_type(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(16) {
            crate::set_fixed(&mut self.protocol_type[..], value);
        }
        self
    }

}
