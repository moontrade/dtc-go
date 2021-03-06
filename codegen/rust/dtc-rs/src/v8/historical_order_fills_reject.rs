// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-27 08:32:53.747629 +0800 WITA m=+0.008238335
use super::*;

pub(crate) const HISTORICAL_ORDER_FILLS_REJECT_VLS_SIZE: usize = 16;

pub(crate) const HISTORICAL_ORDER_FILLS_REJECT_FIXED_SIZE: usize = 104;

/// size         u16     = HistoricalOrderFillsRejectVLSSize  (16)
/// type         u16     = HISTORICAL_ORDER_FILLS_REJECT  (308)
/// base_size    u16     = HistoricalOrderFillsRejectVLSSize  (16)
/// request_id   i32     = 0
/// reject_text  string  = ""
pub(crate) const HISTORICAL_ORDER_FILLS_REJECT_VLS_DEFAULT: [u8; 16] =
    [16, 0, 52, 1, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

/// size         u16       = HistoricalOrderFillsRejectFixedSize  (104)
/// type         u16       = HISTORICAL_ORDER_FILLS_REJECT  (308)
/// request_id   i32       = 0
/// reject_text  string96  = ""
pub(crate) const HISTORICAL_ORDER_FILLS_REJECT_FIXED_DEFAULT: [u8; 104] = [
    104, 0, 52, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0,
];

/// If the Server is unable to serve the request for a HistoricalOrderFillsRequestVLS
/// message received, for a reason other than there not being any historical
/// order fills, then send this message to the Client.
pub trait HistoricalOrderFillsReject: Message {
    type Safe: HistoricalOrderFillsReject;
    type Unsafe: HistoricalOrderFillsReject;

    /// This is set to the RequestID field sent in the HistoricalOrderFillsRequestVLS
    /// message.
    fn request_id(&self) -> i32;

    /// Free-form text indicating the reason for rejection.
    fn reject_text(&self) -> &str;

    /// This is set to the RequestID field sent in the HistoricalOrderFillsRequestVLS
    /// message.
    fn set_request_id(&mut self, value: i32) -> &mut Self;

    /// Free-form text indicating the reason for rejection.
    fn set_reject_text(&mut self, value: &str) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl HistoricalOrderFillsReject) {
        to.set_request_id(self.request_id());
        to.set_reject_text(self.reject_text());
    }

    #[inline]
    fn parse<F: Fn(Parsed<Self::Safe, Self::Unsafe>) -> Result<(), crate::Error>>(
        data: &[u8],
        f: F,
    ) -> Result<(), crate::Error> {
        if data.len() < 6 {
            return Err(crate::Error::Malformed("need more data"));
        }
        let size = unsafe { u16::from_le(*(data.as_ptr() as *const u16)) };
        let base_size = if Self::BASE_SIZE_OFFSET == 0 {
            size
        } else {
            let base_size = unsafe {
                u16::from_le(*(data.as_ptr().offset(Self::BASE_SIZE_OFFSET) as *const u16))
            };
            if base_size > size {
                return Err(crate::Error::Malformed("base_size is greater than size"));
            }
            base_size
        };
        if (base_size as usize) >= Self::BASE_SIZE {
            let msg = unsafe { Self::Safe::from_raw_parts(data.as_ptr(), size as usize) };
            let r = f(Parsed::Safe(&msg));
            core::mem::forget(msg);
            r
        } else {
            let msg = unsafe { Self::Unsafe::from_raw_parts(data.as_ptr(), size as usize) };
            let r = f(Parsed::Unsafe(&msg));
            core::mem::forget(msg);
            r
        }
    }
}

/// If the Server is unable to serve the request for a HistoricalOrderFillsRequestVLS
/// message received, for a reason other than there not being any historical
/// order fills, then send this message to the Client.
pub struct HistoricalOrderFillsRejectVLS {
    data: *const HistoricalOrderFillsRejectVLSData,
    capacity: usize,
}

pub struct HistoricalOrderFillsRejectVLSUnsafe {
    data: *const HistoricalOrderFillsRejectVLSData,
    capacity: usize,
}

#[repr(packed(8), C)]
pub struct HistoricalOrderFillsRejectVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    request_id: i32,
    reject_text: VLS,
}

