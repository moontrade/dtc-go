// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-23 21:07:20.84591 +0800 WITA m=+0.009860584
use super::constants::*;
use super::alias::*;
use super::enums::*;
use crate::message::*;

const HISTORICAL_PRICE_DATA_REJECT_VLS_SIZE: usize = 20;

const HISTORICAL_PRICE_DATA_REJECT_FIXED_SIZE: usize = 108;

/// size                   u16                                      = HistoricalPriceDataRejectVLSSize  (20)
/// r#type                 u16                                      = HISTORICAL_PRICE_DATA_REJECT  (802)
/// base_size              u16                                      = HistoricalPriceDataRejectVLSSize  (20)
/// request_id             i32                                      = 0
/// reject_text            string                                   = ""
/// reject_reason_code     HistoricalPriceDataRejectReasonCodeEnum  = HPDR_UNSET  (0)
/// retry_time_in_seconds  u16                                      = 0
const HISTORICAL_PRICE_DATA_REJECT_VLS_DEFAULT: [u8; 20] = [20, 0, 34, 3, 20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

/// size                   u16                                      = HistoricalPriceDataRejectFixedSize  (108)
/// r#type                 u16                                      = HISTORICAL_PRICE_DATA_REJECT  (802)
/// request_id             i32                                      = 0
/// reject_text            string96                                 = ""
/// reject_reason_code     HistoricalPriceDataRejectReasonCodeEnum  = HPDR_UNSET  (0)
/// retry_time_in_seconds  u16                                      = 0
const HISTORICAL_PRICE_DATA_REJECT_FIXED_DEFAULT: [u8; 108] = [108, 0, 34, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

/// When the Server rejects a historical price data request from the Client,
/// a HistoricalPriceDataRejectVLS message will be sent.
///
/// This message is never compressed.
pub trait HistoricalPriceDataReject {
    /// The numeric identifier from the historical price data request that this
    /// response is in response to.
    fn request_id(&self) -> i32;

    /// Text reason for rejection.
    fn reject_text(&self) -> &str;

    /// Integer identifier identifying the reason for the rejection. For the text
    /// reason, refer to the RejectText field.
    fn reject_reason_code(&self) -> HistoricalPriceDataRejectReasonCodeEnum;

    /// This is an optional field from the Server. This field will normally be
    /// zero.
    ///
    /// If a retry is intended to be performed, the server may give an indication
    /// of how long to wait in seconds. This field indicates that.
    ///
    /// This field is not recommended to be used. If it is used, it is really
    /// an indication of a substandard Server.
    fn retry_time_in_seconds(&self) -> u16;

    /// The numeric identifier from the historical price data request that this
    /// response is in response to.
    fn set_request_id(&mut self, value: i32) -> &mut Self;

    /// Text reason for rejection.
    fn set_reject_text(&mut self, value: &str) -> &mut Self;

    /// Integer identifier identifying the reason for the rejection. For the text
    /// reason, refer to the RejectText field.
    fn set_reject_reason_code(&mut self, value: HistoricalPriceDataRejectReasonCodeEnum) -> &mut Self;

    /// This is an optional field from the Server. This field will normally be
    /// zero.
    ///
    /// If a retry is intended to be performed, the server may give an indication
    /// of how long to wait in seconds. This field indicates that.
    ///
    /// This field is not recommended to be used. If it is used, it is really
    /// an indication of a substandard Server.
    fn set_retry_time_in_seconds(&mut self, value: u16) -> &mut Self;

    fn copy_to(&self, to: &mut impl HistoricalPriceDataReject) {
        to.set_request_id(self.request_id());
        to.set_reject_text(self.reject_text());
        to.set_reject_reason_code(self.reject_reason_code());
        to.set_retry_time_in_seconds(self.retry_time_in_seconds());
    }
}

/// When the Server rejects a historical price data request from the Client,
/// a HistoricalPriceDataRejectVLS message will be sent.
///
/// This message is never compressed.
pub struct HistoricalPriceDataRejectVLS {
    data: *const HistoricalPriceDataRejectVLSData,
    capacity: usize
}

pub struct HistoricalPriceDataRejectVLSUnsafe {
    data: *const HistoricalPriceDataRejectVLSData,
    capacity: usize
}

#[repr(packed, C)]
pub struct HistoricalPriceDataRejectVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    request_id: i32,
    reject_text: VLS,
    reject_reason_code: HistoricalPriceDataRejectReasonCodeEnum,
    retry_time_in_seconds: u16,
}

/// When the Server rejects a historical price data request from the Client,
/// a HistoricalPriceDataRejectVLS message will be sent.
///
/// This message is never compressed.
pub struct HistoricalPriceDataRejectFixed {
    data: *const HistoricalPriceDataRejectFixedData
}

pub struct HistoricalPriceDataRejectFixedUnsafe {
    data: *const HistoricalPriceDataRejectFixedData
}

#[repr(packed, C)]
pub struct HistoricalPriceDataRejectFixedData {
    size: u16,
    r#type: u16,
    request_id: i32,
    reject_text: [u8; 96],
    reject_reason_code: HistoricalPriceDataRejectReasonCodeEnum,
    retry_time_in_seconds: u16,
}

impl HistoricalPriceDataRejectVLSData {
    pub fn new() -> Self {
        Self {
            size: 20u16.to_le(),
            r#type: HISTORICAL_PRICE_DATA_REJECT.to_le(),
            base_size: 20u16.to_le(),
            request_id: 0,
            reject_text: crate::message::VLS::new(),
            reject_reason_code: HistoricalPriceDataRejectReasonCodeEnum::HpdrUnset.to_le(),
            retry_time_in_seconds: 0,
        }
    }
}

impl HistoricalPriceDataRejectFixedData {
    pub fn new() -> Self {
        Self {
            size: 108u16.to_le(),
            r#type: HISTORICAL_PRICE_DATA_REJECT.to_le(),
            request_id: 0,
            reject_text: [0; 96],
            reject_reason_code: HistoricalPriceDataRejectReasonCodeEnum::HpdrUnset.to_le(),
            retry_time_in_seconds: 0,
        }
    }
}

unsafe impl Send for HistoricalPriceDataRejectFixed {}
unsafe impl Send for HistoricalPriceDataRejectFixedUnsafe {}
unsafe impl Send for HistoricalPriceDataRejectFixedData {}
unsafe impl Send for HistoricalPriceDataRejectVLS {}
unsafe impl Send for HistoricalPriceDataRejectVLSUnsafe {}
unsafe impl Send for HistoricalPriceDataRejectVLSData {}

impl Drop for HistoricalPriceDataRejectFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for HistoricalPriceDataRejectFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for HistoricalPriceDataRejectVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for HistoricalPriceDataRejectVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for HistoricalPriceDataRejectFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for HistoricalPriceDataRejectFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for HistoricalPriceDataRejectVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for HistoricalPriceDataRejectVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for HistoricalPriceDataRejectFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for HistoricalPriceDataRejectFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for HistoricalPriceDataRejectVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for HistoricalPriceDataRejectVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for HistoricalPriceDataRejectFixed {
    type Target = HistoricalPriceDataRejectFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalPriceDataRejectFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for HistoricalPriceDataRejectFixedUnsafe {
    type Target = HistoricalPriceDataRejectFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalPriceDataRejectFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for HistoricalPriceDataRejectVLS {
    type Target = HistoricalPriceDataRejectVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalPriceDataRejectVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for HistoricalPriceDataRejectVLSUnsafe {
    type Target = HistoricalPriceDataRejectVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalPriceDataRejectVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for HistoricalPriceDataRejectFixed {
    type Safe = HistoricalPriceDataRejectFixed;
    type Unsafe = HistoricalPriceDataRejectFixedUnsafe;
    type Data = HistoricalPriceDataRejectFixedData;
    const BASE_SIZE: usize = 108;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, HistoricalPriceDataRejectFixedData::new())
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
            data: data as *const HistoricalPriceDataRejectFixedData
        }
    }

}

impl crate::Message for HistoricalPriceDataRejectFixedUnsafe {
    type Safe = HistoricalPriceDataRejectFixed;
    type Unsafe = HistoricalPriceDataRejectFixedUnsafe;
    type Data = HistoricalPriceDataRejectFixedData;
    const BASE_SIZE: usize = 108;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, HistoricalPriceDataRejectFixedData::new())
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
            data: data as *const HistoricalPriceDataRejectFixedData
        }
    }

}

