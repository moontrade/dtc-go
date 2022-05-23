// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-23 21:07:20.84591 +0800 WITA m=+0.009860584
use super::constants::*;
use super::alias::*;
use super::enums::*;
use crate::message::*;

const TRADE_ACCOUNTS_REQUEST_FIXED_SIZE: usize = 8;

/// size        u16  = TradeAccountsRequestFixedSize  (8)
/// r#type      u16  = TRADE_ACCOUNTS_REQUEST  (400)
/// request_id  i32  = 0
const TRADE_ACCOUNTS_REQUEST_FIXED_DEFAULT: [u8; 8] = [8, 0, 144, 1, 0, 0, 0, 0];

/// This is a message from the Client to the Server to request all of the
/// account identifiers for the logged in Username.
///
/// If there are no accounts available, then the Server needs to respond with
/// at least one TradeAccountResponseVLS message containing an empty Trade
/// Account.
pub trait TradeAccountsRequest {
    /// A unique request identifier for this request.
    fn request_id(&self) -> i32;

    /// A unique request identifier for this request.
    fn set_request_id(&mut self, value: i32) -> &mut Self;

    fn copy_to(&self, to: &mut impl TradeAccountsRequest) {
        to.set_request_id(self.request_id());
    }
}

/// This is a message from the Client to the Server to request all of the
/// account identifiers for the logged in Username.
///
/// If there are no accounts available, then the Server needs to respond with
/// at least one TradeAccountResponseVLS message containing an empty Trade
/// Account.
pub struct TradeAccountsRequestFixed {
    data: *const TradeAccountsRequestFixedData
}

pub struct TradeAccountsRequestFixedUnsafe {
    data: *const TradeAccountsRequestFixedData
}

#[repr(packed, C)]
pub struct TradeAccountsRequestFixedData {
    size: u16,
    r#type: u16,
    request_id: i32,
}

impl TradeAccountsRequestFixedData {
    pub fn new() -> Self {
        Self {
            size: 8u16.to_le(),
            r#type: TRADE_ACCOUNTS_REQUEST.to_le(),
            request_id: 0,
        }
    }
}

unsafe impl Send for TradeAccountsRequestFixed {}
unsafe impl Send for TradeAccountsRequestFixedUnsafe {}
unsafe impl Send for TradeAccountsRequestFixedData {}

impl Drop for TradeAccountsRequestFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for TradeAccountsRequestFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for TradeAccountsRequestFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for TradeAccountsRequestFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for TradeAccountsRequestFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for TradeAccountsRequestFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for TradeAccountsRequestFixed {
    type Target = TradeAccountsRequestFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for TradeAccountsRequestFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for TradeAccountsRequestFixedUnsafe {
    type Target = TradeAccountsRequestFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for TradeAccountsRequestFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for TradeAccountsRequestFixed {
    type Safe = TradeAccountsRequestFixed;
    type Unsafe = TradeAccountsRequestFixedUnsafe;
    type Data = TradeAccountsRequestFixedData;
    const BASE_SIZE: usize = 8;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, TradeAccountsRequestFixedData::new())
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
            data: data as *const TradeAccountsRequestFixedData
        }
    }

}

impl crate::Message for TradeAccountsRequestFixedUnsafe {
    type Safe = TradeAccountsRequestFixed;
    type Unsafe = TradeAccountsRequestFixedUnsafe;
    type Data = TradeAccountsRequestFixedData;
    const BASE_SIZE: usize = 8;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, TradeAccountsRequestFixedData::new())
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
            data: data as *const TradeAccountsRequestFixedData
        }
    }

}

/// This is a message from the Client to the Server to request all of the
/// account identifiers for the logged in Username.
///
/// If there are no accounts available, then the Server needs to respond with
/// at least one TradeAccountResponseVLS message containing an empty Trade
/// Account.
impl TradeAccountsRequest for TradeAccountsRequestFixed {
    /// A unique request identifier for this request.
    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    /// A unique request identifier for this request.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }

}

/// This is a message from the Client to the Server to request all of the
/// account identifiers for the logged in Username.
///
/// If there are no accounts available, then the Server needs to respond with
/// at least one TradeAccountResponseVLS message containing an empty Trade
/// Account.
impl TradeAccountsRequest for TradeAccountsRequestFixedUnsafe {
    /// A unique request identifier for this request.
    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(8) {
            0
        } else {
            i32::from_le(self.request_id)
        }
    }

    /// A unique request identifier for this request.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.request_id = value.to_le();
        }
        self
    }

}