/// If the Server is unable to serve the request for a HistoricalOrderFillsRequestVLS
/// message received, for a reason other than there not being any historical
/// order fills, then send this message to the Client.
pub struct HistoricalOrderFillsRejectFixed {
    data: *const HistoricalOrderFillsRejectFixedData,
}

pub struct HistoricalOrderFillsRejectFixedUnsafe {
    data: *const HistoricalOrderFillsRejectFixedData,
}

#[repr(packed(8), C)]
pub struct HistoricalOrderFillsRejectFixedData {
    size: u16,
    r#type: u16,
    request_id: i32,
    reject_text: [u8; 96],
}

impl HistoricalOrderFillsRejectVLSData {
    pub fn new() -> Self {
        Self {
            size: 16u16.to_le(),
            r#type: HISTORICAL_ORDER_FILLS_REJECT.to_le(),
            base_size: 16u16.to_le(),
            request_id: 0i32,
            reject_text: crate::message::VLS::new(),
        }
    }
}

impl HistoricalOrderFillsRejectFixedData {
    pub fn new() -> Self {
        Self {
            size: 104u16.to_le(),
            r#type: HISTORICAL_ORDER_FILLS_REJECT.to_le(),
            request_id: 0i32,
            reject_text: [0; 96],
        }
    }
}

unsafe impl Send for HistoricalOrderFillsRejectFixed {}
unsafe impl Send for HistoricalOrderFillsRejectFixedUnsafe {}
unsafe impl Send for HistoricalOrderFillsRejectFixedData {}
unsafe impl Send for HistoricalOrderFillsRejectVLS {}
unsafe impl Send for HistoricalOrderFillsRejectVLSUnsafe {}
unsafe impl Send for HistoricalOrderFillsRejectVLSData {}

impl Drop for HistoricalOrderFillsRejectFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for HistoricalOrderFillsRejectFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for HistoricalOrderFillsRejectVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for HistoricalOrderFillsRejectVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for HistoricalOrderFillsRejectFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for HistoricalOrderFillsRejectFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for HistoricalOrderFillsRejectVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for HistoricalOrderFillsRejectVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for HistoricalOrderFillsRejectFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for HistoricalOrderFillsRejectFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for HistoricalOrderFillsRejectVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for HistoricalOrderFillsRejectVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for HistoricalOrderFillsRejectFixed {
    type Target = HistoricalOrderFillsRejectFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalOrderFillsRejectFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for HistoricalOrderFillsRejectFixedUnsafe {
    type Target = HistoricalOrderFillsRejectFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalOrderFillsRejectFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for HistoricalOrderFillsRejectVLS {
    type Target = HistoricalOrderFillsRejectVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalOrderFillsRejectVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for HistoricalOrderFillsRejectVLSUnsafe {
    type Target = HistoricalOrderFillsRejectVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalOrderFillsRejectVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::fmt::Display for HistoricalOrderFillsRejectFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("HistoricalOrderFillsRejectFixed(size: {}, type: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.request_id(), self.reject_text()).as_str())
    }
}

impl core::fmt::Debug for HistoricalOrderFillsRejectFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("HistoricalOrderFillsRejectFixed(size: {}, type: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.request_id(), self.reject_text()).as_str())
    }
}

impl core::fmt::Display for HistoricalOrderFillsRejectFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("HistoricalOrderFillsRejectFixedUnsafe(size: {}, type: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.request_id(), self.reject_text()).as_str())
    }
}

impl core::fmt::Debug for HistoricalOrderFillsRejectFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("HistoricalOrderFillsRejectFixedUnsafe(size: {}, type: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.request_id(), self.reject_text()).as_str())
    }
}

impl core::fmt::Display for HistoricalOrderFillsRejectVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("HistoricalOrderFillsRejectVLS(size: {}, type: {}, base_size: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.reject_text()).as_str())
    }
}

impl core::fmt::Debug for HistoricalOrderFillsRejectVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("HistoricalOrderFillsRejectVLS(size: {}, type: {}, base_size: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.reject_text()).as_str())
    }
}

impl core::fmt::Display for HistoricalOrderFillsRejectVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("HistoricalOrderFillsRejectVLSUnsafe(size: {}, type: {}, base_size: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.reject_text()).as_str())
    }
}

impl core::fmt::Debug for HistoricalOrderFillsRejectVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("HistoricalOrderFillsRejectVLSUnsafe(size: {}, type: {}, base_size: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.reject_text()).as_str())
    }
}