impl crate::Message for HistoricalPriceDataRejectVLS {
    type Safe = HistoricalPriceDataRejectVLS;
    type Unsafe = HistoricalPriceDataRejectVLSUnsafe;
    type Data = HistoricalPriceDataRejectVLSData;
    const BASE_SIZE: usize = 20;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, HistoricalPriceDataRejectVLSData::new()),
            capacity: Self::DEFAULT_CAPACITY
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
        u16::from_le(self.base_size)
    }

    #[inline]
    fn capacity(&self) -> u16 {
        self.capacity as u16
    }

    #[inline]
    unsafe fn as_ptr(&self) -> *const Self::Data {
        self.data
    }

    #[inline]
    unsafe fn from_raw_parts(data: *const u8, capacity: usize) -> Self {
        Self {
            data: data as *const HistoricalPriceDataRejectVLSData,
            capacity
        }
    }

}

impl crate::VLSMessage for HistoricalPriceDataRejectVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const HistoricalPriceDataRejectVLSData;
    }

    #[inline]
    fn set_capacity(&mut self, capacity: u16) {
        self.capacity = capacity as usize;
    }

    #[inline]
    fn set_size(&mut self, size: u16) {
        self.size = size.to_le();
    }

}impl crate::Message for HistoricalPriceDataRejectVLSUnsafe {
    type Safe = HistoricalPriceDataRejectVLS;
    type Unsafe = HistoricalPriceDataRejectVLSUnsafe;
    type Data = HistoricalPriceDataRejectVLSData;
    const BASE_SIZE: usize = 20;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, HistoricalPriceDataRejectVLSData::new()),
            capacity: Self::DEFAULT_CAPACITY
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
        u16::from_le(self.base_size)
    }

    #[inline]
    fn capacity(&self) -> u16 {
        self.capacity as u16
    }

    #[inline]
    unsafe fn as_ptr(&self) -> *const Self::Data {
        self.data
    }

    #[inline]
    unsafe fn from_raw_parts(data: *const u8, capacity: usize) -> Self {
        Self {
            data: data as *const HistoricalPriceDataRejectVLSData,
            capacity
        }
    }

}

