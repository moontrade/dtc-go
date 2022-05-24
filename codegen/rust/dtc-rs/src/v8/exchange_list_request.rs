// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-24 10:12:33.526761 +0800 WITA m=+0.004576126
use super::*;
use crate::message::*;

const EXCHANGE_LIST_REQUEST_FIXED_SIZE: usize = 8;

/// size        u16  = ExchangeListRequestFixedSize  (8)
/// type        u16  = EXCHANGE_LIST_REQUEST  (500)
/// request_id  i32  = 0
const EXCHANGE_LIST_REQUEST_FIXED_DEFAULT: [u8; 8] = [8, 0, 244, 1, 0, 0, 0, 0];

/// This is a message from the Client to the Server to request a list of all
/// available exchanges from the Server.
///
/// The server will respond with a separate ExchangeListResponseVLS message
/// for each exchange.
///
/// In the case where the Server does not specify an exchange with its symbols,
/// then the Server should provide a single response with an empty Exchange.
/// then the Server should provide a single response with an empty Exchange.
pub trait ExchangeListRequest {
    /// The unique identifier for this request. This same identifier will be returned
    /// in the ExchangeListResponseVLS message.
    fn request_id(&self) -> i32;

    /// The unique identifier for this request. This same identifier will be returned
    /// in the ExchangeListResponseVLS message.
    fn set_request_id(&mut self, value: i32) -> &mut Self;

    fn copy_to(&self, to: &mut impl ExchangeListRequest) {
        to.set_request_id(self.request_id());
    }
}

/// This is a message from the Client to the Server to request a list of all
/// available exchanges from the Server.
///
/// The server will respond with a separate ExchangeListResponseVLS message
/// for each exchange.
///
/// In the case where the Server does not specify an exchange with its symbols,
/// then the Server should provide a single response with an empty Exchange.
/// then the Server should provide a single response with an empty Exchange.
pub struct ExchangeListRequestFixed {
    data: *const ExchangeListRequestFixedData
}

pub struct ExchangeListRequestFixedUnsafe {
    data: *const ExchangeListRequestFixedData
}

#[repr(packed, C)]
pub struct ExchangeListRequestFixedData {
    size: u16,
    r#type: u16,
    request_id: i32,
}

impl ExchangeListRequestFixedData {
    pub fn new() -> Self {
        Self {
            size: 8u16.to_le(),
            r#type: EXCHANGE_LIST_REQUEST.to_le(),
            request_id: 0i32,
        }
    }
}

unsafe impl Send for ExchangeListRequestFixed {}
unsafe impl Send for ExchangeListRequestFixedUnsafe {}
unsafe impl Send for ExchangeListRequestFixedData {}

impl Drop for ExchangeListRequestFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for ExchangeListRequestFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for ExchangeListRequestFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for ExchangeListRequestFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for ExchangeListRequestFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for ExchangeListRequestFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for ExchangeListRequestFixed {
    type Target = ExchangeListRequestFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for ExchangeListRequestFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for ExchangeListRequestFixedUnsafe {
    type Target = ExchangeListRequestFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for ExchangeListRequestFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for ExchangeListRequestFixed {
    type Safe = ExchangeListRequestFixed;
    type Unsafe = ExchangeListRequestFixedUnsafe;
    type Data = ExchangeListRequestFixedData;
    const BASE_SIZE: usize = 8;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, ExchangeListRequestFixedData::new())
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
            data: data as *const ExchangeListRequestFixedData
        }
    }

}

impl crate::Message for ExchangeListRequestFixedUnsafe {
    type Safe = ExchangeListRequestFixed;
    type Unsafe = ExchangeListRequestFixedUnsafe;
    type Data = ExchangeListRequestFixedData;
    const BASE_SIZE: usize = 8;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, ExchangeListRequestFixedData::new())
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
            data: data as *const ExchangeListRequestFixedData
        }
    }

}

/// This is a message from the Client to the Server to request a list of all
/// available exchanges from the Server.
///
/// The server will respond with a separate ExchangeListResponseVLS message
/// for each exchange.
///
/// In the case where the Server does not specify an exchange with its symbols,
/// then the Server should provide a single response with an empty Exchange.
/// then the Server should provide a single response with an empty Exchange.
impl ExchangeListRequest for ExchangeListRequestFixed {
    /// The unique identifier for this request. This same identifier will be returned
    /// in the ExchangeListResponseVLS message.
    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    /// The unique identifier for this request. This same identifier will be returned
    /// in the ExchangeListResponseVLS message.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }

}

/// This is a message from the Client to the Server to request a list of all
/// available exchanges from the Server.
///
/// The server will respond with a separate ExchangeListResponseVLS message
/// for each exchange.
///
/// In the case where the Server does not specify an exchange with its symbols,
/// then the Server should provide a single response with an empty Exchange.
/// then the Server should provide a single response with an empty Exchange.
impl ExchangeListRequest for ExchangeListRequestFixedUnsafe {
    /// The unique identifier for this request. This same identifier will be returned
    /// in the ExchangeListResponseVLS message.
    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(8) {
            0i32
        } else {
            i32::from_le(self.request_id)
        }
    }

    /// The unique identifier for this request. This same identifier will be returned
    /// in the ExchangeListResponseVLS message.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.request_id = value.to_le();
        }
        self
    }

}
