// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-24 10:12:33.526761 +0800 WITA m=+0.004576126
use super::*;
use crate::message::*;

const JOURNAL_ENTRIES_REQUEST_FIXED_SIZE: usize = 16;

/// size             u16       = JournalEntriesRequestFixedSize  (16)
/// type             u16       = JOURNAL_ENTRIES_REQUEST  (704)
/// request_id       i32       = 0
/// start_date_time  DateTime  = 0
const JOURNAL_ENTRIES_REQUEST_FIXED_DEFAULT: [u8; 16] = [16, 0, 192, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

pub trait JournalEntriesRequest {
    fn request_id(&self) -> i32;

    fn start_date_time(&self) -> DateTime;

    fn set_request_id(&mut self, value: i32) -> &mut Self;

    fn set_start_date_time(&mut self, value: DateTime) -> &mut Self;

    fn copy_to(&self, to: &mut impl JournalEntriesRequest) {
        to.set_request_id(self.request_id());
        to.set_start_date_time(self.start_date_time());
    }
}

pub struct JournalEntriesRequestFixed {
    data: *const JournalEntriesRequestFixedData
}

pub struct JournalEntriesRequestFixedUnsafe {
    data: *const JournalEntriesRequestFixedData
}

#[repr(packed, C)]
pub struct JournalEntriesRequestFixedData {
    size: u16,
    r#type: u16,
    request_id: i32,
    start_date_time: DateTime,
}

impl JournalEntriesRequestFixedData {
    pub fn new() -> Self {
        Self {
            size: 16u16.to_le(),
            r#type: JOURNAL_ENTRIES_REQUEST.to_le(),
            request_id: 0i32,
            start_date_time: 0i64,
        }
    }
}

unsafe impl Send for JournalEntriesRequestFixed {}
unsafe impl Send for JournalEntriesRequestFixedUnsafe {}
unsafe impl Send for JournalEntriesRequestFixedData {}

impl Drop for JournalEntriesRequestFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for JournalEntriesRequestFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for JournalEntriesRequestFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for JournalEntriesRequestFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for JournalEntriesRequestFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for JournalEntriesRequestFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for JournalEntriesRequestFixed {
    type Target = JournalEntriesRequestFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for JournalEntriesRequestFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for JournalEntriesRequestFixedUnsafe {
    type Target = JournalEntriesRequestFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for JournalEntriesRequestFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for JournalEntriesRequestFixed {
    type Safe = JournalEntriesRequestFixed;
    type Unsafe = JournalEntriesRequestFixedUnsafe;
    type Data = JournalEntriesRequestFixedData;
    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, JournalEntriesRequestFixedData::new())
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
            data: data as *const JournalEntriesRequestFixedData
        }
    }

}

impl crate::Message for JournalEntriesRequestFixedUnsafe {
    type Safe = JournalEntriesRequestFixed;
    type Unsafe = JournalEntriesRequestFixedUnsafe;
    type Data = JournalEntriesRequestFixedData;
    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, JournalEntriesRequestFixedData::new())
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
            data: data as *const JournalEntriesRequestFixedData
        }
    }

}

impl JournalEntriesRequest for JournalEntriesRequestFixed {
    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    fn start_date_time(&self) -> DateTime {
        i64::from_le(self.start_date_time)
    }

    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }


    fn set_start_date_time(&mut self, value: DateTime) -> &mut Self {
        self.start_date_time = value.to_le();
        self
    }

}

impl JournalEntriesRequest for JournalEntriesRequestFixedUnsafe {
    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(8) {
            0i32
        } else {
            i32::from_le(self.request_id)
        }
    }

    fn start_date_time(&self) -> DateTime {
        if self.is_out_of_bounds(16) {
            0i64
        } else {
            i64::from_le(self.start_date_time)
        }
    }

    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.request_id = value.to_le();
        }
        self
    }


    fn set_start_date_time(&mut self, value: DateTime) -> &mut Self {
        if !self.is_out_of_bounds(16) {
            self.start_date_time = value.to_le();
        }
        self
    }

}