impl crate::VLSMessage for HistoricalPriceDataRejectVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const HistoricalPriceDataRejectVLSData;
    }

    #[inline]
    fn set_capacity(&mut self, capacity: u16) {
        self.capacity = capacity as usize;
    }

    #[inline]
    fn set_size(&mut self, size: u16) {
        self.size = size.to_le();
    }

}/// When the Server rejects a historical price data request from the Client,
/// a HistoricalPriceDataRejectVLS message will be sent.
///
/// This message is never compressed.
impl HistoricalPriceDataReject for HistoricalPriceDataRejectVLS {
    /// The numeric identifier from the historical price data request that this
    /// response is in response to.
    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    /// Text reason for rejection.
    fn reject_text(&self) -> &str {
        crate::get_vls(self, self.reject_text)
    }

    /// Integer identifier identifying the reason for the rejection. For the text
    /// reason, refer to the RejectText field.
    fn reject_reason_code(&self) -> HistoricalPriceDataRejectReasonCodeEnum {
        HistoricalPriceDataRejectReasonCodeEnum::from_le(self.reject_reason_code)
    }

    /// This is an optional field from the Server. This field will normally be
    /// zero.
    ///
    /// If a retry is intended to be performed, the server may give an indication
    /// of how long to wait in seconds. This field indicates that.
    ///
    /// This field is not recommended to be used. If it is used, it is really
    /// an indication of a substandard Server.
    fn retry_time_in_seconds(&self) -> u16 {
        u16::from_le(self.retry_time_in_seconds)
    }

    /// The numeric identifier from the historical price data request that this
    /// response is in response to.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }


    /// Text reason for rejection.
    fn set_reject_text(&mut self, value: &str) -> &mut Self {
        self.reject_text = crate::set_vls(self, self.reject_text, value);
        self
    }


    /// Integer identifier identifying the reason for the rejection. For the text
    /// reason, refer to the RejectText field.
    fn set_reject_reason_code(&mut self, value: HistoricalPriceDataRejectReasonCodeEnum) -> &mut Self {
        self.reject_reason_code = unsafe { core::mem::transmute((value as i16).to_le()) };
        self
    }