impl crate::Message for HistoricalOrderFillsRejectFixed {
    type Data = HistoricalOrderFillsRejectFixedData;

    const TYPE: u16 = HISTORICAL_ORDER_FILLS_REJECT;
    const BASE_SIZE: usize = 104;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, HistoricalOrderFillsRejectFixedData::new()),
        }
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
            data: data as *const HistoricalOrderFillsRejectFixedData,
        }
    }
}
impl crate::Message for HistoricalOrderFillsRejectFixedUnsafe {
    type Data = HistoricalOrderFillsRejectFixedData;

    const TYPE: u16 = HISTORICAL_ORDER_FILLS_REJECT;
    const BASE_SIZE: usize = 104;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, HistoricalOrderFillsRejectFixedData::new()),
        }
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
            data: data as *const HistoricalOrderFillsRejectFixedData,
        }
    }
}
impl crate::Message for HistoricalOrderFillsRejectVLS {
    type Data = HistoricalOrderFillsRejectVLSData;

    const TYPE: u16 = HISTORICAL_ORDER_FILLS_REJECT;
    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, HistoricalOrderFillsRejectVLSData::new()),
            capacity: Self::DEFAULT_CAPACITY,
        }
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
            data: data as *const HistoricalOrderFillsRejectVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for HistoricalOrderFillsRejectVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const HistoricalOrderFillsRejectVLSData;
    }

    #[inline]
    fn set_capacity(&mut self, capacity: u16) {
        self.capacity = capacity as usize;
    }

    #[inline]
    fn set_size(&mut self, size: u16) {
        self.size = size.to_le();
    }
}
impl crate::Message for HistoricalOrderFillsRejectVLSUnsafe {
    type Data = HistoricalOrderFillsRejectVLSData;

    const TYPE: u16 = HISTORICAL_ORDER_FILLS_REJECT;
    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, HistoricalOrderFillsRejectVLSData::new()),
            capacity: Self::DEFAULT_CAPACITY,
        }
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
            data: data as *const HistoricalOrderFillsRejectVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for HistoricalOrderFillsRejectVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const HistoricalOrderFillsRejectVLSData;
    }

    #[inline]
    fn set_capacity(&mut self, capacity: u16) {
        self.capacity = capacity as usize;
    }

    #[inline]
    fn set_size(&mut self, size: u16) {
        self.size = size.to_le();
    }
}
/// If the Server is unable to serve the request for a HistoricalOrderFillsRequestVLS
/// message received, for a reason other than there not being any historical
/// order fills, then send this message to the Client.
impl HistoricalOrderFillsReject for HistoricalOrderFillsRejectVLS {
    type Safe = HistoricalOrderFillsRejectVLS;
    type Unsafe = HistoricalOrderFillsRejectVLSUnsafe;

    /// This is set to the RequestID field sent in the HistoricalOrderFillsRequestVLS
    /// message.
    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    /// Free-form text indicating the reason for rejection.
    fn reject_text(&self) -> &str {
        get_vls(self, self.reject_text)
    }

    /// This is set to the RequestID field sent in the HistoricalOrderFillsRequestVLS
    /// message.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }

    /// Free-form text indicating the reason for rejection.
    fn set_reject_text(&mut self, value: &str) -> &mut Self {
        self.reject_text = set_vls(self, self.reject_text, value);
        self
    }

    #[inline]
    fn clone_safe(&self) -> Self::Safe {
        let mut s = Self::Safe::new();
        self.copy_to(&mut s);
        s
    }

    #[inline]
    fn to_safe(self) -> Self::Safe {
        self
    }
}

/// If the Server is unable to serve the request for a HistoricalOrderFillsRequestVLS
/// message received, for a reason other than there not being any historical
/// order fills, then send this message to the Client.
impl HistoricalOrderFillsReject for HistoricalOrderFillsRejectVLSUnsafe {
    type Safe = HistoricalOrderFillsRejectVLS;
    type Unsafe = HistoricalOrderFillsRejectVLSUnsafe;

