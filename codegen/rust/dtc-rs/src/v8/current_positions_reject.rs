// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-27 08:32:53.747629 +0800 WITA m=+0.008238335
use super::*;

pub(crate) const CURRENT_POSITIONS_REJECT_VLS_SIZE: usize = 16;

pub(crate) const CURRENT_POSITIONS_REJECT_FIXED_SIZE: usize = 104;

/// size         u16     = CurrentPositionsRejectVLSSize  (16)
/// type         u16     = CURRENT_POSITIONS_REJECT  (307)
/// base_size    u16     = CurrentPositionsRejectVLSSize  (16)
/// request_id   i32     = 0
/// reject_text  string  = ""
pub(crate) const CURRENT_POSITIONS_REJECT_VLS_DEFAULT: [u8; 16] =
    [16, 0, 51, 1, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

/// size         u16       = CurrentPositionsRejectFixedSize  (104)
/// type         u16       = CURRENT_POSITIONS_REJECT  (307)
/// request_id   i32       = 0
/// reject_text  string96  = ""
pub(crate) const CURRENT_POSITIONS_REJECT_FIXED_DEFAULT: [u8; 104] = [
    104, 0, 51, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0,
];

/// If the Server is unable to serve the request for an CurrentPositionsRequestVLS
/// message received, for a reason other than there not being any current
/// Trade positions, then send this message to the Client.
///
/// This must never be sent when there are actually no Trade Positions in
/// the account or accounts requested.
pub trait CurrentPositionsReject: Message {
    type Safe: CurrentPositionsReject;
    type Unsafe: CurrentPositionsReject;

    /// This is set to the RequestID field sent in the CurrentPositionsRequestVLS
    /// message.
    fn request_id(&self) -> i32;

    /// Free-form text indicating the reason for the rejection.
    ///
    fn reject_text(&self) -> &str;

    /// This is set to the RequestID field sent in the CurrentPositionsRequestVLS
    /// message.
    fn set_request_id(&mut self, value: i32) -> &mut Self;

    /// Free-form text indicating the reason for the rejection.
    ///
    fn set_reject_text(&mut self, value: &str) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl CurrentPositionsReject) {
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

/// If the Server is unable to serve the request for an CurrentPositionsRequestVLS
/// message received, for a reason other than there not being any current
/// Trade positions, then send this message to the Client.
///
/// This must never be sent when there are actually no Trade Positions in
/// the account or accounts requested.
pub struct CurrentPositionsRejectVLS {
    data: *const CurrentPositionsRejectVLSData,
    capacity: usize,
}

pub struct CurrentPositionsRejectVLSUnsafe {
    data: *const CurrentPositionsRejectVLSData,
    capacity: usize,
}

#[repr(packed(8), C)]
pub struct CurrentPositionsRejectVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    request_id: i32,
    reject_text: VLS,
}

/// If the Server is unable to serve the request for an CurrentPositionsRequestVLS
/// message received, for a reason other than there not being any current
/// Trade positions, then send this message to the Client.
///
/// This must never be sent when there are actually no Trade Positions in
/// the account or accounts requested.
pub struct CurrentPositionsRejectFixed {
    data: *const CurrentPositionsRejectFixedData,
}

pub struct CurrentPositionsRejectFixedUnsafe {
    data: *const CurrentPositionsRejectFixedData,
}

#[repr(packed(8), C)]
pub struct CurrentPositionsRejectFixedData {
    size: u16,
    r#type: u16,
    request_id: i32,
    reject_text: [u8; 96],
}

impl CurrentPositionsRejectVLSData {
    pub fn new() -> Self {
        Self {
            size: 16u16.to_le(),
            r#type: CURRENT_POSITIONS_REJECT.to_le(),
            base_size: 16u16.to_le(),
            request_id: 0i32,
            reject_text: crate::message::VLS::new(),
        }
    }
}

impl CurrentPositionsRejectFixedData {
    pub fn new() -> Self {
        Self {
            size: 104u16.to_le(),
            r#type: CURRENT_POSITIONS_REJECT.to_le(),
            request_id: 0i32,
            reject_text: [0; 96],
        }
    }
}

unsafe impl Send for CurrentPositionsRejectFixed {}
unsafe impl Send for CurrentPositionsRejectFixedUnsafe {}
unsafe impl Send for CurrentPositionsRejectFixedData {}
unsafe impl Send for CurrentPositionsRejectVLS {}
unsafe impl Send for CurrentPositionsRejectVLSUnsafe {}
unsafe impl Send for CurrentPositionsRejectVLSData {}

impl Drop for CurrentPositionsRejectFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for CurrentPositionsRejectFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for CurrentPositionsRejectVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for CurrentPositionsRejectVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for CurrentPositionsRejectFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for CurrentPositionsRejectFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for CurrentPositionsRejectVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for CurrentPositionsRejectVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for CurrentPositionsRejectFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for CurrentPositionsRejectFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for CurrentPositionsRejectVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for CurrentPositionsRejectVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for CurrentPositionsRejectFixed {
    type Target = CurrentPositionsRejectFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for CurrentPositionsRejectFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for CurrentPositionsRejectFixedUnsafe {
    type Target = CurrentPositionsRejectFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for CurrentPositionsRejectFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for CurrentPositionsRejectVLS {
    type Target = CurrentPositionsRejectVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for CurrentPositionsRejectVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for CurrentPositionsRejectVLSUnsafe {
    type Target = CurrentPositionsRejectVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for CurrentPositionsRejectVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::fmt::Display for CurrentPositionsRejectFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("CurrentPositionsRejectFixed(size: {}, type: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.request_id(), self.reject_text()).as_str())
    }
}

impl core::fmt::Debug for CurrentPositionsRejectFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("CurrentPositionsRejectFixed(size: {}, type: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.request_id(), self.reject_text()).as_str())
    }
}

impl core::fmt::Display for CurrentPositionsRejectFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("CurrentPositionsRejectFixedUnsafe(size: {}, type: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.request_id(), self.reject_text()).as_str())
    }
}

impl core::fmt::Debug for CurrentPositionsRejectFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("CurrentPositionsRejectFixedUnsafe(size: {}, type: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.request_id(), self.reject_text()).as_str())
    }
}

impl core::fmt::Display for CurrentPositionsRejectVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("CurrentPositionsRejectVLS(size: {}, type: {}, base_size: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.reject_text()).as_str())
    }
}