    /// This is an optional field from the Server. This field will normally be
    /// zero.
    ///
    /// If a retry is intended to be performed, the server may give an indication
    /// of how long to wait in seconds. This field indicates that.
    ///
    /// This field is not recommended to be used. If it is used, it is really
    /// an indication of a substandard Server.
    fn set_retry_time_in_seconds(&mut self, value: u16) -> &mut Self {
        self.retry_time_in_seconds = value.to_le();
        self
    }

}

/// When the Server rejects a historical price data request from the Client,
/// a HistoricalPriceDataRejectVLS message will be sent.
///
/// This message is never compressed.
impl HistoricalPriceDataReject for HistoricalPriceDataRejectVLSUnsafe {
    /// The numeric identifier from the historical price data request that this
    /// response is in response to.
    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(12) {
            0
        } else {
            i32::from_le(self.request_id)
        }
    }

    /// Text reason for rejection.
    fn reject_text(&self) -> &str {
        if self.is_out_of_bounds(16) {
            ""
        } else {
            crate::get_vls(self, self.reject_text)
        }
    }

    /// Integer identifier identifying the reason for the rejection. For the text
    /// reason, refer to the RejectText field.
    fn reject_reason_code(&self) -> HistoricalPriceDataRejectReasonCodeEnum {
        if self.is_out_of_bounds(18) {
            HistoricalPriceDataRejectReasonCodeEnum::HpdrUnset.to_le()
        } else {
            HistoricalPriceDataRejectReasonCodeEnum::from_le(self.reject_reason_code)
        }
    }

    /// This is an optional field from the Server. This field will normally be
    /// zero.
    ///
    /// If a retry is intended to be performed, the server may give an indication
    /// of how long to wait in seconds. This field indicates that.
    ///
    /// This field is not recommended to be used. If it is used, it is really
    /// an indication of a substandard Server.
    fn retry_time_in_seconds(&self) -> u16 {
        if self.is_out_of_bounds(20) {
            0
        } else {
            u16::from_le(self.retry_time_in_seconds)
        }
    }

    /// The numeric identifier from the historical price data request that this
    /// response is in response to.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(12) {
            self.request_id = value.to_le();
        }
        self
    }


    /// Text reason for rejection.
    fn set_reject_text(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(16) {
            self.reject_text = crate::set_vls(self, self.reject_text, value);
        }
        self
    }


    /// Integer identifier identifying the reason for the rejection. For the text
    /// reason, refer to the RejectText field.
    fn set_reject_reason_code(&mut self, value: HistoricalPriceDataRejectReasonCodeEnum) -> &mut Self {
        if !self.is_out_of_bounds(18) {
            self.reject_reason_code = unsafe { core::mem::transmute((value as i16).to_le()) };
        }
        self
    }


    /// This is an optional field from the Server. This field will normally be
    /// zero.
    ///
    /// If a retry is intended to be performed, the server may give an indication
    /// of how long to wait in seconds. This field indicates that.
    ///
    /// This field is not recommended to be used. If it is used, it is really
    /// an indication of a substandard Server.
    fn set_retry_time_in_seconds(&mut self, value: u16) -> &mut Self {
        if !self.is_out_of_bounds(20) {
            self.retry_time_in_seconds = value.to_le();
        }
        self
    }

}

/// When the Server rejects a historical price data request from the Client,
/// a HistoricalPriceDataRejectVLS message will be sent.
///
/// This message is never compressed.
impl HistoricalPriceDataReject for HistoricalPriceDataRejectFixed {
    /// The numeric identifier from the historical price data request that this
    /// response is in response to.
    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    /// Text reason for rejection.
    fn reject_text(&self) -> &str {
        crate::get_fixed(&self.reject_text[..])
    }

    /// Integer identifier identifying the reason for the rejection. For the text
    /// reason, refer to the RejectText field.
    fn reject_reason_code(&self) -> HistoricalPriceDataRejectReasonCodeEnum {
        HistoricalPriceDataRejectReasonCodeEnum::from_le(self.reject_reason_code)
    }