    /// This is set to the RequestID field sent in the HistoricalOrderFillsRequestVLS
    /// message.
    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(12) {
            0i32
        } else {
            i32::from_le(self.request_id)
        }
    }

    /// Free-form text indicating the reason for rejection.
    fn reject_text(&self) -> &str {
        if self.is_out_of_bounds(16) {
            ""
        } else {
            get_vls(self, self.reject_text)
        }
    }

    /// This is set to the RequestID field sent in the HistoricalOrderFillsRequestVLS
    /// message.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(12) {
            self.request_id = value.to_le();
        }
        self
    }

    /// Free-form text indicating the reason for rejection.
    fn set_reject_text(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(16) {
            self.reject_text = set_vls(self, self.reject_text, value);
        }
        self
    }

    #[inline]
    fn clone_safe(&self) -> Self::Safe {
        let mut s = Self::Safe::new();
        self.copy_to(&mut s);
        s
    }

    #[inline]
    fn to_safe(self) -> Self::Safe {
        let mut s = Self::Safe::new();
        self.copy_to(&mut s);
        s
    }
}

/// If the Server is unable to serve the request for a HistoricalOrderFillsRequestVLS
/// message received, for a reason other than there not being any historical
/// order fills, then send this message to the Client.
impl HistoricalOrderFillsReject for HistoricalOrderFillsRejectFixed {
    type Safe = HistoricalOrderFillsRejectFixed;
    type Unsafe = HistoricalOrderFillsRejectFixedUnsafe;

    /// This is set to the RequestID field sent in the HistoricalOrderFillsRequestVLS
    /// message.
    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    /// Free-form text indicating the reason for rejection.
    fn reject_text(&self) -> &str {
        get_fixed(&self.reject_text[..])
    }

    /// This is set to the RequestID field sent in the HistoricalOrderFillsRequestVLS
    /// message.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }

    /// Free-form text indicating the reason for rejection.
    fn set_reject_text(&mut self, value: &str) -> &mut Self {
        set_fixed(&mut self.reject_text[..], value);
        self
    }

    #[inline]
    fn clone_safe(&self) -> Self::Safe {
        let mut s = Self::Safe::new();
        self.copy_to(&mut s);
        s
    }

    #[inline]
    fn to_safe(self) -> Self::Safe {
        self
    }
}

/// If the Server is unable to serve the request for a HistoricalOrderFillsRequestVLS
/// message received, for a reason other than there not being any historical
/// order fills, then send this message to the Client.
impl HistoricalOrderFillsReject for HistoricalOrderFillsRejectFixedUnsafe {
    type Safe = HistoricalOrderFillsRejectFixed;
    type Unsafe = HistoricalOrderFillsRejectFixedUnsafe;

    /// This is set to the RequestID field sent in the HistoricalOrderFillsRequestVLS
    /// message.
    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(8) {
            0i32
        } else {
            i32::from_le(self.request_id)
        }
    }

    /// Free-form text indicating the reason for rejection.
    fn reject_text(&self) -> &str {
        if self.is_out_of_bounds(104) {
            ""
        } else {
            get_fixed(&self.reject_text[..])
        }
    }

    /// This is set to the RequestID field sent in the HistoricalOrderFillsRequestVLS
    /// message.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.request_id = value.to_le();
        }
        self
    }

    /// Free-form text indicating the reason for rejection.
    fn set_reject_text(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(104) {
            set_fixed(&mut self.reject_text[..], value);
        }
        self
    }

    #[inline]
    fn clone_safe(&self) -> Self::Safe {
        let mut s = Self::Safe::new();
        self.copy_to(&mut s);
        s
    }

    #[inline]
    fn to_safe(self) -> Self::Safe {
        let mut s = Self::Safe::new();
        self.copy_to(&mut s);
        s
    }
}

#[cfg(test)]
pub(crate) mod tests {
    use super::*;