impl core::fmt::Debug for CurrentPositionsRejectVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("CurrentPositionsRejectVLS(size: {}, type: {}, base_size: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.reject_text()).as_str())
    }
}

impl core::fmt::Display for CurrentPositionsRejectVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("CurrentPositionsRejectVLSUnsafe(size: {}, type: {}, base_size: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.reject_text()).as_str())
    }
}

impl core::fmt::Debug for CurrentPositionsRejectVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("CurrentPositionsRejectVLSUnsafe(size: {}, type: {}, base_size: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.reject_text()).as_str())
    }
}

impl crate::Message for CurrentPositionsRejectFixed {
    type Data = CurrentPositionsRejectFixedData;

    const TYPE: u16 = CURRENT_POSITIONS_REJECT;
    const BASE_SIZE: usize = 104;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, CurrentPositionsRejectFixedData::new()),
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
            data: data as *const CurrentPositionsRejectFixedData,
        }
    }
}
impl crate::Message for CurrentPositionsRejectFixedUnsafe {
    type Data = CurrentPositionsRejectFixedData;

    const TYPE: u16 = CURRENT_POSITIONS_REJECT;
    const BASE_SIZE: usize = 104;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, CurrentPositionsRejectFixedData::new()),
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
            data: data as *const CurrentPositionsRejectFixedData,
        }
    }
}
impl crate::Message for CurrentPositionsRejectVLS {
    type Data = CurrentPositionsRejectVLSData;

    const TYPE: u16 = CURRENT_POSITIONS_REJECT;
    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, CurrentPositionsRejectVLSData::new()),
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
            data: data as *const CurrentPositionsRejectVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for CurrentPositionsRejectVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const CurrentPositionsRejectVLSData;
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
impl crate::Message for CurrentPositionsRejectVLSUnsafe {
    type Data = CurrentPositionsRejectVLSData;

    const TYPE: u16 = CURRENT_POSITIONS_REJECT;
    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, CurrentPositionsRejectVLSData::new()),
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
            data: data as *const CurrentPositionsRejectVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for CurrentPositionsRejectVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const CurrentPositionsRejectVLSData;
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
/// If the Server is unable to serve the request for an CurrentPositionsRequestVLS
/// message received, for a reason other than there not being any current
/// Trade positions, then send this message to the Client.
///
/// This must never be sent when there are actually no Trade Positions in
/// the account or accounts requested.
impl CurrentPositionsReject for CurrentPositionsRejectVLS {
    type Safe = CurrentPositionsRejectVLS;
    type Unsafe = CurrentPositionsRejectVLSUnsafe;

    /// This is set to the RequestID field sent in the CurrentPositionsRequestVLS
    /// message.
    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    /// Free-form text indicating the reason for the rejection.
    ///
    fn reject_text(&self) -> &str {
        get_vls(self, self.reject_text)
    }

    /// This is set to the RequestID field sent in the CurrentPositionsRequestVLS
    /// message.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }

    /// Free-form text indicating the reason for the rejection.
    ///
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

/// If the Server is unable to serve the request for an CurrentPositionsRequestVLS
/// message received, for a reason other than there not being any current
/// Trade positions, then send this message to the Client.
///
/// This must never be sent when there are actually no Trade Positions in
/// the account or accounts requested.
impl CurrentPositionsReject for CurrentPositionsRejectVLSUnsafe {
    type Safe = CurrentPositionsRejectVLS;
    type Unsafe = CurrentPositionsRejectVLSUnsafe;