    /// This is an optional field from the Server. This field will normally be
    /// zero.
    ///
    /// If a retry is intended to be performed, the server may give an indication
    /// of how long to wait in seconds. This field indicates that.
    ///
    /// This field is not recommended to be used. If it is used, it is really
    /// an indication of a substandard Server.
    fn retry_time_in_seconds(&self) -> u16 {
        u16::from_le(self.retry_time_in_seconds)
    }

    /// The numeric identifier from the historical price data request that this
    /// response is in response to.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }


    /// Text reason for rejection.
    fn set_reject_text(&mut self, value: &str) -> &mut Self {
        crate::set_fixed(&mut self.reject_text[..], value);
        self
    }


    /// Integer identifier identifying the reason for the rejection. For the text
    /// reason, refer to the RejectText field.
    fn set_reject_reason_code(&mut self, value: HistoricalPriceDataRejectReasonCodeEnum) -> &mut Self {
        self.reject_reason_code = unsafe { core::mem::transmute((value as i16).to_le()) };
        self
    }


    /// This is an optional field from the Server. This field will normally be
    /// zero.
    ///
    /// If a retry is intended to be performed, the server may give an indication
    /// of how long to wait in seconds. This field indicates that.
    ///
    /// This field is not recommended to be used. If it is used, it is really
    /// an indication of a substandard Server.
    fn set_retry_time_in_seconds(&mut self, value: u16) -> &mut Self {
        self.retry_time_in_seconds = value.to_le();
        self
    }

}

/// When the Server rejects a historical price data request from the Client,
/// a HistoricalPriceDataRejectVLS message will be sent.
///
/// This message is never compressed.
impl HistoricalPriceDataReject for HistoricalPriceDataRejectFixedUnsafe {
    /// The numeric identifier from the historical price data request that this
    /// response is in response to.
    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(8) {
            0
        } else {
            i32::from_le(self.request_id)
        }
    }

    /// Text reason for rejection.
    fn reject_text(&self) -> &str {
        if self.is_out_of_bounds(104) {
            ""
        } else {
            crate::get_fixed(&self.reject_text[..])
        }
    }

    /// Integer identifier identifying the reason for the rejection. For the text
    /// reason, refer to the RejectText field.
    fn reject_reason_code(&self) -> HistoricalPriceDataRejectReasonCodeEnum {
        if self.is_out_of_bounds(106) {
            HistoricalPriceDataRejectReasonCodeEnum::HpdrUnset.to_le()
        } else {
            HistoricalPriceDataRejectReasonCodeEnum::from_le(self.reject_reason_code)
        }
    }

    /// This is an optional field from the Server. This field will normally be
    /// zero.
    ///
    /// If a retry is intended to be performed, the server may give an indication
    /// of how long to wait in seconds. This field indicates that.
    ///
    /// This field is not recommended to be used. If it is used, it is really
    /// an indication of a substandard Server.
    fn retry_time_in_seconds(&self) -> u16 {
        if self.is_out_of_bounds(108) {
            0
        } else {
            u16::from_le(self.retry_time_in_seconds)
        }
    }

    /// The numeric identifier from the historical price data request that this
    /// response is in response to.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.request_id = value.to_le();
        }
        self
    }


    /// Text reason for rejection.
    fn set_reject_text(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(104) {
            crate::set_fixed(&mut self.reject_text[..], value);
        }
        self
    }


    /// Integer identifier identifying the reason for the rejection. For the text
    /// reason, refer to the RejectText field.
    fn set_reject_reason_code(&mut self, value: HistoricalPriceDataRejectReasonCodeEnum) -> &mut Self {
        if !self.is_out_of_bounds(106) {
            self.reject_reason_code = unsafe { core::mem::transmute((value as i16).to_le()) };
        }
        self
    }


    /// This is an optional field from the Server. This field will normally be
    /// zero.
    ///
    /// If a retry is intended to be performed, the server may give an indication
    /// of how long to wait in seconds. This field indicates that.
    ///
    /// This field is not recommended to be used. If it is used, it is really
    /// an indication of a substandard Server.
    fn set_retry_time_in_seconds(&mut self, value: u16) -> &mut Self {
        if !self.is_out_of_bounds(108) {
            self.retry_time_in_seconds = value.to_le();
        }
        self
    }

}