    #[test]
    pub(crate) fn layout() {
        unsafe {
            assert_eq!(
                104usize,
                core::mem::size_of::<HistoricalOrderFillsRejectFixedData>(),
                "HistoricalOrderFillsRejectFixedData sizeof expected {:} but was {:}",
                104usize,
                core::mem::size_of::<HistoricalOrderFillsRejectFixedData>()
            );
            assert_eq!(
                104u16,
                HistoricalOrderFillsRejectFixed::new().size(),
                "HistoricalOrderFillsRejectFixed sizeof expected {:} but was {:}",
                104u16,
                HistoricalOrderFillsRejectFixed::new().size(),
            );
            assert_eq!(
                HISTORICAL_ORDER_FILLS_REJECT,
                HistoricalOrderFillsRejectFixed::new().r#type(),
                "HistoricalOrderFillsRejectFixed type expected {:} but was {:}",
                HISTORICAL_ORDER_FILLS_REJECT,
                HistoricalOrderFillsRejectFixed::new().r#type(),
            );
            assert_eq!(
                308u16,
                HistoricalOrderFillsRejectFixed::new().r#type(),
                "HistoricalOrderFillsRejectFixed type expected {:} but was {:}",
                308u16,
                HistoricalOrderFillsRejectFixed::new().r#type(),
            );
            let d = HistoricalOrderFillsRejectFixedData::new();
            let p = (&d as *const _ as *const u8).offset(0) as usize;
            assert_eq!(
                0usize,
                (core::ptr::addr_of!(d.size) as usize) - p,
                "size offset expected {:} but was {:}",
                0usize,
                (core::ptr::addr_of!(d.size) as usize) - p,
            );
            assert_eq!(
                2usize,
                (core::ptr::addr_of!(d.r#type) as usize) - p,
                "type offset expected {:} but was {:}",
                2usize,
                (core::ptr::addr_of!(d.r#type) as usize) - p,
            );
            assert_eq!(
                4usize,
                (core::ptr::addr_of!(d.request_id) as usize) - p,
                "request_id offset expected {:} but was {:}",
                4usize,
                (core::ptr::addr_of!(d.request_id) as usize) - p,
            );
            assert_eq!(
                8usize,
                (core::ptr::addr_of!(d.reject_text) as usize) - p,
                "reject_text offset expected {:} but was {:}",
                8usize,
                (core::ptr::addr_of!(d.reject_text) as usize) - p,
            );
        }
        unsafe {
            assert_eq!(
                16usize,
                core::mem::size_of::<HistoricalOrderFillsRejectVLSData>(),
                "HistoricalOrderFillsRejectVLSData sizeof expected {:} but was {:}",
                16usize,
                core::mem::size_of::<HistoricalOrderFillsRejectVLSData>()
            );
            assert_eq!(
                16u16,
                HistoricalOrderFillsRejectVLS::new().size(),
                "HistoricalOrderFillsRejectVLS sizeof expected {:} but was {:}",
                16u16,
                HistoricalOrderFillsRejectVLS::new().size(),
            );
            assert_eq!(
                HISTORICAL_ORDER_FILLS_REJECT,
                HistoricalOrderFillsRejectVLS::new().r#type(),
                "HistoricalOrderFillsRejectVLS type expected {:} but was {:}",
                HISTORICAL_ORDER_FILLS_REJECT,
                HistoricalOrderFillsRejectVLS::new().r#type(),
            );
            assert_eq!(
                308u16,
                HistoricalOrderFillsRejectVLS::new().r#type(),
                "HistoricalOrderFillsRejectVLS type expected {:} but was {:}",
                308u16,
                HistoricalOrderFillsRejectVLS::new().r#type(),
            );
            let d = HistoricalOrderFillsRejectVLSData::new();
            let p = (&d as *const _ as *const u8).offset(0) as usize;
            assert_eq!(
                0usize,
                (core::ptr::addr_of!(d.size) as usize) - p,
                "size offset expected {:} but was {:}",
                0usize,
                (core::ptr::addr_of!(d.size) as usize) - p,
            );
            assert_eq!(
                2usize,
                (core::ptr::addr_of!(d.r#type) as usize) - p,
                "type offset expected {:} but was {:}",
                2usize,
                (core::ptr::addr_of!(d.r#type) as usize) - p,
            );
            assert_eq!(
                4usize,
                (core::ptr::addr_of!(d.base_size) as usize) - p,
                "base_size offset expected {:} but was {:}",
                4usize,
                (core::ptr::addr_of!(d.base_size) as usize) - p,
            );
            assert_eq!(
                8usize,
                (core::ptr::addr_of!(d.request_id) as usize) - p,
                "request_id offset expected {:} but was {:}",
                8usize,
                (core::ptr::addr_of!(d.request_id) as usize) - p,
            );
            assert_eq!(
                12usize,
                (core::ptr::addr_of!(d.reject_text) as usize) - p,
                "reject_text offset expected {:} but was {:}",
                12usize,
                (core::ptr::addr_of!(d.reject_text) as usize) - p,
            );
        }
    }
}