    /// This is set to the RequestID field sent in the CurrentPositionsRequestVLS
    /// message.
    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(12) {
            0i32
        } else {
            i32::from_le(self.request_id)
        }
    }

    /// Free-form text indicating the reason for the rejection.
    ///
    fn reject_text(&self) -> &str {
        if self.is_out_of_bounds(16) {
            ""
        } else {
            get_vls(self, self.reject_text)
        }
    }

    /// This is set to the RequestID field sent in the CurrentPositionsRequestVLS
    /// message.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(12) {
            self.request_id = value.to_le();
        }
        self
    }

    /// Free-form text indicating the reason for the rejection.
    ///
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

/// If the Server is unable to serve the request for an CurrentPositionsRequestVLS
/// message received, for a reason other than there not being any current
/// Trade positions, then send this message to the Client.
///
/// This must never be sent when there are actually no Trade Positions in
/// the account or accounts requested.
impl CurrentPositionsReject for CurrentPositionsRejectFixed {
    type Safe = CurrentPositionsRejectFixed;
    type Unsafe = CurrentPositionsRejectFixedUnsafe;

    /// This is set to the RequestID field sent in the CurrentPositionsRequestVLS
    /// message.
    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    /// Free-form text indicating the reason for the rejection.
    ///
    fn reject_text(&self) -> &str {
        get_fixed(&self.reject_text[..])
    }

    /// This is set to the RequestID field sent in the CurrentPositionsRequestVLS
    /// message.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }

    /// Free-form text indicating the reason for the rejection.
    ///
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

/// If the Server is unable to serve the request for an CurrentPositionsRequestVLS
/// message received, for a reason other than there not being any current
/// Trade positions, then send this message to the Client.
///
/// This must never be sent when there are actually no Trade Positions in
/// the account or accounts requested.
impl CurrentPositionsReject for CurrentPositionsRejectFixedUnsafe {
    type Safe = CurrentPositionsRejectFixed;
    type Unsafe = CurrentPositionsRejectFixedUnsafe;

    /// This is set to the RequestID field sent in the CurrentPositionsRequestVLS
    /// message.
    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(8) {
            0i32
        } else {
            i32::from_le(self.request_id)
        }
    }

    /// Free-form text indicating the reason for the rejection.
    ///
    fn reject_text(&self) -> &str {
        if self.is_out_of_bounds(104) {
            ""
        } else {
            get_fixed(&self.reject_text[..])
        }
    }

    /// This is set to the RequestID field sent in the CurrentPositionsRequestVLS
    /// message.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.request_id = value.to_le();
        }
        self
    }

    /// Free-form text indicating the reason for the rejection.
    ///
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
                core::mem::size_of::<CurrentPositionsRejectFixedData>(),
                "CurrentPositionsRejectFixedData sizeof expected {:} but was {:}",
                104usize,
                core::mem::size_of::<CurrentPositionsRejectFixedData>()
            );
            assert_eq!(
                104u16,
                CurrentPositionsRejectFixed::new().size(),
                "CurrentPositionsRejectFixed sizeof expected {:} but was {:}",
                104u16,
                CurrentPositionsRejectFixed::new().size(),
            );
            assert_eq!(
                CURRENT_POSITIONS_REJECT,
                CurrentPositionsRejectFixed::new().r#type(),
                "CurrentPositionsRejectFixed type expected {:} but was {:}",
                CURRENT_POSITIONS_REJECT,
                CurrentPositionsRejectFixed::new().r#type(),
            );
            assert_eq!(
                307u16,
                CurrentPositionsRejectFixed::new().r#type(),
                "CurrentPositionsRejectFixed type expected {:} but was {:}",
                307u16,
                CurrentPositionsRejectFixed::new().r#type(),
            );
            let d = CurrentPositionsRejectFixedData::new();
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
                core::mem::size_of::<CurrentPositionsRejectVLSData>(),
                "CurrentPositionsRejectVLSData sizeof expected {:} but was {:}",
                16usize,
                core::mem::size_of::<CurrentPositionsRejectVLSData>()
            );
            assert_eq!(
                16u16,
                CurrentPositionsRejectVLS::new().size(),
                "CurrentPositionsRejectVLS sizeof expected {:} but was {:}",
                16u16,
                CurrentPositionsRejectVLS::new().size(),
            );
            assert_eq!(
                CURRENT_POSITIONS_REJECT,
                CurrentPositionsRejectVLS::new().r#type(),
                "CurrentPositionsRejectVLS type expected {:} but was {:}",
                CURRENT_POSITIONS_REJECT,
                CurrentPositionsRejectVLS::new().r#type(),
            );
            assert_eq!(
                307u16,
                CurrentPositionsRejectVLS::new().r#type(),
                "CurrentPositionsRejectVLS type expected {:} but was {:}",
                307u16,
                CurrentPositionsRejectVLS::new().r#type(),
            );
            let d = CurrentPositionsRejectVLSData::new();
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
