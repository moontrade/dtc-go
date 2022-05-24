// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-24 10:12:33.526761 +0800 WITA m=+0.004576126
use super::*;
use crate::message::*;

const HISTORICAL_PRICE_DATA_RESPONSE_TRAILER_FIXED_SIZE: usize = 16;

/// size                         u16                          = HistoricalPriceDataResponseTrailerFixedSize  (16)
/// type                         u16                          = HISTORICAL_PRICE_DATA_RESPONSE_TRAILER  (807)
/// request_id                   i32                          = 0
/// final_record_last_date_time  DateTimeWithMicrosecondsInt  = 0
const HISTORICAL_PRICE_DATA_RESPONSE_TRAILER_FIXED_DEFAULT: [u8; 16] = [16, 0, 39, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

pub trait HistoricalPriceDataResponseTrailer {
    fn request_id(&self) -> i32;

    fn final_record_last_date_time(&self) -> DateTimeWithMicrosecondsInt;

    fn set_request_id(&mut self, value: i32) -> &mut Self;

    fn set_final_record_last_date_time(&mut self, value: DateTimeWithMicrosecondsInt) -> &mut Self;

    fn copy_to(&self, to: &mut impl HistoricalPriceDataResponseTrailer) {
        to.set_request_id(self.request_id());
        to.set_final_record_last_date_time(self.final_record_last_date_time());
    }
}

pub struct HistoricalPriceDataResponseTrailerFixed {
    data: *const HistoricalPriceDataResponseTrailerFixedData
}

pub struct HistoricalPriceDataResponseTrailerFixedUnsafe {
    data: *const HistoricalPriceDataResponseTrailerFixedData
}

#[repr(packed, C)]
pub struct HistoricalPriceDataResponseTrailerFixedData {
    size: u16,
    r#type: u16,
    request_id: i32,
    final_record_last_date_time: DateTimeWithMicrosecondsInt,
}

impl HistoricalPriceDataResponseTrailerFixedData {
    pub fn new() -> Self {
        Self {
            size: 16u16.to_le(),
            r#type: HISTORICAL_PRICE_DATA_RESPONSE_TRAILER.to_le(),
            request_id: 0i32,
            final_record_last_date_time: 0i64,
        }
    }
}

unsafe impl Send for HistoricalPriceDataResponseTrailerFixed {}
unsafe impl Send for HistoricalPriceDataResponseTrailerFixedUnsafe {}
unsafe impl Send for HistoricalPriceDataResponseTrailerFixedData {}

impl Drop for HistoricalPriceDataResponseTrailerFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for HistoricalPriceDataResponseTrailerFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for HistoricalPriceDataResponseTrailerFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for HistoricalPriceDataResponseTrailerFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for HistoricalPriceDataResponseTrailerFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for HistoricalPriceDataResponseTrailerFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for HistoricalPriceDataResponseTrailerFixed {
    type Target = HistoricalPriceDataResponseTrailerFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalPriceDataResponseTrailerFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for HistoricalPriceDataResponseTrailerFixedUnsafe {
    type Target = HistoricalPriceDataResponseTrailerFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalPriceDataResponseTrailerFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for HistoricalPriceDataResponseTrailerFixed {
    type Safe = HistoricalPriceDataResponseTrailerFixed;
    type Unsafe = HistoricalPriceDataResponseTrailerFixedUnsafe;
    type Data = HistoricalPriceDataResponseTrailerFixedData;
    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, HistoricalPriceDataResponseTrailerFixedData::new())
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
            data: data as *const HistoricalPriceDataResponseTrailerFixedData
        }
    }

}

impl crate::Message for HistoricalPriceDataResponseTrailerFixedUnsafe {
    type Safe = HistoricalPriceDataResponseTrailerFixed;
    type Unsafe = HistoricalPriceDataResponseTrailerFixedUnsafe;
    type Data = HistoricalPriceDataResponseTrailerFixedData;
    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, HistoricalPriceDataResponseTrailerFixedData::new())
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
            data: data as *const HistoricalPriceDataResponseTrailerFixedData
        }
    }

}

impl HistoricalPriceDataResponseTrailer for HistoricalPriceDataResponseTrailerFixed {
    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    fn final_record_last_date_time(&self) -> DateTimeWithMicrosecondsInt {
        i64::from_le(self.final_record_last_date_time)
    }

    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }


    fn set_final_record_last_date_time(&mut self, value: DateTimeWithMicrosecondsInt) -> &mut Self {
        self.final_record_last_date_time = value.to_le();
        self
    }

}

impl HistoricalPriceDataResponseTrailer for HistoricalPriceDataResponseTrailerFixedUnsafe {
    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(8) {
            0i32
        } else {
            i32::from_le(self.request_id)
        }
    }

    fn final_record_last_date_time(&self) -> DateTimeWithMicrosecondsInt {
        if self.is_out_of_bounds(16) {
            0i64
        } else {
            i64::from_le(self.final_record_last_date_time)
        }
    }

    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.request_id = value.to_le();
        }
        self
    }


    fn set_final_record_last_date_time(&mut self, value: DateTimeWithMicrosecondsInt) -> &mut Self {
        if !self.is_out_of_bounds(16) {
            self.final_record_last_date_time = value.to_le();
        }
        self
    }

}
